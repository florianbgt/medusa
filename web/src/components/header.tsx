import Image from "next/image";
import Button from "@/components/ui/button";
import { logout } from "@/api/api";
import Link from "next/link";

export default function Header() {
    return (
      <div className="flex justify-center bg-secondary shadow-lg">
        <div className="container flex justify-between items-center py-2 mr-2">
          <Link href="/" className="flex items-center">
            <Image src="/medusa.png" height="60" width="60" alt="logo" />
            <div className="text-4xl font-bold">
              Medusa
            </div>
          </Link>
          <div className="flex items-center gap-5">
            <Link href="/settings" className="hover:underline">Settings</Link>
            <Button onClick={logout} type="button" color="light" size="sm">Log out</Button>
          </div>
        </div>
      </div>
    )
  }