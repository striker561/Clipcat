import { useRef, useState } from "react";
import { WindowIsMaximised, WindowMinimise, WindowUnmaximise, WindowMaximise, Quit } from "../../wailsjs/runtime/runtime";
import { useClips } from "@/context/ClipContext";
import gsap from "gsap";
import { useGSAP } from "@gsap/react";

export default function WindowControls() {
    const [fullScreen, setFullScreen] = useState<boolean>(false);
    const { soundOn, setSoundOn } = useClips();
    const { hideContent, setHideContent } = useClips();
    const settingBtnRef = useRef<HTMLButtonElement>(null);
    const settingDialogRef = useRef<HTMLDivElement>(null);
    const [dialogOpen, setDialogOpen] = useState<boolean>(false);

    const MenuSwitch = (isOn: boolean, toggleFunction: () => void): React.JSX.Element => {
        return (
            <button onClick={toggleFunction} className="menu-switch-container block h-6 shrink-0">
                {isOn ? <img src="/on.png" alt="" className='block h-full' /> : <img src="/off.png" alt="" className='block h-full' />}
            </button>
        );
    };

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
                y: 5,
                rotation: -2,
                duration: .3,
                ease: "power2.out"
            }).set(
                settingDialogRef.current, {
                display: "block"
            }).fromTo(
                settingDialogRef.current, {
                opacity: 0,
                y: -15,
                rotation: -3,
                scale: 0.92
            }, {
                opacity: 1,
                y: 0,
                rotation: 0,
                scale: 1,
                duration: 0.5,
                ease: "back.out(1.2)"
            }
            )
        } else {
            setDialogOpen(false)
            tl.to(
                settingDialogRef.current, {
                opacity: 0,
                y: -10,
                rotation: 2,
                scale: 0.95,
                duration: 0.35,
                ease: "power2.in"
            }
            )
                .to(settingBtnRef.current, {
                    y: 0,
                    rotation: 0,
                    duration: .3,
                    ease: "elastic.out(1, 0.5)"
                }).set(
                    settingDialogRef.current, {
                    display: "none",
                    rotation: 0,
                    scale: 1
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

    return (
        <div className="flex flex-row-reverse items-center fixed z-10 top-0 right-0 md:mr-[2%] md:pt-3 pt-2 mr-2 gap-6">
            <div className="mt-1 relative z-10">
                <button onClick={handleSettingsClick} ref={settingBtnRef} className="relative z-10">
                    <img src="/settings.png" alt="close" className="h-5 shadow-md/30" />
                </button>
                <div ref={settingDialogRef} className="setting-dialog p-4 absolute h-fit min-w-40 aspect-square right-0 top-5">
                    <h2 className="text-lg text-center">Settings</h2>
                    <Separator />
                    <div className="flex items-center gap-3 justify-between py-2">
                        <p className="text-base p-0!">Sound</p>
                        {MenuSwitch(soundOn, toggleSound)}
                    </div>
                    <Separator />
                    <div className="flex items-center gap-3 justify-between py-2">
                        <p className="text-base p-0!">Hide Clipboard Content</p>
                        {MenuSwitch(hideContent, toggleHideContent)}
                    </div>
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
