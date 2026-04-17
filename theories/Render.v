
(** * Rendering: pixel helpers, sprites, font, board drawing. *)

From Corelib Require Import PrimString.
From Stdlib Require Import Lists.List.
Import ListNotations.
From Stdlib Require Import Reals Rtrigo Ratan.
From RocqmanGoGame Require Import GoSDL2 Types Game.
From ITree Require Import Core.ITreeDefinition.
Import Types Game.

Local Open Scope pstring_scope.
Local Open Scope itree_scope.

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

Definition ghost_color_index (idx : nat) (gm : ghost_mode) : nat :=
  match gm with
  | Frightened => 4
  | Chase => idx
  end.

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

(** * Pacman sprite rendering *)

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
