(* Copyright 2025 Bloomberg Finance L.P. *)
(* Distributed under the terms of the GNU LGPL v2.1 license. *)

(** * Game loop: event handling, frame processing, initialization. *)

From Corelib Require Import PrimString.
From Stdlib Require Import Lists.List.
Import ListNotations.
From RocqmanGoGame Require Import GoSDL2 Types Game Render.
From ITree Require Import Core.ITreeDefinition.
Import Types Game Render.

Local Open Scope pstring_scope.
Local Open Scope itree_scope.

Import ITreeNotations.

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
