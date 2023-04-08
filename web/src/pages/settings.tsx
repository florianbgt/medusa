import Main from "@/components/layouts/main"
import PasswordChange from "@/components/settings/passwordChange";
import SystemInfo from "@/components/settings/systemInfo";

export default function Settings() {
  return (
    <Main>
      <div>
        <SystemInfo/>
        <PasswordChange/>
      </div>
    </Main>
  )
}

