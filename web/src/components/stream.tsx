import { api, baseUrl } from "../api";
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
          <img
            className="scale-x-[-1]"
            width={10000}
            alt="stream"
            src={streamUrl}
          />
        </>
      )}
    </>
  );
}
