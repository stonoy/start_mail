import React, { useEffect } from 'react'
import { getFav } from '../feature/email/emailSlice'
import { useDispatch, useSelector } from 'react-redux'
import { ImSpinner9 } from 'react-icons/im'
import EmailPreview from './EmailPreview'
import Pagination from './Pagination'

const Favourite = () => {
  const {fav, loading, submitting, page, numOfPages} = useSelector(state => state.email)
  const dispatch = useDispatch()


  useEffect(() => {
    
    dispatch(getFav("getallfavuser"))
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
        {
          fav?.map(fav => <EmailPreview key={fav.id} {...fav} noFav={true}/>)
        }
        <div>
        <Pagination numOfPages={numOfPages} page={page} func={getFav} path={"getallfavuser"}/>
        </div>
      </div>
    }
    </>
  )
}

export default Favourite