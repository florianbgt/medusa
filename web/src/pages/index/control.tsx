import { useState } from "react"
import Stream from "../../components/stream"
import Button from "../../components/ui/button"
import Input from "../../components/ui/input"

export default function Control() {
    const [loading, setLoading] = useState<boolean>(false)
    const [asisStep, setAxisStep] = useState<0.1 | 1 | 10 | 100>(0.1)
    const [extruderStep, setExtruderStep] = useState<number| 1 | 10 | 100>(0.1)

    type move = {x: number, y: number, z: number, e: number}

    async function sendGCode({x, y, z, e}: move) {
        setLoading(true)
        try {
            // TODO implement call
            console.log(x, y, z, e)
        } catch (error) {
            // TODO handle error
            throw error
        } finally {
            setInterval(() => setLoading(false), 500)
        }
    }

    type home = {x: boolean, y: boolean, z: boolean}

    async function goHome({x, y, z}: home) {
        setLoading(true)
        try {
            // TODO implement call
            console.log(x, y, z)
        } catch (error) {
            // TODO handle error
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
                        onClick={() => sendGCode({x: 0, y: asisStep, z: 0, e: 0})}
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
                        onClick={() => sendGCode({x: -asisStep, y: 0, z: 0, e: 0})}
                        color="primary"
                        size="sm"
                        disabled={loading}
                    >
                        <svg width="50px" height="50px" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg">
                            <path fill="#f5f5f5" d="M672 192 288 511.936 672 832z"/>
                        </svg>
                    </Button>
                    <Button
                        onClick={() => goHome({x: true, y: true, z: false})}
                        color="primary"
                        size="sm"
                        disabled={loading}
                    >
                        <svg width="50px" height="50px" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg">
                            <path fill="#f5f5f5" d="M512 128 128 447.936V896h255.936V640H640v256h255.936V447.936z"/>
                        </svg>
                    </Button>
                    <Button
                        onClick={() => sendGCode({x: asisStep, y: 0, z: 0, e: 0})}
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
                        onClick={() => sendGCode({x: 0, y: -asisStep, z: 0, e: 0})}
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
                        onClick={() => sendGCode({x: 0, y: 0, z: asisStep, e: 0})}
                        color="primary"
                        size="sm"
                        disabled={loading}
                    >
                        <svg width="50px" height="50px" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg">
                            <path fill="#f5f5f5" d="M512 320 192 704h639.936z"/>
                        </svg>       
                    </Button>
                    <Button
                        onClick={() => goHome({x: false, y: false, z: true})}
                        color="primary"
                        size="sm"
                        disabled={loading}
                    >
                        <svg width="50px" height="50px" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg">
                            <path fill="#f5f5f5" d="M512 128 128 447.936V896h255.936V640H640v256h255.936V447.936z"/>
                        </svg>
                    </Button>
                    <Button
                        onClick={() => sendGCode({x: 0, y: 0, z: -asisStep, e: 0})}
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
                        onClick={() => sendGCode({x: 0, y: 0, z: 0, e: extruderStep})}
                        color="primary"
                        size="md"
                        disabled={loading}
                    >
                        Extrude
                    </Button>
                    <Button
                        onClick={() => sendGCode({x: 0, y: 0, z: 0, e: -extruderStep})}
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