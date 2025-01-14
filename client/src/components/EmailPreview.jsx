import React from 'react'
import dayjs from "dayjs"
import { FaRegHeart } from "react-icons/fa";
import { useDispatch, useSelector } from 'react-redux';
import { createFav, deleteFav } from '../feature/email/emailSlice';
import { FaTrash } from "react-icons/fa";
import { Link, useNavigate } from 'react-router-dom';

const EmailPreview = ({id, sender, receiver, subject, body, created_at, other_side_email, emailId, noFav}) => {
    const {submitting} = useSelector(state => state.email)
    const dispatch = useDispatch()
    const navigate = useNavigate()

    const handleDeleteFav = () => {
        dispatch(deleteFav(id)).then(({type}) => {
            if (type == "email/deleteFav/fulfilled"){
              navigate(0)
            }
          })
    }

  return (
    <Link to={`/email/${noFav ? emailId : id}`} className='flex justify-between items-center border-b-2 border-slate-300 shadow-md mb-2 py-2 px-4'>
        <div className=' '>
        <h1 className='text-lg font-semibold'>{other_side_email}</h1>
        <p className='text-sm text-slate-500'>{dayjs(created_at).toString()}</p>
        <h1 className='text-lg '>{subject}</h1>
        <p className='text-sm text-gray-700'>{body.split("").slice(0,100).join("")}...</p>
    </div>
    {!noFav ? <button onClick={() => dispatch(createFav(id))} className={submitting ? "animate-pulse" : ""}>
    <FaRegHeart  />
    </button>
    :
    <button onClick={handleDeleteFav} className={submitting ? "animate-pulse" : ""}>
    <FaTrash  /> 
    </button>  
}
    </Link>
    
  )
}

export default EmailPreview