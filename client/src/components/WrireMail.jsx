import React, { useRef } from 'react'
import { ImSpinner9 } from 'react-icons/im'
import { useDispatch, useSelector } from 'react-redux'
import { writeEmail } from '../feature/email/emailSlice'

const WrireMail = ({ setShowWriteMail}) => {
    const {submitting} = useSelector(state => state.email)
    const dispatch = useDispatch()
    const formRef = useRef(null)

    const handleEmailSent = (e) => {
        e.preventDefault()

        const formData = new FormData(formRef.current)
        const data = Object.fromEntries(formData)

        dispatch(writeEmail(data)).then(({type}) => {
            if (type == "email/writeEmail/fulfilled"){
              formRef.current.reset()
              localStorage.removeItem("writeEmail")
              setShowWriteMail(false)
            }
          })
    }

    const handleCancel = () => {
        const formData = new FormData(formRef.current)
        const data = Object.fromEntries(formData)

        localStorage.setItem("writeEmail", JSON.stringify(data))
        setShowWriteMail(false)
    }

  return (
    <section className='w-full h-screen flex bg-slate-300 justify-center items-center'>
        <div className='w-4/5 md:w-2/3 p-2 bg-white'>
        <form ref={formRef} onSubmit={handleEmailSent} className='flex flex-col gap-2'>
        <div>
            <p className='text-slate-900 my-2'>Recipient Email</p>
            <input type='text' name="recipient" defaultValue={JSON.parse(localStorage.getItem("writeEmail"))?.recipient || ""} className='border-2 w-full p-1' />
        </div>
        <div>
            <p className='text-slate-900 my-2'>Subject</p>
            <input type='text' name="subject" defaultValue={JSON.parse(localStorage.getItem("writeEmail"))?.subject || ""} className='border-2 w-full p-1' />
        </div>
        <div>
            <p className='text-slate-900 my-2'>Message</p>
            <textarea rows={4} name="body" defaultValue={JSON.parse(localStorage.getItem("writeEmail"))?.body || ""} className='resize-none w-full border-2 p-1' />
        </div>
        <div className='flex my-2 justify-end items-center gap-2 text-white'>
            <button type='button' onClick={handleCancel} className='py-1.5 px-4 bg-slate-500 rounded-md'>Cancel</button>
            <button type='submit' className='py-1.5 px-4 bg-slate-500 rounded-md'>
                {
                              submitting && 
                              <ImSpinner9 className='animate-spin'/>
                            } 
                              <span>Sent</span>
            </button>
        </div>
        </form>
        </div>
    </section>
  )
}

export default WrireMail