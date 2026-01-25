import { Copy, Pin, Trash2 } from "lucide-react"
import { useState, useRef } from "react"
import type { Clip } from '../../types/clip'
import { TogglePin, Delete } from "../../wailsjs/go/main/App"
import { useClips } from "@/context/ClipContext"
import { playSound } from "@/helpers/playSound"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogDescription } from "@/components/ui/dialog"
import { useRelativeTime } from "@/hooks/use-relative-time"
import { ScrollArea } from "./ui/scroll-area-white"
import { copyBase64ImageToClipboard } from "@/helpers/copyBase64Image"



interface ClipCardProps {
    clip: Clip
    type: "pinned" | "recent"
}

export default function ClipCard({ clip, type }: ClipCardProps) {
    const [copied, setCopied] = useState(false)
    const [dialogOpen, setDialogOpen] = useState(false)
    const cardRef = useRef<HTMLDivElement>(null)
    const { getClips, soundOn, clips, setClips, hideContent } = useClips()

    const handleCopy = async () => {
        playSound("/sounds/paper-copy.wav", soundOn, 1)
        try {
            if (clip.type === "image") {
                copyBase64ImageToClipboard(`data:image/png;base64,${clip.image}`)
                return
            }
            if (clip.content === undefined) return
            await navigator.clipboard.writeText(clip.content)
            setCopied(true)
            setTimeout(() => setCopied(false), 2000)
        } catch (err) {
            console.error("Failed to copy:", err)
        }
    }

    const handlePin = async () => {
        const clipId = Number(clip.id.replace('clip_', ''))
        playSound("/sounds/clipboard-slap.mp3", soundOn, 1)
        await TogglePin(clipId).catch((err) => {
            console.error("Failed to toggle pin:", err)
        }).finally(() => {
            getClips()
        })
    }

    const handleDelete = async () => {
        const clipId = Number(clip.id.replace('clip_', ''))
        playSound("/sounds/paper-rip.mp3", soundOn, .5)
        if (clips.pinned.length <= 1 && clips.recent.length <= 1) {
            setClips({ pinned: [], recent: [] })
        }
        await Delete(clipId).catch((err) => {
            console.error("Failed to delete clip:", err)
        }).then(() => {
            getClips()
        })
    }

    const handleViewClip = () => {
        setDialogOpen(true)
    }

    return (
        <div
            ref={cardRef}
            className={"hand-drawn lined thin p-3 bg-[#F9F5E6] relative group"}
        >   {/* Header with icon and timestamp */}
            {type == "pinned" &&
                <div className="h-10 -top-5 right-[40%] absolute">
                    <img src={"pin.png"} alt="pin-img" className="h-full" />
                    <div className="absolute h-6 w-5 rounded-full shadow-lg/80 top-2 right-3" />
                </div>}
            <div className="mb-3 flex items-start justify-between">
                <span className="text-xl"></span>
                <span className="text-xs text-muted-foreground md:hidden">{useRelativeTime(clip.createdAt)}</span>
            </div>

            {/* Content */}
            <div className={`mb-4 flex-1 overflow-hidden cursor-pointer hover:scale-95 transition-transform  ${hideContent ? "hard-to-read" : ""}`} onClick={handleViewClip}>
                {clip.type === "image" && clip.image ? (
                    <img
                        src={`data:image/png;base64,${clip.image}`}
                        alt="Clip image"
                        className="w-full h-auto object-contain max-h-48 rounded"
                    />
                ) : (
                    <p className={`line-clamp-4 text-sm text-foreground md:line-clamp-8`}>{clip.content}</p>
                )}
            </div>

            {/* Footer with time and actions */}
            <div className="flex items-center justify-between">
                <span className="hidden text-xs text-muted-foreground md:block">{useRelativeTime(clip.createdAt)}</span>
                <div className="flex gap-2 opacity-0 transition-opacity group-hover:opacity-100">
                    <button
                        onClick={handleCopy}
                        className={`rounded p-1.5 transition-colors ${copied ? "bg-green-100 text-green-700" : "bg-foreground/5 text-foreground hover:bg-foreground/10"
                            }`}
                        title="Copy to clipboard"
                    >
                        <Copy className="h-4 w-4" />
                    </button>
                    <button
                        onClick={handlePin}
                        className={`rounded p-1.5 transition-colors ${clip.isPinned
                            ? "bg-yellow-100 text-yellow-700 hover:bg-red-100 hover:text-red-700"
                            : "bg-foreground/5 text-foreground hover:bg-yellow-100 hover:text-yellow-700"
                            }`}
                        title={clip.isPinned ? "Unpin clip" : "Pin clip"}
                    >
                        <Pin className={`h-4 w-4 ${clip.isPinned ? "fill-current" : ""}`} />
                    </button>
                    <button
                        onClick={handleDelete}
                        className="rounded p-1.5 bg-foreground/5 text-foreground transition-colors hover:bg-red-100 hover:text-red-700"
                        title="Delete clip"
                    >
                        <Trash2 className="h-4 w-4" />
                    </button>
                </div>
            </div>

            <Dialog open={dialogOpen} onOpenChange={setDialogOpen}>
                {
                    clip.type === "image" && clip.image ? (
                        <DialogContent className="px-3 border-0 rounded-sm max-w-2xl bg-[url(/board-texture.avif)] bg-cover  h-[90vh]! max-h-125">
                            {/* clip image */}
                            <ScrollArea className=" overflow-auto ">
                                <img
                                    src={`data:image/png;base64,${clip.image}`}
                                    alt="Clip image"
                                    className="w-full h-auto object-contain rounded"
                                />
                            </ScrollArea>
                        </DialogContent>

                    )

                        :

                        (<DialogContent className="px-3 rounded-sm max-w-2xl bg-[url(/board-texture.avif)] bg-cover border-0 h-[90vh]! max-h-125">
                            {/* clip image */}
                            <div className="w-fit absolute h-[20%] top-[-7%] left-0 mx-auto right-0 z-10">
                                <div className="absolute border-black h-2 left-0 right-0 w-[90%] mx-auto bottom-0 shadow-md/65"></div>
                                <img src="/clip.png" className="h-full" alt="" />
                            </div>

                            <div className="page rounded-none! overflow-x-scroll shadow-md/50">
                                <div className="margin"></div>
                                <DialogHeader className="sm:pt-7 pb-0!">
                                    <DialogTitle>Clip Content</DialogTitle>
                                    <DialogDescription>Created {useRelativeTime(clip.createdAt)}</DialogDescription>
                                    <img src="/seperator.png" alt="" className="w-full -mt-6" />
                                </DialogHeader>
                                <div className="overflow-y-auto max-h-[60vh] pr-4 overflow-x-hidden">
                                    {clip.type === "image" && clip.image ? (
                                        <img
                                            src={`data:image/png;base64,${clip.image}`}
                                            alt="Clip image"
                                            className="w-full h-auto object-contain rounded"
                                        />
                                    ) : (
                                        <p className={`whitespace-pre-wrap wrap-break-word text-sm ${hideContent ? "hard-to-read" : ""}`}>{clip.content}</p>
                                    )}
                                </div>
                            </div>
                        </DialogContent>
                        )
                }
            </Dialog>
        </div>
    )
}
