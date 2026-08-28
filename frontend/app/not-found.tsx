import React from 'react'

function notFound() {
  return (
    <div className='relative text-4xl sm:text-6xl w-full h-screen md:text-9xl flex justify-center items-center'>
        <img src="/404.jpg" alt="bg" className='object-cover h-screen w-screen' />
        <h1 className='absolute font-[Caacupe_One] text-white'>404 Not Found  : (</h1>
    </div>
  )
}

export default notFound
