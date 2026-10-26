import { Dispatch, SetStateAction } from "react"

interface ClipsProp {
    id:number
    setSelected:Dispatch<SetStateAction<number>>
    selected:number
    title:string
}

export default function ClipsBar({id,setSelected,selected,title}:ClipsProp) {
    const handleClipSelect = ()=>{
        setSelected(id)
    }

    return (
    <div className={`h-20 w-full ${selected==id?"bg-blue-500/40":"bg-blue-300/30"} border-b px-4 flex cursor-pointer`}
    onClick={handleClipSelect}
    >
        <div className="w-10 mt-2 shrink-0 h-full text-2xl font-[Caacupe_One] flex justify-center">
            {id+1}.
        </div>

        <div className="min-w-0 flex-1 ml-2 mt-[0.4rem] text-2xl font-[Caacupe_One]">
            <p className="line-clamp-2">
                {title}
            </p> 
        </div>
    </div>
  )
}
