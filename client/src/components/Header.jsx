import React from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { useNavigate } from 'react-router-dom'
import { logout } from '../feature/user/userSlice'

const Header = () => {
  const {token, user} = useSelector(state => state.user)
  const navigate = useNavigate()
  const dispatch = useDispatch()
  console.log("header")

  const handleLogout = () => {
    dispatch(logout())
    navigate("/login")
  }

  return (
    <div className='p-4 flex gap-4 justify-end items-center border-b-2 border-slate-500 shadow-md'>
      <h1 className='text-slate-500 text-lg font-semibold'>Welcome <span className='text-green-500 text-md'>{user?.name}</span></h1>
      <button onClick={handleLogout} className='py-1.5 px-4 bg-slate-400 text-white rounded-md'>Logout</button>
    </div>
  )
}

export default Header