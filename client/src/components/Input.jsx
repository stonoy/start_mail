import React from 'react'

const Input = ({type, label,name}) => {
  return (
    <div >
        <label>{label}</label>
        <input type={type} name={name} className='w-full p-1 border-2'/>
    </div>
  )
}

export default Input