import React, { useRef } from 'react'
import { useDispatch } from 'react-redux'
import { useLocation } from 'react-router-dom'
import { getEmail } from '../feature/email/emailSlice'

const TimeFilter = () => {
    const formRef = useRef(null)
    const location = useLocation()
    const dispatch = useDispatch()

    const handleTimeSearch = (e) => {
        e.preventDefault()

        const formData = new FormData(formRef.current)
        const {start_time, end_time} = Object.fromEntries(formData)

        // console.log(data)

        const searchPath = location?.pathname.split("/").slice(-1)[0] == "sentbox" ? `sentboxemails?start_time=${start_time}&end_time=${end_time}&page=1` : `inboxemails?start_time=${start_time}&end_time=${end_time}&page=1`

        dispatch(getEmail(searchPath))
    }

  return (
    <form ref={formRef} onSubmit={handleTimeSearch} className='flex gap-2'>
        <div className='flex gap-2'>
            <span>Start</span>
            <input type='datetime-local' name='start_time' defaultValue={"2024-10-13T17:15"} className='border-2 p-1'/>
        </div>
        <div className='flex gap-2'>
            <span>End</span>
            <input type='datetime-local' name='end_time' defaultValue={"2025-10-13T13:37"} className='border-2 p-1'/>
        </div>
        <button className='py-1 px-1.5 bg-slate-500 text-white rounded-md'>Search</button>
    </form>
  )
}

export default TimeFilter