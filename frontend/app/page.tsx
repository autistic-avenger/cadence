"use client"
import Hero from "@/ui/components/Hero";
import Navbar from "@/ui/components/Navbar";
import axios from "axios";
import { useSearchParams } from "next/navigation";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import toast, { Toaster } from "react-hot-toast";

export default function Home() {
  const searchParams = useSearchParams()
  const router       = useRouter()
  const [isLoggedIn,setLoggedIn] = useState<boolean>(false)
  useEffect(()=>{
    let err = searchParams.get("error")
    if (err == "loginFail"){
      toast.error("Login Failed")
      router.replace("/")
    } 

    const checkJWT = async ()=>{      
      try{
        await axios.get(process.env.NEXT_PUBLIC_API_URL+"/auth/verify",{
          withCredentials :true
        })
        setLoggedIn(true)
      }catch(err){
        if (axios.isAxiosError(err)){
          if (err.response?.status != axios.HttpStatusCode.BadRequest){
            toast.error("Bad Token!")
          }
        }
        setLoggedIn(false)
      }
    }
    checkJWT()
  
  },[])
  return (
    <div>
      <Toaster position="top-center"/>
      <Navbar isLoggedin={isLoggedIn} />
      <div className="relative w-full h-screen ">
        
        <div className="absolute bottom-0 w-full h-[50%] bg-linear-to-t from-white from-15% to-transparent z-2"></div>
        <img src="/404.jpg" alt="hero-bg" className=" absolute -z-9 object-cover h-full w-full select-none [-webkit-user-drag:none]"/>  


        <Hero/>
      </div>
    </div>
  );
}
