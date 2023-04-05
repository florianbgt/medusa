import { api } from "@/api/api"
import Main from "@/components/layouts/main"

export default function Home() {
  api({ url: "/private", method: "GET"})
  return (
    <Main>
      <div>
        index
      </div>
    </Main>
  )
}
