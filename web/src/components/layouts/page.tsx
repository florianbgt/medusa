import Header from '../header'

interface Props {
    children: React.ReactElement<any, any>
}

export default function Layout({ children }: Props) {
    return (
      <div className='min-h-screen flex flex-col'>
        <Header/>
        <div className='grow flex flex-col container self-center mx-3 my-5'>
          {children}
        </div>
      </div>
    )
  }