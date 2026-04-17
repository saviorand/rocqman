package main

type cell int

const (
	Wall cell = iota
	Empty
	Dot
	PowerPellet
)

type direction int

const (
	Up direction = iota
	Down
	Left
	Right
	DirNone
)

type ghost_mode int

const (
	Chase ghost_mode = iota
	Frightened
)

type phase int

const (
	Playing phase = iota
	Paused
	DeathPause
	GameOverScreen
	WinScreen
)

type position struct {
	prow uint
	pcol uint
}

func mkPos(prow uint, pcol uint) position {
	return position{prow: prow, pcol: pcol}
}

type ghost_state struct {
	gpos  position
	gdir  direction
	gmode ghost_mode
}

func mkGhost(gpos position, gdir direction, gmode ghost_mode) ghost_state {
	return ghost_state{gpos: gpos, gdir: gdir, gmode: gmode}
}

type game_state struct {
	board       list[list[cell]]
	pacpos      position
	pacdir      direction
	desired_dir direction
	ghosts      list[ghost_state]
	score       uint
	lives       uint
	dots_left   uint
	power_timer uint
	game_over   bool
	game_won    bool
}

func mkState(board list[list[cell]], pacpos position, pacdir direction, desired_dir direction, ghosts list[ghost_state], score uint, lives uint, dots_left uint, power_timer uint, game_over bool, game_won bool) game_state {
	return game_state{board: board, pacpos: pacpos, pacdir: pacdir, desired_dir: desired_dir, ghosts: ghosts, score: score, lives: lives, dots_left: dots_left, power_timer: power_timer, game_over: game_over, game_won: game_won}
}
