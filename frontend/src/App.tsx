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
      <main className="page min-h-screen md:pt-16">
        <div className="app-boundary">
          <Page />
        </div>
      </main>
    </>
  )
}

export default App
