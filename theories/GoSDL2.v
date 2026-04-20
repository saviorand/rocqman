
(** Go-specific SDL2 bindings for Crane extraction.

    Defines all SDL2 types, event types, and operations needed by the
    Rocqman game, and provides Go extraction mappings for each.

    Also provides:
    - nat -> uint mapping (for efficient integer arithmetic)
    - R  -> float64 mapping (for trig / real arithmetic)
    - sumbool -> bool mapping
    - void type (unit-like for effectful computations)
*)

From Corelib Require Import PrimString.
From Stdlib Require Import Lists.List Bool Reals Rtrigo Ratan.
From Stdlib Require Import PeanoNat.
From Crane Require Extraction.
From Crane.Mapping Require GoStd.
From Crane Require Import Monads.ITree.

(* ------------------------------------------------------------------ *)
(** Void type (unit-like, used as return of effectful SDL2 calls) *)
(* ------------------------------------------------------------------ *)

Axiom void : Type.
Axiom ghost : void.

Crane Extract Inlined Constant void => "struct{}".
Crane Extract Inlined Constant ghost => "struct{}{}".

(* ------------------------------------------------------------------ *)
(** sumbool -> bool (for Rlt_dec, Rle_dec decision procedures) *)
(* ------------------------------------------------------------------ *)

From Stdlib Require Import Sumbool.

Crane Extract Inductive sumbool =>
  "bool"
  [ "true" "false" ]
  "func() any { if %scrut { return %br0 } else { return %br1 } }()".

(* ------------------------------------------------------------------ *)
(** nat -> uint (efficient Peano natural extraction) *)
(* ------------------------------------------------------------------ *)

Crane Extract Inductive nat =>
  "uint"
  [ "0" "(%a0 + 1)" ]
  "func() any { if %scrut == 0 { return %br0 } else { %b1a0 := %scrut - 1; return %br1 } }()".

Crane Extract Numeral nat => "%nu".

Crane Extract Inlined Constant Nat.add    => "(%a0 + %a1)".
Crane Extract Inlined Constant Nat.mul    => "(%a0 * %a1)".
Crane Extract Inlined Constant Nat.sub    =>
  "func() uint { if %a0 > %a1 { return %a0 - %a1 }; return 0 }()".
Crane Extract Inlined Constant Nat.div    =>
  "func() uint { if %a1 == 0 { return 0 }; return %a0 / %a1 }()".
Crane Extract Inlined Constant Nat.modulo =>
  "func() uint { if %a1 == 0 { return %a0 }; return %a0 %% %a1 }()".
Crane Extract Inlined Constant Nat.pred   =>
  "func() uint { if %a0 > 0 { return %a0 - 1 }; return 0 }()".
Crane Extract Inlined Constant Nat.double => "(%a0 + %a0)".
Crane Extract Inlined Constant Nat.eqb    => "(%a0 == %a1)".
Crane Extract Inlined Constant Nat.ltb    => "(%a0 < %a1)".
Crane Extract Inlined Constant Nat.leb    => "(%a0 <= %a1)".
Crane Extract Inlined Constant Nat.max    => "func() uint { if %a0 > %a1 { return %a0 }; return %a1 }()".
Crane Extract Inlined Constant Nat.min    => "func() uint { if %a0 < %a1 { return %a0 }; return %a1 }()".
Crane Extract Inlined Constant Nat.sqrt   => "uint(math.Sqrt(float64(%a0)))".

Crane Extract Inlined Constant PeanoNat.Nat.add    => "(%a0 + %a1)".
Crane Extract Inlined Constant PeanoNat.Nat.mul    => "(%a0 * %a1)".
Crane Extract Inlined Constant PeanoNat.Nat.sub    =>
  "func() uint { if %a0 > %a1 { return %a0 - %a1 }; return 0 }()".
Crane Extract Inlined Constant PeanoNat.Nat.div    =>
  "func() uint { if %a1 == 0 { return 0 }; return %a0 / %a1 }()".
Crane Extract Inlined Constant PeanoNat.Nat.modulo =>
  "func() uint { if %a1 == 0 { return %a0 }; return %a0 %% %a1 }()".
Crane Extract Inlined Constant PeanoNat.Nat.pred   =>
  "func() uint { if %a0 > 0 { return %a0 - 1 }; return 0 }()".
Crane Extract Inlined Constant PeanoNat.Nat.double => "(%a0 + %a0)".
Crane Extract Inlined Constant PeanoNat.Nat.eqb    => "(%a0 == %a1)".
Crane Extract Inlined Constant PeanoNat.Nat.ltb    => "(%a0 < %a1)".
Crane Extract Inlined Constant PeanoNat.Nat.leb    => "(%a0 <= %a1)".
Crane Extract Inlined Constant PeanoNat.Nat.even   => "(%a0 %% 2 == 0)".
Crane Extract Inlined Constant PeanoNat.Nat.odd    => "(%a0 %% 2 == 1)".
Crane Extract Inlined Constant PeanoNat.Nat.div2   => "(%a0 / 2)".
Crane Extract Inlined Constant PeanoNat.Nat.testbit => "((%a0 >> uint(%a1)) & 1 != 0)".

(* ------------------------------------------------------------------ *)
(** R -> float64 (Rocq real numbers mapped to Go float64) *)
(* ------------------------------------------------------------------ *)

Crane Extract Inlined Constant R  => "float64".
Crane Extract Inlined Constant R0 => "0.0".
Crane Extract Inlined Constant R1 => "1.0".

Crane Extract Inlined Constant Rplus  => "(%a0 + %a1)".
Crane Extract Inlined Constant Rminus => "(%a0 - %a1)".
Crane Extract Inlined Constant Rmult  => "(%a0 * %a1)".
Crane Extract Inlined Constant Rdiv   => "(%a0 / %a1)".
Crane Extract Inlined Constant Ropp   => "(-%a0)".
Crane Extract Inlined Constant Rinv   => "(1.0 / %a0)".
Crane Extract Inlined Constant Rsqr   => "(%a0 * %a0)".
Crane Extract Inlined Constant Rabs   => "math.Abs(%a0)".
Crane Extract Inlined Constant Rmax   => "math.Max(%a0, %a1)".
Crane Extract Inlined Constant Rmin   => "math.Min(%a0, %a1)".

Crane Extract Inlined Constant PI  => "math.Pi".
Crane Extract Inlined Constant sin => "math.Sin(%a0)".
Crane Extract Inlined Constant cos => "math.Cos(%a0)".
Crane Extract Inlined Constant tan => "math.Tan(%a0)".
Crane Extract Inlined Constant atan => "math.Atan(%a0)".
Crane Extract Inlined Constant asin => "math.Asin(%a0)".
Crane Extract Inlined Constant acos => "math.Acos(%a0)".
Crane Extract Inlined Constant sqrt => "math.Sqrt(%a0)".
Crane Extract Inlined Constant pow  => "math.Pow(%a0, %a1)".

Crane Extract Inlined Constant Rlt_dec => "(%a0 < %a1)".
Crane Extract Inlined Constant Rle_dec => "(%a0 <= %a1)".
Crane Extract Inlined Constant Req_EM_T => "(%a0 == %a1)".

Crane Extract Inlined Constant INR  => "float64(%a0)".
Crane Extract Inlined Constant IZR  => "float64(%a0)".

(* Short-circuit constructive real infrastructure *)
Crane Extract Inlined Constant Rabst => "%a0".
Crane Extract Inlined Constant Rrepr => "%a0".
Crane Extract Skip Module Rdefinitions.RbaseSymbolsImpl.
Crane Extract Skip Module Rdefinitions.RinvImpl.

(* ------------------------------------------------------------------ *)
(** SDL2 opaque types *)
(* ------------------------------------------------------------------ *)

Axiom sdl_window   : Type.
Axiom sdl_renderer : Type.
Axiom sdl_texture  : Type.

Crane Extract Inlined Constant sdl_window   => "RocqSDLWindow".
Crane Extract Inlined Constant sdl_renderer => "RocqSDLRenderer".
Crane Extract Inlined Constant sdl_texture  => "RocqSDLTexture".

(* ------------------------------------------------------------------ *)
(** SDL2 key codes *)
(* ------------------------------------------------------------------ *)

Inductive sdl_key : Type :=
  | KeyEscape
  | KeyQ
  | KeyUp
  | KeyDown
  | KeyLeft
  | KeyRight
  | KeyW
  | KeyA
  | KeyS
  | KeyD
  | KeySpace.

(* ------------------------------------------------------------------ *)
(** SDL2 event type *)
(* ------------------------------------------------------------------ *)

Inductive sdl_event : Type :=
  | EventQuit    : sdl_event
  | EventKeyDown : sdl_key -> sdl_event.

(* ------------------------------------------------------------------ *)
(** SDL2 effect type (phantom, never extracted as a value) *)
(* ------------------------------------------------------------------ *)

Inductive sdlE : Type -> Type :=.

Crane Extract Skip sdlE.

(* ------------------------------------------------------------------ *)
(** SDL2 effectful operations (axioms, extracted to Go function calls) *)
(* ------------------------------------------------------------------ *)

Axiom sdl_create_window   : PrimString.string -> nat -> nat -> itree sdlE sdl_window.
Axiom sdl_create_renderer : sdl_window -> itree sdlE sdl_renderer.
Axiom sdl_load_texture    : sdl_renderer -> PrimString.string -> itree sdlE sdl_texture.
Axiom sdl_destroy         : sdl_renderer -> sdl_window -> itree sdlE void.

Axiom sdl_get_ticks : itree sdlE nat.
Axiom sdl_delay     : nat -> itree sdlE void.
Axiom sdl_poll_event : itree sdlE sdl_event.

Axiom sdl_set_draw_color : sdl_renderer -> nat -> nat -> nat -> itree sdlE void.
Axiom sdl_fill_rect      : sdl_renderer -> nat -> nat -> nat -> nat -> itree sdlE void.
Axiom sdl_draw_point     : sdl_renderer -> nat -> nat -> itree sdlE void.
Axiom sdl_clear          : sdl_renderer -> itree sdlE void.
Axiom sdl_present        : sdl_renderer -> itree sdlE void.

(** Render a texture at (x,y) sized (w,h) rotated by [angle] degrees.
    [flip_h] flips horizontally. *)
Axiom sdl_render_texture_rotated :
  sdl_renderer -> sdl_texture -> nat -> nat -> nat -> nat -> nat -> bool -> itree sdlE void.

Axiom sdl_play_sound : PrimString.string -> itree sdlE void.

(** Go extraction mappings for SDL2 operations *)

Crane Extract Inlined Constant sdl_create_window =>
  "rocqSDLCreateWindow(%a0, %a1, %a2)".
Crane Extract Inlined Constant sdl_create_renderer =>
  "rocqSDLCreateRenderer(%a0)".
Crane Extract Inlined Constant sdl_load_texture =>
  "rocqSDLLoadTexture(%a0, %a1)".
Crane Extract Inlined Constant sdl_destroy =>
  "rocqSDLDestroy(%a0, %a1)".

Crane Extract Inlined Constant sdl_get_ticks =>
  "rocqSDLGetTicks()".
Crane Extract Inlined Constant sdl_delay =>
  "rocqSDLDelay(%a0)".
Crane Extract Inlined Constant sdl_poll_event =>
  "rocqSDLPollEvent()".

Crane Extract Inlined Constant sdl_set_draw_color =>
  "rocqSDLSetDrawColor(%a0, %a1, %a2, %a3)".
Crane Extract Inlined Constant sdl_fill_rect =>
  "rocqSDLFillRect(%a0, %a1, %a2, %a3, %a4)".
Crane Extract Inlined Constant sdl_draw_point =>
  "rocqSDLDrawPoint(%a0, %a1, %a2)".
Crane Extract Inlined Constant sdl_clear =>
  "rocqSDLClear(%a0)".
Crane Extract Inlined Constant sdl_present =>
  "rocqSDLPresent(%a0)".
Crane Extract Inlined Constant sdl_render_texture_rotated =>
  "rocqSDLRenderTextureRotated(%a0, %a1, %a2, %a3, %a4, %a5, %a6, %a7)".
Crane Extract Inlined Constant sdl_play_sound =>
  "rocqSDLPlaySound(%a0)".
