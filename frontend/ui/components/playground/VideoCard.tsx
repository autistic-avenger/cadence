import React from 'react'
interface VideoCard{
    thumbnailUrl:string
    title:string
} 



export default function VideoCard({thumbnailUrl,title}:VideoCard) {
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
                    <div className='h-full flex justify-center items-end w-full'>
                        <div className='h-12 flex justify-end w-full'>
                            <div className='h-full cursor-pointer active:bg-green-400 flex justify-center items-center w-25 rounded-xl bg-green-300 active:scale-95 duration-200'>
                                <span className='font-[Caacupe_One] text-2xl  select-none '>Create</span>
                            </div>
                        </div>
                    </div>
                </div>

            </div>
        </div>
    </div>
  )
}
