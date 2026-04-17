package main

import "math"

type listImpl struct {
	_v     int
	_c1_d0 any
	_c1_d1 any
}

type list[T1 any] = *listImpl

var nil0 *listImpl = &listImpl{_v: 0}

func cons(a0 any, a1 any) *listImpl {
	return &listImpl{_v: 1, _c1_d0: a0, _c1_d1: a1}
}

func length(l any) uint {
	return (func() uint {
		var _box any = l
		_scrut1 := _box.(*listImpl)
		switch _scrut1._v {
		case 0:
			return uint(0)
		case 1:
			l_ := _scrut1._c1_d1
			return uint((length((l_).(list[any])) + 1))
		}
		panic("unreachable")
	})()
}

func app(l any, m any) list[any] {
	return (func() list[any] {
		var _box any = l
		_scrut2 := _box.(*listImpl)
		switch _scrut2._v {
		case 0:
			return (m).(list[any])
		case 1:
			a := _scrut2._c1_d0
			l1 := _scrut2._c1_d1
			return cons(a, app((l1).(list[any]), (m).(list[any])))
		}
		panic("unreachable")
	})()
}

func add(n uint, m uint) uint {
	return (func() any {
		var _sni_ any = n
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return m
		} else {
			p := _su_ - 1
			return uint((add(p, m) + 1))
		}
	}()).(uint)
}

func mul(n uint, m uint) uint {
	return (func() any {
		var _sni_ any = n
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return uint(0)
		} else {
			p := _su_ - 1
			return add(m, mul(p, m))
		}
	}()).(uint)
}

func sub(n uint, m uint) uint {
	return (func() any {
		var _sni_ any = n
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return n
		} else {
			k := _su_ - 1
			return (func() any {
				var _sni_ any = m
				_su_ := _sni_.(uint)
				if _su_ == 0 {
					return n
				} else {
					l := _su_ - 1
					return sub(k, l)
				}
			}()).(uint)
		}
	}()).(uint)
}

func map0[T1 any, T2 any](f func(T1) T2, l any) list[T2] {
	return (func() list[T2] {
		var _box any = l
		_scrut3 := _box.(*listImpl)
		switch _scrut3._v {
		case 0:
			return nil0
		case 1:
			a := _scrut3._c1_d0
			l0 := _scrut3._c1_d1
			return cons(f((a).(T1)), map0(f, (l0).(list[any])))
		}
		panic("unreachable")
	})()
}

func nth(n any, l any, default0 any) any {
	return func() any {
		var _sni_ any = n
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return (func() any {
				var _box any = l
				_scrut4 := _box.(*listImpl)
				switch _scrut4._v {
				case 0:
					return default0
				case 1:
					x := _scrut4._c1_d0
					return x
				}
				panic("unreachable")
			})()
		} else {
			m := _su_ - 1
			return (func() any {
				var _box any = l
				_scrut5 := _box.(*listImpl)
				switch _scrut5._v {
				case 0:
					return default0
				case 1:
					l_ := _scrut5._c1_d1
					return nth(m, (l_).(list[any]), default0)
				}
				panic("unreachable")
			})()
		}
	}()
}

type sdl_key int

const (
	KeyEscape sdl_key = iota
	KeyQ
	KeyUp
	KeyDown
	KeyLeft
	KeyRight
	KeyW
	KeyA
	KeyS
	KeyD
	KeySpace
)

type sdl_eventImpl struct {
	_v     int
	_c1_d0 any
}

type sdl_event = *sdl_eventImpl

var EventQuit *sdl_eventImpl = &sdl_eventImpl{_v: 0}

func EventKeyDown(a0 any) *sdl_eventImpl {
	return &sdl_eventImpl{_v: 1, _c1_d0: a0}
}

func real_atan2(y float64, x float64) float64 {
	return (func() any {
		if 0.0 < x {
			return math.Atan((y / x))
		} else {
			return (func() any {
				if x < 0.0 {
					return (func() any {
						if 0.0 <= y {
							return (math.Atan((y / x)) + math.Pi)
						} else {
							return (math.Atan((y / x)) - math.Pi)
						}
					}()).(float64)
				} else {
					return (func() any {
						if 0.0 < y {
							return (math.Pi / float64(uint((uint((uint(0) + 1)) + 1))))
						} else {
							return (func() any {
								if y < 0.0 {
									return (-(math.Pi / float64(uint((uint((uint(0) + 1)) + 1)))))
								} else {
									return 0.0
								}
							}()).(float64)
						}
					}()).(float64)
				}
			}()).(float64)
		}
	}()).(float64)
}

func nat_testbit(a0 uint, a1 uint) bool {
	return ((a0>>uint(a1))&1 != 0)
}

type cell int

const (
	Wall cell = iota
	Empty
	Dot
	PowerPellet
)

func cell_rect[T1 any](f T1, f0 T1, f1 T1, f2 T1, c cell) T1 {
	return (func() T1 {
		_scrut6 := c
		switch _scrut6 {
		case Wall:
			return f
		case Empty:
			return f0
		case Dot:
			return f1
		case PowerPellet:
			return f2
		}
		panic("unreachable")
	})()
}

func cell_rec[T1 any](f T1, f0 T1, f1 T1, f2 T1, c cell) T1 {
	return (func() T1 {
		_scrut7 := c
		switch _scrut7 {
		case Wall:
			return f
		case Empty:
			return f0
		case Dot:
			return f1
		case PowerPellet:
			return f2
		}
		panic("unreachable")
	})()
}

type direction int

const (
	Up direction = iota
	Down
	Left
	Right
	DirNone
)

func direction_rect[T1 any](f T1, f0 T1, f1 T1, f2 T1, f3 T1, d direction) T1 {
	return (func() T1 {
		_scrut8 := d
		switch _scrut8 {
		case Up:
			return f
		case Down:
			return f0
		case Left:
			return f1
		case Right:
			return f2
		case DirNone:
			return f3
		}
		panic("unreachable")
	})()
}

func direction_rec[T1 any](f T1, f0 T1, f1 T1, f2 T1, f3 T1, d direction) T1 {
	return (func() T1 {
		_scrut9 := d
		switch _scrut9 {
		case Up:
			return f
		case Down:
			return f0
		case Left:
			return f1
		case Right:
			return f2
		case DirNone:
			return f3
		}
		panic("unreachable")
	})()
}

type ghost_mode int

const (
	Chase ghost_mode = iota
	Frightened
)

func ghost_mode_rect[T1 any](f T1, f0 T1, g ghost_mode) T1 {
	return (func() T1 {
		_scrut10 := g
		switch _scrut10 {
		case Chase:
			return f
		case Frightened:
			return f0
		}
		panic("unreachable")
	})()
}

func ghost_mode_rec[T1 any](f T1, f0 T1, g ghost_mode) T1 {
	return (func() T1 {
		_scrut11 := g
		switch _scrut11 {
		case Chase:
			return f
		case Frightened:
			return f0
		}
		panic("unreachable")
	})()
}

type phase int

const (
	Playing phase = iota
	Paused
	DeathPause
	GameOverScreen
	WinScreen
)

func phase_rect[T1 any](f T1, f0 T1, f1 T1, f2 T1, f3 T1, p phase) T1 {
	return (func() T1 {
		_scrut12 := p
		switch _scrut12 {
		case Playing:
			return f
		case Paused:
			return f0
		case DeathPause:
			return f1
		case GameOverScreen:
			return f2
		case WinScreen:
			return f3
		}
		panic("unreachable")
	})()
}

func phase_rec[T1 any](f T1, f0 T1, f1 T1, f2 T1, f3 T1, p phase) T1 {
	return (func() T1 {
		_scrut13 := p
		switch _scrut13 {
		case Playing:
			return f
		case Paused:
			return f0
		case DeathPause:
			return f1
		case GameOverScreen:
			return f2
		case WinScreen:
			return f3
		}
		panic("unreachable")
	})()
}

type position struct {
	prow uint
	pcol uint
}

func mkPos(prow uint, pcol uint) position {
	return position{prow: prow, pcol: pcol}
}

func prow(p position) uint {
	return (func() uint {
		_scrut14 := p
		prow0 := _scrut14.prow
		return prow0
	})()
}

func pcol(p position) uint {
	return (func() uint {
		_scrut15 := p
		pcol0 := _scrut15.pcol
		return pcol0
	})()
}

type ghost_state struct {
	gpos  position
	gdir  direction
	gmode ghost_mode
}

func mkGhost(gpos position, gdir direction, gmode ghost_mode) ghost_state {
	return ghost_state{gpos: gpos, gdir: gdir, gmode: gmode}
}

func gpos(g ghost_state) position {
	return (func() position {
		_scrut16 := g
		gpos0 := _scrut16.gpos
		return gpos0
	})()
}

func gdir(g ghost_state) direction {
	return (func() direction {
		_scrut17 := g
		gdir0 := _scrut17.gdir
		return gdir0
	})()
}

func gmode(g ghost_state) ghost_mode {
	return (func() ghost_mode {
		_scrut18 := g
		gmode0 := _scrut18.gmode
		return gmode0
	})()
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

func board(g game_state) list[list[cell]] {
	return (func() list[list[cell]] {
		_scrut19 := g
		board0 := _scrut19.board
		return board0
	})()
}

func pacpos(g game_state) position {
	return (func() position {
		_scrut20 := g
		pacpos0 := _scrut20.pacpos
		return pacpos0
	})()
}

func pacdir(g game_state) direction {
	return (func() direction {
		_scrut21 := g
		pacdir0 := _scrut21.pacdir
		return pacdir0
	})()
}

func desired_dir(g game_state) direction {
	return (func() direction {
		_scrut22 := g
		desired_dir0 := _scrut22.desired_dir
		return desired_dir0
	})()
}

func ghosts(g game_state) list[ghost_state] {
	return (func() list[ghost_state] {
		_scrut23 := g
		ghosts0 := _scrut23.ghosts
		return ghosts0
	})()
}

func score(g game_state) uint {
	return (func() uint {
		_scrut24 := g
		score0 := _scrut24.score
		return score0
	})()
}

func lives(g game_state) uint {
	return (func() uint {
		_scrut25 := g
		lives0 := _scrut25.lives
		return lives0
	})()
}

func dots_left(g game_state) uint {
	return (func() uint {
		_scrut26 := g
		dots_left0 := _scrut26.dots_left
		return dots_left0
	})()
}

func power_timer(g game_state) uint {
	return (func() uint {
		_scrut27 := g
		power_timer0 := _scrut27.power_timer
		return power_timer0
	})()
}

func game_over(g game_state) bool {
	return (func() bool {
		_scrut28 := g
		game_over0 := _scrut28.game_over
		return game_over0
	})()
}

func game_won(g game_state) bool {
	return (func() bool {
		_scrut29 := g
		game_won0 := _scrut29.game_won
		return game_won0
	})()
}

var board_height uint = uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))

var board_width uint = uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))

var cell_size uint = uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))

var status_height uint = uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))

var win_width uint = mul(board_width, cell_size)

var win_height uint = add(mul(board_height, cell_size), status_height)

var tick_ms uint = uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))

var frame_ms uint = uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))

var W cell = Wall

var E cell = Empty

var D cell = Dot

var P cell = PowerPellet

var initial_board list[list[cell]] = cons(cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, nil0))))))))))))))))))), cons(cons(W, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(W, nil0))))))))))))))))))), cons(cons(W, cons(D, cons(W, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(W, cons(D, cons(W, nil0))))))))))))))))))), cons(cons(W, cons(P, cons(W, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(W, cons(P, cons(W, nil0))))))))))))))))))), cons(cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, nil0))))))))))))))))))), cons(cons(W, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(W, nil0))))))))))))))))))), cons(cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, nil0))))))))))))))))))), cons(cons(W, cons(D, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(D, cons(W, nil0))))))))))))))))))), cons(cons(W, cons(D, cons(W, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(W, cons(D, cons(W, nil0))))))))))))))))))), cons(cons(W, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(W, nil0))))))))))))))))))), cons(cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, nil0))))))))))))))))))), cons(cons(W, cons(P, cons(W, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(W, cons(P, cons(W, nil0))))))))))))))))))), cons(cons(W, cons(D, cons(W, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(W, cons(D, cons(W, nil0))))))))))))))))))), cons(cons(W, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(W, nil0))))))))))))))))))), cons(cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, nil0))))))))))))))))))), nil0)))))))))))))))

func count_row(r list[cell]) uint {
	return (func() uint {
		var _box any = r
		_scrut30 := _box.(*listImpl)
		switch _scrut30._v {
		case 0:
			return uint(0)
		case 1:
			c := _scrut30._c1_d0
			rest := _scrut30._c1_d1
			return (func() uint {
				_scrut31 := c
				switch _scrut31 {
				case Wall:
					return count_row((rest).(list[cell]))
				case Empty:
					return count_row((rest).(list[cell]))
				default:
					return add(uint((uint(0) + 1)), count_row((rest).(list[cell])))
				}
				panic("unreachable")
			})()
		}
		panic("unreachable")
	})()
}

func count_dots(b list[list[cell]]) uint {
	return (func() uint {
		var _box any = b
		_scrut32 := _box.(*listImpl)
		switch _scrut32._v {
		case 0:
			return uint(0)
		case 1:
			row := _scrut32._c1_d0
			rest := _scrut32._c1_d1
			return add(count_row((row).(list[cell])), count_dots((rest).(list[list[cell]])))
		}
		panic("unreachable")
	})()
}

func get_cell(row uint, col uint, b list[list[cell]]) cell {
	return (nth(col, (nth(row, b, nil0)).(list[any]), Wall)).(cell)
}

func replace_nth[T1 any](n uint, l list[T1], x T1) list[T1] {
	return (func() list[T1] {
		var _box any = l
		_scrut33 := _box.(*listImpl)
		switch _scrut33._v {
		case 0:
			return nil0
		case 1:
			h := _scrut33._c1_d0
			t := _scrut33._c1_d1
			return (func() any {
				var _sni_ any = n
				_su_ := _sni_.(uint)
				if _su_ == 0 {
					return cons(x, t)
				} else {
					n_ := _su_ - 1
					return cons(h, replace_nth(n_, (t).(list[any]), x))
				}
			}()).(list[T1])
		}
		panic("unreachable")
	})()
}

func set_cell(row uint, col uint, c cell, b list[list[cell]]) list[list[cell]] {
	r := (nth(row, b, nil0)).(list[cell])
	return replace_nth(row, b, replace_nth(col, r, c))
}

func is_wall(row uint, col uint, b list[list[cell]]) bool {
	return (func() bool {
		_scrut34 := get_cell(row, col, b)
		switch _scrut34 {
		case Wall:
			return true
		default:
			return false
		}
		panic("unreachable")
	})()
}

func move_pos(d direction, p position) position {
	return (func() position {
		_scrut35 := d
		switch _scrut35 {
		case Up:
			return mkPos(func() uint {
				if (func() uint {
					_scrut36 := p
					prow0 := _scrut36.prow
					return prow0
				})() > 0 {
					return (func() uint {
						_scrut36 := p
						prow0 := _scrut36.prow
						return prow0
					})() - 1
				}
				return 0
			}(), (func() uint {
				_scrut37 := p
				pcol0 := _scrut37.pcol
				return pcol0
			})())
		case Down:
			return mkPos(uint(((func() uint {
				_scrut38 := p
				prow0 := _scrut38.prow
				return prow0
			})() + 1)), (func() uint {
				_scrut39 := p
				pcol0 := _scrut39.pcol
				return pcol0
			})())
		case Left:
			return mkPos((func() uint {
				_scrut40 := p
				prow0 := _scrut40.prow
				return prow0
			})(), func() uint {
				if (func() uint {
					_scrut41 := p
					pcol0 := _scrut41.pcol
					return pcol0
				})() > 0 {
					return (func() uint {
						_scrut41 := p
						pcol0 := _scrut41.pcol
						return pcol0
					})() - 1
				}
				return 0
			}())
		case Right:
			return mkPos((func() uint {
				_scrut42 := p
				prow0 := _scrut42.prow
				return prow0
			})(), uint(((func() uint {
				_scrut43 := p
				pcol0 := _scrut43.pcol
				return pcol0
			})() + 1)))
		case DirNone:
			return p
		}
		panic("unreachable")
	})()
}

func can_move(d direction, p position, b list[list[cell]]) bool {
	new_p := move_pos(d, p)
	return !(is_wall((func() uint {
		_scrut44 := new_p
		prow0 := _scrut44.prow
		return prow0
	})(), (func() uint {
		_scrut45 := new_p
		pcol0 := _scrut45.pcol
		return pcol0
	})(), b))
}

func update_ghost_modes(gs list[ghost_state], powered bool) list[ghost_state] {
	return (func() list[ghost_state] {
		var _box any = gs
		_scrut46 := _box.(*listImpl)
		switch _scrut46._v {
		case 0:
			return nil0
		case 1:
			g := _scrut46._c1_d0
			rest := _scrut46._c1_d1
			mode := (func() any {
				if powered {
					return Frightened
				} else {
					return Chase
				}
			}()).(ghost_mode)
			return cons(mkGhost((func() position {
				var _rbox any = g
				_scrut47 := _rbox.(ghost_state)
				gpos0 := _scrut47.gpos
				return gpos0
			})(), (func() direction {
				var _rbox any = g
				_scrut48 := _rbox.(ghost_state)
				gdir0 := _scrut48.gdir
				return gdir0
			})(), mode), update_ghost_modes((rest).(list[ghost_state]), powered))
		}
		panic("unreachable")
	})()
}

func set_direction(d direction, gs game_state) game_state {
	return mkState((func() list[list[cell]] {
		_scrut49 := gs
		board0 := _scrut49.board
		return board0
	})(), (func() position {
		_scrut50 := gs
		pacpos0 := _scrut50.pacpos
		return pacpos0
	})(), (func() direction {
		_scrut51 := gs
		pacdir0 := _scrut51.pacdir
		return pacdir0
	})(), d, (func() list[ghost_state] {
		_scrut52 := gs
		ghosts0 := _scrut52.ghosts
		return ghosts0
	})(), (func() uint {
		_scrut53 := gs
		score0 := _scrut53.score
		return score0
	})(), (func() uint {
		_scrut54 := gs
		lives0 := _scrut54.lives
		return lives0
	})(), (func() uint {
		_scrut55 := gs
		dots_left0 := _scrut55.dots_left
		return dots_left0
	})(), (func() uint {
		_scrut56 := gs
		power_timer0 := _scrut56.power_timer
		return power_timer0
	})(), (func() bool {
		_scrut57 := gs
		game_over0 := _scrut57.game_over
		return game_over0
	})(), (func() bool {
		_scrut58 := gs
		game_won0 := _scrut58.game_won
		return game_won0
	})())
}

func move_pacman(gs game_state) game_state {
	return (func() any {
		if (func() bool {
			_scrut59 := gs
			game_over0 := _scrut59.game_over
			return game_over0
		})() || (func() bool {
			_scrut60 := gs
			game_won0 := _scrut60.game_won
			return game_won0
		})() {
			return gs
		} else {
			return (func() game_state {
				try_dir0 := (func() any {
					if can_move((func() direction {
						_scrut61 := gs
						desired_dir0 := _scrut61.desired_dir
						return desired_dir0
					})(), (func() position {
						_scrut62 := gs
						pacpos0 := _scrut62.pacpos
						return pacpos0
					})(), (func() list[list[cell]] {
						_scrut63 := gs
						board0 := _scrut63.board
						return board0
					})()) {
						return (func() direction {
							_scrut64 := gs
							desired_dir0 := _scrut64.desired_dir
							return desired_dir0
						})()
					} else {
						return (func() direction {
							_scrut65 := gs
							pacdir0 := _scrut65.pacdir
							return pacdir0
						})()
					}
				}()).(direction)
				return (func() game_state {
					new_pos := move_pos(try_dir0, (func() position {
						_scrut66 := gs
						pacpos0 := _scrut66.pacpos
						return pacpos0
					})())
					return (func() any {
						if is_wall((func() uint {
							_scrut67 := new_pos
							prow0 := _scrut67.prow
							return prow0
						})(), (func() uint {
							_scrut68 := new_pos
							pcol0 := _scrut68.pcol
							return pcol0
						})(), (func() list[list[cell]] {
							_scrut69 := gs
							board0 := _scrut69.board
							return board0
						})()) {
							return gs
						} else {
							return (func() game_state {
								cell0 := get_cell((func() uint {
									_scrut70 := new_pos
									prow0 := _scrut70.prow
									return prow0
								})(), (func() uint {
									_scrut71 := new_pos
									pcol0 := _scrut71.pcol
									return pcol0
								})(), (func() list[list[cell]] {
									_scrut72 := gs
									board0 := _scrut72.board
									return board0
								})())
								return (func() game_state {
									new_board := (func() list[list[cell]] {
										_scrut73 := cell0
										switch _scrut73 {
										case Wall:
											return (func() list[list[cell]] {
												_scrut74 := gs
												board0 := _scrut74.board
												return board0
											})()
										case Empty:
											return (func() list[list[cell]] {
												_scrut75 := gs
												board0 := _scrut75.board
												return board0
											})()
										default:
											return set_cell((func() uint {
												_scrut76 := new_pos
												prow0 := _scrut76.prow
												return prow0
											})(), (func() uint {
												_scrut77 := new_pos
												pcol0 := _scrut77.pcol
												return pcol0
											})(), Empty, (func() list[list[cell]] {
												_scrut78 := gs
												board0 := _scrut78.board
												return board0
											})())
										}
										panic("unreachable")
									})()
									return (func() game_state {
										add_score := (func() uint {
											_scrut79 := cell0
											switch _scrut79 {
											case Dot:
												return uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))
											case PowerPellet:
												return uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))
											default:
												return uint(0)
											}
											panic("unreachable")
										})()
										return (func() game_state {
											new_dots := (func() uint {
												_scrut80 := cell0
												switch _scrut80 {
												case Wall:
													return (func() uint {
														_scrut81 := gs
														dots_left0 := _scrut81.dots_left
														return dots_left0
													})()
												case Empty:
													return (func() uint {
														_scrut82 := gs
														dots_left0 := _scrut82.dots_left
														return dots_left0
													})()
												default:
													return func() uint {
														if (func() uint {
															_scrut83 := gs
															dots_left0 := _scrut83.dots_left
															return dots_left0
														})() > 0 {
															return (func() uint {
																_scrut83 := gs
																dots_left0 := _scrut83.dots_left
																return dots_left0
															})() - 1
														}
														return 0
													}()
												}
												panic("unreachable")
											})()
											return (func() game_state {
												new_power := (func() uint {
													_scrut84 := cell0
													switch _scrut84 {
													case PowerPellet:
														return uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))
													default:
														return (func() uint {
															_scrut85 := gs
															power_timer0 := _scrut85.power_timer
															return power_timer0
														})()
													}
													panic("unreachable")
												})()
												return (func() game_state {
													new_ghosts := (func() list[ghost_state] {
														_scrut86 := cell0
														switch _scrut86 {
														case PowerPellet:
															return update_ghost_modes((func() list[ghost_state] {
																_scrut87 := gs
																ghosts0 := _scrut87.ghosts
																return ghosts0
															})(), true)
														default:
															return (func() list[ghost_state] {
																_scrut88 := gs
																ghosts0 := _scrut88.ghosts
																return ghosts0
															})()
														}
														panic("unreachable")
													})()
													return (func() game_state {
														won := (new_dots == uint(0))
														return mkState(new_board, new_pos, try_dir0, (func() direction {
															_scrut89 := gs
															desired_dir0 := _scrut89.desired_dir
															return desired_dir0
														})(), new_ghosts, add((func() uint {
															_scrut90 := gs
															score0 := _scrut90.score
															return score0
														})(), add_score), (func() uint {
															_scrut91 := gs
															lives0 := _scrut91.lives
															return lives0
														})(), new_dots, new_power, won, won)
													})()
												})()
											})()
										})()
									})()
								})()
							})()
						}
					}()).(game_state)
				})()
			})()
		}
	}()).(game_state)
}

func abs_diff(a uint, b uint) uint {
	return (func() any {
		if a <= b {
			return sub(b, a)
		} else {
			return sub(a, b)
		}
	}()).(uint)
}

func manhattan(p1 position, p2 position) uint {
	return add(abs_diff((func() uint {
		_scrut92 := p1
		prow0 := _scrut92.prow
		return prow0
	})(), (func() uint {
		_scrut93 := p2
		prow0 := _scrut93.prow
		return prow0
	})()), abs_diff((func() uint {
		_scrut94 := p1
		pcol0 := _scrut94.pcol
		return pcol0
	})(), (func() uint {
		_scrut95 := p2
		pcol0 := _scrut95.pcol
		return pcol0
	})()))
}

func is_opposite(d1 direction, d2 direction) bool {
	return (func() bool {
		_scrut96 := d1
		switch _scrut96 {
		case Up:
			return (func() bool {
				_scrut97 := d2
				switch _scrut97 {
				case Down:
					return true
				default:
					return false
				}
				panic("unreachable")
			})()
		case Down:
			return (func() bool {
				_scrut98 := d2
				switch _scrut98 {
				case Up:
					return true
				default:
					return false
				}
				panic("unreachable")
			})()
		case Left:
			return (func() bool {
				_scrut99 := d2
				switch _scrut99 {
				case Right:
					return true
				default:
					return false
				}
				panic("unreachable")
			})()
		case Right:
			return (func() bool {
				_scrut100 := d2
				switch _scrut100 {
				case Left:
					return true
				default:
					return false
				}
				panic("unreachable")
			})()
		case DirNone:
			return false
		}
		panic("unreachable")
	})()
}

func try_dir(d direction, g ghost_state, target position, b list[list[cell]], best_d direction, best_dist uint) struct {
	fst direction
	snd uint
} {
	return (func() any {
		if !(can_move(d, (func() position {
			_scrut101 := g
			gpos0 := _scrut101.gpos
			return gpos0
		})(), b)) {
			return struct {
				fst direction
				snd uint
			}{fst: best_d, snd: best_dist}
		} else {
			return (func() any {
				if is_opposite(d, (func() direction {
					_scrut102 := g
					gdir0 := _scrut102.gdir
					return gdir0
				})()) {
					return struct {
						fst direction
						snd uint
					}{fst: best_d, snd: best_dist}
				} else {
					return (func() struct {
						fst direction
						snd uint
					} {
						new_pos := move_pos(d, (func() position {
							_scrut103 := g
							gpos0 := _scrut103.gpos
							return gpos0
						})())
						return (func() struct {
							fst direction
							snd uint
						} {
							dist := manhattan(new_pos, target)
							return (func() any {
								if dist < best_dist {
									return struct {
										fst direction
										snd uint
									}{fst: d, snd: dist}
								} else {
									return struct {
										fst direction
										snd uint
									}{fst: best_d, snd: best_dist}
								}
							}()).(struct {
								fst direction
								snd uint
							})
						})()
					})()
				}
			}()).(struct {
				fst direction
				snd uint
			})
		}
	}()).(struct {
		fst direction
		snd uint
	})
}

func choose_ghost_dir(g ghost_state, target position, b list[list[cell]]) direction {
	return (func() any {
		d1 := try_dir(Up, g, target, b, DirNone, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))).fst
		dist1 := try_dir(Up, g, target, b, DirNone, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))).snd
		return (func() any {
			d2 := try_dir(Down, g, target, b, d1, dist1).fst
			dist2 := try_dir(Down, g, target, b, d1, dist1).snd
			return (func() any {
				d3 := try_dir(Left, g, target, b, d2, dist2).fst
				dist3 := try_dir(Left, g, target, b, d2, dist2).snd
				return (func() any {
					d4 := try_dir(Right, g, target, b, d3, dist3).fst
					_ = try_dir(Right, g, target, b, d3, dist3).snd
					return d4
				}()).(direction)
			}()).(direction)
		}()).(direction)
	}()).(direction)
}

func move_one_ghost(g ghost_state, pac position, b list[list[cell]]) ghost_state {
	target := (func() position {
		_scrut104 := (func() ghost_mode {
			_scrut105 := g
			gmode0 := _scrut105.gmode
			return gmode0
		})()
		switch _scrut104 {
		case Chase:
			return pac
		case Frightened:
			return mkPos(uint(0), uint(0))
		}
		panic("unreachable")
	})()
	dir := choose_ghost_dir(g, target, b)
	return (func() any {
		if can_move(dir, (func() position {
			_scrut106 := g
			gpos0 := _scrut106.gpos
			return gpos0
		})(), b) {
			return mkGhost(move_pos(dir, (func() position {
				_scrut107 := g
				gpos0 := _scrut107.gpos
				return gpos0
			})()), dir, (func() ghost_mode {
				_scrut108 := g
				gmode0 := _scrut108.gmode
				return gmode0
			})())
		} else {
			return g
		}
	}()).(ghost_state)
}

func move_ghosts_list(gs list[ghost_state], pac position, b list[list[cell]]) list[ghost_state] {
	return (func() list[ghost_state] {
		var _box any = gs
		_scrut109 := _box.(*listImpl)
		switch _scrut109._v {
		case 0:
			return nil0
		case 1:
			g := _scrut109._c1_d0
			rest := _scrut109._c1_d1
			return cons(move_one_ghost((g).(ghost_state), pac, b), move_ghosts_list((rest).(list[ghost_state]), pac, b))
		}
		panic("unreachable")
	})()
}

func move_ghosts(gs game_state) game_state {
	return mkState((func() list[list[cell]] {
		_scrut110 := gs
		board0 := _scrut110.board
		return board0
	})(), (func() position {
		_scrut111 := gs
		pacpos0 := _scrut111.pacpos
		return pacpos0
	})(), (func() direction {
		_scrut112 := gs
		pacdir0 := _scrut112.pacdir
		return pacdir0
	})(), (func() direction {
		_scrut113 := gs
		desired_dir0 := _scrut113.desired_dir
		return desired_dir0
	})(), move_ghosts_list((func() list[ghost_state] {
		_scrut114 := gs
		ghosts0 := _scrut114.ghosts
		return ghosts0
	})(), (func() position {
		_scrut115 := gs
		pacpos0 := _scrut115.pacpos
		return pacpos0
	})(), (func() list[list[cell]] {
		_scrut116 := gs
		board0 := _scrut116.board
		return board0
	})()), (func() uint {
		_scrut117 := gs
		score0 := _scrut117.score
		return score0
	})(), (func() uint {
		_scrut118 := gs
		lives0 := _scrut118.lives
		return lives0
	})(), (func() uint {
		_scrut119 := gs
		dots_left0 := _scrut119.dots_left
		return dots_left0
	})(), (func() uint {
		_scrut120 := gs
		power_timer0 := _scrut120.power_timer
		return power_timer0
	})(), (func() bool {
		_scrut121 := gs
		game_over0 := _scrut121.game_over
		return game_over0
	})(), (func() bool {
		_scrut122 := gs
		game_won0 := _scrut122.game_won
		return game_won0
	})())
}

func ghost_at_pos(row uint, col uint, gs list[ghost_state]) *ghost_state {
	return (func() *ghost_state {
		var _box any = gs
		_scrut123 := _box.(*listImpl)
		switch _scrut123._v {
		case 0:
			return nil
		case 1:
			g := _scrut123._c1_d0
			rest := _scrut123._c1_d1
			return (func() any {
				if ((func() uint {
					_scrut124 := (func() position {
						var _rbox any = g
						_scrut125 := _rbox.(ghost_state)
						gpos0 := _scrut125.gpos
						return gpos0
					})()
					prow0 := _scrut124.prow
					return prow0
				})() == row) && ((func() uint {
					_scrut126 := (func() position {
						var _rbox any = g
						_scrut127 := _rbox.(ghost_state)
						gpos0 := _scrut127.gpos
						return gpos0
					})()
					pcol0 := _scrut126.pcol
					return pcol0
				})() == col) {
					return rocqSome(g)
				} else {
					return ghost_at_pos(row, col, (rest).(list[ghost_state]))
				}
			}()).(*ghost_state)
		}
		panic("unreachable")
	})()
}

func respawn_ghost(_ ghost_state) ghost_state {
	return mkGhost(mkPos(uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))), DirNone, Chase)
}

func check_collisions(gs game_state) game_state {
	return (func() any {
		if ghost_at_pos((func() uint {
			_scrut128 := (func() position {
				_scrut129 := gs
				pacpos0 := _scrut129.pacpos
				return pacpos0
			})()
			prow0 := _scrut128.prow
			return prow0
		})(), (func() uint {
			_scrut130 := (func() position {
				_scrut131 := gs
				pacpos0 := _scrut131.pacpos
				return pacpos0
			})()
			pcol0 := _scrut130.pcol
			return pcol0
		})(), (func() list[ghost_state] {
			_scrut132 := gs
			ghosts0 := _scrut132.ghosts
			return ghosts0
		})()) != nil {
			g := ghost_at_pos((func() uint {
				_scrut128 := (func() position {
					_scrut129 := gs
					pacpos0 := _scrut129.pacpos
					return pacpos0
				})()
				prow0 := _scrut128.prow
				return prow0
			})(), (func() uint {
				_scrut130 := (func() position {
					_scrut131 := gs
					pacpos0 := _scrut131.pacpos
					return pacpos0
				})()
				pcol0 := _scrut130.pcol
				return pcol0
			})(), (func() list[ghost_state] {
				_scrut132 := gs
				ghosts0 := _scrut132.ghosts
				return ghosts0
			})())
			return (func() game_state {
				_scrut133 := (func() ghost_mode {
					_scrut134 := g
					gmode0 := _scrut134.gmode
					return gmode0
				})()
				switch _scrut133 {
				case Chase:
					new_lives := func() uint {
						if (func() uint {
							_scrut135 := gs
							lives0 := _scrut135.lives
							return lives0
						})() > 0 {
							return (func() uint {
								_scrut135 := gs
								lives0 := _scrut135.lives
								return lives0
							})() - 1
						}
						return 0
					}()
					dead := (new_lives == uint(0))
					return mkState((func() list[list[cell]] {
						_scrut136 := gs
						board0 := _scrut136.board
						return board0
					})(), mkPos(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))), DirNone, DirNone, (func() list[ghost_state] {
						_scrut137 := gs
						ghosts0 := _scrut137.ghosts
						return ghosts0
					})(), (func() uint {
						_scrut138 := gs
						score0 := _scrut138.score
						return score0
					})(), new_lives, (func() uint {
						_scrut139 := gs
						dots_left0 := _scrut139.dots_left
						return dots_left0
					})(), uint(0), dead, (func() bool {
						_scrut140 := gs
						game_won0 := _scrut140.game_won
						return game_won0
					})())
				case Frightened:
					new_ghosts := map0(func(g_ any) any {
						return func() any {
							if ((func() uint {
								_scrut141 := (func() position {
									var _rbox any = g_
									_scrut142 := _rbox.(ghost_state)
									gpos0 := _scrut142.gpos
									return gpos0
								})()
								prow0 := _scrut141.prow
								return prow0
							})() == (func() uint {
								_scrut143 := (func() position {
									_scrut144 := gs
									pacpos0 := _scrut144.pacpos
									return pacpos0
								})()
								prow0 := _scrut143.prow
								return prow0
							})()) && ((func() uint {
								_scrut145 := (func() position {
									var _rbox any = g_
									_scrut146 := _rbox.(ghost_state)
									gpos0 := _scrut146.gpos
									return gpos0
								})()
								pcol0 := _scrut145.pcol
								return pcol0
							})() == (func() uint {
								_scrut147 := (func() position {
									_scrut148 := gs
									pacpos0 := _scrut148.pacpos
									return pacpos0
								})()
								pcol0 := _scrut147.pcol
								return pcol0
							})()) {
								return respawn_ghost((g_).(ghost_state))
							} else {
								return g_
							}
						}()
					}, (func() list[any] {
						_scrut149 := gs
						ghosts0 := _scrut149.ghosts
						return ghosts0
					})())
					return mkState((func() list[list[cell]] {
						_scrut150 := gs
						board0 := _scrut150.board
						return board0
					})(), (func() position {
						_scrut151 := gs
						pacpos0 := _scrut151.pacpos
						return pacpos0
					})(), (func() direction {
						_scrut152 := gs
						pacdir0 := _scrut152.pacdir
						return pacdir0
					})(), (func() direction {
						_scrut153 := gs
						desired_dir0 := _scrut153.desired_dir
						return desired_dir0
					})(), new_ghosts, add((func() uint {
						_scrut154 := gs
						score0 := _scrut154.score
						return score0
					})(), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))), (func() uint {
						_scrut155 := gs
						lives0 := _scrut155.lives
						return lives0
					})(), (func() uint {
						_scrut156 := gs
						dots_left0 := _scrut156.dots_left
						return dots_left0
					})(), (func() uint {
						_scrut157 := gs
						power_timer0 := _scrut157.power_timer
						return power_timer0
					})(), (func() bool {
						_scrut158 := gs
						game_over0 := _scrut158.game_over
						return game_over0
					})(), (func() bool {
						_scrut159 := gs
						game_won0 := _scrut159.game_won
						return game_won0
					})())
				}
				panic("unreachable")
			})()
		} else {
			return gs
		}
	}()).(game_state)
}

func tick_power(gs game_state) game_state {
	return (func() any {
		var _sni_ any = (func() uint {
			_scrut160 := gs
			power_timer0 := _scrut160.power_timer
			return power_timer0
		})()
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return gs
		} else {
			n := _su_ - 1
			return (func() game_state {
				powered := !(n == uint(0))
				return (func() game_state {
					new_ghosts := update_ghost_modes((func() list[ghost_state] {
						_scrut161 := gs
						ghosts0 := _scrut161.ghosts
						return ghosts0
					})(), powered)
					return mkState((func() list[list[cell]] {
						_scrut162 := gs
						board0 := _scrut162.board
						return board0
					})(), (func() position {
						_scrut163 := gs
						pacpos0 := _scrut163.pacpos
						return pacpos0
					})(), (func() direction {
						_scrut164 := gs
						pacdir0 := _scrut164.pacdir
						return pacdir0
					})(), (func() direction {
						_scrut165 := gs
						desired_dir0 := _scrut165.desired_dir
						return desired_dir0
					})(), new_ghosts, (func() uint {
						_scrut166 := gs
						score0 := _scrut166.score
						return score0
					})(), (func() uint {
						_scrut167 := gs
						lives0 := _scrut167.lives
						return lives0
					})(), (func() uint {
						_scrut168 := gs
						dots_left0 := _scrut168.dots_left
						return dots_left0
					})(), n, (func() bool {
						_scrut169 := gs
						game_over0 := _scrut169.game_over
						return game_over0
					})(), (func() bool {
						_scrut170 := gs
						game_won0 := _scrut170.game_won
						return game_won0
					})())
				})()
			})()
		}
	}()).(game_state)
}

func tick(gs game_state) game_state {
	return (func() any {
		if (func() bool {
			_scrut171 := gs
			game_over0 := _scrut171.game_over
			return game_over0
		})() || (func() bool {
			_scrut172 := gs
			game_won0 := _scrut172.game_won
			return game_won0
		})() {
			return gs
		} else {
			return (func() game_state {
				gs1 := move_pacman(gs)
				return (func() game_state {
					gs2 := move_ghosts(gs1)
					return tick_power(gs2)
				})()
			})()
		}
	}()).(game_state)
}

var initial_ghosts list[ghost_state] = cons(mkGhost(mkPos(uint((uint(0)+1)), uint((uint(0)+1))), DirNone, Chase), cons(mkGhost(mkPos(uint((uint(0)+1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))), DirNone, Chase), cons(mkGhost(mkPos(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), uint((uint(0)+1))), DirNone, Chase), cons(mkGhost(mkPos(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))), DirNone, Chase), nil0))))

var initial_state game_state = (func() game_state {
	b := set_cell(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), Empty, initial_board)
	return mkState(b, mkPos(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))), DirNone, DirNone, initial_ghosts, uint(0), uint((uint((uint((uint(0) + 1)) + 1)) + 1)), count_dots(b), uint(0), false, false)
})()

func filled_circle_rows(ren RocqSDLRenderer, cx uint, base_y uint, radius uint, i uint, count uint) struct{} {
	return (func() any {
		var _sni_ any = count
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return struct{}{}
		} else {
			count_ := _su_ - 1
			return (func() struct{} {
				dist := abs_diff(i, radius)
				return (func() struct{} {
					d2 := mul(dist, dist)
					return (func() struct{} {
						r2 := mul(radius, radius)
						return (func() struct{} {
							_ = func() any {
								if d2 <= r2 {
									return (func() any {
										half := uint(math.Sqrt(float64(sub(r2, d2))))
										return rocqSDLFillRect(ren, sub(cx, half), add(base_y, i), add(add(half, half), uint((uint(0)+1))), uint((uint(0) + 1)))
									})()
								} else {
									return struct{}{}
								}
							}()
							return filled_circle_rows(ren, cx, base_y, radius, uint((i + 1)), count_)
						})()
					})()
				})()
			})()
		}
	}()).(struct{})
}

func draw_filled_circle(ren RocqSDLRenderer, cx uint, cy uint, radius uint) struct{} {
	return filled_circle_rows(ren, cx, sub(cy, radius), radius, uint(0), add(add(radius, radius), uint((uint(0)+1))))
}

func semicircle_top_rows(ren RocqSDLRenderer, cx uint, base_y uint, radius uint, i uint, count uint) struct{} {
	return (func() any {
		var _sni_ any = count
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return struct{}{}
		} else {
			count_ := _su_ - 1
			return (func() struct{} {
				dist := sub(radius, i)
				return (func() struct{} {
					half := uint(math.Sqrt(float64(sub(mul(radius, radius), mul(dist, dist)))))
					return (func() struct{} {
						_ = rocqSDLFillRect(ren, sub(cx, half), add(base_y, i), add(add(half, half), uint((uint(0)+1))), uint((uint(0) + 1)))
						return semicircle_top_rows(ren, cx, base_y, radius, uint((i + 1)), count_)
					})()
				})()
			})()
		}
	}()).(struct{})
}

func draw_top_semicircle(ren RocqSDLRenderer, cx uint, cy uint, radius uint) struct{} {
	return semicircle_top_rows(ren, cx, sub(cy, radius), radius, uint(0), add(radius, uint((uint(0)+1))))
}

func ghost_body_color(color_idx uint) struct {
	fst struct {
		fst uint
		snd uint
	}
	snd uint
} {
	return (func() any {
		var _sni_ any = color_idx
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return struct {
				fst struct {
					fst uint
					snd uint
				}
				snd uint
			}{fst: struct {
				fst uint
				snd uint
			}{fst: uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), snd: uint(0)}, snd: uint(0)}
		} else {
			n := _su_ - 1
			return (func() any {
				var _sni_ any = n
				_su_ := _sni_.(uint)
				if _su_ == 0 {
					return struct {
						fst struct {
							fst uint
							snd uint
						}
						snd uint
					}{fst: struct {
						fst uint
						snd uint
					}{fst: uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), snd: uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))}, snd: uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))}
				} else {
					n0 := _su_ - 1
					return (func() any {
						var _sni_ any = n0
						_su_ := _sni_.(uint)
						if _su_ == 0 {
							return struct {
								fst struct {
									fst uint
									snd uint
								}
								snd uint
							}{fst: struct {
								fst uint
								snd uint
							}{fst: uint(0), snd: uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))}, snd: uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))}
						} else {
							n1 := _su_ - 1
							return (func() any {
								var _sni_ any = n1
								_su_ := _sni_.(uint)
								if _su_ == 0 {
									return struct {
										fst struct {
											fst uint
											snd uint
										}
										snd uint
									}{fst: struct {
										fst uint
										snd uint
									}{fst: uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), snd: uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))}, snd: uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))}
								} else {
									_ = _su_ - 1
									return struct {
										fst struct {
											fst uint
											snd uint
										}
										snd uint
									}{fst: struct {
										fst uint
										snd uint
									}{fst: uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), snd: uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))}, snd: uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))}
								}
							}()).(struct {
								fst struct {
									fst uint
									snd uint
								}
								snd uint
							})
						}
					}()).(struct {
						fst struct {
							fst uint
							snd uint
						}
						snd uint
					})
				}
			}()).(struct {
				fst struct {
					fst uint
					snd uint
				}
				snd uint
			})
		}
	}()).(struct {
		fst struct {
			fst uint
			snd uint
		}
		snd uint
	})
}

func draw_ghost_body(ren RocqSDLRenderer, cx uint, cy uint, radius uint, cr uint, cg uint, cb uint) struct{} {
	_ = rocqSDLSetDrawColor(ren, cr, cg, cb)
	_ = draw_top_semicircle(ren, cx, cy, radius)
	_ = rocqSDLFillRect(ren, sub(cx, radius), cy, add(add(radius, radius), uint((uint(0)+1))), radius)
	seg := func() uint {
		if uint((uint((uint((uint(0) + 1)) + 1)) + 1)) == 0 {
			return 0
		}
		return add(radius, radius) / uint((uint((uint((uint(0) + 1)) + 1)) + 1))
	}()
	sr := func() uint {
		if uint((uint((uint(0) + 1)) + 1)) == 0 {
			return 0
		}
		return seg / uint((uint((uint(0) + 1)) + 1))
	}()
	bx0 := add(sub(cx, radius), func() uint {
		if uint((uint((uint(0) + 1)) + 1)) == 0 {
			return 0
		}
		return seg / uint((uint((uint(0) + 1)) + 1))
	}())
	_ = draw_filled_circle(ren, bx0, add(cy, radius), sr)
	_ = draw_filled_circle(ren, add(bx0, seg), add(cy, radius), sr)
	return draw_filled_circle(ren, add(add(bx0, seg), seg), add(cy, radius), sr)
}

func draw_ghost_eyes(ren RocqSDLRenderer, cx uint, cy uint, radius uint) struct{} {
	eye_offset := func() uint {
		if uint((uint((uint((uint(0) + 1)) + 1)) + 1)) == 0 {
			return 0
		}
		return radius / uint((uint((uint((uint(0) + 1)) + 1)) + 1))
	}()
	_ = rocqSDLSetDrawColor(ren, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
	_ = draw_filled_circle(ren, sub(cx, eye_offset), sub(cy, uint((uint((uint(0)+1))+1))), uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)))
	_ = draw_filled_circle(ren, add(cx, eye_offset), sub(cy, uint((uint((uint(0)+1))+1))), uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)))
	_ = rocqSDLSetDrawColor(ren, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
	_ = draw_filled_circle(ren, add(sub(cx, eye_offset), uint((uint(0)+1))), sub(cy, uint((uint(0)+1))), uint((uint((uint(0) + 1)) + 1)))
	return draw_filled_circle(ren, add(add(cx, eye_offset), uint((uint(0)+1))), sub(cy, uint((uint(0)+1))), uint((uint((uint(0) + 1)) + 1)))
}

func draw_ghost_bottom(ren RocqSDLRenderer, cx uint, cy uint) struct{} {
	_ = rocqSDLSetDrawColor(ren, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, sub(cx, uint((uint(0)+1))), sub(cy, uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))), uint((uint((uint((uint(0) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
	return rocqSDLFillRect(ren, sub(cx, uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))), add(cy, uint((uint((uint(0)+1))+1))), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint(0) + 1)) + 1)) + 1)))
}

func draw_ghost_sprite(ren RocqSDLRenderer, px uint, py uint, color_idx uint) struct{} {
	radius := uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))
	return (func() any {
		p := ghost_body_color(color_idx).fst
		cb := ghost_body_color(color_idx).snd
		return (func() any {
			cr := p.fst
			cg := p.snd
			return (func() struct{} {
				_ = draw_ghost_body(ren, px, py, radius, cr, cg, cb)
				return draw_ghost_bottom(ren, px, py)
			})()
		}()).(struct{})
	}()).(struct{})
}

func pac_row_pixels(ren RocqSDLRenderer, sx uint, sy_row uint, radius uint, fdy float64, dir_ang float64, mouth float64, dy_sq uint, r2 uint, dx uint, count uint) struct{} {
	return (func() any {
		var _sni_ any = count
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return struct{}{}
		} else {
			count_ := _su_ - 1
			return (func() struct{} {
				dx_off := abs_diff(dx, radius)
				return (func() struct{} {
					dist_sq := add(mul(dx_off, dx_off), dy_sq)
					return (func() struct{} {
						_ = func() any {
							if dist_sq <= r2 {
								return (func() any {
									fdx := (float64(dx) - float64(radius))
									return (func() any {
										ang := real_atan2(fdy, fdx)
										return (func() any {
											rel0 := (ang - dir_ang)
											return (func() any {
												rel1 := (func() any {
													if math.Pi < rel0 {
														return (rel0 - (math.Pi + math.Pi))
													} else {
														return rel0
													}
												}()).(float64)
												return (func() any {
													rel := (func() any {
														if rel1 < (-math.Pi) {
															return (rel1 + (math.Pi + math.Pi))
														} else {
															return rel1
														}
													}()).(float64)
													return func() any {
														if math.Abs(rel) < mouth {
															return struct{}{}
														} else {
															return rocqSDLDrawPoint(ren, add(sx, dx), sy_row)
														}
													}()
												})()
											})()
										})()
									})()
								})()
							} else {
								return struct{}{}
							}
						}()
						return pac_row_pixels(ren, sx, sy_row, radius, fdy, dir_ang, mouth, dy_sq, r2, uint((dx + 1)), count_)
					})()
				})()
			})()
		}
	}()).(struct{})
}

func pac_rows(ren RocqSDLRenderer, sx uint, base_y uint, radius uint, dir_ang float64, mouth float64, r2 uint, diameter uint, dy uint, count uint) struct{} {
	return (func() any {
		var _sni_ any = count
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return struct{}{}
		} else {
			count_ := _su_ - 1
			return (func() struct{} {
				dy_off := abs_diff(dy, radius)
				return (func() struct{} {
					dy_sq := mul(dy_off, dy_off)
					return (func() struct{} {
						fdy := (float64(dy) - float64(radius))
						return (func() struct{} {
							_ = pac_row_pixels(ren, sx, add(base_y, dy), radius, fdy, dir_ang, mouth, dy_sq, r2, uint(0), diameter)
							return pac_rows(ren, sx, base_y, radius, dir_ang, mouth, r2, diameter, uint((dy + 1)), count_)
						})()
					})()
				})()
			})()
		}
	}()).(struct{})
}

func dir_angle(dir uint) float64 {
	return (func() any {
		var _sni_ any = dir
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return float64(uint(0))
		} else {
			n := _su_ - 1
			return (func() any {
				var _sni_ any = n
				_su_ := _sni_.(uint)
				if _su_ == 0 {
					return math.Pi
				} else {
					n0 := _su_ - 1
					return (func() any {
						var _sni_ any = n0
						_su_ := _sni_.(uint)
						if _su_ == 0 {
							return (-(math.Pi / float64(uint((uint((uint(0) + 1)) + 1)))))
						} else {
							n1 := _su_ - 1
							return (func() any {
								var _sni_ any = n1
								_su_ := _sni_.(uint)
								if _su_ == 0 {
									return (math.Pi / float64(uint((uint((uint(0) + 1)) + 1))))
								} else {
									_ = _su_ - 1
									return float64(uint(0))
								}
							}()).(float64)
						}
					}()).(float64)
				}
			}()).(float64)
		}
	}()).(float64)
}

func compute_mouth_angle(time_ms uint) float64 {
	ftime := float64(time_ms)
	phase0 := (ftime / float64(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))))
	arg := ((phase0 * float64(uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))) * math.Pi)
	s := math.Abs(math.Sin(arg))
	f015 := (float64(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))) / float64(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))))
	f035 := (float64(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))) / float64(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))))
	return (f015 + (f035 * s))
}

func draw_pacman_sprite(ren RocqSDLRenderer, px uint, py uint, dir uint, time_ms uint) struct{} {
	radius := uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))
	diameter := add(add(radius, radius), uint((uint(0) + 1)))
	r2 := mul(radius, radius)
	dir_a := dir_angle(dir)
	mouth := compute_mouth_angle(time_ms)
	_ = rocqSDLSetDrawColor(ren, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint(0))
	return pac_rows(ren, sub(px, radius), sub(py, radius), radius, dir_a, mouth, r2, diameter, uint(0), diameter)
}

func glyph_row_data(g uint, row uint) uint {
	return (nth(row, (nth(g, cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint(0)+1)), cons(uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint(0)+1)), cons(uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1)), cons(uint((uint(0)+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint(0)+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint(0)+1))+1)), cons(uint((uint((uint(0)+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint(0)+1)), cons(uint((uint(0)+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint(0)+1)), cons(uint((uint((uint(0)+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint(0)+1)), cons(uint((uint((uint(0)+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint(0), cons(uint(0), cons(uint(0), cons(uint(0), cons(uint(0), cons(uint(0), cons(uint(0), nil0))))))), cons(cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint(0)+1))+1)), cons(uint((uint((uint(0)+1))+1)), cons(uint((uint((uint(0)+1))+1)), cons(uint((uint((uint(0)+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint(0)+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), nil0))))))), cons(cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint(0)+1)), cons(uint((uint((uint(0)+1))+1)), cons(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))), nil0))))))))))))))))))))))))))))))))))))), nil0)).(list[any]), uint(0))).(uint)
}

func draw_glyph_row(ren RocqSDLRenderer, sx uint, sy uint, row_bits uint, dx uint, count uint, s uint) struct{} {
	return (func() any {
		var _sni_ any = count
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return struct{}{}
		} else {
			count_ := _su_ - 1
			return (func() struct{} {
				_ = func() any {
					if nat_testbit(row_bits, sub(uint((uint((uint((uint((uint(0)+1))+1))+1))+1)), dx)) {
						return rocqSDLFillRect(ren, add(sx, mul(dx, s)), sy, s, s)
					} else {
						return struct{}{}
					}
				}()
				return draw_glyph_row(ren, sx, sy, row_bits, uint((dx + 1)), count_, s)
			})()
		}
	}()).(struct{})
}

func draw_glyph_rows(ren RocqSDLRenderer, sx uint, sy uint, g uint, row uint, count uint, s uint) struct{} {
	return (func() any {
		var _sni_ any = count
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return struct{}{}
		} else {
			count_ := _su_ - 1
			return (func() struct{} {
				bits := glyph_row_data(g, row)
				return (func() struct{} {
					_ = draw_glyph_row(ren, sx, add(sy, mul(row, s)), bits, uint(0), uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)), s)
					return draw_glyph_rows(ren, sx, sy, g, uint((row + 1)), count_, s)
				})()
			})()
		}
	}()).(struct{})
}

func draw_one_glyph(ren RocqSDLRenderer, sx uint, sy uint, g uint, s uint) struct{} {
	return draw_glyph_rows(ren, sx, sy, g, uint(0), uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), s)
}

func draw_glyphs(ren RocqSDLRenderer, sx uint, sy uint, s uint, glyphs list[uint]) struct{} {
	return (func() struct{} {
		var _box any = glyphs
		_scrut173 := _box.(*listImpl)
		switch _scrut173._v {
		case 0:
			return struct{}{}
		case 1:
			g := _scrut173._c1_d0
			rest := _scrut173._c1_d1
			_ = draw_one_glyph(ren, sx, sy, (g).(uint), s)
			return draw_glyphs(ren, add(sx, mul(uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1)), s)), sy, s, (rest).(list[uint]))
		}
		panic("unreachable")
	})()
}

func draw_number_digits(ren RocqSDLRenderer, sx uint, sy uint, digits list[uint]) struct{} {
	return (func() struct{} {
		var _box any = digits
		_scrut174 := _box.(*listImpl)
		switch _scrut174._v {
		case 0:
			return struct{}{}
		case 1:
			d := _scrut174._c1_d0
			rest := _scrut174._c1_d1
			_ = draw_one_glyph(ren, sx, sy, (d).(uint), uint((uint((uint((uint(0) + 1)) + 1)) + 1)))
			return draw_number_digits(ren, add(sx, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))), sy, (rest).(list[uint]))
		}
		panic("unreachable")
	})()
}

func nat_digits_aux(n uint, acc list[uint], fuel uint) list[uint] {
	return (func() any {
		var _sni_ any = fuel
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return acc
		} else {
			fuel_ := _su_ - 1
			return (func() list[uint] {
				d := func() uint {
					if uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) == 0 {
						return n
					}
					return n % uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))
				}()
				return (func() list[uint] {
					rest := func() uint {
						if uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) == 0 {
							return 0
						}
						return n / uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))
					}()
					return (func() any {
						if rest == uint(0) {
							return cons(d, acc)
						} else {
							return nat_digits_aux(rest, cons(d, acc), fuel_)
						}
					}()).(list[uint])
				})()
			})()
		}
	}()).(list[uint])
}

func nat_digit_list(n uint) list[uint] {
	return (func() any {
		var _sni_ any = n
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return cons(uint(0), nil0)
		} else {
			_ = _su_ - 1
			return nat_digits_aux(n, nil0, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
		}
	}()).(list[uint])
}

func draw_number_sprite(ren RocqSDLRenderer, n uint, sx uint, sy uint) struct{} {
	_ = rocqSDLSetDrawColor(ren, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
	return draw_number_digits(ren, sx, sy, nat_digit_list(n))
}

var msg_game_over list[uint] = cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0)))))))))

var msg_you_win list[uint] = cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0)))))))

func msg_lives_left(n uint) list[uint] {
	return app(nat_digit_list(n), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))))))))
}

var msg_paused list[uint] = cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), cons(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), nil0))))))

func draw_message_screen(ren RocqSDLRenderer, msg list[uint]) struct{} {
	s := uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1))
	glyph_w := mul(uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), s)
	text_w := mul(length(msg), glyph_w)
	sx := func() uint {
		if uint((uint((uint(0) + 1)) + 1)) == 0 {
			return 0
		}
		return sub(win_width, text_w) / uint((uint((uint(0) + 1)) + 1))
	}()
	sy := sub(func() uint {
		if uint((uint((uint(0) + 1)) + 1)) == 0 {
			return 0
		}
		return win_height / uint((uint((uint(0) + 1)) + 1))
	}(), func() uint {
		if uint((uint((uint(0) + 1)) + 1)) == 0 {
			return 0
		}
		return mul(uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1)), s) / uint((uint((uint(0) + 1)) + 1))
	}())
	_ = rocqSDLSetDrawColor(ren, uint(0), uint(0), uint(0))
	_ = rocqSDLClear(ren)
	_ = rocqSDLSetDrawColor(ren, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
	_ = draw_glyphs(ren, sx, sy, s, msg)
	return rocqSDLPresent(ren)
}

func draw_status_message(ren RocqSDLRenderer, msg list[uint]) struct{} {
	s := uint((uint((uint((uint(0) + 1)) + 1)) + 1))
	glyph_w := mul(uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), s)
	text_w := mul(length(msg), glyph_w)
	sx := func() uint {
		if uint((uint((uint(0) + 1)) + 1)) == 0 {
			return 0
		}
		return sub(win_width, text_w) / uint((uint((uint(0) + 1)) + 1))
	}()
	sy := add(mul(board_height, cell_size), func() uint {
		if uint((uint((uint(0) + 1)) + 1)) == 0 {
			return 0
		}
		return sub(status_height, mul(uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1)), s)) / uint((uint((uint(0) + 1)) + 1))
	}())
	_ = rocqSDLSetDrawColor(ren, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
	return draw_glyphs(ren, sx, sy, s, msg)
}

func draw_life_icon(ren RocqSDLRenderer, x uint, y uint) struct{} {
	_ = rocqSDLSetDrawColor(ren, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, sub(x, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))), sub(y, uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))), uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, add(x, uint((uint((uint((uint(0)+1))+1))+1))), sub(y, uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))), uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, sub(x, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))), y, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, sub(x, uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))), add(y, uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, sub(x, uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))), add(y, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, sub(x, uint((uint((uint((uint(0)+1))+1))+1))), add(y, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))), uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint(0) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, sub(x, uint((uint(0)+1))), add(y, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))), uint((uint((uint(0) + 1)) + 1)), uint((uint((uint(0) + 1)) + 1)))
	_ = rocqSDLSetDrawColor(ren, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, sub(x, uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))), sub(y, uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))), uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, add(x, uint((uint((uint((uint(0)+1))+1))+1))), sub(y, uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))), uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, sub(x, uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))), sub(y, uint((uint(0)+1))), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, sub(x, uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))), add(y, uint((uint((uint((uint((uint(0)+1))+1))+1))+1))), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, sub(x, uint((uint((uint((uint(0)+1))+1))+1))), add(y, uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))), uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)))
	_ = rocqSDLSetDrawColor(ren, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, sub(x, uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))), sub(y, uint((uint((uint((uint((uint(0)+1))+1))+1))+1))), uint((uint((uint(0) + 1)) + 1)), uint((uint((uint(0) + 1)) + 1)))
	return rocqSDLFillRect(ren, add(x, uint((uint((uint((uint(0)+1))+1))+1))), sub(y, uint((uint((uint((uint((uint(0)+1))+1))+1))+1))), uint((uint((uint(0) + 1)) + 1)), uint((uint((uint(0) + 1)) + 1)))
}

var half_cell uint = uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))

func cell_center_x(col uint) uint {
	return add(mul(col, cell_size), half_cell)
}

func cell_center_y(row uint) uint {
	return add(mul(row, cell_size), half_cell)
}

func lerp(from_v uint, to_v uint, num uint, den uint) uint {
	return (func() any {
		if den <= uint(0) {
			return to_v
		} else {
			return (func() any {
				if from_v <= to_v {
					return add(from_v, func() uint {
						if den == 0 {
							return 0
						}
						return mul(sub(to_v, from_v), num) / den
					}())
				} else {
					return sub(from_v, func() uint {
						if den == 0 {
							return 0
						}
						return mul(sub(from_v, to_v), num) / den
					}())
				}
			}()).(uint)
		}
	}()).(uint)
}

func dir_to_nat(d direction) uint {
	return (func() uint {
		_scrut175 := d
		switch _scrut175 {
		case Up:
			return uint((uint((uint(0) + 1)) + 1))
		case Down:
			return uint((uint((uint((uint(0) + 1)) + 1)) + 1))
		case Left:
			return uint((uint(0) + 1))
		default:
			return uint(0)
		}
		panic("unreachable")
	})()
}

func ghost_color_index(idx uint, gm ghost_mode) uint {
	return (func() uint {
		_scrut176 := gm
		switch _scrut176 {
		case Chase:
			return idx
		case Frightened:
			return uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1))
		}
		panic("unreachable")
	})()
}

func draw_dot_check(ren RocqSDLRenderer, cx uint, cy uint) struct{} {
	_ = rocqSDLSetDrawColor(ren, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, sub(cx, uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))), add(cy, uint((uint(0)+1))), uint((uint((uint((uint(0) + 1)) + 1)) + 1)), uint((uint((uint((uint(0) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, sub(cx, uint((uint((uint((uint(0)+1))+1))+1))), add(cy, uint((uint((uint((uint((uint(0)+1))+1))+1))+1))), uint((uint((uint((uint(0) + 1)) + 1)) + 1)), uint((uint((uint((uint(0) + 1)) + 1)) + 1)))
	_ = rocqSDLFillRect(ren, cx, add(cy, uint((uint(0)+1))), uint((uint((uint((uint(0) + 1)) + 1)) + 1)), uint((uint((uint((uint(0) + 1)) + 1)) + 1)))
	return rocqSDLFillRect(ren, add(cx, uint((uint((uint((uint(0)+1))+1))+1))), sub(cy, uint((uint((uint(0)+1))+1))), uint((uint((uint((uint(0) + 1)) + 1)) + 1)), uint((uint((uint((uint(0) + 1)) + 1)) + 1)))
}

func draw_row_cells(ren RocqSDLRenderer, row uint, col uint, cells list[cell], pellet_phase uint) struct{} {
	return (func() struct{} {
		var _box any = cells
		_scrut177 := _box.(*listImpl)
		switch _scrut177._v {
		case 0:
			return struct{}{}
		case 1:
			c := _scrut177._c1_d0
			rest := _scrut177._c1_d1
			_ = (func() any {
				_scrut178 := c
				switch _scrut178 {
				case Wall:
					_ = rocqSDLSetDrawColor(ren, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
					return rocqSDLFillRect(ren, add(mul(col, cell_size), uint((uint(0)+1))), add(mul(row, cell_size), uint((uint(0)+1))), sub(cell_size, uint((uint((uint(0)+1))+1))), sub(cell_size, uint((uint((uint(0)+1))+1))))
				case Empty:
					return struct{}{}
				case Dot:
					_ = rocqSDLSetDrawColor(ren, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)))
					return draw_filled_circle(ren, cell_center_x(col), cell_center_y(row), uint((uint((uint(0) + 1)) + 1)))
				case PowerPellet:
					return draw_dot_check(ren, cell_center_x(col), cell_center_y(row))
				}
				panic("unreachable")
			})()
			return draw_row_cells(ren, row, uint((col + 1)), (rest).(list[cell]), pellet_phase)
		}
		panic("unreachable")
	})()
}

func draw_board_rows(ren RocqSDLRenderer, row uint, rows list[list[cell]], pellet_phase uint) struct{} {
	return (func() struct{} {
		var _box any = rows
		_scrut179 := _box.(*listImpl)
		switch _scrut179._v {
		case 0:
			return struct{}{}
		case 1:
			cells := _scrut179._c1_d0
			rest := _scrut179._c1_d1
			_ = draw_row_cells(ren, row, uint(0), (cells).(list[cell]), pellet_phase)
			return draw_board_rows(ren, uint((row + 1)), (rest).(list[list[cell]]), pellet_phase)
		}
		panic("unreachable")
	})()
}

func draw_board_sdl(ren RocqSDLRenderer, gs game_state, pellet_phase uint) struct{} {
	return draw_board_rows(ren, uint(0), (func() list[list[cell]] {
		_scrut180 := gs
		board0 := _scrut180.board
		return board0
	})(), pellet_phase)
}

func draw_ghosts_aux(ren RocqSDLRenderer, idx uint, gs list[ghost_state], prev_gs list[ghost_state], t_num uint, t_den uint, time_ms uint) struct{} {
	return (func() struct{} {
		var _box any = gs
		_scrut181 := _box.(*listImpl)
		switch _scrut181._v {
		case 0:
			return struct{}{}
		case 1:
			g := _scrut181._c1_d0
			rest := _scrut181._c1_d1
			prev_g := (func() ghost_state {
				var _box any = prev_gs
				_scrut182 := _box.(*listImpl)
				switch _scrut182._v {
				case 0:
					return (g).(ghost_state)
				case 1:
					pg := _scrut182._c1_d0
					return (pg).(ghost_state)
				}
				panic("unreachable")
			})()
			prev_rest := (func() list[ghost_state] {
				var _box any = prev_gs
				_scrut183 := _box.(*listImpl)
				switch _scrut183._v {
				case 0:
					return nil0
				case 1:
					pr := _scrut183._c1_d1
					return (pr).(list[ghost_state])
				}
				panic("unreachable")
			})()
			px := lerp(cell_center_x((func() uint {
				_scrut184 := (func() position {
					_scrut185 := prev_g
					gpos0 := _scrut185.gpos
					return gpos0
				})()
				pcol0 := _scrut184.pcol
				return pcol0
			})()), cell_center_x((func() uint {
				_scrut186 := (func() position {
					var _rbox any = g
					_scrut187 := _rbox.(ghost_state)
					gpos0 := _scrut187.gpos
					return gpos0
				})()
				pcol0 := _scrut186.pcol
				return pcol0
			})()), t_num, t_den)
			py := lerp(cell_center_y((func() uint {
				_scrut188 := (func() position {
					_scrut189 := prev_g
					gpos0 := _scrut189.gpos
					return gpos0
				})()
				prow0 := _scrut188.prow
				return prow0
			})()), cell_center_y((func() uint {
				_scrut190 := (func() position {
					var _rbox any = g
					_scrut191 := _rbox.(ghost_state)
					gpos0 := _scrut191.gpos
					return gpos0
				})()
				prow0 := _scrut190.prow
				return prow0
			})()), t_num, t_den)
			col := ghost_color_index(func() uint {
				if uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) == 0 {
					return idx
				}
				return idx % uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1))
			}(), (func() ghost_mode {
				var _rbox any = g
				_scrut192 := _rbox.(ghost_state)
				gmode0 := _scrut192.gmode
				return gmode0
			})())
			_ = draw_ghost_sprite(ren, px, py, col)
			return draw_ghosts_aux(ren, uint((idx + 1)), (rest).(list[ghost_state]), prev_rest, t_num, t_den, time_ms)
		}
		panic("unreachable")
	})()
}

func dir_to_degrees(d direction) uint {
	return (func() uint {
		_scrut193 := d
		switch _scrut193 {
		case Up:
			return uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))
		case Down:
			return uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))
		default:
			return uint(0)
		}
		panic("unreachable")
	})()
}

func dir_flip_h(d direction) bool {
	return (func() bool {
		_scrut194 := d
		switch _scrut194 {
		case Left:
			return true
		default:
			return false
		}
		panic("unreachable")
	})()
}

func draw_player_sdl(ren RocqSDLRenderer, tex RocqSDLTexture, gs game_state, prev_pos position, t_num uint, t_den uint) struct{} {
	px := lerp(cell_center_x((func() uint {
		_scrut195 := prev_pos
		pcol0 := _scrut195.pcol
		return pcol0
	})()), cell_center_x((func() uint {
		_scrut196 := (func() position {
			_scrut197 := gs
			pacpos0 := _scrut197.pacpos
			return pacpos0
		})()
		pcol0 := _scrut196.pcol
		return pcol0
	})()), t_num, t_den)
	py := lerp(cell_center_y((func() uint {
		_scrut198 := prev_pos
		prow0 := _scrut198.prow
		return prow0
	})()), cell_center_y((func() uint {
		_scrut199 := (func() position {
			_scrut200 := gs
			pacpos0 := _scrut200.pacpos
			return pacpos0
		})()
		prow0 := _scrut199.prow
		return prow0
	})()), t_num, t_den)
	return rocqSDLRenderTextureRotated(ren, tex, px, py, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), dir_to_degrees((func() direction {
		_scrut201 := gs
		pacdir0 := _scrut201.pacdir
		return pacdir0
	})()), dir_flip_h((func() direction {
		_scrut202 := gs
		pacdir0 := _scrut202.pacdir
		return pacdir0
	})()))
}

func draw_lives_aux(ren RocqSDLRenderer, n uint, i uint) struct{} {
	return (func() any {
		var _sni_ any = n
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return struct{}{}
		} else {
			n_ := _su_ - 1
			return (func() struct{} {
				_ = draw_life_icon(ren, sub(sub(win_width, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))), mul(i, uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)))), add(add(mul(board_height, cell_size), uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))))
				return draw_lives_aux(ren, n_, uint((i + 1)))
			})()
		}
	}()).(struct{})
}

func draw_status_bar(ren RocqSDLRenderer, gs game_state) struct{} {
	_ = draw_number_sprite(ren, (func() uint {
		_scrut203 := gs
		score0 := _scrut203.score
		return score0
	})(), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), add(mul(board_height, cell_size), uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))))
	return draw_lives_aux(ren, (func() uint {
		_scrut204 := gs
		lives0 := _scrut204.lives
		return lives0
	})(), uint(0))
}

func render_frame(ren RocqSDLRenderer, tex RocqSDLTexture, gs game_state, prev_pac position, prev_ghosts list[ghost_state], t_num uint, t_den uint, time_ms uint) struct{} {
	_ = rocqSDLSetDrawColor(ren, uint(0), uint(0), uint(0))
	_ = rocqSDLClear(ren)
	_ = draw_board_sdl(ren, gs, time_ms)
	_ = draw_ghosts_aux(ren, uint(0), (func() list[ghost_state] {
		_scrut205 := gs
		ghosts0 := _scrut205.ghosts
		return ghosts0
	})(), prev_ghosts, t_num, t_den, time_ms)
	_ = draw_player_sdl(ren, tex, gs, prev_pac, t_num, t_den)
	_ = draw_status_bar(ren, gs)
	return rocqSDLPresent(ren)
}

func render_paused_frame(ren RocqSDLRenderer, tex RocqSDLTexture, gs game_state, prev_pac position, prev_ghosts list[ghost_state], t_num uint, t_den uint, time_ms uint) struct{} {
	_ = rocqSDLSetDrawColor(ren, uint(0), uint(0), uint(0))
	_ = rocqSDLClear(ren)
	_ = draw_board_sdl(ren, gs, time_ms)
	_ = draw_ghosts_aux(ren, uint(0), (func() list[ghost_state] {
		_scrut206 := gs
		ghosts0 := _scrut206.ghosts
		return ghosts0
	})(), prev_ghosts, t_num, t_den, time_ms)
	_ = draw_player_sdl(ren, tex, gs, prev_pac, t_num, t_den)
	_ = draw_status_bar(ren, gs)
	_ = draw_status_message(ren, msg_paused)
	return rocqSDLPresent(ren)
}

func handle_key_down(key sdl_key, gs game_state) struct {
	fst bool
	snd game_state
} {
	return (func() struct {
		fst bool
		snd game_state
	} {
		_scrut207 := key
		switch _scrut207 {
		case KeyEscape:
			return struct {
				fst bool
				snd game_state
			}{fst: true, snd: gs}
		case KeyQ:
			return struct {
				fst bool
				snd game_state
			}{fst: true, snd: gs}
		case KeyUp:
			return struct {
				fst bool
				snd game_state
			}{fst: false, snd: set_direction(Up, gs)}
		case KeyDown:
			return struct {
				fst bool
				snd game_state
			}{fst: false, snd: set_direction(Down, gs)}
		case KeyLeft:
			return struct {
				fst bool
				snd game_state
			}{fst: false, snd: set_direction(Left, gs)}
		case KeyW:
			return struct {
				fst bool
				snd game_state
			}{fst: false, snd: set_direction(Up, gs)}
		case KeyA:
			return struct {
				fst bool
				snd game_state
			}{fst: false, snd: set_direction(Left, gs)}
		case KeyS:
			return struct {
				fst bool
				snd game_state
			}{fst: false, snd: set_direction(Down, gs)}
		case KeySpace:
			return struct {
				fst bool
				snd game_state
			}{fst: false, snd: gs}
		default:
			return struct {
				fst bool
				snd game_state
			}{fst: false, snd: set_direction(Right, gs)}
		}
		panic("unreachable")
	})()
}

func handle_event(ev sdl_event, gs game_state) struct {
	fst bool
	snd game_state
} {
	return (func() struct {
		fst bool
		snd game_state
	} {
		var _box any = ev
		_scrut208 := _box.(*sdl_eventImpl)
		switch _scrut208._v {
		case 0:
			return struct {
				fst bool
				snd game_state
			}{fst: true, snd: gs}
		case 1:
			key := _scrut208._c1_d0
			return handle_key_down((key).(sdl_key), gs)
		}
		panic("unreachable")
	})()
}

func find_pixel_collision(px uint, py uint, gs list[ghost_state], prev_gs list[ghost_state], t_num uint, t_den uint, threshold uint, idx uint) *struct {
	fst uint
	snd ghost_mode
} {
	return (func() *struct {
		fst uint
		snd ghost_mode
	} {
		var _box any = gs
		_scrut209 := _box.(*listImpl)
		switch _scrut209._v {
		case 0:
			return nil
		case 1:
			g := _scrut209._c1_d0
			rest := _scrut209._c1_d1
			prev_g := (func() ghost_state {
				var _box any = prev_gs
				_scrut210 := _box.(*listImpl)
				switch _scrut210._v {
				case 0:
					return (g).(ghost_state)
				case 1:
					pg := _scrut210._c1_d0
					return (pg).(ghost_state)
				}
				panic("unreachable")
			})()
			prev_rest := (func() list[ghost_state] {
				var _box any = prev_gs
				_scrut211 := _box.(*listImpl)
				switch _scrut211._v {
				case 0:
					return nil0
				case 1:
					pr := _scrut211._c1_d1
					return (pr).(list[ghost_state])
				}
				panic("unreachable")
			})()
			gx := lerp(cell_center_x((func() uint {
				_scrut212 := (func() position {
					_scrut213 := prev_g
					gpos0 := _scrut213.gpos
					return gpos0
				})()
				pcol0 := _scrut212.pcol
				return pcol0
			})()), cell_center_x((func() uint {
				_scrut214 := (func() position {
					var _rbox any = g
					_scrut215 := _rbox.(ghost_state)
					gpos0 := _scrut215.gpos
					return gpos0
				})()
				pcol0 := _scrut214.pcol
				return pcol0
			})()), t_num, t_den)
			gy := lerp(cell_center_y((func() uint {
				_scrut216 := (func() position {
					_scrut217 := prev_g
					gpos0 := _scrut217.gpos
					return gpos0
				})()
				prow0 := _scrut216.prow
				return prow0
			})()), cell_center_y((func() uint {
				_scrut218 := (func() position {
					var _rbox any = g
					_scrut219 := _rbox.(ghost_state)
					gpos0 := _scrut219.gpos
					return gpos0
				})()
				prow0 := _scrut218.prow
				return prow0
			})()), t_num, t_den)
			dx := abs_diff(px, gx)
			dy := abs_diff(py, gy)
			return (func() any {
				if add(mul(dx, dx), mul(dy, dy)) <= mul(threshold, threshold) {
					return rocqSome(struct {
						fst uint
						snd ghost_mode
					}{fst: idx, snd: (func() ghost_mode {
						var _rbox any = g
						_scrut220 := _rbox.(ghost_state)
						gmode0 := _scrut220.gmode
						return gmode0
					})()})
				} else {
					return find_pixel_collision(px, py, (rest).(list[ghost_state]), prev_rest, t_num, t_den, threshold, uint((idx + 1)))
				}
			}()).(*struct {
				fst uint
				snd ghost_mode
			})
		}
		panic("unreachable")
	})()
}

func eat_ghost_idx(idx uint, gs game_state) game_state {
	default_g := mkGhost(mkPos(uint(0), uint(0)), DirNone, Chase)
	new_ghosts := replace_nth(idx, (func() list[any] {
		_scrut221 := gs
		ghosts0 := _scrut221.ghosts
		return ghosts0
	})(), respawn_ghost((nth(idx, (func() list[any] {
		_scrut222 := gs
		ghosts0 := _scrut222.ghosts
		return ghosts0
	})(), default_g)).(ghost_state)))
	return mkState((func() list[list[cell]] {
		_scrut223 := gs
		board0 := _scrut223.board
		return board0
	})(), (func() position {
		_scrut224 := gs
		pacpos0 := _scrut224.pacpos
		return pacpos0
	})(), (func() direction {
		_scrut225 := gs
		pacdir0 := _scrut225.pacdir
		return pacdir0
	})(), (func() direction {
		_scrut226 := gs
		desired_dir0 := _scrut226.desired_dir
		return desired_dir0
	})(), new_ghosts, add((func() uint {
		_scrut227 := gs
		score0 := _scrut227.score
		return score0
	})(), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))), (func() uint {
		_scrut228 := gs
		lives0 := _scrut228.lives
		return lives0
	})(), (func() uint {
		_scrut229 := gs
		dots_left0 := _scrut229.dots_left
		return dots_left0
	})(), (func() uint {
		_scrut230 := gs
		power_timer0 := _scrut230.power_timer
		return power_timer0
	})(), (func() bool {
		_scrut231 := gs
		game_over0 := _scrut231.game_over
		return game_over0
	})(), (func() bool {
		_scrut232 := gs
		game_won0 := _scrut232.game_won
		return game_won0
	})())
}

func lose_one_life(gs game_state) game_state {
	new_lives := func() uint {
		if (func() uint {
			_scrut233 := gs
			lives0 := _scrut233.lives
			return lives0
		})() > 0 {
			return (func() uint {
				_scrut233 := gs
				lives0 := _scrut233.lives
				return lives0
			})() - 1
		}
		return 0
	}()
	return mkState((func() list[list[cell]] {
		_scrut234 := gs
		board0 := _scrut234.board
		return board0
	})(), mkPos(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))), DirNone, DirNone, (func() list[ghost_state] {
		_scrut235 := gs
		ghosts0 := _scrut235.ghosts
		return ghosts0
	})(), (func() uint {
		_scrut236 := gs
		score0 := _scrut236.score
		return score0
	})(), new_lives, (func() uint {
		_scrut237 := gs
		dots_left0 := _scrut237.dots_left
		return dots_left0
	})(), uint(0), (new_lives == uint(0)), (func() bool {
		_scrut238 := gs
		game_won0 := _scrut238.game_won
		return game_won0
	})())
}

func apply_direction(key sdl_key, gs game_state) game_state {
	return (func() game_state {
		_scrut239 := key
		switch _scrut239 {
		case KeyUp:
			return set_direction(Up, gs)
		case KeyDown:
			return set_direction(Down, gs)
		case KeyLeft:
			return set_direction(Left, gs)
		case KeyRight:
			return set_direction(Right, gs)
		case KeyW:
			return set_direction(Up, gs)
		case KeyA:
			return set_direction(Left, gs)
		case KeyS:
			return set_direction(Down, gs)
		case KeyD:
			return set_direction(Right, gs)
		default:
			return gs
		}
		panic("unreachable")
	})()
}

type loop_state struct {
	ls_game        game_state
	ls_prev_pac    position
	ls_prev_ghosts list[ghost_state]
	ls_last_tick   uint
	ls_start_time  uint
	ls_texture     RocqSDLTexture
	ls_phase       phase
	ls_phase_time  uint
	ls_quit        bool
}

func mkLoop(ls_game game_state, ls_prev_pac position, ls_prev_ghosts list[ghost_state], ls_last_tick uint, ls_start_time uint, ls_texture RocqSDLTexture, ls_phase phase, ls_phase_time uint, ls_quit bool) loop_state {
	return loop_state{ls_game: ls_game, ls_prev_pac: ls_prev_pac, ls_prev_ghosts: ls_prev_ghosts, ls_last_tick: ls_last_tick, ls_start_time: ls_start_time, ls_texture: ls_texture, ls_phase: ls_phase, ls_phase_time: ls_phase_time, ls_quit: ls_quit}
}

func ls_game(l loop_state) game_state {
	return (func() game_state {
		_scrut240 := l
		ls_game0 := _scrut240.ls_game
		return ls_game0
	})()
}

func ls_prev_pac(l loop_state) position {
	return (func() position {
		_scrut241 := l
		ls_prev_pac0 := _scrut241.ls_prev_pac
		return ls_prev_pac0
	})()
}

func ls_prev_ghosts(l loop_state) list[ghost_state] {
	return (func() list[ghost_state] {
		_scrut242 := l
		ls_prev_ghosts0 := _scrut242.ls_prev_ghosts
		return ls_prev_ghosts0
	})()
}

func ls_last_tick(l loop_state) uint {
	return (func() uint {
		_scrut243 := l
		ls_last_tick0 := _scrut243.ls_last_tick
		return ls_last_tick0
	})()
}

func ls_start_time(l loop_state) uint {
	return (func() uint {
		_scrut244 := l
		ls_start_time0 := _scrut244.ls_start_time
		return ls_start_time0
	})()
}

func ls_texture(l loop_state) RocqSDLTexture {
	return (func() RocqSDLTexture {
		_scrut245 := l
		ls_texture0 := _scrut245.ls_texture
		return ls_texture0
	})()
}

func ls_phase(l loop_state) phase {
	return (func() phase {
		_scrut246 := l
		ls_phase0 := _scrut246.ls_phase
		return ls_phase0
	})()
}

func ls_phase_time(l loop_state) uint {
	return (func() uint {
		_scrut247 := l
		ls_phase_time0 := _scrut247.ls_phase_time
		return ls_phase_time0
	})()
}

func ls_quit(l loop_state) bool {
	return (func() bool {
		_scrut248 := l
		ls_quit0 := _scrut248.ls_quit
		return ls_quit0
	})()
}

func frame_delay(frame_start uint) struct{} {
	now2 := rocqSDLGetTicks()
	elapsed := sub(now2, frame_start)
	return (func() any {
		if elapsed < frame_ms {
			return rocqSDLDelay(sub(frame_ms, elapsed))
		} else {
			return struct{}{}
		}
	}()).(struct{})
}

var collision_threshold uint = uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))

var snd_checkmark string = "assets/checkmark.mp3"

var snd_game_over string = "assets/game-over.mp3"

var snd_kill_ghost string = "assets/kill-ghost.mp3"

var snd_lose_life string = "assets/lose-life.mp3"

var snd_tap string = "assets/tap.mp3"

var snd_win string = "assets/win.mp3"

func play_cell_sound(c cell) struct{} {
	return (func() struct{} {
		_scrut249 := c
		switch _scrut249 {
		case Dot:
			return rocqSDLPlaySound(snd_tap)
		case PowerPellet:
			return rocqSDLPlaySound(snd_checkmark)
		default:
			return struct{}{}
		}
		panic("unreachable")
	})()
}

func process_frame(ren RocqSDLRenderer, ls loop_state) struct {
	fst bool
	snd loop_state
} {
	ev := rocqSDLPollEvent()
	return (func() struct {
		fst bool
		snd loop_state
	} {
		var _box any = ev
		_scrut250 := _box.(*sdl_eventImpl)
		switch _scrut250._v {
		case 0:
			return struct {
				fst bool
				snd loop_state
			}{fst: true, snd: ls}
		case 1:
			now := rocqSDLGetTicks()
			time_ms := sub(now, (func() uint {
				_scrut251 := ls
				ls_start_time0 := _scrut251.ls_start_time
				return ls_start_time0
			})())
			return (func() struct {
				fst bool
				snd loop_state
			} {
				_scrut252 := (func() phase {
					_scrut253 := ls
					ls_phase0 := _scrut253.ls_phase
					return ls_phase0
				})()
				switch _scrut252 {
				case Playing:
					return (func() struct {
						fst bool
						snd loop_state
					} {
						var _box any = ev
						_scrut254 := _box.(*sdl_eventImpl)
						switch _scrut254._v {
						case 0:
							gs1 := (func() game_state {
								_scrut255 := ls
								ls_game0 := _scrut255.ls_game
								return ls_game0
							})()
							elapsed := sub(now, (func() uint {
								_scrut256 := ls
								ls_last_tick0 := _scrut256.ls_last_tick
								return ls_last_tick0
							})())
							do_tick := (tick_ms <= elapsed)
							gs2 := (func() any {
								if do_tick {
									return tick(gs1)
								} else {
									return gs1
								}
							}()).(game_state)
							new_prev_pac := (func() any {
								if do_tick {
									return (func() position {
										_scrut257 := gs1
										pacpos0 := _scrut257.pacpos
										return pacpos0
									})()
								} else {
									return (func() position {
										_scrut258 := ls
										ls_prev_pac0 := _scrut258.ls_prev_pac
										return ls_prev_pac0
									})()
								}
							}()).(position)
							new_prev_ghosts := (func() any {
								if do_tick {
									return (func() list[ghost_state] {
										_scrut259 := gs1
										ghosts0 := _scrut259.ghosts
										return ghosts0
									})()
								} else {
									return (func() list[ghost_state] {
										_scrut260 := ls
										ls_prev_ghosts0 := _scrut260.ls_prev_ghosts
										return ls_prev_ghosts0
									})()
								}
							}()).(list[ghost_state])
							new_last_tick := (func() any {
								if do_tick {
									return now
								} else {
									return (func() uint {
										_scrut261 := ls
										ls_last_tick0 := _scrut261.ls_last_tick
										return ls_last_tick0
									})()
								}
							}()).(uint)
							eaten_cell := (func() any {
								if do_tick {
									return get_cell((func() uint {
										_scrut262 := (func() position {
											_scrut263 := gs2
											pacpos0 := _scrut263.pacpos
											return pacpos0
										})()
										prow0 := _scrut262.prow
										return prow0
									})(), (func() uint {
										_scrut264 := (func() position {
											_scrut265 := gs2
											pacpos0 := _scrut265.pacpos
											return pacpos0
										})()
										pcol0 := _scrut264.pcol
										return pcol0
									})(), (func() list[list[cell]] {
										_scrut266 := gs1
										board0 := _scrut266.board
										return board0
									})())
								} else {
									return Empty
								}
							}()).(cell)
							t_num := sub(now, new_last_tick)
							ppx := lerp(cell_center_x((func() uint {
								_scrut267 := new_prev_pac
								pcol0 := _scrut267.pcol
								return pcol0
							})()), cell_center_x((func() uint {
								_scrut268 := (func() position {
									_scrut269 := gs2
									pacpos0 := _scrut269.pacpos
									return pacpos0
								})()
								pcol0 := _scrut268.pcol
								return pcol0
							})()), t_num, tick_ms)
							ppy := lerp(cell_center_y((func() uint {
								_scrut270 := new_prev_pac
								prow0 := _scrut270.prow
								return prow0
							})()), cell_center_y((func() uint {
								_scrut271 := (func() position {
									_scrut272 := gs2
									pacpos0 := _scrut272.pacpos
									return pacpos0
								})()
								prow0 := _scrut271.prow
								return prow0
							})()), t_num, tick_ms)
							return (func() any {
								if (func() bool {
									_scrut273 := gs2
									game_won0 := _scrut273.game_won
									return game_won0
								})() {
									return (func() struct {
										fst bool
										snd loop_state
									} {
										_ = rocqSDLPlaySound(snd_win)
										return (func() struct {
											fst bool
											snd loop_state
										} {
											_ = render_frame(ren, (func() RocqSDLTexture {
												_scrut274 := ls
												ls_texture0 := _scrut274.ls_texture
												return ls_texture0
											})(), gs2, new_prev_pac, new_prev_ghosts, t_num, tick_ms, time_ms)
											return struct {
												fst bool
												snd loop_state
											}{fst: false, snd: mkLoop(gs2, new_prev_pac, new_prev_ghosts, new_last_tick, (func() uint {
												_scrut275 := ls
												ls_start_time0 := _scrut275.ls_start_time
												return ls_start_time0
											})(), (func() RocqSDLTexture {
												_scrut276 := ls
												ls_texture0 := _scrut276.ls_texture
												return ls_texture0
											})(), WinScreen, now, false)}
										})()
									})()
								} else {
									return (func() any {
										if find_pixel_collision(ppx, ppy, (func() list[ghost_state] {
											_scrut277 := gs2
											ghosts0 := _scrut277.ghosts
											return ghosts0
										})(), new_prev_ghosts, t_num, tick_ms, collision_threshold, uint(0)) != nil {
											p := find_pixel_collision(ppx, ppy, (func() list[ghost_state] {
												_scrut277 := gs2
												ghosts0 := _scrut277.ghosts
												return ghosts0
											})(), new_prev_ghosts, t_num, tick_ms, collision_threshold, uint(0))
											return (func() any {
												idx := p.fst
												g := p.snd
												return (func() struct {
													fst bool
													snd loop_state
												} {
													_scrut278 := g
													switch _scrut278 {
													case Chase:
														gs3 := lose_one_life(gs2)
														next_pac := (func() position {
															_scrut279 := gs3
															pacpos0 := _scrut279.pacpos
															return pacpos0
														})()
														next_ghosts := (func() list[ghost_state] {
															_scrut280 := gs3
															ghosts0 := _scrut280.ghosts
															return ghosts0
														})()
														return (func() any {
															if (func() uint {
																_scrut281 := gs3
																lives0 := _scrut281.lives
																return lives0
															})() == uint(0) {
																return (func() struct {
																	fst bool
																	snd loop_state
																} {
																	next_ls := mkLoop(gs3, next_pac, next_ghosts, now, (func() uint {
																		_scrut282 := ls
																		ls_start_time0 := _scrut282.ls_start_time
																		return ls_start_time0
																	})(), (func() RocqSDLTexture {
																		_scrut283 := ls
																		ls_texture0 := _scrut283.ls_texture
																		return ls_texture0
																	})(), GameOverScreen, now, false)
																	return (func() struct {
																		fst bool
																		snd loop_state
																	} {
																		_ = rocqSDLPlaySound(snd_game_over)
																		return struct {
																			fst bool
																			snd loop_state
																		}{fst: false, snd: next_ls}
																	})()
																})()
															} else {
																return (func() struct {
																	fst bool
																	snd loop_state
																} {
																	next_ls := mkLoop(gs3, next_pac, next_ghosts, now, (func() uint {
																		_scrut284 := ls
																		ls_start_time0 := _scrut284.ls_start_time
																		return ls_start_time0
																	})(), (func() RocqSDLTexture {
																		_scrut285 := ls
																		ls_texture0 := _scrut285.ls_texture
																		return ls_texture0
																	})(), DeathPause, now, false)
																	return (func() struct {
																		fst bool
																		snd loop_state
																	} {
																		_ = rocqSDLPlaySound(snd_lose_life)
																		return struct {
																			fst bool
																			snd loop_state
																		}{fst: false, snd: next_ls}
																	})()
																})()
															}
														}()).(struct {
															fst bool
															snd loop_state
														})
													case Frightened:
														gs3 := eat_ghost_idx(idx, gs2)
														next_ls := mkLoop(gs3, new_prev_pac, new_prev_ghosts, new_last_tick, (func() uint {
															_scrut286 := ls
															ls_start_time0 := _scrut286.ls_start_time
															return ls_start_time0
														})(), (func() RocqSDLTexture {
															_scrut287 := ls
															ls_texture0 := _scrut287.ls_texture
															return ls_texture0
														})(), Playing, uint(0), false)
														_ = rocqSDLPlaySound(snd_kill_ghost)
														_ = render_frame(ren, (func() RocqSDLTexture {
															_scrut288 := ls
															ls_texture0 := _scrut288.ls_texture
															return ls_texture0
														})(), gs3, (func() position {
															_scrut289 := next_ls
															ls_prev_pac0 := _scrut289.ls_prev_pac
															return ls_prev_pac0
														})(), (func() list[ghost_state] {
															_scrut290 := next_ls
															ls_prev_ghosts0 := _scrut290.ls_prev_ghosts
															return ls_prev_ghosts0
														})(), t_num, tick_ms, time_ms)
														_ = frame_delay(now)
														return struct {
															fst bool
															snd loop_state
														}{fst: false, snd: next_ls}
													}
													panic("unreachable")
												})()
											}()).(struct {
												fst bool
												snd loop_state
											})
										} else {
											return (func() struct {
												fst bool
												snd loop_state
											} {
												_ = play_cell_sound(eaten_cell)
												return (func() struct {
													fst bool
													snd loop_state
												} {
													_ = render_frame(ren, (func() RocqSDLTexture {
														_scrut291 := ls
														ls_texture0 := _scrut291.ls_texture
														return ls_texture0
													})(), gs2, new_prev_pac, new_prev_ghosts, t_num, tick_ms, time_ms)
													return (func() struct {
														fst bool
														snd loop_state
													} {
														_ = frame_delay(now)
														return struct {
															fst bool
															snd loop_state
														}{fst: false, snd: mkLoop(gs2, new_prev_pac, new_prev_ghosts, new_last_tick, (func() uint {
															_scrut292 := ls
															ls_start_time0 := _scrut292.ls_start_time
															return ls_start_time0
														})(), (func() RocqSDLTexture {
															_scrut293 := ls
															ls_texture0 := _scrut293.ls_texture
															return ls_texture0
														})(), Playing, uint(0), false)}
													})()
												})()
											})()
										}
									}()).(struct {
										fst bool
										snd loop_state
									})
								}
							}()).(struct {
								fst bool
								snd loop_state
							})
						case 1:
							key := _scrut254._c1_d0
							return (func() struct {
								fst bool
								snd loop_state
							} {
								_scrut294 := key
								switch _scrut294 {
								case KeySpace:
									_ = render_paused_frame(ren, (func() RocqSDLTexture {
										_scrut295 := ls
										ls_texture0 := _scrut295.ls_texture
										return ls_texture0
									})(), (func() game_state {
										_scrut296 := ls
										ls_game0 := _scrut296.ls_game
										return ls_game0
									})(), (func() position {
										_scrut297 := ls
										ls_prev_pac0 := _scrut297.ls_prev_pac
										return ls_prev_pac0
									})(), (func() list[ghost_state] {
										_scrut298 := ls
										ls_prev_ghosts0 := _scrut298.ls_prev_ghosts
										return ls_prev_ghosts0
									})(), uint(0), uint((uint(0) + 1)), time_ms)
									_ = frame_delay(now)
									return struct {
										fst bool
										snd loop_state
									}{fst: false, snd: mkLoop((func() game_state {
										_scrut299 := ls
										ls_game0 := _scrut299.ls_game
										return ls_game0
									})(), (func() position {
										_scrut300 := ls
										ls_prev_pac0 := _scrut300.ls_prev_pac
										return ls_prev_pac0
									})(), (func() list[ghost_state] {
										_scrut301 := ls
										ls_prev_ghosts0 := _scrut301.ls_prev_ghosts
										return ls_prev_ghosts0
									})(), now, (func() uint {
										_scrut302 := ls
										ls_start_time0 := _scrut302.ls_start_time
										return ls_start_time0
									})(), (func() RocqSDLTexture {
										_scrut303 := ls
										ls_texture0 := _scrut303.ls_texture
										return ls_texture0
									})(), Paused, now, false)}
								default:
									return (func() any {
										quit := handle_key_down((key).(sdl_key), (func() game_state {
											_scrut304 := ls
											ls_game0 := _scrut304.ls_game
											return ls_game0
										})()).fst
										gs1 := handle_key_down((key).(sdl_key), (func() game_state {
											_scrut304 := ls
											ls_game0 := _scrut304.ls_game
											return ls_game0
										})()).snd
										return (func() any {
											if quit {
												return struct {
													fst bool
													snd loop_state
												}{fst: true, snd: ls}
											} else {
												return (func() struct {
													fst bool
													snd loop_state
												} {
													elapsed := sub(now, (func() uint {
														_scrut305 := ls
														ls_last_tick0 := _scrut305.ls_last_tick
														return ls_last_tick0
													})())
													return (func() struct {
														fst bool
														snd loop_state
													} {
														do_tick := (tick_ms <= elapsed)
														return (func() struct {
															fst bool
															snd loop_state
														} {
															gs2 := (func() any {
																if do_tick {
																	return tick(gs1)
																} else {
																	return gs1
																}
															}()).(game_state)
															return (func() struct {
																fst bool
																snd loop_state
															} {
																new_prev_pac := (func() any {
																	if do_tick {
																		return (func() position {
																			_scrut306 := gs1
																			pacpos0 := _scrut306.pacpos
																			return pacpos0
																		})()
																	} else {
																		return (func() position {
																			_scrut307 := ls
																			ls_prev_pac0 := _scrut307.ls_prev_pac
																			return ls_prev_pac0
																		})()
																	}
																}()).(position)
																return (func() struct {
																	fst bool
																	snd loop_state
																} {
																	new_prev_ghosts := (func() any {
																		if do_tick {
																			return (func() list[ghost_state] {
																				_scrut308 := gs1
																				ghosts0 := _scrut308.ghosts
																				return ghosts0
																			})()
																		} else {
																			return (func() list[ghost_state] {
																				_scrut309 := ls
																				ls_prev_ghosts0 := _scrut309.ls_prev_ghosts
																				return ls_prev_ghosts0
																			})()
																		}
																	}()).(list[ghost_state])
																	return (func() struct {
																		fst bool
																		snd loop_state
																	} {
																		new_last_tick := (func() any {
																			if do_tick {
																				return now
																			} else {
																				return (func() uint {
																					_scrut310 := ls
																					ls_last_tick0 := _scrut310.ls_last_tick
																					return ls_last_tick0
																				})()
																			}
																		}()).(uint)
																		return (func() struct {
																			fst bool
																			snd loop_state
																		} {
																			eaten_cell := (func() any {
																				if do_tick {
																					return get_cell((func() uint {
																						_scrut311 := (func() position {
																							_scrut312 := gs2
																							pacpos0 := _scrut312.pacpos
																							return pacpos0
																						})()
																						prow0 := _scrut311.prow
																						return prow0
																					})(), (func() uint {
																						_scrut313 := (func() position {
																							_scrut314 := gs2
																							pacpos0 := _scrut314.pacpos
																							return pacpos0
																						})()
																						pcol0 := _scrut313.pcol
																						return pcol0
																					})(), (func() list[list[cell]] {
																						_scrut315 := gs1
																						board0 := _scrut315.board
																						return board0
																					})())
																				} else {
																					return Empty
																				}
																			}()).(cell)
																			return (func() struct {
																				fst bool
																				snd loop_state
																			} {
																				t_num := sub(now, new_last_tick)
																				return (func() struct {
																					fst bool
																					snd loop_state
																				} {
																					ppx := lerp(cell_center_x((func() uint {
																						_scrut316 := new_prev_pac
																						pcol0 := _scrut316.pcol
																						return pcol0
																					})()), cell_center_x((func() uint {
																						_scrut317 := (func() position {
																							_scrut318 := gs2
																							pacpos0 := _scrut318.pacpos
																							return pacpos0
																						})()
																						pcol0 := _scrut317.pcol
																						return pcol0
																					})()), t_num, tick_ms)
																					return (func() struct {
																						fst bool
																						snd loop_state
																					} {
																						ppy := lerp(cell_center_y((func() uint {
																							_scrut319 := new_prev_pac
																							prow0 := _scrut319.prow
																							return prow0
																						})()), cell_center_y((func() uint {
																							_scrut320 := (func() position {
																								_scrut321 := gs2
																								pacpos0 := _scrut321.pacpos
																								return pacpos0
																							})()
																							prow0 := _scrut320.prow
																							return prow0
																						})()), t_num, tick_ms)
																						return (func() any {
																							if (func() bool {
																								_scrut322 := gs2
																								game_won0 := _scrut322.game_won
																								return game_won0
																							})() {
																								return (func() struct {
																									fst bool
																									snd loop_state
																								} {
																									_ = rocqSDLPlaySound(snd_win)
																									return (func() struct {
																										fst bool
																										snd loop_state
																									} {
																										_ = render_frame(ren, (func() RocqSDLTexture {
																											_scrut323 := ls
																											ls_texture0 := _scrut323.ls_texture
																											return ls_texture0
																										})(), gs2, new_prev_pac, new_prev_ghosts, t_num, tick_ms, time_ms)
																										return struct {
																											fst bool
																											snd loop_state
																										}{fst: false, snd: mkLoop(gs2, new_prev_pac, new_prev_ghosts, new_last_tick, (func() uint {
																											_scrut324 := ls
																											ls_start_time0 := _scrut324.ls_start_time
																											return ls_start_time0
																										})(), (func() RocqSDLTexture {
																											_scrut325 := ls
																											ls_texture0 := _scrut325.ls_texture
																											return ls_texture0
																										})(), WinScreen, now, false)}
																									})()
																								})()
																							} else {
																								return (func() any {
																									if find_pixel_collision(ppx, ppy, (func() list[ghost_state] {
																										_scrut326 := gs2
																										ghosts0 := _scrut326.ghosts
																										return ghosts0
																									})(), new_prev_ghosts, t_num, tick_ms, collision_threshold, uint(0)) != nil {
																										p := find_pixel_collision(ppx, ppy, (func() list[ghost_state] {
																											_scrut326 := gs2
																											ghosts0 := _scrut326.ghosts
																											return ghosts0
																										})(), new_prev_ghosts, t_num, tick_ms, collision_threshold, uint(0))
																										return (func() any {
																											idx := p.fst
																											g := p.snd
																											return (func() struct {
																												fst bool
																												snd loop_state
																											} {
																												_scrut327 := g
																												switch _scrut327 {
																												case Chase:
																													gs3 := lose_one_life(gs2)
																													next_pac := (func() position {
																														_scrut328 := gs3
																														pacpos0 := _scrut328.pacpos
																														return pacpos0
																													})()
																													next_ghosts := (func() list[ghost_state] {
																														_scrut329 := gs3
																														ghosts0 := _scrut329.ghosts
																														return ghosts0
																													})()
																													return (func() any {
																														if (func() uint {
																															_scrut330 := gs3
																															lives0 := _scrut330.lives
																															return lives0
																														})() == uint(0) {
																															return (func() struct {
																																fst bool
																																snd loop_state
																															} {
																																next_ls := mkLoop(gs3, next_pac, next_ghosts, now, (func() uint {
																																	_scrut331 := ls
																																	ls_start_time0 := _scrut331.ls_start_time
																																	return ls_start_time0
																																})(), (func() RocqSDLTexture {
																																	_scrut332 := ls
																																	ls_texture0 := _scrut332.ls_texture
																																	return ls_texture0
																																})(), GameOverScreen, now, false)
																																return (func() struct {
																																	fst bool
																																	snd loop_state
																																} {
																																	_ = rocqSDLPlaySound(snd_game_over)
																																	return struct {
																																		fst bool
																																		snd loop_state
																																	}{fst: false, snd: next_ls}
																																})()
																															})()
																														} else {
																															return (func() struct {
																																fst bool
																																snd loop_state
																															} {
																																next_ls := mkLoop(gs3, next_pac, next_ghosts, now, (func() uint {
																																	_scrut333 := ls
																																	ls_start_time0 := _scrut333.ls_start_time
																																	return ls_start_time0
																																})(), (func() RocqSDLTexture {
																																	_scrut334 := ls
																																	ls_texture0 := _scrut334.ls_texture
																																	return ls_texture0
																																})(), DeathPause, now, false)
																																return (func() struct {
																																	fst bool
																																	snd loop_state
																																} {
																																	_ = rocqSDLPlaySound(snd_lose_life)
																																	return struct {
																																		fst bool
																																		snd loop_state
																																	}{fst: false, snd: next_ls}
																																})()
																															})()
																														}
																													}()).(struct {
																														fst bool
																														snd loop_state
																													})
																												case Frightened:
																													gs3 := eat_ghost_idx(idx, gs2)
																													next_ls := mkLoop(gs3, new_prev_pac, new_prev_ghosts, new_last_tick, (func() uint {
																														_scrut335 := ls
																														ls_start_time0 := _scrut335.ls_start_time
																														return ls_start_time0
																													})(), (func() RocqSDLTexture {
																														_scrut336 := ls
																														ls_texture0 := _scrut336.ls_texture
																														return ls_texture0
																													})(), Playing, uint(0), false)
																													_ = rocqSDLPlaySound(snd_kill_ghost)
																													_ = render_frame(ren, (func() RocqSDLTexture {
																														_scrut337 := ls
																														ls_texture0 := _scrut337.ls_texture
																														return ls_texture0
																													})(), gs3, (func() position {
																														_scrut338 := next_ls
																														ls_prev_pac0 := _scrut338.ls_prev_pac
																														return ls_prev_pac0
																													})(), (func() list[ghost_state] {
																														_scrut339 := next_ls
																														ls_prev_ghosts0 := _scrut339.ls_prev_ghosts
																														return ls_prev_ghosts0
																													})(), t_num, tick_ms, time_ms)
																													_ = frame_delay(now)
																													return struct {
																														fst bool
																														snd loop_state
																													}{fst: false, snd: next_ls}
																												}
																												panic("unreachable")
																											})()
																										}()).(struct {
																											fst bool
																											snd loop_state
																										})
																									} else {
																										return (func() struct {
																											fst bool
																											snd loop_state
																										} {
																											_ = play_cell_sound(eaten_cell)
																											return (func() struct {
																												fst bool
																												snd loop_state
																											} {
																												_ = render_frame(ren, (func() RocqSDLTexture {
																													_scrut340 := ls
																													ls_texture0 := _scrut340.ls_texture
																													return ls_texture0
																												})(), gs2, new_prev_pac, new_prev_ghosts, t_num, tick_ms, time_ms)
																												return (func() struct {
																													fst bool
																													snd loop_state
																												} {
																													_ = frame_delay(now)
																													return struct {
																														fst bool
																														snd loop_state
																													}{fst: false, snd: mkLoop(gs2, new_prev_pac, new_prev_ghosts, new_last_tick, (func() uint {
																														_scrut341 := ls
																														ls_start_time0 := _scrut341.ls_start_time
																														return ls_start_time0
																													})(), (func() RocqSDLTexture {
																														_scrut342 := ls
																														ls_texture0 := _scrut342.ls_texture
																														return ls_texture0
																													})(), Playing, uint(0), false)}
																												})()
																											})()
																										})()
																									}
																								}()).(struct {
																									fst bool
																									snd loop_state
																								})
																							}
																						}()).(struct {
																							fst bool
																							snd loop_state
																						})
																					})()
																				})()
																			})()
																		})()
																	})()
																})()
															})()
														})()
													})()
												})()
											}
										}()).(struct {
											fst bool
											snd loop_state
										})
									}()).(struct {
										fst bool
										snd loop_state
									})
								}
								panic("unreachable")
							})()
						}
						panic("unreachable")
					})()
				case Paused:
					return (func() struct {
						fst bool
						snd loop_state
					} {
						var _box any = ev
						_scrut343 := _box.(*sdl_eventImpl)
						switch _scrut343._v {
						case 0:
							_ = render_paused_frame(ren, (func() RocqSDLTexture {
								_scrut344 := ls
								ls_texture0 := _scrut344.ls_texture
								return ls_texture0
							})(), (func() game_state {
								_scrut345 := ls
								ls_game0 := _scrut345.ls_game
								return ls_game0
							})(), (func() position {
								_scrut346 := ls
								ls_prev_pac0 := _scrut346.ls_prev_pac
								return ls_prev_pac0
							})(), (func() list[ghost_state] {
								_scrut347 := ls
								ls_prev_ghosts0 := _scrut347.ls_prev_ghosts
								return ls_prev_ghosts0
							})(), uint(0), uint((uint(0) + 1)), time_ms)
							_ = frame_delay(now)
							return struct {
								fst bool
								snd loop_state
							}{fst: false, snd: ls}
						case 1:
							key := _scrut343._c1_d0
							return (func() struct {
								fst bool
								snd loop_state
							} {
								_scrut348 := key
								switch _scrut348 {
								case KeySpace:
									gs := (func() game_state {
										_scrut349 := ls
										ls_game0 := _scrut349.ls_game
										return ls_game0
									})()
									return struct {
										fst bool
										snd loop_state
									}{fst: false, snd: mkLoop(gs, (func() position {
										_scrut350 := gs
										pacpos0 := _scrut350.pacpos
										return pacpos0
									})(), (func() list[ghost_state] {
										_scrut351 := gs
										ghosts0 := _scrut351.ghosts
										return ghosts0
									})(), now, (func() uint {
										_scrut352 := ls
										ls_start_time0 := _scrut352.ls_start_time
										return ls_start_time0
									})(), (func() RocqSDLTexture {
										_scrut353 := ls
										ls_texture0 := _scrut353.ls_texture
										return ls_texture0
									})(), Playing, uint(0), false)}
								default:
									return (func() any {
										quit := handle_key_down((key).(sdl_key), (func() game_state {
											_scrut354 := ls
											ls_game0 := _scrut354.ls_game
											return ls_game0
										})()).fst
										_ = handle_key_down((key).(sdl_key), (func() game_state {
											_scrut354 := ls
											ls_game0 := _scrut354.ls_game
											return ls_game0
										})()).snd
										return (func() any {
											if quit {
												return struct {
													fst bool
													snd loop_state
												}{fst: true, snd: ls}
											} else {
												return (func() struct {
													fst bool
													snd loop_state
												} {
													_ = render_paused_frame(ren, (func() RocqSDLTexture {
														_scrut355 := ls
														ls_texture0 := _scrut355.ls_texture
														return ls_texture0
													})(), (func() game_state {
														_scrut356 := ls
														ls_game0 := _scrut356.ls_game
														return ls_game0
													})(), (func() position {
														_scrut357 := ls
														ls_prev_pac0 := _scrut357.ls_prev_pac
														return ls_prev_pac0
													})(), (func() list[ghost_state] {
														_scrut358 := ls
														ls_prev_ghosts0 := _scrut358.ls_prev_ghosts
														return ls_prev_ghosts0
													})(), uint(0), uint((uint(0) + 1)), time_ms)
													return (func() struct {
														fst bool
														snd loop_state
													} {
														_ = frame_delay(now)
														return struct {
															fst bool
															snd loop_state
														}{fst: false, snd: ls}
													})()
												})()
											}
										}()).(struct {
											fst bool
											snd loop_state
										})
									}()).(struct {
										fst bool
										snd loop_state
									})
								}
								panic("unreachable")
							})()
						}
						panic("unreachable")
					})()
				case DeathPause:
					return (func() any {
						if uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) <= sub(now, (func() uint {
							_scrut359 := ls
							ls_phase_time0 := _scrut359.ls_phase_time
							return ls_phase_time0
						})()) {
							return struct {
								fst bool
								snd loop_state
							}{fst: false, snd: mkLoop((func() game_state {
								_scrut360 := ls
								ls_game0 := _scrut360.ls_game
								return ls_game0
							})(), (func() position {
								_scrut361 := (func() game_state {
									_scrut362 := ls
									ls_game0 := _scrut362.ls_game
									return ls_game0
								})()
								pacpos0 := _scrut361.pacpos
								return pacpos0
							})(), (func() list[ghost_state] {
								_scrut363 := (func() game_state {
									_scrut364 := ls
									ls_game0 := _scrut364.ls_game
									return ls_game0
								})()
								ghosts0 := _scrut363.ghosts
								return ghosts0
							})(), now, (func() uint {
								_scrut365 := ls
								ls_start_time0 := _scrut365.ls_start_time
								return ls_start_time0
							})(), (func() RocqSDLTexture {
								_scrut366 := ls
								ls_texture0 := _scrut366.ls_texture
								return ls_texture0
							})(), Playing, uint(0), false)}
						} else {
							return (func() struct {
								fst bool
								snd loop_state
							} {
								_ = draw_message_screen(ren, msg_lives_left((func() uint {
									_scrut367 := (func() game_state {
										_scrut368 := ls
										ls_game0 := _scrut368.ls_game
										return ls_game0
									})()
									lives0 := _scrut367.lives
									return lives0
								})()))
								return (func() struct {
									fst bool
									snd loop_state
								} {
									_ = frame_delay(now)
									return struct {
										fst bool
										snd loop_state
									}{fst: false, snd: ls}
								})()
							})()
						}
					}()).(struct {
						fst bool
						snd loop_state
					})
				case GameOverScreen:
					return (func() any {
						if uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) <= sub(now, (func() uint {
							_scrut369 := ls
							ls_phase_time0 := _scrut369.ls_phase_time
							return ls_phase_time0
						})()) {
							return struct {
								fst bool
								snd loop_state
							}{fst: true, snd: ls}
						} else {
							return (func() struct {
								fst bool
								snd loop_state
							} {
								_ = draw_message_screen(ren, msg_game_over)
								return (func() struct {
									fst bool
									snd loop_state
								} {
									_ = frame_delay(now)
									return struct {
										fst bool
										snd loop_state
									}{fst: false, snd: ls}
								})()
							})()
						}
					}()).(struct {
						fst bool
						snd loop_state
					})
				case WinScreen:
					return (func() any {
						if uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) <= sub(now, (func() uint {
							_scrut370 := ls
							ls_phase_time0 := _scrut370.ls_phase_time
							return ls_phase_time0
						})()) {
							return struct {
								fst bool
								snd loop_state
							}{fst: true, snd: ls}
						} else {
							return (func() struct {
								fst bool
								snd loop_state
							} {
								_ = draw_message_screen(ren, msg_you_win)
								return (func() struct {
									fst bool
									snd loop_state
								} {
									_ = frame_delay(now)
									return struct {
										fst bool
										snd loop_state
									}{fst: false, snd: ls}
								})()
							})()
						}
					}()).(struct {
						fst bool
						snd loop_state
					})
				}
				panic("unreachable")
			})()
		}
		panic("unreachable")
	})()
}

var init_game struct {
	fst struct {
		fst RocqSDLWindow
		snd RocqSDLRenderer
	}
	snd loop_state
} = (func() struct {
	fst struct {
		fst RocqSDLWindow
		snd RocqSDLRenderer
	}
	snd loop_state
} {
	win := rocqSDLCreateWindow("Rocqman", win_width, win_height)
	return (func() struct {
		fst struct {
			fst RocqSDLWindow
			snd RocqSDLRenderer
		}
		snd loop_state
	} {
		ren := rocqSDLCreateRenderer(win)
		return (func() struct {
			fst struct {
				fst RocqSDLWindow
				snd RocqSDLRenderer
			}
			snd loop_state
		} {
			tex := rocqSDLLoadTexture(ren, "assets/rocq.svg")
			return (func() struct {
				fst struct {
					fst RocqSDLWindow
					snd RocqSDLRenderer
				}
				snd loop_state
			} {
				t0 := rocqSDLGetTicks()
				return (func() struct {
					fst struct {
						fst RocqSDLWindow
						snd RocqSDLRenderer
					}
					snd loop_state
				} {
					ls := mkLoop(initial_state, (func() position {
						_scrut371 := initial_state
						pacpos0 := _scrut371.pacpos
						return pacpos0
					})(), (func() list[ghost_state] {
						_scrut372 := initial_state
						ghosts0 := _scrut372.ghosts
						return ghosts0
					})(), t0, t0, tex, Playing, uint(0), false)
					return struct {
						fst struct {
							fst RocqSDLWindow
							snd RocqSDLRenderer
						}
						snd loop_state
					}{fst: struct {
						fst RocqSDLWindow
						snd RocqSDLRenderer
					}{fst: win, snd: ren}, snd: ls}
				})()
			})()
		})()
	})()
})()

func cleanup(a0 RocqSDLRenderer, a1 RocqSDLWindow) struct{} {
	return rocqSDLDestroy(a0, a1)
}

func exit_game(win RocqSDLWindow, ren RocqSDLRenderer) struct{} {
	return cleanup(ren, win)
}

func run_game(win RocqSDLWindow, ren RocqSDLRenderer, ls loop_state) struct{} {
	res := process_frame(ren, ls)
	return (func() any {
		quit := res.fst
		new_ls := res.snd
		return (func() any {
			if quit {
				return exit_game(win, ren)
			} else {
				return run_game(win, ren, new_ls)
			}
		}()).(struct{})
	}()).(struct{})
}
