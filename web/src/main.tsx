import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'
import Login from './pages/login';
import Settings from './pages/settings';
import Index from './pages/index';
import Temperature from './pages/index/temperature';
import Control from './pages/index/control';
import GCode from './pages/index/gcode/gcode';
import { Navigate } from "react-router-dom";
import NotFound from './pages/404';
import Viewer from './pages/index/gcode/Viewer';

interface Props {
  children: React.ReactNode
}


function RequireAuth({children}: Props) {
  if (!localStorage.getItem('access')) {
    return <Navigate to="/login" replace />
  }
  return <>{children}</>
}

function RequireNotAuth({children}: Props) {
  if (!!localStorage.getItem('access')) {
    return <Navigate to="/" replace />
  }
  return <>{children}</>
}

const router = createBrowserRouter([
  {
    path: '/',
    element: <RequireAuth>
      <Index/>
    </RequireAuth>,
    children: [
      {
        path: 'temperature',
        element: <Temperature/>,
      },
      {
        path: 'control',
        element: <Control/>,
      },
      {
        path: 'gcode',
        element: <GCode/>,
      }
    ]
  },
  {
    path: '/settings',
    element: <RequireAuth><Settings/></RequireAuth>
  },
  {
    path: '/login',
    element: <RequireNotAuth><Login /></RequireNotAuth>,
  },
  {
    path: '*',
    element: <NotFound />,
  }
])

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <div className='bg-dark text-light min-h-screen'>
      <RouterProvider router={router}/>
    </div>
  </React.StrictMode>,
)
