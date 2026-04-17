//go:build !rocq

package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

// ---------------------------------------------------------------------------
// Type aliases used by extracted Rocq code
// ---------------------------------------------------------------------------

type RocqSDLWindow = *sdl.Window
type RocqSDLRenderer = *sdl.Renderer
type RocqSDLTexture = *sdl.Texture

// ---------------------------------------------------------------------------
// SDL2 lifecycle
// ---------------------------------------------------------------------------

func rocqSDLCreateWindow(title string, w uint, h uint) RocqSDLWindow {
	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_AUDIO); err != nil {
		fmt.Fprintf(os.Stderr, "SDL Init error: %v\n", err)
		os.Exit(1)
	}
	if err := img.Init(img.INIT_PNG); err != nil {
		fmt.Fprintf(os.Stderr, "SDL_image Init error: %v\n", err)
		os.Exit(1)
	}
	if err := mix.Init(mix.INIT_OGG | mix.INIT_MP3); err != nil {
		fmt.Fprintf(os.Stderr, "SDL_mixer Init error: %v\n", err)
		os.Exit(1)
	}
	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		fmt.Fprintf(os.Stderr, "SDL_mixer OpenAudio error: %v\n", err)
		// non-fatal: continue without audio
	}
	win, err := sdl.CreateWindow(title, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		int32(w), int32(h), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "CreateWindow error: %v\n", err)
		os.Exit(1)
	}
	return win
}

func rocqSDLCreateRenderer(win RocqSDLWindow) RocqSDLRenderer {
	ren, err := sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		fmt.Fprintf(os.Stderr, "CreateRenderer error: %v\n", err)
		os.Exit(1)
	}
	return ren
}

func rocqSDLLoadTexture(ren RocqSDLRenderer, path string) RocqSDLTexture {
	tex, err := img.LoadTexture(ren, path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LoadTexture(%q) error: %v\n", path, err)
		os.Exit(1)
	}
	return tex
}

func rocqSDLDestroy(ren RocqSDLRenderer, win RocqSDLWindow) struct{} {
	ren.Destroy()
	win.Destroy()
	mix.CloseAudio()
	mix.Quit()
	img.Quit()
	sdl.Quit()
	return struct{}{}
}

// ---------------------------------------------------------------------------
// Timing
// ---------------------------------------------------------------------------

func rocqSDLGetTicks() uint {
	return uint(sdl.GetTicks())
}

func rocqSDLDelay(ms uint) struct{} {
	sdl.Delay(uint32(ms))
	return struct{}{}
}

// ---------------------------------------------------------------------------
// Event handling
// ---------------------------------------------------------------------------

// lastDirKey holds the most recent directional key pressed.  When no SDL
// event is pending we replay it so Pac-Man keeps moving in the last
// direction instead of triggering a pause (KeySpace would pause the game).
var lastDirKey sdl_key = KeyRight

func isDirKey(k sdl_key) bool {
	switch k {
	case KeyUp, KeyDown, KeyLeft, KeyRight, KeyW, KeyA, KeyS, KeyD:
		return true
	}
	return false
}

// rocqSDLPollEvent returns one SDL event translated to a Rocq sdl_event.
// When no event is pending it returns the last directional key so the game
// loop keeps ticking without unintended state transitions.
func rocqSDLPollEvent() sdl_event {
	for {
		ev := sdl.PollEvent()
		if ev == nil {
			// No real event: replay last direction so Pac-Man keeps moving.
			return EventKeyDown(lastDirKey)
		}
		switch e := ev.(type) {
		case *sdl.QuitEvent:
			return EventQuit
		case *sdl.KeyboardEvent:
			if e.Type == sdl.KEYDOWN {
				if key, ok := translateKey(e.Keysym.Sym); ok {
					if isDirKey(key) {
						lastDirKey = key
					}
					return EventKeyDown(key)
				}
			}
		}
	}
}

func translateKey(sym sdl.Keycode) (sdl_key, bool) {
	switch sym {
	case sdl.K_ESCAPE:
		return KeyEscape, true
	case sdl.K_q:
		return KeyQ, true
	case sdl.K_UP:
		return KeyUp, true
	case sdl.K_DOWN:
		return KeyDown, true
	case sdl.K_LEFT:
		return KeyLeft, true
	case sdl.K_RIGHT:
		return KeyRight, true
	case sdl.K_w:
		return KeyW, true
	case sdl.K_a:
		return KeyA, true
	case sdl.K_s:
		return KeyS, true
	case sdl.K_d:
		return KeyD, true
	case sdl.K_SPACE:
		return KeySpace, true
	default:
		return 0, false
	}
}

// ---------------------------------------------------------------------------
// Drawing
// ---------------------------------------------------------------------------

func rocqSDLSetDrawColor(ren RocqSDLRenderer, r uint, g uint, b uint) struct{} {
	ren.SetDrawColor(uint8(r), uint8(g), uint8(b), 255)
	return struct{}{}
}

func rocqSDLFillRect(ren RocqSDLRenderer, x uint, y uint, w uint, h uint) struct{} {
	ren.FillRect(&sdl.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)})
	return struct{}{}
}

func rocqSDLDrawPoint(ren RocqSDLRenderer, x uint, y uint) struct{} {
	ren.DrawPoint(int32(x), int32(y))
	return struct{}{}
}

func rocqSDLClear(ren RocqSDLRenderer) struct{} {
	ren.Clear()
	return struct{}{}
}

func rocqSDLPresent(ren RocqSDLRenderer) struct{} {
	ren.Present()
	return struct{}{}
}

func rocqSDLRenderTextureRotated(
	ren RocqSDLRenderer, tex RocqSDLTexture,
	x uint, y uint, w uint, h uint,
	angle uint, flipH bool,
) struct{} {
	dst := &sdl.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)}
	flip := sdl.FLIP_NONE
	if flipH {
		flip = sdl.FLIP_HORIZONTAL
	}
	ren.CopyEx(tex, nil, dst, float64(angle), nil, sdl.RendererFlip(flip))
	return struct{}{}
}

// ---------------------------------------------------------------------------
// Sound
// ---------------------------------------------------------------------------

// soundCache avoids re-loading the same chunk repeatedly.
var soundCache = map[string]*mix.Chunk{}

func rocqSDLPlaySound(path string) struct{} {
	chunk, ok := soundCache[path]
	if !ok {
		var err error
		chunk, err = mix.LoadWAV(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "LoadWAV(%q) error: %v\n", path, err)
			return struct{}{}
		}
		soundCache[path] = chunk
	}
	chunk.Play(-1, 0)
	return struct{}{}
}
