"use client"

import axios from "axios";
import { useState } from "react";
import toast from "react-hot-toast";
import VideoCard from "./VideoCard";

interface InfoResponse {
    title: string
		thumbnail: string
		author: string
		url: string

}

export default function InputBox({setVideoArr}:any) {
    const [link ,setLink] = useState<string>("")
    const YoutubeURL = "https://www.youtube.com/oembed?url=https://www.youtube.com/watch?v=K78QplOC8QQ&format=json"

    async function HandleClick(){
        let linkCopy = link
        if (linkCopy == ""){
            return
        }
        setLink("")
        let regexp:RegExp = /^(?:(?:https?:)?\/\/)?(?:(?:www|m)\.)?(?:youtube\.com|youtu\.be)\/(?:[\w-]+\?v=|embed\/|v\/)?([\w-]+)/;
        
        const vidID = linkCopy.match(regexp)
        if (vidID?.[1] == null){
            toast.error("Invalid Link")
            return
        }

        try{
          const response = await axios.get(process.env.NEXT_PUBLIC_API_URL+"/api/video",{
            params:{
              uid:vidID[1]
            }
          })

          const jsonRes = response.data as InfoResponse
          setVideoArr((a:any)=>{
            return [<VideoCard key={vidID[1]} thumbnailUrl={jsonRes.thumbnail} title={jsonRes.title} />,...a]
          })

        }catch(error){
          toast.error(`Error Adding Video${error}`)
        }
        


    }
    return (
    <div className="w-full h-20 flex items-end justify-center">
      <div className=" flex gap-1 bg-amber-100/20 p-1 backdrop-blur-sm h-13 w-180 rounded-2xl mx-2">
        <input type="text" placeholder={`Enter youtube link "https://youtu.be/s3a4OQR-10M" `} className="bg-blue-200 pl-4 focus:outline-none rounded-l-2xl h-full w-full" value={link} onChange={(e)=>{
            setLink(e.target.value)
        }}>

        </input>
        <button className="h-full bg-green-300 active:bg-green-400 cursor-pointer w-15 rounded-r-2xl" onClick={HandleClick}>Add</button>
      </div>
    </div>
  );
}
