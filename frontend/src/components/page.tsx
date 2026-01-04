import { useState, useRef, useEffect } from "react"
import { Search } from "lucide-react"
import ClipCard from "./ui/clip-card"
import { useClips } from "../context/ClipContext"
import gsap from 'gsap';
import { useGSAP } from '@gsap/react';
import { playSound } from "@/helpers/playSound";
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog";
import { GetVersion } from "../../wailsjs/go/main/App";
import { WindowIsMaximised, WindowMinimise, WindowUnmaximise, WindowMaximise, Quit } from "../../wailsjs/runtime";

function PageContent() {
    const [searchQuery, setSearchQuery] = useState("")
    const [version, setVersion] = useState("")
    const [fullScreen, setFullScreen] = useState<boolean>()
    const { clips } = useClips()
    const searchInputRef = useRef<HTMLInputElement>(null)
    const tl = gsap.timeline();


    const handleWindowScreen = async () => {
        const isMax = await WindowIsMaximised();
        if (isMax) {
            WindowUnmaximise();
            setFullScreen(false);
        } else {
            WindowMaximise();
            setFullScreen(true);
        }
    }

    useGSAP(() => {
        tl.to('.paper-curtain-1', { left: "-53vw", duration: 1.5, ease: "steps(12)", onStart: () => playSound('paper-curtain-sound.mp3', true, 1) })
            .to('.paper-curtain-2', { right: '-53vw', duration: 1.5, ease: "steps(9)", }, '-=1.5')
            .from('.pussy', { y: '100%', xPercent: -100, ease: "steps(12)" }, "-=0.5")
    })

    useEffect(() => {
        GetVersion().then(setVersion).catch(err => console.error("Failed to get version:", err))
    }, [])


    const filteredClips = () => {
        const query = searchQuery.toLowerCase()

        if (!query) {
            return {
                pinned: clips.pinned,
                recent: clips.recent,
            }
        }

        return {
            pinned: clips.pinned.filter(
                (clip) =>
                    clip.content.toLowerCase().includes(query),
            ),
            recent: clips.recent.filter(
                (clip) =>
                    clip.content.toLowerCase().includes(query),
            ),
        }
    }

    useEffect(() => {
        const handleKeyDown = (e: KeyboardEvent) => {
            if (e.ctrlKey && e.key === 'f') {
                e.preventDefault()
                searchInputRef.current?.focus()
            }
        }

        window.addEventListener('keydown', handleKeyDown)
        return () => window.removeEventListener('keydown', handleKeyDown)
    }, [])

    return (
        <main className="min-h-screen bg-background p-6 md:p-10">
            {/* Draggable title bar */}
            <div className="fixed z-10 top-0 left-0 right-0 h-10 cursor-grab" style={{ '--wails-draggable': 'drag' } as React.CSSProperties}></div>
            <div className="fixed z-10 top-0 right-0 md:mr-[2%] md:pt-3 pt-2 mr-2 flex items-center gap-1" style={{ '--wails-draggable': 'no-drag' } as React.CSSProperties}>
                <button onClick={() => WindowMinimise()}>
                    <img src="/minimize.png" alt="minimize" className="h-5 shadow-md hover:shadow-lg" />
                </button>
                <button onClick={() => handleWindowScreen()}>
                    <img src={fullScreen ? "/unmaximise.png" : "/maximize.png"} alt="maximize" className="h-5 shadow-md hover:shadow-lg" />
                </button>
                <button onClick={() => Quit()}>
                    <img src="/close.png" alt="close" className="h-5 shadow-md hover:shadow-lg" />
                </button>
            </div>
            <img src="/paper-curtain.png" className="paper-curtain-1 h-screen fixed w-[53vw] left-0 top-0 bottom-0 z-10 " />
            <img src="/paper-curtain.png" className="paper-curtain-2 h-screen fixed w-[53vw] -right-8 top-0 bottom-0 z-10 " />
            {/* // pussy cat image */}
            <div className="h-[20vh] pussy fixed bottom-0 -left-6 z-1">
                <img src="/pussy.png" alt="pussy" className="block h-full" />
            </div>
            <div className="margin"></div>
            <div className="mx-auto max-w-6xl">
                {/* Header */}
                <div className="mb-10 flex items-center gap-8 justify-between">
                    <div className="flex items-center gap-2">
                        <h1 className="font-serif text-4xl font-bold italic text-foreground md:text-5xl">Clipussy</h1>
                        <Dialog>
                            <DialogTrigger asChild>
                                <button className="heartbeat text-2xl hover:opacity-70 transition-opacity cursor-pointer font-bold" title="About">
                                    ⓘ
                                </button>
                            </DialogTrigger>
                            <DialogContent className="!bg-[transparent] shadow-none border-0 pt-9">
                                <div className="absolute h-[calc(100%+2rem)] w-full -z-1">
                                    <img src="/dialog-bg.png" alt="" className=" h-full w-full" />
                                </div>
                                <DialogHeader>
                                    <DialogTitle className="text-2xl font-serif italic">About Clipussy</DialogTitle>
                                    <DialogDescription className="text-base pt-4 space-y-3">
                                        <p>
                                            <strong>Clipussy</strong> is a creative clipboard manager that helps you keep track of your copied content with style.
                                        </p>
                                        <p>
                                            Created with 💜 by <strong>Onyekwelu Jesse</strong> (
                                            <a
                                                href="https://github.com/d3uceY"
                                                target="_blank"
                                                rel="noopener noreferrer"
                                                className="text-blue-600 hover:underline"
                                            >
                                                @d3uceY
                                            </a>
                                            )
                                        </p>
                                        <p className="text-sm text-muted-foreground pt-2">
                                            Built with Wails, React, TypeScript, and Go
                                        </p>
                                        {version && (
                                            <p className="text-xs text-muted-foreground pt-1">
                                                Version: {version}
                                            </p>
                                        )}
                                    </DialogDescription>
                                </DialogHeader>
                            </DialogContent>
                        </Dialog>
                    </div>
                    <div className="relative w-full max-w-md torn-input">
                        <div className="tape-1 absolute -top-3 left-0 h-12 w-4 bg-yellow-200/40 rotate-45 rounded-sm shadow-sm"></div>
                        <div className="tape-2 absolute -top-3 right-0 h-12 w-4 bg-yellow-200/40 -rotate-45 rounded-sm shadow-sm"></div>
                        <input
                            ref={searchInputRef}
                            type="text"
                            placeholder="Search (Ctrl+F)"
                            value={searchQuery}
                            onChange={(e) => setSearchQuery(e.target.value)}
                            className="w-full px-4 pt-2 text-foreground placeholder-gray-500 focus:outline-none"
                        />
                        <Search className="absolute right-3 top-1/2 h-5 w-5 -translate-y-1/2 text-[#c0bdbd]" />
                    </div>
                </div>

                {/* Pinned Section */}
                {filteredClips().pinned.length > 0 && (
                    <section className="mb-12">
                        <h2 className="mb-6 flex items-center gap-2 text-2xl font-bold text-foreground">
                            <span className="text-2xl">📌</span>
                            <span className="italic">Pinned</span>
                        </h2>
                        <div className="free-form-grid-container">
                            {filteredClips().pinned.map((clip) => (
                                <ClipCard key={clip.id} clip={clip} type="pinned" />
                            ))}
                        </div>
                    </section>
                )}

                {/* Recent Section */}
                {filteredClips().recent.length > 0 && (
                    <section>
                        <h2 className="mb-6 flex items-center gap-2 text-2xl font-bold text-foreground">
                            <span className="text-2xl">📝</span>
                            <span className="italic">Recent</span>
                        </h2>
                        <div className="free-form-grid-container">
                            {filteredClips().recent.map((clip) => (
                                <ClipCard key={clip.id} clip={clip} type="recent" />
                            ))}
                        </div>
                    </section>
                )}

                {/* Empty State */}
                {filteredClips().pinned.length === 0 && filteredClips().recent.length === 0 && (
                    <div className="flex h-64 items-center justify-center text-center">
                        <p className="text-lg text-muted-foreground">
                            {searchQuery ? "No clips found matching your search" : "No clips yet. Start copying!"}
                        </p>
                    </div>
                )}
            </div>
        </main>
    )
}

import { ClipProvider } from "../context/ClipContext"

export default function Page() {
    return (
        <ClipProvider>
            <PageContent />
        </ClipProvider>
    )
}
