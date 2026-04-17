
(** * Game mechanics: board layout, movement, ghost AI, collision, tick. *)

From Stdlib Require Import Lists.List.
Import ListNotations.
From Stdlib Require Import Bool.
From RocqmanGoGame Require Import GoSDL2 Types.
Import Types.

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

(** * Board helpers *)

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
