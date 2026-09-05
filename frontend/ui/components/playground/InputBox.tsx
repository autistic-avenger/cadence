"use client"

import { useState } from "react";
import toast from "react-hot-toast";


export default function InputBox(setVideoArr:any) {
    const [link ,setLink] = useState<string>("")
    
    function HandleClick(){
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
