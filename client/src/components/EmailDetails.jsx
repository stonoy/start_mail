import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { getEmailDetail } from '../feature/email/emailSlice'
import { useParams } from 'react-router-dom'
import { ImSpinner9 } from 'react-icons/im'
import dayjs from 'dayjs'

const EmailDetails = () => {
    const {loading, theEmail} = useSelector(state => state.email)
    const dispatch = useDispatch()
    const {emailId} = useParams()

    useEffect(() => {
        dispatch(getEmailDetail(emailId))
    }, [])

  return (
    <>
        {
            loading ?
            <div className='h-screen flex justify-center items-center bg-slate-300'>
                    <ImSpinner9 className='animate-spin'/>
                  </div>
            :
            <div className='p-2 flex flex-col'>
        <div className='flex gap-4 text-lg font-semibold text-slate-700'>
            <span className=''>From : </span>
            <h1 >{theEmail?.sender_email}</h1>
        </div>
        <div className='flex gap-4 text-lg text-slate-700'>
            <span>To : </span>
            <h1>{theEmail?.reciver_email}</h1>
        </div>
        <div className='flex gap-4 text-md text-slate-700'>
            <span>Timing : </span>
            <h1>{dayjs(theEmail?.created_at).toString()}</h1>
        </div>
        <div className='flex gap-4 text-md font-semibold text-slate-700'>
            <span>Subject : </span>
            <h1>{theEmail?.subject}</h1>
        </div>
        <div className='text-gray-800'>
            <span>Body : </span>
            <h1>{theEmail?.body}</h1>
        </div>
    </div>
        }
    </>
  )
}

export default EmailDetails