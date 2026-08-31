"use client"
import Navbar from "@/ui/components/Navbar";
import { useSearchParams } from "next/navigation";
import { useRouter } from "next/navigation";
import { useEffect } from "react";
import toast, { Toaster } from "react-hot-toast";

export default function Home() {
  const searchParams = useSearchParams()
  const router       = useRouter()
  useEffect(()=>{
    let err = searchParams.get("error")
    if (err == "loginFail"){
      toast.error("Login Failed")
      router.replace("/")
    } 
  },[])
  return (
    <div>
      <Toaster position="top-center"/>
      <Navbar/>
      <div className="relative w-full h-screen ">
        
        <div className="absolute bottom-0 w-full h-[50%] bg-linear-to-t from-white from-15% to-transparent z-2"></div>

        <img src="/hero_bg.jpg" alt="hero-bg" className="object-cover h-full w-full select-none [-webkit-user-drag:none]"/>      
      </div>
    </div>
  );
}
