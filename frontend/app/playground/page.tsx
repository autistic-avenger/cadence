"use client"
import InputBox from '@/ui/components/playground/InputBox'
import VideoCard from '@/ui/components/playground/VideoCard'
import { useState } from 'react'
import { Toaster } from 'react-hot-toast'

export interface cardInfo{
  thumbnailUrl:string
  title:string
  videoId:string
  creator:string
}


export default function Playground() {
  const [videoLinks,setVideoLinks] = useState<cardInfo[]>([])

  return (
    <div>
      <Toaster position="top-center" />
      <img
        src="/404.jpg"
        alt="hero-bg"
        className="fixed top-0 -z-9 object-cover h-full w-full select-none [-webkit-user-drag:none]"
      />
      <InputBox setVideoLinks={setVideoLinks} videoLinks={videoLinks} />
      

      {videoLinks.length == 0 && <div className='w-full h-140 text-white/80 flex justify-center items-center'>
          <h1>
            Nothing to Show.
          </h1>
        </div>}


      {videoLinks.map((id:cardInfo)=>{
        return <VideoCard key={id.videoId} thumbnailUrl={id.thumbnailUrl} title={id.title} videoId={id.videoId} creator={id.creator}  />
      })}
    </div>
  );
}
