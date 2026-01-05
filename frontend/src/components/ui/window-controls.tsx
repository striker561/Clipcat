import { useState } from "react";
import { WindowIsMaximised, WindowMinimise, WindowUnmaximise, WindowMaximise, Quit } from "../../../wailsjs/runtime";

export default function WindowControls() {
    const [fullScreen, setFullScreen] = useState<boolean>(false);

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
    );
}
