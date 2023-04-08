import { api } from "@/api/api";
import { useEffect, useState } from "react";

export default function SystemInfo() {
    const [systemInfo, setSystemInfo] = useState({
        wifi: "",
        host: "",
        arch: "",
    })

    useEffect(function() {
        getSystemInfo()
    }, [])

    async function getSystemInfo() {
        const url = "/system"
        try {
            type Response = {
                ssid: string,
                host: string,
                arch: string,
            }
            let {ssid, host, arch}: Response = await api({ url, method: "GET"})
            setSystemInfo({
                wifi: ssid,
                host: host,
                arch: arch,
            })
        } catch (error: any) {
        throw error
        }
    }
    
    
    return (
        <div>
            <div>
                Wifi: <span className="font-bold">{systemInfo.wifi}</span>
            </div>
            <div>
                Host: <span className="font-bold">{systemInfo.host}</span>
            </div>
            <div>
                Architecture: <span className="font-bold">{systemInfo.arch}</span>
            </div>
        </div>
        
    )
}
