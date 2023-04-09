import { useEffect, useState } from "react";
import Button from "../ui/button"
import Section from "./section";

export default function Current() {
    const [state, setState] = useState({
        file: "test.gcode",
        timelaspse: "",
        filament: "",
        totalTime: "",
        currentTime: "",
        leftTime: "",
    })

    useEffect(() => {
        let timer = setTimeout(pollState, 0)
        function pollState() {
            console.log("poll state")
            timer = setTimeout(pollState, 1000)
        }
        return () => {
            clearTimeout(timer)
        }
    }, [])

    function print() {
        console.log("print")
    }

    function pause() {
        console.log("pause")
    }

    function cancel() {
        console.log("cancel")
    }

    return (
        <Section name="Job">
            <div className="flex flex-col gap-2 my-2">
                <div className="flex flex-col gap-1">
                    <div>
                        File: <span className="font-bold">{state.file}</span>
                    </div>
                    <div>
                        Timelapse: <span className="font-bold">{state.timelaspse}</span>
                    </div>
                    <div>
                        Filament: <span className="font-bold">{state.filament}</span>
                    </div>
                    <div>
                        Total Time: <span className="font-bold">{state.totalTime}</span>
                    </div>
                    <div>
                        Current Time: <span className="font-bold">{state.currentTime}</span>
                    </div>
                    <div>
                        Left Time: <span className="font-bold">{state.leftTime}</span>
                    </div>
                </div>
                <div className="flex gap-2">
                    <Button onClick={print} color="primary" size="sm" className="w-full">
                        Print
                    </Button>
                    <Button onClick={pause} color="light" size="sm" className="w-full">
                        Pause
                    </Button>
                    <Button onClick={cancel} color="accent" size="sm" className="w-full">
                        Cancel
                    </Button>
                </div>
            </div>
        </Section>
    );
}
