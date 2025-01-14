import React, { useRef } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import Input from './Input'
import {Link, useNavigate} from "react-router-dom"
import { registerUser } from '../feature/user/userSlice'
import { ImSpinner9 } from 'react-icons/im'

const Register = () => {
  const {submitting} = useSelector(state => state.user)
  const dispatch = useDispatch()
  const formRef = useRef(null)
  const navigate = useNavigate()

  const handleSubmit = (e) => {
    e.preventDefault()

    const formData = new FormData(formRef.current)
    const data = Object.fromEntries(formData)
   
    dispatch(registerUser(data)).then(({type}) => {
      if (type == "user/register/fulfilled"){
        formRef.current.reset()
        navigate("/login")
      }
    })
  }

  return (
    <main className='h-screen w-full flex justify-center items-center'>
      <div className='w-1/3 p-2 flex flex-col  bg-slate-300 rounded-md shadow-lg'>
        <h1 className='text-xl my-2'>Register</h1>
        <form ref={formRef} onSubmit={handleSubmit} className='flex flex-col gap-2 my-2'>
          <Input type="text" label="Name" name="name" />
          <Input type="text" label="Email" name="email" />
          <Input type="password" label="Password" name="password" />
          <button type='submit' className='ml-auto bg-slate-500 text-white rounded-md py-1.5 px-4 mt-2 flex gap-2'>
            {
              submitting && 
              <ImSpinner9 className='animate-spin'/>
            } 
              <span>Register</span>
            
            
          </button>
        </form>
        <p className='my-1 w-fit mx-auto'>Already a user <Link className='text-slate-600' to="/login">Login</Link></p>
      </div>
    </main>
  )
}

export default Register