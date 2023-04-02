import { useRouter } from "next/router"
import { useEffect, useState } from "react"

interface Props {
  children: React.ReactElement<any, any>
}

export default function Input({ children }: Props) {
  const [authenticated, setAuthenticated] = useState(false);
  
  const router = useRouter()

  useEffect(function () {
    authCheck(router.asPath)
    const hideContent = function() {setAuthenticated(false)}
    router.events.on('routeChangeStart', hideContent)
    router.events.on('routeChangeComplete', authCheck)
    return function() {
      router.events.off('routeChangeStart', hideContent);
      router.events.off('routeChangeComplete', authCheck);
    }
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  function authCheck(url: string) {
    const publicPaths = ['/login'];
    const path = url.split('?')[0];
    const access_token = localStorage.getItem("access")
    if (!access_token && !publicPaths.includes(path)) {
      setAuthenticated(false);
        router.push({
            pathname: '/login',
            query: { returnUrl: router.asPath }
        });
    } else if (access_token && publicPaths.includes(path)) {
        router.push('/');
    } else {
      setAuthenticated(true);
    }
  }

  return authenticated ? children : <></>
}