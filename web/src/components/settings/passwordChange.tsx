import { api } from "@/api/api";
import Button from "@/components/ui/button"
import Input from "@/components/ui/input"
import { useEffect, useRef, useState } from "react";

export default function PasswordChange() {
  const [isFormVisible, setIsFormVisible] = useState<boolean>(false);

  const [isSuccessMessageVisible, setIsSuccessMessageVisible] = useState<boolean>(false);

  const [oldPassword, setOldPassword] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [password2, setPassword2] = useState<string>("");

  const oldPasswordRef = useRef<HTMLInputElement>(null);
  const passwordRef = useRef<HTMLInputElement>(null);
  const password2Ref = useRef<HTMLInputElement>(null);

  useEffect(function() {
    oldPasswordRef.current?.setCustomValidity("")
  }, [oldPassword])

  useEffect(function() {
    passwordRef.current?.setCustomValidity("")
  }, [password])

  useEffect(function() {
    password2Ref.current?.setCustomValidity("")
  }, [password2])

  async function handleSubmit(formEvent: React.FormEvent<HTMLFormElement>) {
    formEvent.preventDefault()

    const payload = {
      old_password: oldPassword,
      password: password,
      password2: password2,
    }
    const url = "/password/change"

    try {
      await api({ url, method: "POST", data: payload })
      setOldPassword("")
      setPassword("")
      setPassword2("")
      setIsFormVisible(false)
      setIsSuccessMessageVisible(true)
      setTimeout(function() {
        setIsSuccessMessageVisible(false)
      }, 3000)
    } catch (error: any) {
      if (error.error === "old_password_incorrect") {
        oldPasswordRef.current?.setCustomValidity("Password is incorrect")
        oldPasswordRef.current?.reportValidity()
        return
      }
      if (error.error === "password2_does_not_match") {
        password2Ref.current?.setCustomValidity("Passwords don't match")
        password2Ref.current?.reportValidity()
        return
      }
      if (error.error === "invalid_password") {
        passwordRef.current?.setCustomValidity("Must be at least 8 characters long, with upper and lower case and contain at least one number and special character")
        passwordRef.current?.reportValidity()
        return
      }
      throw error
    }
  }
  
  return (
    <div className="max-w-lg">
      <div className="flex items-center gap-5">
        <Button
          className="mt-2"
          color="primary"
          size="lg"
          type="button"
          onClick={function() {
            setIsFormVisible(!isFormVisible)
            setOldPassword("")
            setPassword("")
            setPassword2("")
          }}
        >
          Change Password {isFormVisible ? "▲" : "▼"}
        </Button>
        {isSuccessMessageVisible && <div className="text-primary font-bold">New password successfully set!</div>}
      </div>
      
      {isFormVisible && (
        <form onSubmit={handleSubmit} className="transition ease-in-out delay-150 bg-secondary flex flex-col gap-1 mt-2 p-3 rounded-xl">
          <Input
            elRef={oldPasswordRef}
            id="old_password"
            value={oldPassword}
            onChange={function(event) {setOldPassword(event.target.value)}}
            label="Current Password"
            type="password"
            required
          />
          <Input
            elRef={passwordRef}
            id="password"
            value={password}
            onChange={function(event) {setPassword(event.target.value)}}
            label="New Password"
            type="password"
            required
          />
          <Input
            elRef={password2Ref}
            id="password2"
            value={password2}
            onChange={function(event) {setPassword2(event.target.value)}}
            label="Repeat New Password"
            type="password"
            required
          />
          <Button className="mt-2" type="submit" color="primary" pill>Change Password</Button>
        </form>
      )}
    </div>
    
  )
}
