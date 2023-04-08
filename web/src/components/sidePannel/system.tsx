import { useEffect, useState } from "react";
import Section from "@/components/sidePannel/section";
import { api } from "@/api/api";

export default function System() {
    const [state, setState] = useState({
        cpuLoad: "0%",
        cpuTemp: "0°C",
    })

    useEffect(() => {
        let timer = setTimeout(getSystemInfo, 0)
        async function getSystemInfo() {
            const url = "/system/metrics"
    
            try {
                let {cpu_load, cpu_temp}: {cpu_load: number, cpu_temp: number} = await api({ url, method: "GET"})
                cpu_load = Math.round(cpu_load*100)/100
                cpu_temp = Math.round(cpu_temp*100)/100
                setState({
                    cpuLoad: `${cpu_load}%`,
                    cpuTemp: `${cpu_temp}°C`,
                })
                timer = setTimeout(getSystemInfo, 3000)
            } catch (error: any) {
            throw error
            }
        }
        return () => {
            clearTimeout(timer)
        }
    }, [])

    return (
        <Section name="System">
            <div className="flex flex-col gap-2 my-2">
                <div>
                    CPU load: <span className="font-bold">{state.cpuLoad}</span>
                </div>
                <div>
                    CPU temp: <span className="font-bold">{state.cpuTemp}</span>
                </div>
            </div>
        </Section>
    );
}