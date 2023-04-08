import AuthGuard from '@/components/authGuard'
import '@/styles/globals.css'
import type { AppProps } from 'next/app'
import Head from 'next/head'

export default function App({ Component, pageProps }: AppProps) {
  return <>
    <Head>
      <title>Medusa</title>
      <meta name="description" content="Monitor and crontrol your 3D printer remotly with your own portable http server" />
      <meta name="viewport" content="width=device-width, initial-scale=1" />
      <link rel="icon" href="/favicon.ico" />
    </Head>
    <AuthGuard>
      <Component {...pageProps} />
    </AuthGuard>
  </>
}
