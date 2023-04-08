import { api } from "@/api/api";
import { useEffect, useState } from "react";

export default function SystemInfo() {
    const [systemInfo, setSystemInfo] = useState({
        wifi: "",
    })

    useEffect(function() {
        getSystemInfo()
    }, [])

    async function getSystemInfo() {
        const url = "/system"
        try {
            let {ssid}: {ssid: string} = await api({ url, method: "GET"})
            setSystemInfo({
                wifi: ssid,
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
        </div>
        
    )
}
