"use client";

import Hero from "@/ui/components/Hero";
import Navbar from "@/ui/components/Navbar";
import axios from "axios";
import { useSearchParams, useRouter } from "next/navigation";
import { Suspense, useEffect, useState } from "react";
import toast, { Toaster } from "react-hot-toast";

function LoginErrorHandler() {
  const searchParams = useSearchParams();
  const router = useRouter();

  useEffect(() => {
    const err = searchParams.get("error");
    if (err === "loginFail") {
      toast.error("Login Failed");
      router.replace("/");
    }
  }, [router, searchParams]);

  return null;
}

function HomeContent() {
  const [isLoggedIn, setLoggedIn] = useState<boolean>(false);

  useEffect(() => {
    const checkJWT = async () => {
      try {
        await axios.get(process.env.NEXT_PUBLIC_API_URL + "/auth/verify", {
          withCredentials: true,
        });
        setLoggedIn(true);
      } catch (err) {
        if (axios.isAxiosError(err)) {
          if (err.response?.status !== axios.HttpStatusCode.BadRequest) {
            toast.error("Bad Token!");
          }
        }
        setLoggedIn(false);
      }
    };
    checkJWT();
  }, []);

  return (
    <div className="h-screen w-full relative overflow-hidden flex flex-col justify-between">
      <Toaster position="top-center" />
      <Suspense fallback={null}>
        <LoginErrorHandler />
      </Suspense>
      <Navbar isLoggedin={isLoggedIn} />
      <div className="relative w-full h-full flex flex-col justify-center">
        <div className="pointer-events-none absolute bottom-0 w-full h-[35%] bg-linear-to-t from-white from-15% to-transparent z-10" />
        <img
          src="/404.jpg"
          alt="hero-bg"
          className="fixed top-0 left-0 -z-10 object-cover h-full w-full select-none [-webkit-user-drag:none]"
        />
        <Hero />
      </div>
    </div>
  );
}

export default function Home() {
  return (
    <Suspense fallback={null}>
      <HomeContent />
    </Suspense>
  );
}
