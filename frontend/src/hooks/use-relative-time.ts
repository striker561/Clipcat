import { useState, useEffect } from 'react'
import { formatTime } from '@/helpers/formatTime'

export const useRelativeTime = (dateString: string) => {
    const [time, setTime] = useState(() => formatTime(dateString))

    useEffect(() => {
        // Update immediately
        setTime(formatTime(dateString))

        // Then update every minute
        const interval = setInterval(() => {
            setTime(formatTime(dateString))
        }, 60000)

        return () => clearInterval(interval)
    }, [dateString])

    return time
}