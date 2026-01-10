import { useEffect } from "react"
import Page from "./components/page"
import { preloadSounds } from "./helpers/playSound"

function App() {

  useEffect(() => {
    // Preload all sound files when app starts
    preloadSounds();
  }, []);

  return (
    <>
    <main className="page min-h-screen main">
      <Page />
    </main>
    </>
  )
}

export default App
