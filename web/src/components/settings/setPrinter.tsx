import { api } from "../../api";
import Button from "../ui/button"
import Input from "../ui/input"
import { useEffect, useState } from "react";

export default function SetPrinter() {
  const [isFormVisible, setIsFormVisible] = useState<boolean>(false);

  const [isSuccessMessageVisible, setIsSuccessMessageVisible] = useState<boolean>(false);

  const [x, setX] = useState<number>(0);
  const [y, setY] = useState<number>(0);
  const [z, setZ] = useState<number>(0);

  async function getPrinterInfo() {
    const url = "/printer"
    const response = await api({ url, method: "GET" })
    setX(response.data.x)
    setY(response.data.y)
    setZ(response.data.z)
  }

  async function handleSubmit(formEvent: React.FormEvent<HTMLFormElement>) {
    formEvent.preventDefault()

    const payload = {
      x,
      y,
      z,
    }
    const url = "/printer"

    await api({ url, method: "POST", data: { x, y, z }})
    setIsFormVisible(false)
    setIsSuccessMessageVisible(true)
    setTimeout(function() {
    setIsSuccessMessageVisible(false)
    }, 3000)
  }

    useEffect(function() {
        getPrinterInfo()
    }, [])
  
  return (
    <div className="max-w-lg">
      <div className="flex items-center gap-5">
        <Button
          className="mt-2"
          color="primary"
          size="lg"
          type="button"
          onClick={function() {setIsFormVisible(!isFormVisible)}}
        >
          Set printer info {isFormVisible ? "▲" : "▼"}
        </Button>
        {isSuccessMessageVisible && <div className="text-primary font-bold">New password successfully set!</div>}
      </div>
      
      {isFormVisible && (
        <form onSubmit={handleSubmit} className="border-4 border-secondary bg-secondary/5 flex flex-col gap-1 mt-2 p-3 rounded-xl">
          <Input
            id="x"
            value={x}
            onChange={function(event) {setX(parseFloat(event.target.value))}}
            label="Max X"
            type="number"
            required
          />
          <Input
            id="y"
            value={y}
            onChange={function(event) {setY(parseFloat(event.target.value))}}
            label="Max Y"
            type="number"
            required
          />
          <Input
            id="z"
            value={z}
            onChange={function(event) {setZ(parseFloat(event.target.value))}}
            label="Max Z"
            type="number"
            required
          />
          <Button className="mt-2" type="submit" color="primary" size="md" pill>Set info</Button>
        </form>
      )}
    </div>
    
  )
}
