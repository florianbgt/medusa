import { api } from "@/api/api"

export default function Home() {
  api({ url: "private", method: "GET"})
  return (
    <>
      <main>
        index
      </main>
    </>
  )
}
