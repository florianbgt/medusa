import Page from "../../components/layouts/page";
import SidePannel from "../../components/sidePannel/sidePannel";
import { Link, Navigate, Outlet, useLocation } from "react-router-dom";

interface Props {
    children: React.ReactNode;
}

export default function Main() {
    const currentPath = useLocation().pathname;

    function getClassName(path: string) {
        const base = "h-full px-4 py-1 text-lg font-bold underline";
        if (currentPath === path) {
            return `${base} bg-primary`;
        }
        return `${base} hover:bg-primary/50`;
    }

    if (currentPath === "/") return <Navigate to="/temperature" replace={true}/>;

    return (
        <Page>
            <div className="grid grid-cols-1 md:grid-cols-3 border-2 border-primary rounded-xl">
                <div className="md:border-r-2 border-primary">
                    <SidePannel/>
                </div>
                <div className="col-span-2 flex flex-col">
                    <div className="flex bg-secondary border-b-2 border-primary sm:rounded-tr-xl">
                        <Link className={getClassName("/temperature")} to="/temperature">
                            Temperature
                        </Link>
                        <Link className={getClassName("/control")} to="/control">
                            Control
                        </Link>
                        <Link className={getClassName("/gcode")} to="gcode">
                            Gcode
                        </Link>
                    </div>
                    <div className="flex flex-col p-1">
                        <Outlet/>
                    </div>
                </div>
            </div>
        </Page>
    );
}
