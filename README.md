# Rocqman-Go

A Pac-Man game formally verified in [Rocq](https://rocq-prover.org/) and extracted to Go via [Crane](https://github.com/bloomberg/crane).

The game logic is written entirely in Rocq and then mechanically translated to Go — the Go compiler just builds and runs the result.

## Playing

```sh
go build ./...
./rocqman-go
```

Arrow keys or WASD to move. Space to pause. Q or Escape to quit.

## Project Structure

```
theories/         Rocq source (the "real" source of truth)
  Types.v         algebraic types: cell, direction, ghost_mode, phase, position, …
  Game.v          board layout, movement, ghost AI, collision, tick
  Render.v        pixel helpers, sprites, font, SDL drawing
  Loop.v          event handling, frame processing, init_game
  GoSDL2.v        SDL2 axioms and crane extraction templates
  RocqmanGo.v     extraction entry point — crane directives live here

entry.go          handwritten: main(), wires up init_game → run_game
sdl.go            handwritten: real SDL2 implementation of the GoSDL2 axioms
```

The following files are **auto-generated** by `crane` from the Rocq sources. Do not edit them directly — changes will be overwritten on the next extraction.

```
types.go          ← theories/Types.v
game.go           ← theories/Game.v
render.go         ← theories/Render.v
loop.go           ← theories/Loop.v
rocqmango.go      ← theories/RocqmanGo.v  (run_game, exit_game)
gosdl2.go         ← theories/GoSDL2.v     (sdl_key enum, sdl_event type)
datatypes.go      ← Rocq stdlib Datatypes  (list)
listdef.go        ← Rocq stdlib List       (nth)
nat.go            ← Rocq stdlib Nat        (pred, add, mul, …)
peanonat.go       ← Rocq stdlib (stub — erased at extraction)
primstring.go     ← Rocq stdlib (stub — erased at extraction)
itreedefinition.go← ITree library (stub — erased at extraction)
```

## Re-generating the Go files

Requires [Rocq](https://rocq-prover.org/) and [Crane](https://github.com/saviorand/crane/tree/go-backend) (Go backend branch) installed via opam.

```sh
# 1. Build and install the crane plugin
cd /path/to/crane/crane
eval $(opam env --switch=rocq900)
OCAMLFIND_CONF=~/.opam/rocq900/lib/findlib.conf dune build -p rocq-crane
OCAMLFIND_CONF=~/.opam/rocq900/lib/findlib.conf dune install -p rocq-crane

# 2. Run extraction
cd /path/to/rocqman-go
eval $(opam env --switch=rocq900)
touch theories/RocqmanGo.v
OCAMLFIND_CONF=~/.opam/rocq900/lib/findlib.conf dune build theories/RocqmanGo.vo

# 3. Copy generated files to project root
cp _build/RocqmanGoGame/types.go \
   _build/RocqmanGoGame/game.go \
   _build/RocqmanGoGame/render.go \
   _build/RocqmanGoGame/loop.go \
   _build/RocqmanGoGame/rocqmango.go \
   _build/RocqmanGoGame/gosdl2.go \
   _build/RocqmanGoGame/datatypes.go \
   _build/RocqmanGoGame/listdef.go \
   _build/RocqmanGoGame/nat.go \
   _build/RocqmanGoGame/peanonat.go \
   _build/RocqmanGoGame/primstring.go \
   _build/RocqmanGoGame/itreedefinition.go \
   .
goimports -w *.go
go build ./...
```

Do **not** copy `_build/RocqmanGoGame/main.go` — that is a stale artifact from before multi-file extraction was added.

## How it works

Crane's Go backend translates Rocq's ML-like extraction to idiomatic-ish Go:

- Inductive types with only nullary constructors become `iota` enums.
- Records become Go structs.
- `CoFixpoint` game loops become Go `for` loops (via `Set Crane Loopify`).
- Small nat literals are inlined as `uint(N)` via `Crane Extract Inlined Constant`.
- Each Rocq source file extracts to its own `.go` file via `Crane Separate Extraction`.
- SDL2 is bridged through axioms in `GoSDL2.v` with Go template strings, implemented in `sdl.go`.

## Dependencies

- [go-sdl2](https://github.com/veandco/go-sdl2) — SDL2 bindings for Go
- SDL2, SDL2_image, SDL2_mixer — system libraries
