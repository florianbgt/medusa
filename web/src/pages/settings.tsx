import Page from "@/components/layouts/page"
import PasswordChange from "@/components/settings/passwordChange";
import SystemInfo from "@/components/settings/systemInfo";

export default function Settings() {
  return (
    <Page>
      <div>
        <SystemInfo/>
        <PasswordChange/>
      </div>
    </Page>
  )
}

