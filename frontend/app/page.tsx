import Navbar from "@/ui/components/Navbar";
import Image from "next/image";

export default function Home() {
  return (
    <div>
      <Navbar/>
      <div className="relative w-full h-screen ">
        
        <div className="absolute bottom-0 w-full h-[50%] bg-linear-to-t from-white from-15% to-transparent z-2"></div>

        <img src="/hero_bg.jpg" alt="hero-bg" className="object-cover h-full w-full select-none [-webkit-user-drag:none]"/>      
      </div>
    </div>
  );
}
