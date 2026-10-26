"use client";

import ClipsBar from "@/ui/components/view/clipsBar";
import axios from "axios";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { FaArrowLeft } from "react-icons/fa6";

interface Clip {
  start:string
  end:string
  title:string
}

function ViewPage() {
  const UrlParams = useParams();
  const router = useRouter();

  const [clips,setClips] = useState<Clip[]|null>(null)
  const [selected, setSelected] = useState<number>(0)

  const handleBackClick = async () => {
    router.back();
  };

  const getClips = async ()=>{
    const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL!}/api/clips?id=${UrlParams?.jobId}`)
    const clips = response.data as Clip[]
    setClips(clips)
    console.log(clips)
  }

  useEffect(()=>{
    getClips()
  },[])

  return (
    <div className="flex">
      <div>
        <img
          src="/404.jpg"
          alt="hero-bg"
          className="fixed top-0 -z-9 object-cover h-full w-full select-none [-webkit-user-drag:none]"
        />
        <div className="h-18 w-full flex items-end justify-start">
          <div
            className="select-none ml-7  w-35 cursor-pointer active:scale-105 duration-400 rounded-4xl h-13 bg-blue-400/30 p-1 backdrop-blur-2xl"
            onClick={handleBackClick}
          >
            <div className="w-full h-full rounded-4xl flex justify-center items-center bg-white/50">
              <FaArrowLeft width={20} height={20} className="mr-2" />
              <h1 className="font-[Caacupe_One] text-2xl">Go Back</h1>
            </div>
          </div>
        </div>

        <div className="aspect-video ml-7 mt-5 w-[57svw] p-1 bg-amber-50/20 rounded-2xl">
          <div className="w-full h-full bg-pink-50/20 backdrop-blur-md overflow-hidden rounded-2xl">
            <video
              key={selected}
              autoPlay
              playsInline
              controls
              preload="auto"
              src={`http://localhost:8000/api/fs/${UrlParams?.jobId}/clips/clip${selected+1}.mp4`}
              className="w-full h-full"
            ></video>
          </div>
        </div>

        <div className="w-[57vw] h-12 mt-2 ml-8 font-[Caacupe_One] text-3xl">
          <h1 className="line-clamp-1 truncate">
            {clips?.[selected]?.title ?? "Loading..." }
          </h1>
        </div>
      </div>

      <div className="flex-1 h-full min-w-0">
        <div className="w-full h-23"></div>
        <div className="h-150 w-[90%] rounded-t-2xl ml-7 bg-blue-500/50 p-1">

          <div className="w-full h-full bg-white/30 rounded-t-2xl">

          <div className="h-10 flex justify-center items-center bg-amber-200/25 rounded-t-2xl w-full border-b">
            <span className="font-[Caacupe_One] text-3xl">{clips?.length==0?"Loading...":`Clips(${clips?.length})`}</span>
          </div>

          <div className="flex-col scrollbar-thin scrollbar-thumb-gray-600 scrollbar-track-transparent h-[93%] overflow-y-scroll">
            {clips?.map((clip, index) => (
              <ClipsBar key={index} id={index} setSelected={setSelected} selected={selected} title={clip.title} />
            )) ?? "Loading..."}
          </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default ViewPage;
