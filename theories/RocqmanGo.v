(** * Rocqman game logic and rendering – Go extraction variant.

    Identical to [Rocqman.v] except:
    - Imports [GoSDL2] instead of the C++ CraneSDL2 / Mapping.Std / Mapping.Real.
    - [run_game] and [exit_game] return [void] (no C++ exit-code idiom).
    - No top-level [main]; the Go entry point is hand-written in [entry.go].
    - Extracts to [main.go] (package main) via Crane Go backend.

    Usage: rocq compile RocqmanGo.v
*)

From Corelib Require Import PrimString.
From Stdlib Require Import Lists.List.
Import ListNotations.
From Stdlib Require Import Bool.
From Stdlib Require Import Reals Rtrigo Ratan.
From RocqmanGoGame Require Import GoSDL2.
From ITree Require Import Core.ITreeDefinition.

Local Open Scope pstring_scope.
Local Open Scope itree_scope.

Module Rocqman.

Import ITreeNotations.

(** Two-argument arctangent with the conventional range [-pi, pi]. *)
Definition real_atan2 (y x : R) : R :=
  if Rlt_dec R0 x then atan (Rdiv y x)
  else if Rlt_dec x R0 then
         if Rle_dec R0 y then Rplus (atan (Rdiv y x)) PI
         else Rminus (atan (Rdiv y x)) PI
       else if Rlt_dec R0 y then Rdiv PI (INR 2)
       else if Rlt_dec y R0 then Ropp (Rdiv PI (INR 2))
       else R0.

(** Test bit [n] of a natural number. *)
Definition nat_testbit : nat -> nat -> bool := Nat.testbit.

(** * Types *)

Inductive cell : Type := Wall | Empty | Dot | PowerPellet.
Inductive direction : Type := Up | Down | Left | Right | DirNone.
Inductive ghost_mode : Type := Chase | Frightened.
Inductive phase : Type := Playing | Paused | DeathPause | GameOverScreen | WinScreen.

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

(** * Constants *)

Definition board_height : nat := 15.
Definition board_width : nat := 19.
Definition cell_size : nat := 32.
Definition status_height : nat := 40.
Definition win_width : nat := board_width * cell_size.
Definition win_height : nat := board_height * cell_size + status_height.
Definition tick_ms : nat := 400.
Definition frame_ms : nat := 16.

(** * Board layout *)

Local Definition W := Wall.
Local Definition E := Empty.
Local Definition D := Dot.
Local Definition P := PowerPellet.

Definition initial_board : list (list cell) :=
  [ [W;W;W;W;W;W;W;W;W;W;W;W;W;W;W;W;W;W;W]
  ; [W;D;D;D;D;D;D;D;D;W;D;D;D;D;D;D;D;D;W]
  ; [W;D;W;W;D;W;W;W;D;W;D;W;W;W;D;W;W;D;W]
  ; [W;P;W;D;D;D;W;D;D;D;D;D;W;D;D;D;W;P;W]
  ; [W;D;W;W;W;D;W;D;W;W;W;D;W;D;W;W;W;D;W]
  ; [W;D;D;D;D;D;D;D;D;D;D;D;D;D;D;D;D;D;W]
  ; [W;W;D;W;D;W;W;D;W;W;W;D;W;W;D;W;D;W;W]
  ; [W;D;D;D;D;W;D;D;D;W;D;D;D;W;D;D;D;D;W]
  ; [W;D;W;W;D;W;W;W;D;W;D;W;W;W;D;W;W;D;W]
  ; [W;D;D;D;D;D;D;D;D;D;D;D;D;D;D;D;D;D;W]
  ; [W;D;W;W;W;D;W;D;W;W;W;D;W;D;W;W;W;D;W]
  ; [W;P;W;D;D;D;W;D;D;D;D;D;W;D;D;D;W;P;W]
  ; [W;D;W;W;D;W;W;W;D;W;D;W;W;W;D;W;W;D;W]
  ; [W;D;D;D;D;D;D;D;D;W;D;D;D;D;D;D;D;D;W]
  ; [W;W;W;W;W;W;W;W;W;W;W;W;W;W;W;W;W;W;W]
  ].

(** * Helper functions *)

Fixpoint count_row (r : list cell) : nat :=
  match r with
  | [] => 0
  | Dot :: rest => 1 + count_row rest
  | PowerPellet :: rest => 1 + count_row rest
  | _ :: rest => count_row rest
  end.

Fixpoint count_dots (b : list (list cell)) : nat :=
  match b with
  | [] => 0
  | row :: rest => count_row row + count_dots rest
  end.

Definition get_cell (row col : nat) (b : list (list cell)) : cell :=
  nth col (nth row b []) Wall.

Fixpoint replace_nth {A : Type} (n : nat) (l : list A) (x : A) : list A :=
  match l with
  | [] => []
  | h :: t => match n with
              | 0 => x :: t
              | S n' => h :: replace_nth n' t x
              end
  end.

Definition set_cell (row col : nat) (c : cell) (b : list (list cell))
  : list (list cell) :=
  let r := nth row b [] in
  replace_nth row b (replace_nth col r c).

Definition is_wall (row col : nat) (b : list (list cell)) : bool :=
  match get_cell row col b with
  | Wall => true
  | _ => false
  end.

Definition move_pos (d : direction) (p : position) : position :=
  match d with
  | Up    => mkPos (Nat.pred (prow p)) (pcol p)
  | Down  => mkPos (S (prow p)) (pcol p)
  | Left  => mkPos (prow p) (Nat.pred (pcol p))
  | Right => mkPos (prow p) (S (pcol p))
  | DirNone => p
  end.

Definition can_move (d : direction) (p : position) (b : list (list cell))
  : bool :=
  let new_p := move_pos d p in
  negb (is_wall (prow new_p) (pcol new_p) b).

(** * Game logic *)

Fixpoint update_ghost_modes (gs : list ghost_state) (powered : bool) : list ghost_state :=
  match gs with
  | [] => []
  | g :: rest =>
    let mode := if powered then Frightened else Chase in
    mkGhost (gpos g) (gdir g) mode :: update_ghost_modes rest powered
  end.

Definition set_direction (d : direction) (gs : game_state) : game_state :=
  mkState (board gs) (pacpos gs) (pacdir gs) d (ghosts gs) (score gs) (lives gs)
          (dots_left gs) (power_timer gs) (game_over gs) (game_won gs).

Definition move_pacman (gs : game_state) : game_state :=
  if game_over gs || game_won gs then gs
  else
    let try_dir := if can_move (desired_dir gs) (pacpos gs) (board gs)
                   then desired_dir gs
                   else pacdir gs in
    let new_pos := move_pos try_dir (pacpos gs) in
    if is_wall (prow new_pos) (pcol new_pos) (board gs) then gs
    else
      let cell := get_cell (prow new_pos) (pcol new_pos) (board gs) in
      let new_board := match cell with
                       | Dot | PowerPellet =>
                         set_cell (prow new_pos) (pcol new_pos) Empty (board gs)
                       | _ => board gs
                       end in
      let add_score := match cell with
                       | Dot => 10
                       | PowerPellet => 50
                       | _ => 0
                       end in
      let new_dots := match cell with
                      | Dot | PowerPellet => Nat.pred (dots_left gs)
                      | _ => dots_left gs
                      end in
      let new_power := match cell with
                       | PowerPellet => 20
                       | _ => power_timer gs
                       end in
      let new_ghosts := match cell with
                        | PowerPellet =>
                          update_ghost_modes (ghosts gs) true
                        | _ => ghosts gs
                        end in
      let won := Nat.eqb new_dots 0 in
      mkState new_board new_pos try_dir (desired_dir gs) new_ghosts
              (score gs + add_score) (lives gs)
              new_dots new_power won won.

(** * Ghost AI *)

Definition abs_diff (a b : nat) : nat :=
  if Nat.leb a b then b - a else a - b.

Definition manhattan (p1 p2 : position) : nat :=
  abs_diff (prow p1) (prow p2) + abs_diff (pcol p1) (pcol p2).

Definition is_opposite (d1 d2 : direction) : bool :=
  match d1, d2 with
  | Up, Down | Down, Up | Left, Right | Right, Left => true
  | _, _ => false
  end.

Definition try_dir (d : direction) (g : ghost_state) (target : position)
                   (b : list (list cell))
                   (best_d : direction) (best_dist : nat)
  : direction * nat :=
  if negb (can_move d (gpos g) b) then (best_d, best_dist)
  else if is_opposite d (gdir g) then (best_d, best_dist)
  else
    let new_pos := move_pos d (gpos g) in
    let dist := manhattan new_pos target in
    if Nat.ltb dist best_dist then (d, dist) else (best_d, best_dist).

Definition choose_ghost_dir (g : ghost_state) (target : position)
                            (b : list (list cell)) : direction :=
  let '(d1, dist1) := try_dir Up    g target b DirNone 999 in
  let '(d2, dist2) := try_dir Down  g target b d1 dist1 in
  let '(d3, dist3) := try_dir Left  g target b d2 dist2 in
  let '(d4, _)     := try_dir Right g target b d3 dist3 in
  d4.

Definition move_one_ghost (g : ghost_state) (pac : position)
                          (b : list (list cell)) : ghost_state :=
  let target := match gmode g with
                | Chase => pac
                | Frightened => mkPos 0 0
                end in
  let dir := choose_ghost_dir g target b in
  if can_move dir (gpos g) b
  then mkGhost (move_pos dir (gpos g)) dir (gmode g)
  else g.

Fixpoint move_ghosts_list (gs : list ghost_state) (pac : position)
                          (b : list (list cell)) : list ghost_state :=
  match gs with
  | [] => []
  | g :: rest => move_one_ghost g pac b :: move_ghosts_list rest pac b
  end.

Definition move_ghosts (gs : game_state) : game_state :=
  mkState (board gs) (pacpos gs) (pacdir gs) (desired_dir gs)
          (move_ghosts_list (ghosts gs) (pacpos gs) (board gs))
          (score gs) (lives gs) (dots_left gs) (power_timer gs)
          (game_over gs) (game_won gs).

(** * Collision detection *)

Fixpoint ghost_at_pos (row col : nat) (gs : list ghost_state) : option ghost_state :=
  match gs with
  | [] => None
  | g :: rest =>
    if Nat.eqb (prow (gpos g)) row && Nat.eqb (pcol (gpos g)) col
    then Some g
    else ghost_at_pos row col rest
  end.

Definition respawn_ghost (g : ghost_state) : ghost_state :=
  mkGhost (mkPos 7 9) DirNone Chase.

Definition check_collisions (gs : game_state) : game_state :=
  match ghost_at_pos (prow (pacpos gs)) (pcol (pacpos gs)) (ghosts gs) with
  | None => gs
  | Some g =>
    match gmode g with
    | Frightened =>
      let new_ghosts := map (fun g' =>
        if Nat.eqb (prow (gpos g')) (prow (pacpos gs)) &&
           Nat.eqb (pcol (gpos g')) (pcol (pacpos gs))
        then respawn_ghost g'
        else g') (ghosts gs) in
      mkState (board gs) (pacpos gs) (pacdir gs) (desired_dir gs) new_ghosts
              (score gs + 200) (lives gs) (dots_left gs) (power_timer gs)
              (game_over gs) (game_won gs)
    | Chase =>
      let new_lives := Nat.pred (lives gs) in
      let dead := Nat.eqb new_lives 0 in
      mkState (board gs) (mkPos 9 9) DirNone DirNone (ghosts gs)
              (score gs) new_lives (dots_left gs) 0
              dead (game_won gs)
    end
  end.

Definition tick_power (gs : game_state) : game_state :=
  match power_timer gs with
  | 0 => gs
  | S n =>
    let powered := negb (Nat.eqb n 0) in
    let new_ghosts := update_ghost_modes (ghosts gs) powered in
    mkState (board gs) (pacpos gs) (pacdir gs) (desired_dir gs) new_ghosts
            (score gs) (lives gs) (dots_left gs) n
            (game_over gs) (game_won gs)
  end.

(** * Main tick *)

Definition tick (gs : game_state) : game_state :=
  if game_over gs || game_won gs then gs
  else
    let gs1 := move_pacman gs in
    let gs2 := move_ghosts gs1 in
    tick_power gs2.

(** * Initial state *)

Definition initial_ghosts : list ghost_state :=
  [ mkGhost (mkPos 1 1) DirNone Chase
  ; mkGhost (mkPos 1 17) DirNone Chase
  ; mkGhost (mkPos 13 1) DirNone Chase
  ; mkGhost (mkPos 13 17) DirNone Chase
  ].

Definition initial_state : game_state :=
  let b := set_cell 9 9 Empty initial_board in
  mkState b
          (mkPos 9 9)
          DirNone
          DirNone
          initial_ghosts
          0
          3
          (count_dots b)
          0
          false
          false.

(** * Drawing primitives *)

Fixpoint filled_circle_rows (ren : sdl_renderer) (cx base_y : nat)
                            (radius i count : nat) : itree sdlE void :=
  match count with
  | 0 => Ret ghost
  | S count' =>
    let dist := abs_diff i radius in
    let d2 := dist * dist in
    let r2 := radius * radius in
    (if Nat.leb d2 r2 then
       let half := Nat.sqrt (r2 - d2) in
       sdl_fill_rect ren (cx - half) (base_y + i) (half + half + 1) 1
     else Ret ghost) ;;
    filled_circle_rows ren cx base_y radius (S i) count'
  end.

Definition draw_filled_circle (ren : sdl_renderer) (cx cy radius : nat)
  : itree sdlE void :=
  filled_circle_rows ren cx (cy - radius) radius 0 (radius + radius + 1).

Fixpoint semicircle_top_rows (ren : sdl_renderer) (cx base_y : nat)
                             (radius i count : nat) : itree sdlE void :=
  match count with
  | 0 => Ret ghost
  | S count' =>
    let dist := radius - i in
    let half := Nat.sqrt (radius * radius - dist * dist) in
    sdl_fill_rect ren (cx - half) (base_y + i) (half + half + 1) 1 ;;
    semicircle_top_rows ren cx base_y radius (S i) count'
  end.

Definition draw_top_semicircle (ren : sdl_renderer) (cx cy radius : nat)
  : itree sdlE void :=
  semicircle_top_rows ren cx (cy - radius) radius 0 (radius + 1).

(** * Ghost sprite rendering *)

Definition ghost_body_color (color_idx : nat)
  : (nat * nat * nat) :=
  match color_idx with
  | 0 => (255, 0, 0)
  | 1 => (255, 184, 222)
  | 2 => (0, 255, 255)
  | 3 => (255, 184, 82)
  | _ => (33, 33, 222)
  end.

Definition draw_ghost_body (ren : sdl_renderer) (cx cy : nat)
                           (radius : nat) (cr cg cb : nat) : itree sdlE void :=
  sdl_set_draw_color ren cr cg cb ;;
  draw_top_semicircle ren cx cy radius ;;
  sdl_fill_rect ren (cx - radius) cy (radius + radius + 1) radius ;;
  let seg := Nat.div (radius + radius) 3 in
  let sr := Nat.div seg 2 in
  let bx0 := cx - radius + Nat.div seg 2 in
  draw_filled_circle ren bx0 (cy + radius) sr ;;
  draw_filled_circle ren (bx0 + seg) (cy + radius) sr ;;
  draw_filled_circle ren (bx0 + seg + seg) (cy + radius) sr.

Definition draw_ghost_eyes (ren : sdl_renderer) (cx cy : nat)
                           (radius : nat) : itree sdlE void :=
  let eye_offset := Nat.div radius 3 in
  sdl_set_draw_color ren 255 255 255 ;;
  draw_filled_circle ren (cx - eye_offset) (cy - 2) 4 ;;
  draw_filled_circle ren (cx + eye_offset) (cy - 2) 4 ;;
  sdl_set_draw_color ren 33 33 222 ;;
  draw_filled_circle ren (cx - eye_offset + 1) (cy - 1) 2 ;;
  draw_filled_circle ren (cx + eye_offset + 1) (cy - 1) 2.

Definition draw_ghost_bottom (ren : sdl_renderer)
                             (cx cy : nat) : itree sdlE void :=
  sdl_set_draw_color ren 255 255 255 ;;
  sdl_fill_rect ren (cx - 1) (cy - 7) 3 11 ;;
  sdl_fill_rect ren (cx - 6) (cy + 2) 13 3.

Definition draw_ghost_sprite (ren : sdl_renderer)
                             (px py color_idx : nat) : itree sdlE void :=
  let radius := 13 in
  let '(cr, cg, cb) := ghost_body_color color_idx in
  draw_ghost_body ren px py radius cr cg cb ;;
  draw_ghost_bottom ren px py.

(** * Rocqman sprite rendering *)

Fixpoint pac_row_pixels (ren : sdl_renderer)
                        (sx sy_row : nat) (radius : nat)
                        (fdy dir_ang mouth : R)
                        (dy_sq r2 : nat) (dx count : nat) : itree sdlE void :=
  match count with
  | 0 => Ret ghost
  | S count' =>
    let dx_off := abs_diff dx radius in
    let dist_sq := dx_off * dx_off + dy_sq in
    (if Nat.leb dist_sq r2 then
       let fdx := Rminus (INR dx) (INR radius) in
       let ang := real_atan2 fdy fdx in
       let rel0 := Rminus ang dir_ang in
       let rel1 := if Rlt_dec PI rel0
                   then Rminus rel0 (PI + PI) else rel0 in
       let rel := if Rlt_dec rel1 (Ropp PI)
                  then Rplus rel1 (PI + PI) else rel1 in
       if Rlt_dec (Rabs rel) mouth
       then Ret ghost
       else sdl_draw_point ren (sx + dx) sy_row
     else Ret ghost) ;;
    pac_row_pixels ren sx sy_row radius fdy dir_ang mouth dy_sq r2 (S dx) count'
  end.

Fixpoint pac_rows (ren : sdl_renderer) (sx base_y : nat) (radius : nat)
                  (dir_ang mouth : R) (r2 diameter : nat)
                  (dy count : nat) : itree sdlE void :=
  match count with
  | 0 => Ret ghost
  | S count' =>
    let dy_off := abs_diff dy radius in
    let dy_sq := dy_off * dy_off in
    let fdy := Rminus (INR dy) (INR radius) in
    pac_row_pixels ren sx (base_y + dy) radius fdy dir_ang mouth dy_sq r2 0 diameter ;;
    pac_rows ren sx base_y radius dir_ang mouth r2 diameter (S dy) count'
  end.

Definition dir_angle (dir : nat) : R :=
  match dir with
  | 0 => INR 0
  | 1 => PI
  | 2 => Ropp (PI / INR 2)
  | 3 => PI / INR 2
  | _ => INR 0
  end.

Definition compute_mouth_angle (time_ms : nat) : R :=
  let ftime := INR time_ms in
  let phase := Rdiv ftime (INR 1000) in
  let arg := Rmult (Rmult phase (INR 8)) PI in
  let s := Rabs (sin arg) in
  let f015 := Rdiv (INR 15) (INR 100) in
  let f035 := Rdiv (INR 35) (INR 100) in
  Rplus f015 (Rmult f035 s).

Definition draw_pacman_sprite (ren : sdl_renderer)
                              (px py dir time_ms : nat) : itree sdlE void :=
  let radius := 14 in
  let diameter := radius + radius + 1 in
  let r2 := radius * radius in
  let dir_a := dir_angle dir in
  let mouth := compute_mouth_angle time_ms in
  sdl_set_draw_color ren 255 255 0 ;;
  pac_rows ren (px - radius) (py - radius) radius dir_a mouth
               r2 diameter 0 diameter.

(** * Bitmap font *)

Definition glyph_row_data (g row : nat) : nat :=
  nth row (nth g
    [ [14;17;19;21;25;17;14]
    ; [4;12;4;4;4;4;14]
    ; [14;17;1;6;8;16;31]
    ; [14;17;1;6;1;17;14]
    ; [2;6;10;18;31;2;2]
    ; [31;16;30;1;1;17;14]
    ; [6;8;16;30;17;17;14]
    ; [31;1;2;4;8;8;8]
    ; [14;17;17;14;17;17;14]
    ; [14;17;17;15;1;2;12]
    ; [0;0;0;0;0;0;0]
    ; [4;10;17;17;31;17;17]
    ; [30;17;17;30;17;17;30]
    ; [14;17;16;16;16;17;14]
    ; [28;18;17;17;17;18;28]
    ; [31;16;16;30;16;16;31]
    ; [31;16;16;30;16;16;16]
    ; [14;17;16;23;17;17;14]
    ; [17;17;17;31;17;17;17]
    ; [14;4;4;4;4;4;14]
    ; [7;2;2;2;2;18;12]
    ; [17;18;20;24;20;18;17]
    ; [16;16;16;16;16;16;31]
    ; [17;27;21;17;17;17;17]
    ; [17;25;21;19;17;17;17]
    ; [14;17;17;17;17;17;14]
    ; [30;17;17;30;16;16;16]
    ; [14;17;17;17;21;18;13]
    ; [30;17;17;30;20;18;17]
    ; [14;17;16;14;1;17;14]
    ; [31;4;4;4;4;4;4]
    ; [17;17;17;17;17;17;14]
    ; [17;17;17;17;10;10;4]
    ; [17;17;17;21;21;21;10]
    ; [17;17;10;4;10;17;17]
    ; [17;17;10;4;4;4;4]
    ; [31;1;2;4;8;16;31]
    ] []) 0.

Fixpoint draw_glyph_row (ren : sdl_renderer)
                         (sx sy : nat) (row_bits : nat)
                         (dx count s : nat) : itree sdlE void :=
  match count with
  | 0 => Ret ghost
  | S count' =>
    (if nat_testbit row_bits (4 - dx) then
       sdl_fill_rect ren (sx + dx * s) sy s s
     else Ret ghost) ;;
    draw_glyph_row ren sx sy row_bits (S dx) count' s
  end.

Fixpoint draw_glyph_rows (ren : sdl_renderer) (sx sy g : nat)
                          (row count s : nat) : itree sdlE void :=
  match count with
  | 0 => Ret ghost
  | S count' =>
    let bits := glyph_row_data g row in
    draw_glyph_row ren sx (sy + row * s) bits 0 5 s ;;
    draw_glyph_rows ren sx sy g (S row) count' s
  end.

Definition draw_one_glyph (ren : sdl_renderer) (sx sy g s : nat) : itree sdlE void :=
  draw_glyph_rows ren sx sy g 0 7 s.

Fixpoint draw_glyphs (ren : sdl_renderer) (sx sy s : nat)
                      (glyphs : list nat) : itree sdlE void :=
  match glyphs with
  | [] => Ret ghost
  | g :: rest =>
    draw_one_glyph ren sx sy g s ;;
    draw_glyphs ren (sx + 6 * s) sy s rest
  end.

Fixpoint draw_number_digits (ren : sdl_renderer) (sx sy : nat)
                            (digits : list nat) : itree sdlE void :=
  match digits with
  | [] => Ret ghost
  | d :: rest =>
    draw_one_glyph ren sx sy d 3 ;;
    draw_number_digits ren (sx + 18) sy rest
  end.

(** Convert [n] to a list of decimal digit indices for the bitmap font.
    E.g. [nat_digit_list 123 = [1;2;3]]. *)
Fixpoint nat_digits_aux (n : nat) (acc : list nat) (fuel : nat) : list nat :=
  match fuel with
  | 0 => acc
  | S fuel' =>
    let d    := Nat.modulo n 10 in
    let rest := Nat.div n 10 in
    if Nat.eqb rest 0 then (d :: acc)
    else nat_digits_aux rest (d :: acc) fuel'
  end.

Definition nat_digit_list (n : nat) : list nat :=
  match n with
  | 0 => [0]
  | _ => nat_digits_aux n [] 12
  end.

Definition draw_number_sprite (ren : sdl_renderer) (n sx sy : nat) : itree sdlE void :=
  sdl_set_draw_color ren 255 255 255 ;;
  draw_number_digits ren sx sy (nat_digit_list n).

(** * Arcade text messages *)

Definition msg_game_over : list nat :=
  [17;11;23;15;10;25;32;15;28].

Definition msg_you_win : list nat :=
  [35;25;31;10;33;19;24].

Definition msg_lives_left (n : nat) : list nat :=
  nat_digit_list n ++ [10;22;19;32;15;29;10;22;15;16;30].

Definition msg_paused : list nat :=
  [26;11;31;29;15;14].

Definition draw_message_screen (ren : sdl_renderer) (msg : list nat)
  : itree sdlE void :=
  let s := 4 in
  let glyph_w := 6 * s in
  let text_w := length msg * glyph_w in
  let sx := Nat.div (win_width - text_w) 2 in
  let sy := Nat.div win_height 2 - Nat.div (7 * s) 2 in
  sdl_set_draw_color ren 0 0 0 ;;
  sdl_clear ren ;;
  sdl_set_draw_color ren 255 255 255 ;;
  draw_glyphs ren sx sy s msg ;;
  sdl_present ren.

Definition draw_status_message (ren : sdl_renderer) (msg : list nat) : itree sdlE void :=
  let s := 3 in
  let glyph_w := 6 * s in
  let text_w := length msg * glyph_w in
  let sx := Nat.div (win_width - text_w) 2 in
  let sy := board_height * cell_size + Nat.div (status_height - 7 * s) 2 in
  sdl_set_draw_color ren 255 255 255 ;;
  draw_glyphs ren sx sy s msg.

(** * Life icon *)

Definition draw_life_icon (ren : sdl_renderer) (x y : nat) : itree sdlE void :=
  sdl_set_draw_color ren 150 10 35 ;;
  sdl_fill_rect ren (x - 9) (y - 6) 6 6 ;;
  sdl_fill_rect ren (x + 3) (y - 6) 6 6 ;;
  sdl_fill_rect ren (x - 11) y 22 6 ;;
  sdl_fill_rect ren (x - 7) (y + 6) 14 4 ;;
  sdl_fill_rect ren (x - 5) (y + 10) 10 4 ;;
  sdl_fill_rect ren (x - 3) (y + 14) 6 3 ;;
  sdl_fill_rect ren (x - 1) (y + 17) 2 2 ;;
  sdl_set_draw_color ren 230 45 85 ;;
  sdl_fill_rect ren (x - 7) (y - 5) 4 4 ;;
  sdl_fill_rect ren (x + 3) (y - 5) 4 4 ;;
  sdl_fill_rect ren (x - 8) (y - 1) 16 5 ;;
  sdl_fill_rect ren (x - 5) (y + 4) 10 4 ;;
  sdl_fill_rect ren (x - 3) (y + 8) 6 4 ;;
  sdl_set_draw_color ren 255 170 190 ;;
  sdl_fill_rect ren (x - 5) (y - 4) 2 2 ;;
  sdl_fill_rect ren (x + 3) (y - 4) 2 2.

(** * Pixel coordinate helpers *)

Definition half_cell : nat := 16.

Definition cell_center_x (col : nat) : nat :=
  col * cell_size + half_cell.

Definition cell_center_y (row : nat) : nat :=
  row * cell_size + half_cell.

Definition lerp (from_v to_v num den : nat) : nat :=
  if Nat.leb den 0 then to_v
  else
    if Nat.leb from_v to_v
    then from_v + Nat.div ((to_v - from_v) * num) den
    else from_v - Nat.div ((from_v - to_v) * num) den.

Definition dir_to_nat (d : direction) : nat :=
  match d with
  | Right => 0
  | Left => 1
  | Up => 2
  | Down => 3
  | DirNone => 0
  end.

Definition ghost_color_index (idx : nat) (gm : ghost_mode) : nat :=
  match gm with
  | Frightened => 4
  | Chase => idx
  end.

(** * SDL rendering *)

Definition draw_dot_check (ren : sdl_renderer) (cx cy : nat) : itree sdlE void :=
  sdl_set_draw_color ren 60 200 90 ;;
  sdl_fill_rect ren (cx - 6) (cy + 1) 3 3 ;;
  sdl_fill_rect ren (cx - 3) (cy + 4) 3 3 ;;
  sdl_fill_rect ren cx (cy + 1) 3 3 ;;
  sdl_fill_rect ren (cx + 3) (cy - 2) 3 3.

Fixpoint draw_row_cells (ren : sdl_renderer) (row col : nat)
                        (cells : list cell) (pellet_phase : nat) : itree sdlE void :=
  match cells with
  | [] => Ret ghost
  | c :: rest =>
    (match c with
     | Wall =>
       sdl_set_draw_color ren 33 33 222 ;;
       sdl_fill_rect ren (col * cell_size + 1) (row * cell_size + 1)
                     (cell_size - 2) (cell_size - 2)
     | Dot =>
       sdl_set_draw_color ren 255 255 255 ;;
       draw_filled_circle ren (cell_center_x col) (cell_center_y row) 2
     | PowerPellet =>
       draw_dot_check ren (cell_center_x col) (cell_center_y row)
     | Empty => Ret ghost
     end) ;;
    draw_row_cells ren row (S col) rest pellet_phase
  end.

Fixpoint draw_board_rows (ren : sdl_renderer) (row : nat)
                         (rows : list (list cell)) (pellet_phase : nat) : itree sdlE void :=
  match rows with
  | [] => Ret ghost
  | cells :: rest =>
    draw_row_cells ren row 0 cells pellet_phase ;;
    draw_board_rows ren (S row) rest pellet_phase
  end.

Definition draw_board_sdl (ren : sdl_renderer) (gs : game_state)
                          (pellet_phase : nat) : itree sdlE void :=
  draw_board_rows ren 0 (board gs) pellet_phase.

Fixpoint draw_ghosts_aux (ren : sdl_renderer) (idx : nat) (gs : list ghost_state)
                         (prev_gs : list ghost_state)
                         (t_num t_den : nat) (time_ms : nat) : itree sdlE void :=
  match gs with
  | [] => Ret ghost
  | g :: rest =>
    let prev_g := match prev_gs with
                  | [] => g
                  | pg :: _ => pg
                  end in
    let prev_rest := match prev_gs with
                     | [] => []
                     | _ :: pr => pr
                     end in
    let px := lerp (cell_center_x (pcol (gpos prev_g)))
                   (cell_center_x (pcol (gpos g))) t_num t_den in
    let py := lerp (cell_center_y (prow (gpos prev_g)))
                   (cell_center_y (prow (gpos g))) t_num t_den in
    let col := ghost_color_index (Nat.modulo idx 4) (gmode g) in
    draw_ghost_sprite ren px py col ;;
    draw_ghosts_aux ren (S idx) rest prev_rest t_num t_den time_ms
  end.

Definition dir_to_degrees (d : direction) : nat :=
  match d with
  | Right   => 0
  | Down    => 90
  | Left    => 0
  | Up      => 270
  | DirNone => 0
  end.

Definition dir_flip_h (d : direction) : bool :=
  match d with
  | Left => true
  | _ => false
  end.

Definition draw_player_sdl (ren : sdl_renderer) (tex : sdl_texture)
                           (gs : game_state) (prev_pos : position)
                           (t_num t_den : nat) : itree sdlE void :=
  let px := lerp (cell_center_x (pcol prev_pos))
                 (cell_center_x (pcol (pacpos gs))) t_num t_den in
  let py := lerp (cell_center_y (prow prev_pos))
                 (cell_center_y (prow (pacpos gs))) t_num t_den in
  sdl_render_texture_rotated ren tex px py 28 28
                             (dir_to_degrees (pacdir gs))
                             (dir_flip_h (pacdir gs)).

Fixpoint draw_lives_aux (ren : sdl_renderer) (n : nat) (i : nat) : itree sdlE void :=
  match n with
  | 0 => Ret ghost
  | S n' =>
    draw_life_icon ren (win_width - 30 - i * 28)
                   (board_height * cell_size + 8 + 12) ;;
    draw_lives_aux ren n' (S i)
  end.

Definition draw_status_bar (ren : sdl_renderer) (gs : game_state) : itree sdlE void :=
  draw_number_sprite ren (score gs) 10 (board_height * cell_size + 8) ;;
  draw_lives_aux ren (lives gs) 0.

Definition render_frame (ren : sdl_renderer) (tex : sdl_texture)
                        (gs : game_state)
                        (prev_pac : position) (prev_ghosts : list ghost_state)
                        (t_num t_den : nat) (time_ms : nat) : itree sdlE void :=
  sdl_set_draw_color ren 0 0 0 ;;
  sdl_clear ren ;;
  draw_board_sdl ren gs time_ms ;;
  draw_ghosts_aux ren 0 (ghosts gs) prev_ghosts t_num t_den time_ms ;;
  draw_player_sdl ren tex gs prev_pac t_num t_den ;;
  draw_status_bar ren gs ;;
  sdl_present ren.

Definition render_paused_frame (ren : sdl_renderer) (tex : sdl_texture)
                        (gs : game_state)
                        (prev_pac : position) (prev_ghosts : list ghost_state)
                        (t_num t_den : nat) (time_ms : nat) : itree sdlE void :=
  sdl_set_draw_color ren 0 0 0 ;;
  sdl_clear ren ;;
  draw_board_sdl ren gs time_ms ;;
  draw_ghosts_aux ren 0 (ghosts gs) prev_ghosts t_num t_den time_ms ;;
  draw_player_sdl ren tex gs prev_pac t_num t_den ;;
  draw_status_bar ren gs ;;
  draw_status_message ren msg_paused ;;
  sdl_present ren.

(** * Event handling *)

Definition handle_key_down (key : sdl_key) (gs : game_state) : (bool * game_state) :=
  match key with
  | KeyEscape | KeyQ => (true, gs)
  | KeyUp | KeyW => (false, set_direction Up gs)
  | KeyDown | KeyS => (false, set_direction Down gs)
  | KeyLeft | KeyA => (false, set_direction Left gs)
  | KeyRight | KeyD => (false, set_direction Right gs)
  | KeySpace => (false, gs)
  end.

Definition handle_event (ev : sdl_event) (gs : game_state) : (bool * game_state) :=
  match ev with
  | EventQuit => (true, gs)
  | EventKeyDown key => handle_key_down key gs
  end.

(** * Pixel-level collision detection *)

Fixpoint find_pixel_collision (px py : nat) (gs : list ghost_state)
    (prev_gs : list ghost_state) (t_num t_den threshold idx : nat)
    : option (nat * ghost_mode) :=
  match gs with
  | [] => None
  | g :: rest =>
    let prev_g := match prev_gs with [] => g | pg :: _ => pg end in
    let prev_rest := match prev_gs with [] => [] | _ :: pr => pr end in
    let gx := lerp (cell_center_x (pcol (gpos prev_g)))
                   (cell_center_x (pcol (gpos g))) t_num t_den in
    let gy := lerp (cell_center_y (prow (gpos prev_g)))
                   (cell_center_y (prow (gpos g))) t_num t_den in
    let dx := abs_diff px gx in
    let dy := abs_diff py gy in
    if Nat.leb (dx * dx + dy * dy) (threshold * threshold) then
      Some (idx, gmode g)
    else
      find_pixel_collision px py rest prev_rest t_num t_den threshold (S idx)
  end.

Definition eat_ghost_idx (idx : nat) (gs : game_state) : game_state :=
  let default_g := mkGhost (mkPos 0 0) DirNone Chase in
  let new_ghosts := replace_nth idx (ghosts gs)
                      (respawn_ghost (nth idx (ghosts gs) default_g)) in
  mkState (board gs) (pacpos gs) (pacdir gs) (desired_dir gs) new_ghosts
          (score gs + 200) (lives gs) (dots_left gs) (power_timer gs)
          (game_over gs) (game_won gs).

Definition lose_one_life (gs : game_state) : game_state :=
  let new_lives := Nat.pred (lives gs) in
  mkState (board gs) (mkPos 9 9) DirNone DirNone (ghosts gs)
          (score gs) new_lives (dots_left gs) 0
          (Nat.eqb new_lives 0) (game_won gs).

Definition apply_direction (key : sdl_key) (gs : game_state) : game_state :=
  match key with
  | KeyUp | KeyW => set_direction Up gs
  | KeyDown | KeyS => set_direction Down gs
  | KeyLeft | KeyA => set_direction Left gs
  | KeyRight | KeyD => set_direction Right gs
  | _ => gs
  end.

(** * Game loop state *)

Record loop_state : Type := mkLoop {
  ls_game : game_state;
  ls_prev_pac : position;
  ls_prev_ghosts : list ghost_state;
  ls_last_tick : nat;
  ls_start_time : nat;
  ls_texture : sdl_texture;
  ls_phase : phase;
  ls_phase_time : nat;
  ls_quit : bool
}.

(** * Process one frame *)

Definition frame_delay (frame_start : nat) : itree sdlE void :=
  now2 <- sdl_get_ticks ;;
  let elapsed := now2 - frame_start in
  if Nat.ltb elapsed frame_ms
  then sdl_delay (frame_ms - elapsed)
  else Ret ghost.

Definition collision_threshold : nat := 22.

Definition snd_checkmark : PrimString.string := "assets/checkmark.mp3".
Definition snd_game_over : PrimString.string := "assets/game-over.mp3".
Definition snd_kill_ghost : PrimString.string := "assets/kill-ghost.mp3".
Definition snd_lose_life : PrimString.string := "assets/lose-life.mp3".
Definition snd_tap : PrimString.string := "assets/tap.mp3".
Definition snd_win : PrimString.string := "assets/win.mp3".

Definition play_cell_sound (c : cell) : itree sdlE void :=
  match c with
  | Dot => sdl_play_sound snd_tap
  | PowerPellet => sdl_play_sound snd_checkmark
  | _ => Ret ghost
  end.

Definition process_frame (ren : sdl_renderer) (ls : loop_state)
  : itree sdlE (bool * loop_state) :=
  ev <- sdl_poll_event ;;
  match ev with
  | EventQuit => Ret (true, ls)
  | EventKeyDown _ =>
    now <- sdl_get_ticks ;;
    let time_ms := now - ls_start_time ls in
    match ls_phase ls with
    | Playing =>
      match ev with
      | EventKeyDown KeySpace =>
        render_paused_frame ren (ls_texture ls) (ls_game ls)
                            (ls_prev_pac ls) (ls_prev_ghosts ls)
                            0 1 time_ms ;;
        frame_delay now ;;
        Ret (false, mkLoop (ls_game ls) (ls_prev_pac ls) (ls_prev_ghosts ls)
                           now (ls_start_time ls) (ls_texture ls)
                           Paused now false)
      | EventKeyDown key =>
        let '(quit, gs1) := handle_key_down key (ls_game ls) in
        if quit then Ret (true, ls) else
        let elapsed := now - ls_last_tick ls in
        let do_tick := Nat.leb tick_ms elapsed in
        let gs2 := if do_tick then tick gs1 else gs1 in
        let new_prev_pac := if do_tick then pacpos gs1 else ls_prev_pac ls in
        let new_prev_ghosts := if do_tick then ghosts gs1 else ls_prev_ghosts ls in
        let new_last_tick := if do_tick then now else ls_last_tick ls in
        let eaten_cell := if do_tick
                          then get_cell (prow (pacpos gs2)) (pcol (pacpos gs2))
                                        (board gs1)
                          else Empty in
        let t_num := now - new_last_tick in
        let ppx := lerp (cell_center_x (pcol new_prev_pac))
                        (cell_center_x (pcol (pacpos gs2))) t_num tick_ms in
        let ppy := lerp (cell_center_y (prow new_prev_pac))
                        (cell_center_y (prow (pacpos gs2))) t_num tick_ms in
        if game_won gs2 then
          sdl_play_sound snd_win ;;
          render_frame ren (ls_texture ls) gs2 new_prev_pac new_prev_ghosts
                       t_num tick_ms time_ms ;;
          Ret (false, mkLoop gs2 new_prev_pac new_prev_ghosts
                             new_last_tick (ls_start_time ls) (ls_texture ls)
                             WinScreen now false)
        else
        match find_pixel_collision ppx ppy (ghosts gs2) new_prev_ghosts
                t_num tick_ms collision_threshold 0 with
        | Some (idx, Frightened) =>
          let gs3 := eat_ghost_idx idx gs2 in
          let next_ls := mkLoop gs3 new_prev_pac new_prev_ghosts
                                new_last_tick (ls_start_time ls) (ls_texture ls)
                                Playing 0 false in
          sdl_play_sound snd_kill_ghost ;;
          render_frame ren (ls_texture ls) gs3 (ls_prev_pac next_ls)
                       (ls_prev_ghosts next_ls)
                       t_num tick_ms time_ms ;;
          frame_delay now ;;
          Ret (false, next_ls)
        | Some (_, Chase) =>
          let gs3 := lose_one_life gs2 in
          let next_pac := pacpos gs3 in
          let next_ghosts := ghosts gs3 in
          if Nat.eqb (lives gs3) 0 then
            let next_ls := mkLoop gs3 next_pac next_ghosts
                                  now (ls_start_time ls) (ls_texture ls)
                                  GameOverScreen now false in
            sdl_play_sound snd_game_over ;;
            Ret (false, next_ls)
          else
            let next_ls := mkLoop gs3 next_pac next_ghosts
                                  now (ls_start_time ls) (ls_texture ls)
                                  DeathPause now false in
            sdl_play_sound snd_lose_life ;;
            Ret (false, next_ls)
        | None =>
          play_cell_sound eaten_cell ;;
          render_frame ren (ls_texture ls) gs2 new_prev_pac new_prev_ghosts
                       t_num tick_ms time_ms ;;
          frame_delay now ;;
          Ret (false, mkLoop gs2 new_prev_pac new_prev_ghosts
                             new_last_tick (ls_start_time ls) (ls_texture ls)
                             Playing 0 false)
        end
      | _ =>
        let gs1 := ls_game ls in
        let elapsed := now - ls_last_tick ls in
        let do_tick := Nat.leb tick_ms elapsed in
        let gs2 := if do_tick then tick gs1 else gs1 in
        let new_prev_pac := if do_tick then pacpos gs1 else ls_prev_pac ls in
        let new_prev_ghosts := if do_tick then ghosts gs1 else ls_prev_ghosts ls in
        let new_last_tick := if do_tick then now else ls_last_tick ls in
        let eaten_cell := if do_tick
                          then get_cell (prow (pacpos gs2)) (pcol (pacpos gs2))
                                        (board gs1)
                          else Empty in
        let t_num := now - new_last_tick in
        let ppx := lerp (cell_center_x (pcol new_prev_pac))
                        (cell_center_x (pcol (pacpos gs2))) t_num tick_ms in
        let ppy := lerp (cell_center_y (prow new_prev_pac))
                        (cell_center_y (prow (pacpos gs2))) t_num tick_ms in
        if game_won gs2 then
          sdl_play_sound snd_win ;;
          render_frame ren (ls_texture ls) gs2 new_prev_pac new_prev_ghosts
                       t_num tick_ms time_ms ;;
          Ret (false, mkLoop gs2 new_prev_pac new_prev_ghosts
                             new_last_tick (ls_start_time ls) (ls_texture ls)
                             WinScreen now false)
        else
        match find_pixel_collision ppx ppy (ghosts gs2) new_prev_ghosts
                t_num tick_ms collision_threshold 0 with
        | Some (idx, Frightened) =>
          let gs3 := eat_ghost_idx idx gs2 in
          let next_ls := mkLoop gs3 new_prev_pac new_prev_ghosts
                                new_last_tick (ls_start_time ls) (ls_texture ls)
                                Playing 0 false in
          sdl_play_sound snd_kill_ghost ;;
          render_frame ren (ls_texture ls) gs3 (ls_prev_pac next_ls)
                       (ls_prev_ghosts next_ls)
                       t_num tick_ms time_ms ;;
          frame_delay now ;;
          Ret (false, next_ls)
        | Some (_, Chase) =>
          let gs3 := lose_one_life gs2 in
          let next_pac := pacpos gs3 in
          let next_ghosts := ghosts gs3 in
          if Nat.eqb (lives gs3) 0 then
            let next_ls := mkLoop gs3 next_pac next_ghosts
                                  now (ls_start_time ls) (ls_texture ls)
                                  GameOverScreen now false in
            sdl_play_sound snd_game_over ;;
            Ret (false, next_ls)
          else
            let next_ls := mkLoop gs3 next_pac next_ghosts
                                  now (ls_start_time ls) (ls_texture ls)
                                  DeathPause now false in
            sdl_play_sound snd_lose_life ;;
            Ret (false, next_ls)
        | None =>
          play_cell_sound eaten_cell ;;
          render_frame ren (ls_texture ls) gs2 new_prev_pac new_prev_ghosts
                       t_num tick_ms time_ms ;;
          frame_delay now ;;
          Ret (false, mkLoop gs2 new_prev_pac new_prev_ghosts
                             new_last_tick (ls_start_time ls) (ls_texture ls)
                             Playing 0 false)
        end
      end

    | Paused =>
      match ev with
      | EventKeyDown KeySpace =>
        let gs := ls_game ls in
        Ret (false, mkLoop gs (pacpos gs) (ghosts gs)
                           now (ls_start_time ls) (ls_texture ls)
                           Playing 0 false)
      | EventKeyDown key =>
        let '(quit, _) := handle_key_down key (ls_game ls) in
        if quit then Ret (true, ls) else
        render_paused_frame ren (ls_texture ls) (ls_game ls)
                            (ls_prev_pac ls) (ls_prev_ghosts ls)
                            0 1 time_ms ;;
        frame_delay now ;;
        Ret (false, ls)
      | _ =>
        render_paused_frame ren (ls_texture ls) (ls_game ls)
                            (ls_prev_pac ls) (ls_prev_ghosts ls)
                            0 1 time_ms ;;
        frame_delay now ;;
        Ret (false, ls)
      end

    | DeathPause =>
      if Nat.leb 2000 (now - ls_phase_time ls) then
        Ret (false, mkLoop (ls_game ls) (pacpos (ls_game ls))
                           (ghosts (ls_game ls)) now (ls_start_time ls)
                           (ls_texture ls) Playing 0 false)
      else
        draw_message_screen ren (msg_lives_left (lives (ls_game ls))) ;;
        frame_delay now ;;
        Ret (false, ls)

    | GameOverScreen =>
      if Nat.leb 3000 (now - ls_phase_time ls) then
        Ret (true, ls)
      else
        draw_message_screen ren msg_game_over ;;
        frame_delay now ;;
        Ret (false, ls)

    | WinScreen =>
      if Nat.leb 3000 (now - ls_phase_time ls) then
        Ret (true, ls)
      else
        draw_message_screen ren msg_you_win ;;
        frame_delay now ;;
        Ret (false, ls)
    end
  end.

(** * Init and cleanup *)

Definition init_game : itree sdlE (sdl_window * sdl_renderer * loop_state) :=
  win <- sdl_create_window "Rocqman" win_width win_height ;;
  ren <- sdl_create_renderer win ;;
  tex <- sdl_load_texture ren "assets/rocq.svg" ;;
  t0 <- sdl_get_ticks ;;
  let gs := initial_state in
  let ls := mkLoop gs (pacpos gs) (ghosts gs) t0 t0 tex Playing 0 false in
  Ret (win, ren, ls).

Definition cleanup (ren : sdl_renderer) (win : sdl_window) : itree sdlE void :=
  sdl_destroy ren win.

End Rocqman.

Import Rocqman.
Import ITreeNotations.

(** Cleanup wrapper — returns void (no C exit-code). *)
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

(** * Go extraction *)

(** Additional nat/Nat mappings not covered by GoSDL2 *)
Crane Extract Inlined Constant Nat.testbit =>
  "((%a0 >> uint(%a1)) & 1 != 0)".

(** Enable loopification: converts run_game CoFixpoint into a for loop. *)
Set Crane Loopify.

(** Extract to [main.go] (package main). *)
Crane Extraction "main" Rocqman run_game exit_game.
