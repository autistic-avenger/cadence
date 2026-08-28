import { FaGoogle } from "react-icons/fa";


function Login() {
    return (
    <div className='relative text-4xl sm:text-6xl w-full h-screen md:text-9xl flex justify-center items-center'>
        <img src="/404.jpg" alt="bg" className='object-cover h-screen w-screen' />
        
        <div className='absolute flex justify-center items-center w-full h-120'>
            <div className='h-110 p-2 w-87 bg-white/20 backdrop-blur-sm mx-2 rounded-2xl '>
                <div className='relative w-full flex overflow-hidden h-full bg-white/40 rounded-2xl'>
                    <img src="/login.jpg" alt="login-bg" className='object-cover' />
                    <div className='absolute bottom-2 px-1.5 w-full h-10'>
                        <div className='w-full h-full flex justify-center items-center cursor-pointer text-xl bg-white/60 font-bold text-black rounded-2xl '>
                            <FaGoogle className="mr-2"/>
                            Continue with Google
                        </div>
                    </div>
                </div>
            </div>
            
        </div>
    </div>
  )
}

export default Login
