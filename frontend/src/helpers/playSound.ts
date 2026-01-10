import { Howl } from "howler";

let currentSound: Howl | null = null;

// Preloaded sound cache
const soundCache = new Map<string, Howl>();

// List of all sound files to preload
const SOUND_FILES = [
  '/sounds/crank.mp3',
  '/sounds/paper-collect.mp3',
  '/sounds/paper-copy.wav',
  '/sounds/clipboard-slap.mp3',
  '/sounds/paper-rip.mp3',
];

// Preload all sounds into memory
export function preloadSounds() {
  SOUND_FILES.forEach(soundSrc => {
    if (!soundCache.has(soundSrc)) {
      const sound = new Howl({
        src: [soundSrc],
        preload: true,
      });
      soundCache.set(soundSrc, sound);
    }
  });
}

export function playSound(soundSrc: string, soundOn = true, volume = .1) {
  if (!soundOn) return;

  // Stop previous sound
  if (currentSound) {
    currentSound.stop();
  }

  // Use cached sound or create new one
  let sound = soundCache.get(soundSrc);
  
  if (!sound) {
    sound = new Howl({
      src: [soundSrc],
      volume: volume,
    });
    soundCache.set(soundSrc, sound);
  } else {
    sound.volume(volume);
  }

  currentSound = sound;
  sound.play();
}