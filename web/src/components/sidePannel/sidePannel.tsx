import Current from "./current";
import Files from "./files";
import System from "./system";

export default function SidePannel() {
  return (
    <div className="flex flex-col gap-2">
        <System/>
        <Current/>
        <Files/>
    </div>
  );
}
