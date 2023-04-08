import Page from "@/components/layouts/page";
import SidePannel from "@/components/sidePannel/sidePannel";
import { useRouter } from "next/router";
import Link from "next/link";

interface Props {
    children: React.ReactNode;
}

export default function Main({children}: Props) {
    const router = useRouter();
    const currentPath = router.pathname;

    function getClassName(path: string) {
        const base = "h-full px-4 py-1 text-lg font-bold underline";
        if (currentPath === path) {
            return `${base} bg-primary`;
        }
        return `${base} hover:bg-primary/50`;
    }

    return (
        <Page>
            <div className="grid grid-cols-1 md:grid-cols-3 border-2 border-primary rounded-xl">
                <div className="md:border-r-2 border-primary">
                    <SidePannel/>
                </div>
                <div className="col-span-2 flex flex-col">
                    <div className="flex bg-secondary border-b-2 border-primary sm:rounded-tr-xl">
                        <Link className={getClassName("/temperature")} href="/temperature">
                            Temperature
                        </Link>
                        <Link className={getClassName("/control")} href="/control">
                            Control
                        </Link>
                        <Link className={getClassName("/gcode")} href="gcode">
                            Gcode
                        </Link>
                    </div>
                    <div className="flex flex-col p-1">
                        {children}
                    </div>
                </div>
            </div>
        </Page>
    );
}
