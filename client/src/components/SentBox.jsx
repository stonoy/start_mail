import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { useCustomHook } from '../customHook'
import { ImSpinner9 } from 'react-icons/im'
import EmailPreview from './EmailPreview'
import { getEmail } from '../feature/email/emailSlice'
import Pagination from './Pagination'
import SearchBox from './SearchBox'
import TimeFilter from './TimeFilter'

const SentBox = () => {
  const {data, loading, submitting, page, numOfPages} = useSelector(state => state.email)
  const dispatch = useDispatch()

  useEffect(() => {
    dispatch(getEmail("sentboxemails"))
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
        <div className='flex flex-col md:flex-row justify-between items-center'>
          <SearchBox />
          <TimeFilter />
        </div>
        {
          data?.map(email => <EmailPreview key={email.id} {...email} />)
        }
        <Pagination numOfPages={numOfPages} page={page} func={getEmail} path={"sentboxemails"}/>
      </div>
    }
    </>
  )
}

export default SentBox