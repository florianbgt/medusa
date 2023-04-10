import { useEffect, useState } from "react";
import { GCodeViewer } from "react-gcode-viewer";
import { useSearchParams } from "react-router-dom";
import { api, baseUrl } from "../../../api";

export default function GCode() {
    interface GCodeInfo {
        file_name: string;
        layer_count: string;
        filament_used: string;
        layer_height: string;
        total_time: string;
        nozzle_temp: string;
        bed_temp: string;
    }

    const [searchParams, _setSearchParams] = useSearchParams();
    const [url, setUrl] = useState<string>("");
    const [visible, setVisible] = useState<number>(1);
    const [hasFile, sethasFile] = useState<boolean>(false);
    const [gcodeInfo, setGcodeInfo] = useState<GCodeInfo>({
        file_name: "",
        layer_count: "",
        filament_used: "",
        layer_height: "",
        total_time: "",
        nozzle_temp: "",
        bed_temp: "",
    })

    async function setParams() {
        try {
            const info: GCodeInfo = await api({url: `/files/${searchParams.get("name")}/gcode/info`})
            setGcodeInfo(info);
            setUrl(`${baseUrl}/files/${searchParams.get("name")}/gcode?token=${localStorage.getItem("access")}`);
            sethasFile(true);
        } catch(err) {
            throw err
        }
    }
    
    useEffect(() => {
        if (searchParams.has("name")) {
            setParams();
        } else {
            sethasFile(false);
        }
    }, [searchParams]);

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
                hasFile ? (
                    <div className="flex flex-col m-2 gap-2">
                        <input
                            type="range"
                            value={visible}
                            onChange={(event) => setVisible(+event.target.value)}
                            min={0}
                            max={1}
                            step={0.01}
                            className="h-2 appearance-none bg-light/20 text-primary rounded-full cursor-pointer focus:outline-none"
                        />
                        <div className="flex flex-wrap gap-5 mt-2">
                            <div>
                                File name: <span className="font-bold">
                                    {gcodeInfo.file_name}
                                </span>
                            </div>
                            <div>
                                Total layer: <span className="font-bold">
                                    {gcodeInfo.layer_count}
                                </span>
                            </div>
                            <div>
                                Filament used: <span className="font-bold">
                                    {gcodeInfo.filament_used}
                                </span>
                            </div>
                            <div>
                                Layer height: <span className="font-bold">
                                    {gcodeInfo.layer_height}mm
                                </span>
                            </div>
                            <div>
                                Total time: <span className="font-bold">
                                    {new Date(parseInt(gcodeInfo.total_time, 10) * 1000).toISOString().substr(11, 8)}
                                </span>
                            </div>
                            <div>
                                Nozzle temp: <span className="font-bold">
                                    {gcodeInfo.nozzle_temp}°C
                                </span>
                            </div>
                            <div>
                                Bed temp: <span className="font-bold">
                                    {gcodeInfo.bed_temp} °C
                                </span>
                            </div>
                        </div>
                    </div>
                    
                ) : (
                    <div className=" mx-2 my-5 text-center text-accent text-xl font-bold">
                        Click on "View G-code" on one of the file
                    </div>
                )
            }
        </div>
    )
}  