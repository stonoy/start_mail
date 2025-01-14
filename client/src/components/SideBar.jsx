import React from 'react'
import { Navlinks } from '../utils'
import NavLinks from './NavLinks'

const SideBar = () => {
  console.log("sidebar")
  return (
    <div className='w-1/5 h-screen flex flex-col gap-2'>
        {
            Navlinks.map((link) => <NavLinks key={link.id} {...link}/>)
        }
    </div>
  )
}

export default SideBar