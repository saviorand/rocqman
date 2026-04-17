(** * Rocqman game – Go extraction entry point.

    This file defines only the top-level game-loop entry points ([run_game],
    [exit_game]) and carries all crane extraction directives.  Game logic,
    rendering, and types live in separate files:

    - [Types.v]  – algebraic and record types
    - [Game.v]   – board layout, game mechanics, ghost AI, tick
    - [Render.v] – pixel helpers, sprites, font, SDL drawing
    - [Loop.v]   – event handling, frame processing, initialisation

    Usage: rocq compile RocqmanGo.v
*)

From Corelib Require Import PrimString.
From Stdlib Require Import Lists.List.
Import ListNotations.
From Stdlib Require Import Reals Rtrigo Ratan.
From RocqmanGoGame Require Import GoSDL2 Types Game Render Loop.
From ITree Require Import Core.ITreeDefinition.
Import Types Game Render Loop.

Local Open Scope pstring_scope.
Local Open Scope itree_scope.

Import ITreeNotations.

(** Cleanup wrapper — returns void. *)
Definition exit_game (win : sdl_window) (ren : sdl_renderer) : itree sdlE void :=
  cleanup ren win.

(** Main game loop (coinductive).
    With [Set Crane Loopify] this becomes a Go [for] loop. *)
CoFixpoint run_game (win : sdl_window) (ren : sdl_renderer)
                    (ls : loop_state) : itree sdlE void :=
  res <- process_frame ren ls ;;
  let '(quit, new_ls) := res in
  if quit then
    exit_game win ren
  else
    Tau (run_game win ren new_ls).

(** * Go extraction configuration *)

(** Additional nat/Nat mappings not covered by GoSDL2 *)
Crane Extract Inlined Constant Nat.testbit =>
  "((%a0 >> uint(%a1)) & 1 != 0)".

(** Map small nat literal constants to Go uint literals to avoid Peano
    S-chains in the output.  These values are inlined at every use site. *)
Crane Extract Inlined Constant board_height  => "uint(15)".
Crane Extract Inlined Constant board_width   => "uint(19)".
Crane Extract Inlined Constant cell_size     => "uint(32)".
Crane Extract Inlined Constant status_height => "uint(40)".
Crane Extract Inlined Constant tick_ms       => "uint(400)".
Crane Extract Inlined Constant frame_ms      => "uint(16)".
Crane Extract Inlined Constant half_cell     => "uint(16)".
Crane Extract Inlined Constant collision_threshold => "uint(22)".

(** Skip auto-generated induction/recursion principles – not used in Go. *)
Crane Extract Skip cell_rect.
Crane Extract Skip cell_rec.
Crane Extract Skip direction_rect.
Crane Extract Skip direction_rec.
Crane Extract Skip ghost_mode_rect.
Crane Extract Skip ghost_mode_rec.
Crane Extract Skip phase_rect.
Crane Extract Skip phase_rec.

(** Enable loopification: converts run_game CoFixpoint into a for loop. *)
Set Crane Loopify.

(** All extracted files belong to package main (single binary). *)
Set Crane Go Package "main".

(** Extract each source file to its own .go file. *)
Crane Separate Extraction run_game exit_game init_game.
