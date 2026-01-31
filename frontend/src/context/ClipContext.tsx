import { createContext, useContext, useState, useEffect } from "react"
import type { ReactNode } from "react"
import { GetClips, AddClip } from "../../wailsjs/go/main/App"
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
}

const ClipContext = createContext<ClipContextType | undefined>(undefined)

export function ClipProvider({ children }: { children: ReactNode }) {
    const [clips, setClips] = useState<{ pinned: Clip[]; recent: Clip[] }>({ pinned: [], recent: [] })
    const [soundOn, setSoundOn] = useState<boolean>(localStorage.getItem("soundOn") !== "false")
    const [hideContent, setHideContent] = useState<boolean>(localStorage.getItem("hideContent") === "true" || false)

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

    useEffect(() => {
        getClips()
    }, [])

    useEffect(() => {
        EventsOn("clipboard:changed", () => {
            getClips()
        })
    }, [])

    return (
        <ClipContext.Provider value={{ clips, setClips, getClips, addClip, soundOn, setSoundOn, hideContent, setHideContent }}>
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
