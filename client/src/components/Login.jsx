import React, { useRef } from 'react'
import Input from './Input'
import {useDispatch, useSelector} from "react-redux"
import { ImSpinner9 } from "react-icons/im";
import { loginUser } from '../feature/user/userSlice';
import { Link, useNavigate } from 'react-router-dom';

const Login = () => {
  const {submitting} = useSelector(state => state.user)
  const dispatch = useDispatch()
  const formRef = useRef(null)
  const navigate = useNavigate()

  const handleSubmit = (e) => {
    e.preventDefault()

    const formData = new FormData(formRef.current)
    const data = Object.fromEntries(formData)
    
    dispatch(loginUser(data)).then(({type}) => {
      if (type == "user/login/fulfilled"){
        formRef.current.reset()
        navigate("/")
      }
    })
  }

  return (
    <main className='h-screen w-full flex justify-center items-center'>
      <div className='w-1/3 p-2 flex flex-col  bg-slate-300 rounded-md shadow-lg'>
        <h1 className='text-xl my-2'>Login</h1>
        <form ref={formRef} onSubmit={handleSubmit} className='flex flex-col gap-2 my-2'>
          {/* <Input type="text" label="Name" name="name" /> */}
          <Input type="text" label="Email" name="email" />
          <Input type="password" label="Password" name="password" />
          <button type='submit' className='ml-auto bg-slate-500 text-white rounded-md py-1.5 px-4 mt-2 flex gap-2'>
            {
              submitting && 
              <ImSpinner9 className='animate-spin'/>
            } 
              <span>Login</span>
            
            
          </button>
        </form>
        <p className='my-1 w-fit mx-auto'>New here <Link className='text-slate-600' to="/register">Register</Link></p>
      </div>
    </main>
  )
}

export default Login