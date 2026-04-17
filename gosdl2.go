package main

type sdl_key int

const (
	KeyEscape sdl_key = iota
	KeyQ
	KeyUp
	KeyDown
	KeyLeft
	KeyRight
	KeyW
	KeyA
	KeyS
	KeyD
	KeySpace
)

type sdl_eventImpl struct {
	_v     int
	_c1_d0 any
}

type sdl_event = *sdl_eventImpl

var EventQuit *sdl_eventImpl = &sdl_eventImpl{_v: 0}

func EventKeyDown(a0 any) *sdl_eventImpl {
	return &sdl_eventImpl{_v: 1, _c1_d0: a0}
}
