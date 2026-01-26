import { gsap } from "gsap";
import { useRef } from "react";

export default function DeleteButton({ onClick }: { onClick: () => void }) {

    const tl = gsap.timeline();
     const deleteBtnRef = useRef<HTMLButtonElement>(null);
     const deleteImgRef = useRef<HTMLImageElement>(null);
     const deleteImgRedRef = useRef<HTMLImageElement>(null);

    // Animation tween for hover effect 
    const tween = () =>
        tl.set(deleteImgRedRef.current, { display: "block" })
            .set(deleteImgRef.current, { display: "none" })
            .to(deleteImgRedRef.current, {
                scale: 1.2,
                duration: 0.2,
                rotate: 10,
                ease: "power1.out"
            });


    return (
        <button onClick={onClick} onMouseEnter={() => tween().play()} onMouseLeave={() => tween().reverse(0.1)} className="py-3  block w-full" ref={deleteBtnRef}>
            <img src="/delete-base.png" alt="" className="h-8 block mx-auto delete-btn" ref={deleteImgRef} />
            <img src="/delete-red.png" alt="" className="h-8 mx-auto delete-btn-red hidden" ref={deleteImgRedRef} />
        </button>
    );
}