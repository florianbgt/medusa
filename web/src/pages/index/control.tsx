import { useState } from "react"
import Stream from "../../components/stream"
import Button from "../../components/ui/button"
import Input from "../../components/ui/input"
import { api } from "../../api"

export default function Control() {
    const [loading, setLoading] = useState<boolean>(false)
    const [asisStep, setAxisStep] = useState<0.1 | 1 | 10 | 100>(0.1)
    const [extruderStep, setExtruderStep] = useState<number| 1 | 10 | 100>(0.1)

    async function move({x, y, z}: {x: number, y: number, z: number}) {
        setLoading(true)
        try {
            await api({url: "/control/move", method: "POST", data: {x, y, z}})
        } catch (error) {
            throw error
        } finally {
            setLoading(false)
        }
    }

    async function extrude(e: number) {
        setLoading(true)
        try {
            await api({url: "/control/extrude", method: "POST", data: {e: e}})
        } catch (error) {
            throw error
        } finally {
            setLoading(false)
        }
    }

    async function goHome() {
        setLoading(true)
        try {
            await api({url: "/control/home", method: "POST"})
        } catch (error) {
            throw error
        } finally {
            setInterval(() => setLoading(false), 500)
        }
    }

    async function autoLevel() {
        setLoading(true)
        try {
            await api({url: "/control/level", method: "POST"})
        } catch (error) {
            throw error
        } finally {
            setInterval(() => setLoading(false), 500)
        }
    }

    return (
        <div className="flex flex-col">
            <Stream/>
            <div className="flex flex-wrap justify-around  items-center m-3 gap-5">
                <div className="grid grid-cols-3 gap-2">
                    <div/>
                    <Button
                        onClick={() => move({x: 0, y: asisStep, z: 0})}
                        color="primary"
                        size="sm"
                        disabled={loading}
                    >
                        <svg width="50px" height="50px" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg">
                            <path fill="#f5f5f5" d="M512 320 192 704h639.936z"/>
                        </svg>
                    </Button>
                    <div/>
                    <Button
                        onClick={() => move({x: -asisStep, y: 0, z: 0})}
                        color="primary"
                        size="sm"
                        disabled={loading}
                    >
                        <svg width="50px" height="50px" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg">
                            <path fill="#f5f5f5" d="M672 192 288 511.936 672 832z"/>
                        </svg>
                    </Button>
                    <Button
                        onClick={() => goHome()}
                        color="primary"
                        size="sm"
                        disabled={loading}
                    >
                        <svg width="50px" height="50px" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg">
                            <path fill="#f5f5f5" d="M512 128 128 447.936V896h255.936V640H640v256h255.936V447.936z"/>
                        </svg>
                    </Button>
                    <Button
                        onClick={() => move({x: asisStep, y: 0, z: 0})}
                        color="primary"
                        size="sm"
                        disabled={loading}
                    >
                        <svg width="50px" height="50px" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg">
                            <path fill="#f5f5f5" d="M384 192v640l384-320.064z"/>
                        </svg>
                    </Button>
                    <div/>
                    <Button
                        onClick={() => move({x: 0, y: -asisStep, z: 0})}
                        color="primary"
                        size="sm"
                        disabled={loading}
                    >
                        <svg width="50px" height="50px" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg">
                            <path fill="#f5f5f5" d="m192 384 320 384 320-384z"/>
                        </svg>
                    </Button>
                    <div/>
                </div>
                <div className="grid grid-cols-1 gap-2">
                    <Button
                        onClick={() => move({x: 0, y: 0, z: asisStep})}
                        color="primary"
                        size="sm"
                        disabled={loading}
                    >
                        <svg width="50px" height="50px" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg">
                            <path fill="#f5f5f5" d="M512 320 192 704h639.936z"/>
                        </svg>       
                    </Button>
                    <Button
                        onClick={() => autoLevel()}
                        color="primary"
                        size="sm"
                        disabled={loading}
                    >
                        <svg fill="#f5f5f5" xmlns="http://www.w3.org/2000/svg" width="50px" height="50px" viewBox="0 0 360.177 360.177">
                            <g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"/>
                            <g id="SVGRepo_iconCarrier">
                                <path d="M0,144.308v71.561h360.177v-71.561H0z M238.05,175.241c0,1.483-1.197,2.688-2.688,2.688H131.647 c-1.483,0-2.688-1.205-2.688-2.688v-22.798c0-1.483,1.205-2.688,2.688-2.688h103.714c1.491,0,2.688,1.205,2.688,2.688V175.241z"/>
                                <rect x="167.387" y="155.131" width="29.958" height="17.422"/>
                                <rect x="202.721" y="155.131" width="29.953" height="17.422"/>
                                <rect x="134.334" y="155.131" width="27.678" height="17.422"/>
                            </g>
                        </svg>
                        {/* <svg width="50px" height="50px" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg">
                            <path fill="#f5f5f5" d="M512 128 128 447.936V896h255.936V640H640v256h255.936V447.936z"/>
                        </svg> */}
                    </Button>
                    <Button
                        onClick={() => move({x: 0, y: 0, z: -asisStep})}
                        color="primary"
                        size="sm"
                        disabled={loading}
                    >
                        <svg width="50px" height="50px" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg">
                            <path fill="#f5f5f5" d="m192 384 320 384 320-384z"/>
                        </svg>      
                    </Button>
                </div>
                <div className="flex flex-col gap-2">
                    <div className="text-lg font-bold">
                        XYZ Step:
                    </div>
                    <Button
                        color="primary"
                        size="md"
                        disabled={asisStep === 0.1}
                        onClick={() => setAxisStep(0.1)}
                    >
                        0.1
                    </Button>
                    <Button
                        color="primary"
                        size="md"
                        disabled={asisStep === 1}
                        onClick={() => setAxisStep(1)}
                    >
                        1
                    </Button>
                    <Button
                        color="primary"
                        size="md"
                        disabled={asisStep === 10}
                        onClick={() => setAxisStep(10)}
                    >
                        10
                    </Button>
                    <Button
                        color="primary"
                        size="md"
                        disabled={asisStep === 100}
                        onClick={() => setAxisStep(100)}
                    >
                        100
                    </Button>
                </div>
                <div className="flex flex-col gap-2">
                    <div className="text-lg font-bold">
                        Extruder Step:
                    </div>
                    <div>
                        <Input
                            id="extrudeStep"
                            type="number"
                            min={0.1}
                            max={10}
                            step={0.1}
                            value={extruderStep}
                            onChange={(e) => setExtruderStep(parseFloat(e.target.value))}
                        />
                    </div>
                    <Button
                        onClick={() => extrude(extruderStep)}
                        color="primary"
                        size="md"
                        disabled={loading}
                    >
                        Extrude
                    </Button>
                    <Button
                        onClick={() => extrude(-extruderStep)}
                        color="primary"
                        size="md"
                        disabled={loading}
                    >
                        Retract
                    </Button>
                </div>
            </div>
        </div>
    )
}  