import Header from '../header'

interface Props {
    children: React.ReactElement<any, any>
}

export default function Layout({ children }: Props) {
    return (
      <div className='min-h-screen flex flex-col'>
        <Header/>
        {children}
      </div>
    )
  }