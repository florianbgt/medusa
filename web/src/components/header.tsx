import Button from "./ui/button";
import { logout } from "../api";
import { Link } from "react-router-dom";

export default function Header() {
    return (
      <div className="flex justify-center bg-secondary shadow-lg">
        <div className="container flex justify-between items-center py-2 mr-2">
          <Link to="/" className="flex items-center">
            <img src="/medusa.png" height="60" width="60" alt="logo" />
            <div className="text-4xl font-bold">
              Medusa
            </div>
          </Link>
          <div className="flex items-center gap-5">
            <Link to="/settings" className="hover:underline">Settings</Link>
            <Button onClick={logout} type="button" color="light" size="sm">Log out</Button>
          </div>
        </div>
      </div>
    )
  }