import { createContext, useContext, useState, useEffect } from "react"
import type { ReactNode } from "react"
import {
    GetClips,
    AddClip,
    MakeMiniClip,
    IsMiniClip,
    IsStartupEnabled,
    EnableStartup,
    DisableStartup
} from "../../wailsjs/go/main/App"
import { EventsOn } from "../../wailsjs/runtime"
import type { Clip } from '../../types/clip'

interface ClipContextType {
    clips: { pinned: Clip[]; recent: Clip[] }
    setClips: React.Dispatch<React.SetStateAction<{ pinned: Clip[]; recent: Clip[] }>>
    getClips: () => Promise<void>
    addClip: (content: string, pinned: boolean) => Promise<void>
    soundOn: boolean
    setSoundOn: React.Dispatch<React.SetStateAction<boolean>>
    hideContent: boolean
    setHideContent: React.Dispatch<React.SetStateAction<boolean>>
    toggleMiniClip: () => Promise<void>
    isMiniClip: boolean
    isStartup: boolean
    toggleStartup: () => Promise<void>
}

const ClipContext = createContext<ClipContextType | undefined>(undefined)

export function ClipProvider({ children }: { children: ReactNode }) {
    const [clips, setClips] = useState<{ pinned: Clip[]; recent: Clip[] }>({ pinned: [], recent: [] })
    const [soundOn, setSoundOn] = useState<boolean>(localStorage.getItem("soundOn") !== "false")
    const [hideContent, setHideContent] = useState<boolean>(localStorage.getItem("hideContent") === "true" || false)
    const [isMiniClip, setIsMiniClip] = useState(false);
    const [isStartup, setIsStartup] = useState(false);


    /* ===============================
        STARTUP FUNCTIONS     
       ===============================
    */   
    const checkStartup = async () => {
        await IsStartupEnabled().then((res) => {
            setIsStartup(res);
        })
    }

    const toggleStartup = async () => {
        if (isStartup) {
            await DisableStartup().then(() => {
                setIsStartup(false);
            })
        } else {
            await EnableStartup().then(() => {
                setIsStartup(true);
            })
        }
    }


    /* ===============================
        MINI CLIP FUNCTIONS     
       ===============================
    */  
    const toggleMiniClip = async () => {
        await MakeMiniClip(!isMiniClip).then(() => {
            setIsMiniClip((prev) => (
                !prev
            ))
        })

        await IsMiniClip().then((res) => {
            setIsMiniClip(res);
        })
    }

    /* ===============================
        CLIP OPS FUNCTIONS     
       ===============================
    */  

    const getClips = async () => {
        return GetClips().then((data) => {
            if (data != null) {
                const pinned = data.filter(clip => clip.isPinned)
                const recent = data.filter(clip => !clip.isPinned)
                setClips({ pinned, recent })
            }
            else {
                setClips({ pinned: [], recent: [] })
            }
        })
    }

    const addClip = async (content: string, pinned: boolean) => {
        await AddClip(content, pinned)
        await getClips()
    }

    /* ===============================
        RUN FUNCTIONS ON APP LOAD 
       ===============================
    */
    useEffect(() => {
        checkStartup()
        getClips()
        EventsOn("clipboard:changed", () => {
            getClips()
        })
    }, [])

    return (
        <ClipContext.Provider value={
            {
                // CLIP OPERATIONS
                clips,
                setClips,
                getClips,
                addClip,
                // SOUND OPERATIONS
                soundOn,
                setSoundOn,
                // PRIVACY OPERATIONS
                hideContent,
                setHideContent,
                // MINI CLIP OPERATIONS
                isMiniClip,
                toggleMiniClip,
                // STARTUP OPERATIONS
                isStartup,
                toggleStartup
            }
        }>
            {children}
        </ClipContext.Provider>
    )
}

export function useClips() {
    const context = useContext(ClipContext)
    if (context === undefined) {
        throw new Error("useClips must be used within a ClipProvider")
    }
    return context
}
