import { api } from "@/api/api";
import Button from "@/components/button";
import Input from "@/components/input";
import { useRouter } from "next/router";
import React from "react";
import { useEffect, useRef } from "react";

export default function Home() {
  const [password, setPassword] = React.useState<string>("");
  const passwordRef = useRef<HTMLInputElement>(null);

  const router = useRouter()

  useEffect(() => {
    passwordRef.current?.setCustomValidity("")
  }, [password])

  async function handleSubmit(FormEvent: React.FormEvent<HTMLFormElement>) {
    FormEvent.preventDefault()

    const payload = {
      password: password,
    }
    const url = "login"

    try {
      const {access_token, refresh_token} = await api({ url, method: "POST", data: payload })
      localStorage.setItem("access", access_token)
      localStorage.setItem("refresh", refresh_token)
      const path = router.query.returnUrl as string || "/"
      router.push(path)
    } catch (error: any) {
      if (error.error === "password_incorrect") {
        passwordRef.current?.setCustomValidity("Password is incorrect")
        passwordRef.current?.reportValidity()
        return
      }
      throw error
    }
  }


  return (
    <div className="flex justify-center items-stretch">
      <div className="w-full max-w-lg border-4 rounded-3xl border-secondary bg-light/10 p-5">
        <div className="text-5xl font-bold text-center">
          Login
        </div>
        <div className="text-xl text-center mt-5">
          Welcome to Medusa, your 3D printer assistant 🛠️
        </div>
        <form onSubmit={handleSubmit} className="mt-5 flex flex-col">
          <Input
            elRef={passwordRef}
            id="password"
            label="Password"
            type="password"
            value={password}
            onChange={(event) => setPassword(event.target.value)}
            placeholder="password"
            required
          />
          <Button className="mt-5" color="primary" size="lg" type="submit">Login</Button>
        </form>
      </div>
    </div>
  )
}
