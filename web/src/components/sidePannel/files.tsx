import { useEffect, useState } from "react";
import Button from "../ui/button"
import Section from "./section";
import FileUpload from "./fileUpload";
import { api } from "../../api";

interface File{
    name: string
    uploaded: string
    size: number
}

export default function Files() {
    const [files, setFiles] = useState<Array<File>>([])

    useEffect(() => {
        fetchFiles()
    }, [])

    function formatFileSize(size: number) {
        if (size < 1024) {
            return `${size} B`
        } else if (size < 1024 * 1024) {
            return `${(size / 1024).toFixed(2)} KB`
        } else if (size < 1024 * 1024 * 1024) {
            return `${(size / 1024 / 1024).toFixed(2)} MB`
        } else {
            return `${(size / 1024 / 1024 / 1024).toFixed(2)} GB`
        }
    }

    async function fetchFiles() {
        const url = "/files"
        try{
            const files: Array<File>  = await api({url})
            setFiles(files)
        } catch (error: any) {
            throw error
        }
    }

    function File({file}: {file: File}) {
        async function deleteFile() {
            const url = `/files/${encodeURIComponent(file.name)}`
            try{
                await api({url, method: "DELETE"})
                await fetchFiles()
            } catch (error: any) {
                throw error
            }
        }

        return (
            <div className="flex flex-col gap-1">
                <div className="text-primary text-xl font-bold">
                    {file.name}
                </div>
                <div>
                    Uploaded: <span className="font-bold">
                        {new Date(file.uploaded).toLocaleDateString()}
                    </span>
                </div>
                <div>
                    Size: <span className="font-bold">
                        {formatFileSize(file.size)}
                    </span>
                </div>
                <div className="flex gap-1">
                    <Button size="sm" color="primary" className="w-full">
                        Load
                    </Button>
                    <Button size="sm" color="light" className="w-full">
                        View G-code
                    </Button>
                    <Button onClick={deleteFile} size="sm" color="accent" className="w-full">
                        Delete
                    </Button>
                </div>
            </div>
        )
    }

    return (
        <Section name="Files">
            <div className="py-4">
                <FileUpload onUploaded={fetchFiles}/>
            </div>
            <div className="max-h-[450px] overflow-y-auto overflow-x-none flex flex-col gap-3 my-2">
                {files.length === 0 ? (
                    <div className="text-center">No files</div>
                ) : (
                    files.map((file) => {return <File key={file.name} file={file}/>})
                )}
            </div>
        </Section>
    );
}
