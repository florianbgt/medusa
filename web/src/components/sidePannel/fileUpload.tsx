import { useRef, useState } from "react";
import { api } from "../../api";

interface Props extends React.InputHTMLAttributes<HTMLInputElement> {
    onUploaded: () => Promise<void>
}

export default function FileUpload({onUploaded}: Props) {
    const inputRef = useRef<HTMLInputElement>(null)
    const [dragActive, setDragActive] = useState<boolean>(false);

    const handleDrag = function(event: React.DragEvent<HTMLLabelElement | HTMLDivElement>) {
        event.preventDefault();
        event.stopPropagation();
        if (event.type === "dragenter" || event.type === "dragover") {
            setDragActive(true);
        } else if (event.type === "dragleave") {
            setDragActive(false);
        }
    };

    const handleDrop = function(event: React.DragEvent<HTMLDivElement>) {
        event.preventDefault();
        event.stopPropagation();
        setDragActive(false);
        const file = event.dataTransfer.files?.item(0);
        if (file) {
            uploadFile(file)
        }
      };

    function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
        const file = e.target.files?.item(0)
        if (file) {
            uploadFile(file)
        }
    }

    async function uploadFile(file: File) {
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
        } catch (error: any) {
            throw error
        }
    }

    return (
        <div className="flex">
            <label
                htmlFor="file-upload"
                className="p-5 border-2 border-dashed border-primary rounded-xl w-full text-center"
                onDragEnter={handleDrag}
            >
                Drag and Drop Here or <button className="underline py-2" onClick={() => inputRef.current?.click()}>browse</button>
            </label>
            <input
                ref={inputRef}
                id="file-upload"
                type="file"
                accept=".gcode"
                onChange={handleChange}
                className="hidden"
            />
            { dragActive && (
                <div
                    className="absolute bottom-0 top-0 left-0 right-0 w-screen h-screen"
                    onDragEnter={handleDrag}
                    onDragLeave={handleDrag}
                    onDragOver={handleDrag}
                    onDrop={handleDrop}
                />
            ) }
        </div>
    );
}
