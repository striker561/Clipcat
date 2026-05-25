import { useEffect, type RefObject } from "react"

const ROW_HEIGHT = 10 // Must match grid-auto-rows in CSS
const ROW_GAP = 16    // Must match gap in CSS

const rafMap = new WeakMap<HTMLElement, number>()

function updateRowSpan(el: HTMLElement): void {
    const cardHeight = el.getBoundingClientRect().height
    const rowSpan = Math.ceil((cardHeight + ROW_GAP) / (ROW_HEIGHT + ROW_GAP))
    el.style.setProperty('--row-span', String(rowSpan))
}

const sharedObserver = new ResizeObserver((entries) => {
    for (const entry of entries) {
        const el = entry.target as HTMLElement
        const existing = rafMap.get(el)
        if (existing !== undefined) cancelAnimationFrame(existing)
        const id = requestAnimationFrame(() => {
            updateRowSpan(el)
            rafMap.delete(el)
        })
        rafMap.set(el, id)
    }
})

export function useCardRowSpan(ref: RefObject<HTMLElement | null>, isMiniClip: boolean): void {
    useEffect(() => {
        const el = ref.current
        if (!el) return

        sharedObserver.observe(el)

        // Initial calculation after a short delay for image clips to settle
        const timerId = setTimeout(() => updateRowSpan(el), 25)

        return () => {
            sharedObserver.unobserve(el)
            clearTimeout(timerId)
            const rafId = rafMap.get(el)
            if (rafId !== undefined) {
                cancelAnimationFrame(rafId)
                rafMap.delete(el)
            }
        }
    }, [isMiniClip]) // eslint-disable-line react-hooks/exhaustive-deps
}
