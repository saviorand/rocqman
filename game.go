package main

var win_width uint = mul(uint(19), uint(32))

var win_height uint = add(mul(uint(15), uint(32)), uint(40))

var W cell = Wall

var D cell = Dot

var P cell = PowerPellet

var initial_board list[list[cell]] = cons(cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, nil_))))))))))))))))))), cons(cons(W, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(W, nil_))))))))))))))))))), cons(cons(W, cons(D, cons(W, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(W, cons(D, cons(W, nil_))))))))))))))))))), cons(cons(W, cons(P, cons(W, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(W, cons(P, cons(W, nil_))))))))))))))))))), cons(cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, nil_))))))))))))))))))), cons(cons(W, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(W, nil_))))))))))))))))))), cons(cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, nil_))))))))))))))))))), cons(cons(W, cons(D, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(D, cons(W, nil_))))))))))))))))))), cons(cons(W, cons(D, cons(W, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(W, cons(D, cons(W, nil_))))))))))))))))))), cons(cons(W, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(W, nil_))))))))))))))))))), cons(cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, nil_))))))))))))))))))), cons(cons(W, cons(P, cons(W, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(W, cons(P, cons(W, nil_))))))))))))))))))), cons(cons(W, cons(D, cons(W, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(D, cons(W, cons(W, cons(W, cons(D, cons(W, cons(W, cons(D, cons(W, nil_))))))))))))))))))), cons(cons(W, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(W, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(D, cons(W, nil_))))))))))))))))))), cons(cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, cons(W, nil_))))))))))))))))))), nil_)))))))))))))))

func count_row(r list[cell]) uint {
	return (func() uint {
		var _box any = r
		_scrut1 := _box.(*listImpl)
		switch _scrut1._v {
		case 0:
			return uint(0)
		case 1:
			c := _scrut1._c1_d0
			rest := _scrut1._c1_d1
			return (func() uint {
				_scrut2 := c
				switch _scrut2 {
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
		_scrut3 := _box.(*listImpl)
		switch _scrut3._v {
		case 0:
			return uint(0)
		case 1:
			row := _scrut3._c1_d0
			rest := _scrut3._c1_d1
			return add(count_row((row).(list[cell])), count_dots((rest).(list[list[cell]])))
		}
		panic("unreachable")
	})()
}

func get_cell(row uint, col uint, b list[list[cell]]) cell {
	return (nth(col, (nth(row, b, nil_)).(list[any]), Wall)).(cell)
}

func replace_nth[T1 any](n uint, l list[T1], x T1) list[T1] {
	return (func() list[T1] {
		var _box any = l
		_scrut4 := _box.(*listImpl)
		switch _scrut4._v {
		case 0:
			return nil_
		case 1:
			h := _scrut4._c1_d0
			t := _scrut4._c1_d1
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
	r := (nth(row, b, nil_)).(list[cell])
	return replace_nth(row, b, replace_nth(col, r, c))
}

func is_wall(row uint, col uint, b list[list[cell]]) bool {
	return (func() bool {
		_scrut5 := get_cell(row, col, b)
		switch _scrut5 {
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
		_scrut6 := d
		switch _scrut6 {
		case Up:
			return mkPos(pred((func() uint {
				_scrut7 := p
				prow := _scrut7.prow
				return prow
			})()), (func() uint {
				_scrut8 := p
				pcol := _scrut8.pcol
				return pcol
			})())
		case Down:
			return mkPos(uint(((func() uint {
				_scrut9 := p
				prow := _scrut9.prow
				return prow
			})() + 1)), (func() uint {
				_scrut10 := p
				pcol := _scrut10.pcol
				return pcol
			})())
		case Left:
			return mkPos((func() uint {
				_scrut11 := p
				prow := _scrut11.prow
				return prow
			})(), pred((func() uint {
				_scrut12 := p
				pcol := _scrut12.pcol
				return pcol
			})()))
		case Right:
			return mkPos((func() uint {
				_scrut13 := p
				prow := _scrut13.prow
				return prow
			})(), uint(((func() uint {
				_scrut14 := p
				pcol := _scrut14.pcol
				return pcol
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
		_scrut15 := new_p
		prow := _scrut15.prow
		return prow
	})(), (func() uint {
		_scrut16 := new_p
		pcol := _scrut16.pcol
		return pcol
	})(), b))
}

func update_ghost_modes(gs list[ghost_state], powered bool) list[ghost_state] {
	return (func() list[ghost_state] {
		var _box any = gs
		_scrut17 := _box.(*listImpl)
		switch _scrut17._v {
		case 0:
			return nil_
		case 1:
			g := _scrut17._c1_d0
			rest := _scrut17._c1_d1
			mode := (func() any {
				if powered {
					return Frightened
				} else {
					return Chase
				}
			}()).(ghost_mode)
			return cons(mkGhost((func() position {
				var _rbox any = g
				_scrut18 := _rbox.(ghost_state)
				gpos := _scrut18.gpos
				return gpos
			})(), (func() direction {
				var _rbox any = g
				_scrut19 := _rbox.(ghost_state)
				gdir := _scrut19.gdir
				return gdir
			})(), mode), update_ghost_modes((rest).(list[ghost_state]), powered))
		}
		panic("unreachable")
	})()
}

func set_direction(d direction, gs game_state) game_state {
	return mkState((func() list[list[cell]] {
		_scrut20 := gs
		board := _scrut20.board
		return board
	})(), (func() position {
		_scrut21 := gs
		pacpos := _scrut21.pacpos
		return pacpos
	})(), (func() direction {
		_scrut22 := gs
		pacdir := _scrut22.pacdir
		return pacdir
	})(), d, (func() list[ghost_state] {
		_scrut23 := gs
		ghosts := _scrut23.ghosts
		return ghosts
	})(), (func() uint {
		_scrut24 := gs
		score := _scrut24.score
		return score
	})(), (func() uint {
		_scrut25 := gs
		lives := _scrut25.lives
		return lives
	})(), (func() uint {
		_scrut26 := gs
		dots_left := _scrut26.dots_left
		return dots_left
	})(), (func() uint {
		_scrut27 := gs
		power_timer := _scrut27.power_timer
		return power_timer
	})(), (func() bool {
		_scrut28 := gs
		game_over := _scrut28.game_over
		return game_over
	})(), (func() bool {
		_scrut29 := gs
		game_won := _scrut29.game_won
		return game_won
	})())
}

func move_pacman(gs game_state) game_state {
	return (func() any {
		if (func() bool {
			_scrut30 := gs
			game_over := _scrut30.game_over
			return game_over
		})() || (func() bool {
			_scrut31 := gs
			game_won := _scrut31.game_won
			return game_won
		})() {
			return gs
		} else {
			return (func() game_state {
				try_dir0 := (func() any {
					if can_move((func() direction {
						_scrut32 := gs
						desired_dir := _scrut32.desired_dir
						return desired_dir
					})(), (func() position {
						_scrut33 := gs
						pacpos := _scrut33.pacpos
						return pacpos
					})(), (func() list[list[cell]] {
						_scrut34 := gs
						board := _scrut34.board
						return board
					})()) {
						return (func() direction {
							_scrut35 := gs
							desired_dir := _scrut35.desired_dir
							return desired_dir
						})()
					} else {
						return (func() direction {
							_scrut36 := gs
							pacdir := _scrut36.pacdir
							return pacdir
						})()
					}
				}()).(direction)
				return (func() game_state {
					new_pos := move_pos(try_dir0, (func() position {
						_scrut37 := gs
						pacpos := _scrut37.pacpos
						return pacpos
					})())
					return (func() any {
						if is_wall((func() uint {
							_scrut38 := new_pos
							prow := _scrut38.prow
							return prow
						})(), (func() uint {
							_scrut39 := new_pos
							pcol := _scrut39.pcol
							return pcol
						})(), (func() list[list[cell]] {
							_scrut40 := gs
							board := _scrut40.board
							return board
						})()) {
							return gs
						} else {
							return (func() game_state {
								cell0 := get_cell((func() uint {
									_scrut41 := new_pos
									prow := _scrut41.prow
									return prow
								})(), (func() uint {
									_scrut42 := new_pos
									pcol := _scrut42.pcol
									return pcol
								})(), (func() list[list[cell]] {
									_scrut43 := gs
									board := _scrut43.board
									return board
								})())
								return (func() game_state {
									new_board := (func() list[list[cell]] {
										_scrut44 := cell0
										switch _scrut44 {
										case Wall:
											return (func() list[list[cell]] {
												_scrut45 := gs
												board := _scrut45.board
												return board
											})()
										case Empty:
											return (func() list[list[cell]] {
												_scrut46 := gs
												board := _scrut46.board
												return board
											})()
										default:
											return set_cell((func() uint {
												_scrut47 := new_pos
												prow := _scrut47.prow
												return prow
											})(), (func() uint {
												_scrut48 := new_pos
												pcol := _scrut48.pcol
												return pcol
											})(), Empty, (func() list[list[cell]] {
												_scrut49 := gs
												board := _scrut49.board
												return board
											})())
										}
										panic("unreachable")
									})()
									return (func() game_state {
										add_score := (func() uint {
											_scrut50 := cell0
											switch _scrut50 {
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
												_scrut51 := cell0
												switch _scrut51 {
												case Wall:
													return (func() uint {
														_scrut52 := gs
														dots_left := _scrut52.dots_left
														return dots_left
													})()
												case Empty:
													return (func() uint {
														_scrut53 := gs
														dots_left := _scrut53.dots_left
														return dots_left
													})()
												default:
													return pred((func() uint {
														_scrut54 := gs
														dots_left := _scrut54.dots_left
														return dots_left
													})())
												}
												panic("unreachable")
											})()
											return (func() game_state {
												new_power := (func() uint {
													_scrut55 := cell0
													switch _scrut55 {
													case PowerPellet:
														return uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1))
													default:
														return (func() uint {
															_scrut56 := gs
															power_timer := _scrut56.power_timer
															return power_timer
														})()
													}
													panic("unreachable")
												})()
												return (func() game_state {
													new_ghosts := (func() list[ghost_state] {
														_scrut57 := cell0
														switch _scrut57 {
														case PowerPellet:
															return update_ghost_modes((func() list[ghost_state] {
																_scrut58 := gs
																ghosts := _scrut58.ghosts
																return ghosts
															})(), true)
														default:
															return (func() list[ghost_state] {
																_scrut59 := gs
																ghosts := _scrut59.ghosts
																return ghosts
															})()
														}
														panic("unreachable")
													})()
													return (func() game_state {
														won := eqb(new_dots, uint(0))
														return mkState(new_board, new_pos, try_dir0, (func() direction {
															_scrut60 := gs
															desired_dir := _scrut60.desired_dir
															return desired_dir
														})(), new_ghosts, add((func() uint {
															_scrut61 := gs
															score := _scrut61.score
															return score
														})(), add_score), (func() uint {
															_scrut62 := gs
															lives := _scrut62.lives
															return lives
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
		if leb(a, b) {
			return sub(b, a)
		} else {
			return sub(a, b)
		}
	}()).(uint)
}

func manhattan(p1 position, p2 position) uint {
	return add(abs_diff((func() uint {
		_scrut63 := p1
		prow := _scrut63.prow
		return prow
	})(), (func() uint {
		_scrut64 := p2
		prow := _scrut64.prow
		return prow
	})()), abs_diff((func() uint {
		_scrut65 := p1
		pcol := _scrut65.pcol
		return pcol
	})(), (func() uint {
		_scrut66 := p2
		pcol := _scrut66.pcol
		return pcol
	})()))
}

func is_opposite(d1 direction, d2 direction) bool {
	return (func() bool {
		_scrut67 := d1
		switch _scrut67 {
		case Up:
			return (func() bool {
				_scrut68 := d2
				switch _scrut68 {
				case Down:
					return true
				default:
					return false
				}
				panic("unreachable")
			})()
		case Down:
			return (func() bool {
				_scrut69 := d2
				switch _scrut69 {
				case Up:
					return true
				default:
					return false
				}
				panic("unreachable")
			})()
		case Left:
			return (func() bool {
				_scrut70 := d2
				switch _scrut70 {
				case Right:
					return true
				default:
					return false
				}
				panic("unreachable")
			})()
		case Right:
			return (func() bool {
				_scrut71 := d2
				switch _scrut71 {
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
			_scrut72 := g
			gpos := _scrut72.gpos
			return gpos
		})(), b)) {
			return struct {
				fst direction
				snd uint
			}{fst: best_d, snd: best_dist}
		} else {
			return (func() any {
				if is_opposite(d, (func() direction {
					_scrut73 := g
					gdir := _scrut73.gdir
					return gdir
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
							_scrut74 := g
							gpos := _scrut74.gpos
							return gpos
						})())
						return (func() struct {
							fst direction
							snd uint
						} {
							dist := manhattan(new_pos, target)
							return (func() any {
								if ltb(dist, best_dist) {
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
		_scrut75 := (func() ghost_mode {
			_scrut76 := g
			gmode := _scrut76.gmode
			return gmode
		})()
		switch _scrut75 {
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
			_scrut77 := g
			gpos := _scrut77.gpos
			return gpos
		})(), b) {
			return mkGhost(move_pos(dir, (func() position {
				_scrut78 := g
				gpos := _scrut78.gpos
				return gpos
			})()), dir, (func() ghost_mode {
				_scrut79 := g
				gmode := _scrut79.gmode
				return gmode
			})())
		} else {
			return g
		}
	}()).(ghost_state)
}

func move_ghosts_list(gs list[ghost_state], pac position, b list[list[cell]]) list[ghost_state] {
	return (func() list[ghost_state] {
		var _box any = gs
		_scrut80 := _box.(*listImpl)
		switch _scrut80._v {
		case 0:
			return nil_
		case 1:
			g := _scrut80._c1_d0
			rest := _scrut80._c1_d1
			return cons(move_one_ghost((g).(ghost_state), pac, b), move_ghosts_list((rest).(list[ghost_state]), pac, b))
		}
		panic("unreachable")
	})()
}

func move_ghosts(gs game_state) game_state {
	return mkState((func() list[list[cell]] {
		_scrut81 := gs
		board := _scrut81.board
		return board
	})(), (func() position {
		_scrut82 := gs
		pacpos := _scrut82.pacpos
		return pacpos
	})(), (func() direction {
		_scrut83 := gs
		pacdir := _scrut83.pacdir
		return pacdir
	})(), (func() direction {
		_scrut84 := gs
		desired_dir := _scrut84.desired_dir
		return desired_dir
	})(), move_ghosts_list((func() list[ghost_state] {
		_scrut85 := gs
		ghosts := _scrut85.ghosts
		return ghosts
	})(), (func() position {
		_scrut86 := gs
		pacpos := _scrut86.pacpos
		return pacpos
	})(), (func() list[list[cell]] {
		_scrut87 := gs
		board := _scrut87.board
		return board
	})()), (func() uint {
		_scrut88 := gs
		score := _scrut88.score
		return score
	})(), (func() uint {
		_scrut89 := gs
		lives := _scrut89.lives
		return lives
	})(), (func() uint {
		_scrut90 := gs
		dots_left := _scrut90.dots_left
		return dots_left
	})(), (func() uint {
		_scrut91 := gs
		power_timer := _scrut91.power_timer
		return power_timer
	})(), (func() bool {
		_scrut92 := gs
		game_over := _scrut92.game_over
		return game_over
	})(), (func() bool {
		_scrut93 := gs
		game_won := _scrut93.game_won
		return game_won
	})())
}

func respawn_ghost(_ ghost_state) ghost_state {
	return mkGhost(mkPos(uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))), DirNone, Chase)
}

func tick_power(gs game_state) game_state {
	return (func() any {
		var _sni_ any = (func() uint {
			_scrut94 := gs
			power_timer := _scrut94.power_timer
			return power_timer
		})()
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return gs
		} else {
			n := _su_ - 1
			return (func() game_state {
				powered := !(eqb(n, uint(0)))
				return (func() game_state {
					new_ghosts := update_ghost_modes((func() list[ghost_state] {
						_scrut95 := gs
						ghosts := _scrut95.ghosts
						return ghosts
					})(), powered)
					return mkState((func() list[list[cell]] {
						_scrut96 := gs
						board := _scrut96.board
						return board
					})(), (func() position {
						_scrut97 := gs
						pacpos := _scrut97.pacpos
						return pacpos
					})(), (func() direction {
						_scrut98 := gs
						pacdir := _scrut98.pacdir
						return pacdir
					})(), (func() direction {
						_scrut99 := gs
						desired_dir := _scrut99.desired_dir
						return desired_dir
					})(), new_ghosts, (func() uint {
						_scrut100 := gs
						score := _scrut100.score
						return score
					})(), (func() uint {
						_scrut101 := gs
						lives := _scrut101.lives
						return lives
					})(), (func() uint {
						_scrut102 := gs
						dots_left := _scrut102.dots_left
						return dots_left
					})(), n, (func() bool {
						_scrut103 := gs
						game_over := _scrut103.game_over
						return game_over
					})(), (func() bool {
						_scrut104 := gs
						game_won := _scrut104.game_won
						return game_won
					})())
				})()
			})()
		}
	}()).(game_state)
}

func tick(gs game_state) game_state {
	return (func() any {
		if (func() bool {
			_scrut105 := gs
			game_over := _scrut105.game_over
			return game_over
		})() || (func() bool {
			_scrut106 := gs
			game_won := _scrut106.game_won
			return game_won
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

var initial_ghosts list[ghost_state] = cons(mkGhost(mkPos(uint((uint(0)+1)), uint((uint(0)+1))), DirNone, Chase), cons(mkGhost(mkPos(uint((uint(0)+1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))), DirNone, Chase), cons(mkGhost(mkPos(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), uint((uint(0)+1))), DirNone, Chase), cons(mkGhost(mkPos(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))+1))), DirNone, Chase), nil_))))

var initial_state game_state = (func() game_state {
	b := set_cell(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)) + 1)), Empty, initial_board)
	return mkState(b, mkPos(uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1)), uint((uint((uint((uint((uint((uint((uint((uint((uint((uint(0)+1))+1))+1))+1))+1))+1))+1))+1))+1))), DirNone, DirNone, initial_ghosts, uint(0), uint((uint((uint((uint(0) + 1)) + 1)) + 1)), count_dots(b), uint(0), false, false)
})()
