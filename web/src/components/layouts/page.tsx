import Header from '../header'

interface Props {
    children: React.ReactElement<any, any>
}

export default function Layout({ children }: Props) {
    return (
      <>
        <Header/>
        <div className="container mx-auto my-5 px-2">
          {children}
        </div>
      </>
    )
  }