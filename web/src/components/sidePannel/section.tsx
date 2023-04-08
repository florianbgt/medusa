import { useState } from "react";
import Button from "@/components/ui/button"

interface Props {
    name: string,
    children: React.ReactNode
}

export default function Section({name, children}: Props) {
    return (
        <div>
            <div className="bg-secondary text-xl font-bold text-center p-2">
                {name}
            </div>
            <div>
                {children}
            </div>
        </div>
    );
}
