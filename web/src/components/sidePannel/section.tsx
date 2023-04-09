import { useState } from "react";
import Button from "../ui/button"

interface Props {
    name: string,
    children: React.ReactNode
}

export default function Section({name, children}: Props) {
    return (
        <>
            <div className="bg-secondary text-lg font-bold text-center py-1 mx-4">
                {name}
            </div>
            <div className="px-4">
                {children}
            </div>
        </>
    );
}
