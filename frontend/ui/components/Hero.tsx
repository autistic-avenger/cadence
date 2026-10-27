"use client";

import DemoShowcase from "@/ui/components/DemoShowcase";

export default function Hero() {
  return (
    <div className="relative w-full flex flex-col items-center">
      <div className="h-20 w-full"></div>

      <div className="w-full pt-4 pb-2 sm:pb-4 flex justify-center z-20">
        <div className="max-w-4xl text-center px-4">
          <h1 className="font-[Caacupe_One] text-4xl sm:text-5xl md:text-6xl lg:text-7xl text-black tracking-tight leading-tight select-none">
            Your best moments, cut down to{" "}
            <span className="font-[Great_Vibes] text-6xl sm:text-7xl md:text-8xl lg:text-9xl text-black block sm:inline">
              seconds.
            </span>
          </h1>
        </div>
      </div>

      <div className="w-full z-20 -mt-6 sm:-mt-8 lg:-mt-10">
        <DemoShowcase />
      </div>
    </div>
  );
}
