
(** * Game type definitions.

    Defines all algebraic data types and record types used by the Rocqman
    game.  Kept in a separate file so that the type-level view of the game
    is easy to read independently of the logic.
*)

From Stdlib Require Import Lists.List.
Import ListNotations.
From RocqmanGoGame Require Import GoSDL2.

(** ** Enumeration types
    All-nullary inductives extract to [type T int] + [const ( … iota )]
    blocks in Go. *)

Inductive cell : Type := Wall | Empty | Dot | PowerPellet.
Inductive direction : Type := Up | Down | Left | Right | DirNone.
Inductive ghost_mode : Type := Chase | Frightened.
Inductive phase : Type := Playing | Paused | DeathPause | GameOverScreen | WinScreen.

(** ** Record types *)

Record position : Type := mkPos { prow : nat; pcol : nat }.

Record ghost_state : Type := mkGhost {
  gpos : position;
  gdir : direction;
  gmode : ghost_mode
}.

Record game_state : Type := mkState {
  board : list (list cell);
  pacpos : position;
  pacdir : direction;
  desired_dir : direction;
  ghosts : list ghost_state;
  score : nat;
  lives : nat;
  dots_left : nat;
  power_timer : nat;
  game_over : bool;
  game_won : bool
}.
