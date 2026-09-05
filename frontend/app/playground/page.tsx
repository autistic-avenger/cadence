"use client"
import InputBox from '@/ui/components/playground/InputBox'
import { useState } from 'react'
import { Toaster } from 'react-hot-toast'



export default function Playground() {
  const [videoArr,setVideoArr] = useState<React.ReactNode[]>([])

  return (
    <div>
      <Toaster position="top-center" />
      <img
        src="/404.jpg"
        alt="hero-bg"
        className="fixed top-0 -z-9 object-cover h-full w-full select-none [-webkit-user-drag:none]"
      />
      <InputBox />

      {videoArr.length == 0 && <div className='w-full h-140 text-white/80 flex justify-center items-center'>
          <h1>
            Nothing to Show.
          </h1>
        </div>}
    </div>
  );
}
