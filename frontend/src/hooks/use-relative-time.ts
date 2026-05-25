import { useState, useEffect } from 'react'
import { formatTime } from '@/helpers/formatTime'

export const useRelativeTime = (dateString: string) => {
    const [time, setTime] = useState(() => formatTime(dateString))

    useEffect(() => {
        // Initial value is already set correctly by useState above.
        // Only set up the interval for future minute-tick updates.
        const interval = setInterval(() => {
            setTime(formatTime(dateString))
        }, 60000)

        return () => clearInterval(interval)
    }, [dateString])

    return time
}