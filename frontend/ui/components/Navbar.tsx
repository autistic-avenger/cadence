import { FaArrowRightLong } from "react-icons/fa6";
import "@/app/globals.css"

export default function Navbar({isLoggedin}:{isLoggedin:boolean}) {
  return (
    <div className="fixed flex justify-center items-center z-50 top-3 w-full h-15 "> 
        <div className="flex pl-5 items-center justify-between h-13 w-255 rounded-2xl bg-white/50 shadow-2xs p-2 backdrop-blur-sm mx-2">
            <h1 className="font-[Caacupe_One] text-3xl text-black/79 text-clip o">
                Cadence
            </h1>
            
            <div className="flex space-x-1">
                
                {!isLoggedin && 
                    <a href="/login" className="h-8 cursor-pointer flex items-center px-3 hover:text-gray-950 duration-200 ease-in-out text-gray-800 font-medium text-sm rounded-md hover:bg-amber-50/70">
                        <h1>
                            Login
                        </h1>
                    </a>
                }


                <a href="/playground" className="group h-8 w-32 px-4 flex gap-2 items-center bg-[#333333] rounded-md text-sm text-white">
                    <h1>
                        Get Started
                    </h1>
                    <FaArrowRightLong className="h-4 w-4 group-hover:translate-x-2 transition-all text-white duration-300"/>
                </a>
            </div>
        </div>
    </div>
  )
}

