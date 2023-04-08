import { useState } from "react";
import { api } from "@/api/api";

interface Props extends React.InputHTMLAttributes<HTMLInputElement> {
    onUploaded: () => Promise<void>
}

export default function FileUpload({onUploaded}: Props) {
    async function handleChange(event: React.ChangeEvent<HTMLInputElement>) {
        if (event.target.files?.length) {
            const file = event.target.files[0]
            const body = new FormData()
            body.append("file", file)
            try{
                await api({
                    url: "/files",
                    method: "POST",
                    body,
                    headers: {
                        "Content-Type": "multipart/form-data"
                    }
                })
                await onUploaded()
                event.target.value = ""
            } catch (error: any) {
                throw error
            }
        }
    }

    return (
        <input
            type="file"
            accept=".gcode"
            onChange={handleChange}
        />
    );
}
