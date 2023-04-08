import { useLocation, useNavigate } from "react-router-dom";
import { api } from "../api";
import Button from "../components/ui/button";
import Input from "../components/ui/input";
import React from "react";
import { useEffect, useRef, useState } from "react";

export default function Home() {
  const [password, setPassword] = useState<string>("");
  const passwordRef = useRef<HTMLInputElement>(null);

  const location = useLocation()
  const navigate = useNavigate()

  useEffect(function() {
    passwordRef.current?.setCustomValidity("")
  }, [password])

  async function handleSubmit(FormEvent: React.FormEvent<HTMLFormElement>) {
    FormEvent.preventDefault()

    const payload = {
      password: password,
    }
    const url = "/login"

    try {
      const {access_token, refresh_token} = await api({ url, method: "POST", data: payload })
      localStorage.setItem("access", access_token)
      localStorage.setItem("refresh", refresh_token)
      return navigate("/")
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
    <div className="min-h-screen flex flex-col justify-center items-center">
      <div className="w-full max-w-lg border-4 rounded-3xl border-secondary bg-secondary/5 p-5">
        <div className="flex flex-col items-center">
          <img src="/medusa.png" width="150" height="150" alt="logo"/>
          <div className="text-3xl mt-5">
            Welcome to Medusa
          </div>
          <div className="text-3xl mt-2">
            your 3D printer assistant 🛠️
          </div>
        </div>
        <form onSubmit={handleSubmit} className="mt-5 flex flex-col">
          <Input
            elRef={passwordRef}
            id="password"
            label="Password"
            type="password"
            value={password}
            onChange={function(event) {setPassword(event.target.value)}}
            placeholder="password"
            required
          />
          <Button className="mt-5" color="primary" size="lg" pill type="submit">Login</Button>
        </form>
      </div>
    </div>
  )
}
