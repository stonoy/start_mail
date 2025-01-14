import React from 'react'
import { NavLink } from 'react-router-dom'
import { ImSpinner5 } from "react-icons/im";
import { useSelector } from 'react-redux';

const NavLinks = ({name, link, getCount}) => {
  const {emailCounts, countsLoading} = useSelector(state => state.email)

  return (
    <div className='w-full flex justify-between items-center p-2 border-b-2 border-slate-500'>
        <NavLink to={link} className={({isActive, isPending}) => isPending ? " text-2xl text-slate-300" : isActive ? " text-2xl text-green-500" : " text-2xl text-black-500"}>{name}</NavLink>
        {
          countsLoading ?
          <ImSpinner5 className='animate-spin'/>
          :
          <h1>{emailCounts[getCount]}</h1>
        }
    </div>
  )
}

export default NavLinks