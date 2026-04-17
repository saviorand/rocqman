
// Go entry point for the Rocqman game.
// The game logic lives in the Rocq-extracted main.go; this file calls it.
package main

// rocqSome wraps a value in a heap-allocated pointer, giving it the same
// representation as the Go extraction of Rocq's [Some] constructor:
// [option T] extracts to [*T], so [Some x] must produce a non-nil [*T].
func rocqSome[T any](v T) *T { return &v }

func main() {
	// init_game is extracted from Rocqman.init_game.
	// It creates the SDL2 window/renderer, loads the texture,
	// and returns the initial window, renderer, and loop state.
	result := init_game

	// prod (prod sdl_window sdl_renderer) loop_state
	// is represented as struct{ fst struct{ fst W; snd R }; snd L }
	win := result.fst.fst
	ren := result.fst.snd
	ls := result.snd

	// run_game runs the game loop.
	// With Set Crane Loopify it compiles to a for loop (no stack overflow).
	run_game(win, ren, ls)
}
