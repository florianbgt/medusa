import Current from "@/components/sidePannel/current";
import Files from "@/components/sidePannel/files";
import System from "@/components/sidePannel/system";

export default function SidePannel() {
  return (
    <div className="m-3 flex flex-col gap-2">
        <System/>
        <Current/>
        <Files/>
    </div>
  );
}
