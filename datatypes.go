package main

type listImpl struct {
	_v     int
	_c1_d0 any
	_c1_d1 any
}

type list[T1 any] = *listImpl

var nil_ *listImpl = &listImpl{_v: 0}

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
