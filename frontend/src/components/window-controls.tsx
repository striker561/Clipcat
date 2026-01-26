import { useRef, useState, useEffect } from "react";
import { WindowIsMaximised, WindowMinimise, WindowUnmaximise, WindowMaximise, Quit } from "../../wailsjs/runtime/runtime";
import { useClips } from "@/context/ClipContext";
import gsap from "gsap";
import { useGSAP } from "@gsap/react";
import { playSound } from "@/helpers/playSound";
import { UpdateStorageLimit, GetStorageLimit } from "../../wailsjs/go/main/App";
import { GetClips } from "../../wailsjs/go/main/App";
import { ScrollArea } from "./ui/scroll-area";
import DeleteButton from "./deleteButton";

export default function WindowControls() {
    const [fullScreen, setFullScreen] = useState<boolean>(false);
    const { soundOn, setSoundOn } = useClips();
    const { hideContent, setHideContent, clips } = useClips();
    const settingBtnRef = useRef<HTMLButtonElement>(null);
    const settingDialogRef = useRef<HTMLDivElement>(null);
    const [dialogOpen, setDialogOpen] = useState<boolean>(false);
    const [limit, setLimit] = useState(100)

    // Load initial limit from Wails on component mount
    useEffect(() => {
        const loadLimit = async () => {
            try {
                const currentLimit = await GetStorageLimit();
                setLimit(currentLimit);
            } catch (error) {
                console.error("Failed to load storage limit:", error);
            }
        };
        loadLimit();
    }, []);

    const incrementLimit = async () => {
        // i know this sound is backwards but it sounds better this way 😂
        playSound('/sounds/switch-off.mp3', soundOn, 1);
        const newLimit = Math.min(limit + 50, 500);
        setLimit(newLimit);
        try {
            await UpdateStorageLimit(newLimit);
            await GetClips();
        } catch (error) {
            console.error("Failed to update storage limit:", error);
        }
    };

    const decrementLimit = async () => {
        // yeah, ik it is backwards
        playSound('/sounds/switch-on.mp3', soundOn, 1);
        const newLimit = Math.max(limit - 50, 100);
        setLimit(newLimit);
        try {
            await UpdateStorageLimit(newLimit);
            await GetClips();
        } catch (error) {
            console.error("Failed to update storage limit:", error);
        }
    };

    const MenuSwitch = (isOn: boolean, toggleFunction: () => void, disabled = false): React.JSX.Element => {
        const handleToggleFunction = () => {
            playSound(isOn ? '/sounds/switch-on.mp3' : '/sounds/switch-off.mp3', soundOn, 1);
            if (!disabled) {
                toggleFunction();
            }
        }
        return (
            <button onClick={handleToggleFunction} className="menu-switch-container block h-6 shrink-0 disabled:opacity-50" disabled={disabled}>
                {isOn ? <img src="/on.png" alt="" className='block h-full' /> : <img src="/off.png" alt="" className='block h-full' />}
            </button>
        );
    };



    const ClipStorageLimitSwitch = () => {
        return (
            <div className="flex items-center flex-col">
                <button
                    className="block w-4 -rotate-90 disabled:opacity-50"
                    onClick={incrementLimit}
                    disabled={limit >= 500}
                >
                    <img src="/arrow.png" alt="increment" className="h-full block" />
                </button>
                <div className="flex items-center justify-center w-fit">
                    <span className="text-center w-full text-sm">{limit}</span>
                </div>
                <button
                    className="block w-4 rotate-90 disabled:opacity-50"
                    onClick={decrementLimit}
                    disabled={limit <= 100}
                >
                    <img src="/arrow.png" alt="decrement" className="h-full block" />
                </button>
            </div>
        )
    }

    useGSAP(() => {
        gsap.set(
            settingDialogRef.current, {
            display: "none"
        })
    }, [])

    const tl = gsap.timeline();

    const handleSettingsClick = () => {
        if (!dialogOpen) {
            setDialogOpen(true)
            tl.to(settingBtnRef.current, {
                y: 10,
                rotation: -2,
                duration: .3,
                ease: "power2.out",
                onStart: () => playSound('/sounds/crank.mp3', soundOn, 1)
            }).set(
                settingDialogRef.current, {
                display: "block"
            }).fromTo(
                settingDialogRef.current, {
                opacity: 0,
                y: -15,
                rotation: -3,
                scale: 0.92,
            }, {
                opacity: 1,
                y: 5,
                rotation: 0,
                scale: 1,
                duration: 0.5,
                ease: "back.out(1.2)",
                onStart: () => playSound('/sounds/paper-collect.mp3', soundOn, 1)
            })
        } else {
            setDialogOpen(false)
            tl.to(
                settingDialogRef.current, {
                opacity: 0,
                y: -10,
                rotation: 2,
                scale: 0.95,
                duration: 0.35,
                ease: "power2.in",
                onStart: () => playSound('/sounds/paper-collect.mp3', soundOn, 1)
            }
            )
                .to(settingBtnRef.current, {
                    y: 0,
                    rotation: 0,
                    duration: .3,
                    ease: "elastic.out(1, 0.5)",
                    onStart: () => playSound('/sounds/crank.mp3', soundOn, 1)

                }).set(
                    settingDialogRef.current, {
                    display: "none",
                    rotation: 0,
                    scale: 1,
                })
        }
    }

    const Separator = () => {
        return (
            <div>
                <img src="/seperator.png" alt="" className="w-125" />
            </div>
        );
    };


    const toggleSound = () => {
        setSoundOn((prev) => {
            localStorage.setItem("soundOn", (!prev).toString());
            return !prev;
        });
    }

    const toggleHideContent = () => {
        setHideContent((prev) => {
            localStorage.setItem("hideContent", (!prev).toString());
            return !prev;
        });
    }

    const handleWindowScreen = async () => {
        const isMax = await WindowIsMaximised();
        if (isMax) {
            WindowUnmaximise();
            setFullScreen(false);
        } else {
            WindowMaximise();
            setFullScreen(true);
        }
    };

    const hasClips = () => {
        return clips.recent.length > 0 || clips.pinned.length > 0;
    }

    return (
        <div className="flex flex-row-reverse items-center fixed z-10 top-0 right-0 md:mr-[2%] md:pt-3 pt-2 mr-2 gap-6">
            <div className="mt-1 relative z-10">
                <button onClick={handleSettingsClick} ref={settingBtnRef} className="relative z-10">
                    <img src="/settings.png" alt="close" className="h-5 shadow-md/30" />
                </button>
                <div ref={settingDialogRef} className="setting-dialog absolute min-w-40 aspect-square right-0 top-5">
                    <ScrollArea className="h-full pt-4 px-4 border">
                        <h2 className="text-lg text-center">Settings</h2>
                        <Separator />
                        <div className="flex items-center gap-3 justify-between py-2">
                            <p className="text-base p-0!">Sound</p>
                            {MenuSwitch(soundOn, toggleSound)}
                        </div>
                        <Separator />
                        <div className="flex items-center gap-3 justify-between py-2">
                            <p className="text-base p-0!">Hide Clipboard Content</p>
                            {MenuSwitch(hideContent, toggleHideContent, !hasClips())}
                        </div>
                        <Separator />
                        <div className="flex items-center gap-3 justify-between py-2">
                            <p className="text-base p-0!">Clipboard Limit</p>
                            <ClipStorageLimitSwitch />
                        </div>
                        <Separator />
                        <DeleteButton onClick={() => {/* Add delete functionality here */}} />
                    </ScrollArea>
                    <img src="/menu-clean.png" alt="" className="settings-bg" />
                </div>
            </div>

            <div className=" flex items-center gap-1" style={{ '--wails-draggable': 'no-drag' } as React.CSSProperties}>
                <button onClick={() => WindowMinimise()}>
                    <img src="/minimize.png" alt="minimize" className="h-5 shadow-md/30" />
                </button>
                <button onClick={() => handleWindowScreen()}>
                    <img src={fullScreen ? "/unmaximise.png" : "/maximize.png"} alt="maximize" className="h-5 shadow-md/30" />
                </button>
                <button onClick={() => Quit()}>
                    <img src="/close.png" alt="close" className="h-5 shadow-md/30" />
                </button>
            </div>
        </div>
    );
}
