import { Html, Head, Main, NextScript } from 'next/document'

export default function Document() {
  return (
    <Html lang="en">
      <Head/>
      <body className="flex flex-col justify-stretch items-center bg-dark text-light text-md">
        <div className="w-full min-h-screen max-w-7xl p-5">
          <Main/>
        </div>
        <NextScript />
      </body>
    </Html>
  )
}
