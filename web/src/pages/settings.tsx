import Page from "../components/layouts/page"
import PasswordChange from "../components/settings/passwordChange";
import SystemInfo from "../components/settings/systemInfo";

export default function Settings() {
  return (
    <Page>
      <div className="flex flex-col gap-2">
        <SystemInfo/>
        <PasswordChange/>
      </div>
    </Page>
  )
}

