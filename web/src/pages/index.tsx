import Stream from "@/components/stream";
import Main from "@/components/layouts/main"
import SidePannel from "@/components/sidePannel/sidePannel";

export default function Home() {
  return (
    <Main>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-2">
        <div className="rounded-xl border-2 border-primary shadow-xl">
          <SidePannel/>
        </div>
        <div className="col-span-2">
          <Stream/>
        </div>
      </div>
    </Main>
  );
}
