package main

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
