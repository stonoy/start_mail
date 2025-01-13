import { createBrowserRouter, RouterProvider } from "react-router-dom"
import HomLayOut from "./components/HomLayOut"
import { Favourite, Inbox, Login, Register, SentBox } from "./components"

const router = createBrowserRouter([
  {
    path: "/",
    element: <HomLayOut />,
    children: [
      {
        index: true,
        element: <Inbox/>,
      },
      {
        path: "sentbox",
        element: <SentBox/>,
      },
      {
        path: "favourite",
        element: <Favourite/>,
      }
    ]
  },
  {
    path: "login",
    element: <Login/>,
  },
  {
    path: "register",
    element: <Register/>,
  }
])

const App = () => {

  return <RouterProvider router={router} />
}

export default App