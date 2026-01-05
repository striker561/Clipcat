import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog";

interface AboutDialogProps {
    version: string;
}

export default function AboutDialog({ version }: AboutDialogProps) {
    return (
        <Dialog>
            <DialogTrigger asChild>
                <button className="heartbeat info text-2xl hover:opacity-70 transition-opacity cursor-pointer font-bold" title="About">
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
    );
}
