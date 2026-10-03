import { cardInfo } from '@/app/playground/page'
import axios from 'axios'
import { useEffect, useState } from 'react'
import { FaSpinner } from 'react-icons/fa6'
import toast from 'react-hot-toast'

type JobStatus = "pending" | "queued" | "downloading" | "transcribing" | "picking" | "completed" | "failed"


export default function VideoCard({thumbnailUrl,title,url,creator,jobId}:cardInfo) {
    const [jobCurrentStatus ,setJobCurrentStatus] = useState<JobStatus>("pending")

    useEffect(() => {
        getStatus()
        const id = setInterval(() => {
            getStatus()
        }, 5000)

        return () => clearInterval(id)
    }, [])



    async function getStatus(){
        try{
            let status = await axios.get<{"status":JobStatus}>(process.env.NEXT_PUBLIC_API_URL+"/api/jobstatus",{
                params:{"jobid":jobId}
            })
            setJobCurrentStatus(status.data.status)
        }catch(error){
            toast.error("Error getting jobStatus!")
        }
    }

    async function handleCreate(){
        try{
            await axios.post(process.env.NEXT_PUBLIC_API_URL+"/api/create",
                {
                    url:url,
                    jobID:jobId
                }
            ) 
            setJobCurrentStatus("queued")
        }catch(error){
            if (axios.isAxiosError(error)){      
                toast.error(error.response?.data.error)
            }
        }
    }
    return (
    <div className='flex justify-center my-4 items-center w-full h-50 '>
        <div className='h-full w-180 bg-blue-300/40 rounded-2xl backdrop-blur-sm p-1'>
            <div className='w-full h-full bg-blue-300 gap-2 p-2 flex rounded-2xl'>
                <div className='w-70 h-full rounded-2xl overflow-hidden'>
                    <img src={thumbnailUrl} alt="thumbnail" className='object-cover h-full w-full'/>
                </div>

                <div className='h-full flex flex-col flex-1'>
                    <div className='max-h-12 line-clamp-1'>
                        <h1 className=' text-2xl font-[Caacupe_One]'>
                            {title}
                        </h1>
                    </div>
                    
                    <div className='h-8 font-[Caacupe_One] text-xl line-clamp-1'>
                        <span className='text-blue-900'>By </span>- {creator}
                    </div>

                    <div className='h-full flex-1 flex flex-col items-center justify-end w-full'>
                        <div className='w-full h-full flex '>
                            <div className='w-full h-6'>
                                <div className={`flex font-[Caacupe_One] text-2xl justify-center items-center rounded-2xl w-fit px-3 py-1 h-full
                                ${jobCurrentStatus=="completed"?"bg-green-300":""}
                                ${jobCurrentStatus=="pending"?"bg-gray-300":""}
                                ${jobCurrentStatus=="queued"?"bg-amber-300":""}
                                ${jobCurrentStatus=="downloading"?"bg-pink-300":""}
                                ${jobCurrentStatus=="transcribing"?"bg-violet-300":""}
                                ${jobCurrentStatus=="picking"?"bg-cyan-300":""}
                                ${jobCurrentStatus=="failed"?"bg-red-400":""}
                                `}>
                                    {jobCurrentStatus}
                                </div>
                            </div>
                        </div>

                        <div className='h-12 flex justify-end w-full shrink-0'>

                            {jobCurrentStatus !="failed" && 
                                <div className={`h-full cursor-pointer active:bg-green-400 flex justify-center items-center w-25 rounded-xl bg-green-300 active:scale-95 duration-200
                                ${jobCurrentStatus=="pending" || jobCurrentStatus == "completed"?"":"pointer-events-none"}
                                `}
                                onClick={handleCreate}
                                >
                                    {jobCurrentStatus != "pending" ? (
                                        jobCurrentStatus == "completed" ? (
                                        <span className="font-[Caacupe_One] text-2xl  select-none ">
                                            View Clip
                                        </span>
                                        ) : (
                                        <span className="animate-spin ">
                                            <FaSpinner className="text-black w-7 h-7" />
                                        </span>
                                        )
                                    ) : (
                                        <span className="font-[Caacupe_One] text-2xl  select-none ">
                                        Create
                                        </span>
                                    )}
                                </div>
                            }
                        </div>
                    </div>
                </div>

            </div>
        </div>
    </div>
  )
}
