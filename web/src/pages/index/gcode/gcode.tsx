import { useEffect, useState } from "react";
import { GCodeViewer } from "react-gcode-viewer";
import { useNavigation, useSearchParams } from "react-router-dom";
import { api, baseUrl } from "../../../api";

export default function GCode() {
    const [authChecked, setAuthChecked] = useState<boolean>(false);
    const [searchParams, _setSearchParams] = useSearchParams();
    const [url, setUrl] = useState<string>("");
    const [visible, setVisible] = useState<number>(1);
    const [hasFileName, sethasFileName] = useState<boolean>(false);

    // We need to trigger an api call to refresh the token if it is expired
    // Else, the gcode src will not load
    const navigation = useNavigation();
    async function checkAuthenticated() {
        await api({url: "/authenticated"})
        setAuthChecked(true)
    }

    useEffect(() => {
        if (!authChecked) {
          checkAuthenticated()
        }
      }, [authChecked]);
    
    useEffect(() => {
        if (authChecked) {
            if (searchParams.has("name")) {
                sethasFileName(true);
                setUrl(`${baseUrl}/files/${searchParams.get("name")}?token=${localStorage.getItem("access")}`);
            } else {
                sethasFileName(false);
            }
        }
    }, [searchParams, authChecked]);

    return (
        <div className="grow flex flex-col justify-stretch items-stretchborder border-primary">
            <GCodeViewer
                url={url}
                quality={1}
                visible={visible} // put a slider for this
                layerColor="#ffff00"
                topLayerColor="#000000"
                showAxes={true}
                orbitControls={true}
                floorProps={{
                    gridWidth: 225,
                    gridLength: 225,
                }}
                className="w-full grow max-h-[400px] md:max-h-[800px] bg-primary"
            />
            {
                hasFileName ? (
                    <input
                        type="range"
                        value={visible}
                        onChange={(event) => setVisible(+event.target.value)}
                        min={0}
                        max={1}
                        step={0.01}
                        className="mx-2 my-5 h-2 appearance-none bg-light/20 text-primary rounded-full cursor-pointer focus:outline-none"
                    />
                ) : (
                    <div className=" mx-2 my-5 text-center text-accent text-xl font-bold">
                        Click on "View G-code" on one of the file
                    </div>
                )
            }
        </div>
    )
}  