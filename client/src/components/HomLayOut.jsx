import React from 'react'
import { Outlet } from 'react-router-dom'

const HomLayOut = () => {
  console.log("home")
  return (
    <>
      <h1>Home</h1>
      <Outlet />
    </>
  )
}

export default HomLayOut