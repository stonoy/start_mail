import React, { useState } from 'react'
import { debounce } from '../utils'
import { useDispatch } from 'react-redux'
import { useLocation } from 'react-router-dom'
import { getEmail } from '../feature/email/emailSlice'

const SearchBox = () => {
    const [search, setSearch] = useState(localStorage.getItem("searchText") || "")
    const dispatch = useDispatch()
    const location = useLocation()

    // console.log(location?.pathname.split("/").slice(-1)[0])

    const handleSearch = (e) => {
        setSearch(e.target.value)


        const searchPath = location?.pathname.split("/").slice(-1)[0] == "sentbox" ? `sentboxemails?body=${e.target.value}&page=1` : `inboxemails?body=${e.target.value}&page=1`

        debounce(() => {
            localStorage.setItem("searchText", e.target.value)
            dispatch(getEmail(searchPath))
        })
    }

  return (
    <div className='border-2 border-slate-500 m-2 w-fit rounded-md'>
        <input className='p-1' placeholder='search...' type='text' name="search" value={search} onChange={handleSearch}/>
    </div>
  )
}

export default SearchBox