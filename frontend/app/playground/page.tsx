"use client"
import InputBox from '@/ui/components/playground/InputBox'
import VideoCard from '@/ui/components/playground/VideoCard'
import axios, { AxiosError } from 'axios'
import { useEffect, useState } from 'react'
import toast, { Toaster } from 'react-hot-toast'

export interface cardInfo{
  thumbnailUrl:string
  title:string
  videoId?:string
  creator:string
  jobId?:string
  email?:string
  url?:string
}

interface GetVideosResponse{
  email:string
  title:string
  thumbnail:string
  url:string
  creator:string
  jobID:string
}



export default function Playground() {
  const [videoLinks,setVideoLinks] = useState<cardInfo[]>([])
  const [loading,setLoading] = useState<boolean>(true)
  
  useEffect(()=>{
      const GetVids = async ()=>{
        const respones = await axios.get(process.env.NEXT_PUBLIC_API_URL+"/api/getVideos",{
          withCredentials:true
        })

        const vids = respones.data as GetVideosResponse[]
        if (vids!=null){
          vids.map((vid:GetVideosResponse)=>{
            let cardInfo:cardInfo = {
              thumbnailUrl :vid.thumbnail,
              title:vid.title,
              creator:vid.creator,
              jobId:vid.jobID,
              email:vid.email,
              url:vid.url
            }
  
            setVideoLinks((prev:cardInfo[])=>{
              return [...prev,cardInfo]
            })
            
          })
        }
      }

      try{
        GetVids()
      }catch(error){
        toast.error("Error fetching Added vids")
      }finally{
        setLoading(false)
      }

      setLoading(false)

  },[])

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
            {loading?"Loading...":"Nothing to Show."}
          </h1>
        </div>}


      {videoLinks.map((id:cardInfo)=>{
        return <VideoCard key={id.videoId} thumbnailUrl={id.thumbnailUrl} title={id.title} videoId={id.videoId} creator={id.creator}  />
      })}
    </div>
  );
}
