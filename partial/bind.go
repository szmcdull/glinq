package partial

func Bind[T1 any](f func(T1), t1 T1) func() {
	return func() {
		f(t1)
	}
}

func Bindr[T1, R any](f func(T1) R, t1 T1) func() R {
	return func() R {
		return f(t1)
	}
}

func Bindr2[T1, R, R2 any](f func(T1) (R, R2), t1 T1) func() (R, R2) {
	return func() (R, R2) {
		return f(t1)
	}
}

func Bind2_1[T1, T2 any](f func(T1, T2), t1 T1) func(T2) {
	return func(t2 T2) {
		f(t1, t2)
	}
}

func Bind2_2[T1, T2 any](f func(T1, T2), t2 T2) func(T1) {
	return func(t1 T1) {
		f(t1, t2)
	}
}

func Bind2_1r[T1, T2, R any](f func(T1, T2) R, t1 T1) func(T2) R {
	return func(t2 T2) R {
		return f(t1, t2)
	}
}

func Bind2_2r[T1, T2, R any](f func(T1, T2) R, t2 T2) func(T1) R {
	return func(t1 T1) R {
		return f(t1, t2)
	}
}

func Bind2_1r2[T1, T2, R, R2 any](f func(T1, T2) (R, R2), t1 T1) func(T2) (R, R2) {
	return func(t2 T2) (R, R2) {
		return f(t1, t2)
	}
}

func Bind2_2r2[T1, T2, R, R2 any](f func(T1, T2) (R, R2), t2 T2) func(T1) (R, R2) {
	return func(t1 T1) (R, R2) {
		return f(t1, t2)
	}
}

func Bind3_1[T1, T2, T3 any](f func(T1, T2, T3), t1 T1) func(T2, T3) {
	return func(t2 T2, t3 T3) {
		f(t1, t2, t3)
	}
}

func Bind3_2[T1, T2, T3 any](f func(T1, T2, T3), t2 T2) func(T1, T3) {
	return func(t1 T1, t3 T3) {
		f(t1, t2, t3)
	}
}

func Bind3_3[T1, T2, T3 any](f func(T1, T2, T3), t3 T3) func(T1, T2) {
	return func(t1 T1, t2 T2) {
		f(t1, t2, t3)
	}
}

func Bind3_1r[T1, T2, T3, R any](f func(T1, T2, T3) R, t1 T1) func(T2, T3) R {
	return func(t2 T2, t3 T3) R {
		return f(t1, t2, t3)
	}
}

func Bind3_2r[T1, T2, T3, R any](f func(T1, T2, T3) R, t2 T2) func(T1, T3) R {
	return func(t1 T1, t3 T3) R {
		return f(t1, t2, t3)
	}
}

func Bind3_3r[T1, T2, T3, R any](f func(T1, T2, T3) R, t3 T3) func(T1, T2) R {
	return func(t1 T1, t2 T2) R {
		return f(t1, t2, t3)
	}
}

func Bind3_1r2[T1, T2, T3, R, R2 any](f func(T1, T2, T3) (R, R2), t1 T1) func(T2, T3) (R, R2) {
	return func(t2 T2, t3 T3) (R, R2) {
		return f(t1, t2, t3)
	}
}

func Bind3_2r2[T1, T2, T3, R, R2 any](f func(T1, T2, T3) (R, R2), t2 T2) func(T1, T3) (R, R2) {
	return func(t1 T1, t3 T3) (R, R2) {
		return f(t1, t2, t3)
	}
}

func Bind3_3r2[T1, T2, T3, R, R2 any](f func(T1, T2, T3) (R, R2), t3 T3) func(T1, T2) (R, R2) {
	return func(t1 T1, t2 T2) (R, R2) {
		return f(t1, t2, t3)
	}
}

func Bind4_1[T1, T2, T3, T4 any](f func(T1, T2, T3, T4), t1 T1) func(T2, T3, T4) {
	return func(t2 T2, t3 T3, t4 T4) {
		f(t1, t2, t3, t4)
	}
}

func Bind4_2[T1, T2, T3, T4 any](f func(T1, T2, T3, T4), t2 T2) func(T1, T3, T4) {
	return func(t1 T1, t3 T3, t4 T4) {
		f(t1, t2, t3, t4)
	}
}

func Bind4_3[T1, T2, T3, T4 any](f func(T1, T2, T3, T4), t3 T3) func(T1, T2, T4) {
	return func(t1 T1, t2 T2, t4 T4) {
		f(t1, t2, t3, t4)
	}
}

func Bind4_4[T1, T2, T3, T4 any](f func(T1, T2, T3, T4), t4 T4) func(T1, T2, T3) {
	return func(t1 T1, t2 T2, t3 T3) {
		f(t1, t2, t3, t4)
	}
}

func Bind4_1r[T1, T2, T3, T4, R any](f func(T1, T2, T3, T4) R, t1 T1) func(T2, T3, T4) R {
	return func(t2 T2, t3 T3, t4 T4) R {
		return f(t1, t2, t3, t4)
	}
}

func Bind4_2r[T1, T2, T3, T4, R any](f func(T1, T2, T3, T4) R, t2 T2) func(T1, T3, T4) R {
	return func(t1 T1, t3 T3, t4 T4) R {
		return f(t1, t2, t3, t4)
	}
}

func Bind4_3r[T1, T2, T3, T4, R any](f func(T1, T2, T3, T4) R, t3 T3) func(T1, T2, T4) R {
	return func(t1 T1, t2 T2, t4 T4) R {
		return f(t1, t2, t3, t4)
	}
}

func Bind4_4r[T1, T2, T3, T4, R any](f func(T1, T2, T3, T4) R, t4 T4) func(T1, T2, T3) R {
	return func(t1 T1, t2 T2, t3 T3) R {
		return f(t1, t2, t3, t4)
	}
}

func Bind4_1r2[T1, T2, T3, T4, R, R2 any](f func(T1, T2, T3, T4) (R, R2), t1 T1) func(T2, T3, T4) (R, R2) {
	return func(t2 T2, t3 T3, t4 T4) (R, R2) {
		return f(t1, t2, t3, t4)
	}
}

func Bind4_2r2[T1, T2, T3, T4, R, R2 any](f func(T1, T2, T3, T4) (R, R2), t2 T2) func(T1, T3, T4) (R, R2) {
	return func(t1 T1, t3 T3, t4 T4) (R, R2) {
		return f(t1, t2, t3, t4)
	}
}

func Bind4_3r2[T1, T2, T3, T4, R, R2 any](f func(T1, T2, T3, T4) (R, R2), t3 T3) func(T1, T2, T4) (R, R2) {
	return func(t1 T1, t2 T2, t4 T4) (R, R2) {
		return f(t1, t2, t3, t4)
	}
}

func Bind4_4r2[T1, T2, T3, T4, R, R2 any](f func(T1, T2, T3, T4) (R, R2), t4 T4) func(T1, T2, T3) (R, R2) {
	return func(t1 T1, t2 T2, t3 T3) (R, R2) {
		return f(t1, t2, t3, t4)
	}
}
