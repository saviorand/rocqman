package main

func nth(n any, l any, default0 any) any {
	return func() any {
		var _sni_ any = n
		_su_ := _sni_.(uint)
		if _su_ == 0 {
			return (func() any {
				var _box any = l
				_scrut1 := _box.(*listImpl)
				switch _scrut1._v {
				case 0:
					return default0
				case 1:
					x := _scrut1._c1_d0
					return x
				}
				panic("unreachable")
			})()
		} else {
			m := _su_ - 1
			return (func() any {
				var _box any = l
				_scrut2 := _box.(*listImpl)
				switch _scrut2._v {
				case 0:
					return default0
				case 1:
					l_ := _scrut2._c1_d1
					return nth(m, (l_).(list[any]), default0)
				}
				panic("unreachable")
			})()
		}
	}()
}
