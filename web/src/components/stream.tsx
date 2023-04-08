import { api, baseUrl } from "@/api/api";
import { useEffect, useState } from "react";

export default function Stream() {
  const [authChecked, setAuthChecked] = useState<boolean>(false);

  // We need to trigger an api call to refresh the token if it is expired
  // Else, the image src will not load
  async function checkAuthenticated() {
    await api({url: "/authenticated", method: "GET"})
    setAuthChecked(true)
  }

  useEffect(() => {
    if (!authChecked) {
      checkAuthenticated()
    }
  }, [authChecked]);

  const streamUrl = `${baseUrl}/stream?token=${localStorage.getItem("access")}`;

  return (
    <>
      {authChecked && (
        <>
          {/* eslint-disable-next-line @next/next/no-img-element */}
          <img
            className="rounded-xl border-2 border-primary shadow-xl scale-x-[-1]"
            width={1000}
            alt="stream"
            src={streamUrl}
          />
        </>
      )}
    </>
  );
}