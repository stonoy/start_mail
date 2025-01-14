import React, { useEffect, useState } from 'react'
import { Outlet, useNavigate } from 'react-router-dom'
import SideBar from './SideBar'
import Header from './Header'
import { useDispatch, useSelector } from 'react-redux'
import { customFetch } from '../utils'
import { getEmailCounts } from '../feature/email/emailSlice'
import WrireMail from './WrireMail'

const HomLayOut = () => {
  const [showWriteMail, setShowWriteMail] = useState(false)
  const {token} = useSelector(state => state.user)
  const { submitting} = useSelector(state => state.email)
  const navigate = useNavigate()
  const dispatch = useDispatch()
  console.log("home.")

  useEffect(() => {
    if (!token){
      navigate("/login")
    } else {
      dispatch(getEmailCounts("getmailboxnums"))
    }
  }, [ submitting])

  return (
    <>
      <Header />
      {
        showWriteMail ?
        <WrireMail showWriteMail={showWriteMail} setShowWriteMail={setShowWriteMail}/>
         :
         <main className='flex h-screen gap-2'>
      <SideBar />
      <div className='w-4/5 h-screen relative border-l-2 border-slate-600'>
          <Outlet/>
          <button onClick={() => setShowWriteMail(true)} className='fixed bottom-20 right-5 bg-green-500 text-white font-semibold rounded-md w-28 h-16 px-2 py-4'>Create Email</button>
      </div>
    </main>
      }
    </>
  )
}

export default HomLayOut