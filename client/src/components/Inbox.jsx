import React, { useEffect } from 'react'
import { useCustomHook } from '../customHook'
import { useDispatch, useSelector } from 'react-redux'
import EmailPreview from './EmailPreview'
import { ImSpinner9 } from 'react-icons/im'
import { getEmail } from '../feature/email/emailSlice'
import Pagination from './Pagination'
import SearchBox from './SearchBox'
import TimeFilter from './TimeFilter'

const Inbox = () => {
  const {data, loading, submitting, page, numOfPages} = useSelector(state => state.email)
  const dispatch = useDispatch()

  useEffect(() => {
    
    dispatch(getEmail("inboxemails"))
  }, [])

  return (
    <>
    {
      loading ?
      <div className='h-screen flex justify-center items-center bg-slate-300'>
        <ImSpinner9 className='animate-spin'/>
      </div>
      :
      <div className='w-full'>
        <div className='flex justify-between items-center'>
          <SearchBox />
          <TimeFilter />
        </div>
        {
          data?.map(email => <EmailPreview key={email.id} {...email} />)
        }
        <Pagination numOfPages={numOfPages} page={page} func={getEmail} path={"inboxemails"}/>
      </div>
    }
    </>
  )
}

export default Inbox