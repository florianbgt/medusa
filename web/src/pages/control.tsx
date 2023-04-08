import Main from "@/components/layouts/main";
import Stream from "@/components/stream";

export default function Home() {
  return (
    <Main>
      <div className="flex flex-col gap-2">
        <Stream/>
        <div>
          Control
        </div>
      </div>
    </Main>
  );
}
