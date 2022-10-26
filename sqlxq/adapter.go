package sqlxq

import (
	"io"
	"runtime"

	"github.com/jmoiron/sqlx"
	"github.com/szmcdull/glinq"
)

type (
	RowsEnumerator[T any] struct {
		*sqlx.Rows
		peeked  bool
		any     bool
		pointer bool
		current T
	}
)

func Queryx[T any](db sqlx.Queryer, query string, args ...any) (glinq.IEnumerable[T], error) {
	rows, err := db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	result := &RowsEnumerator[T]{
		Rows: rows,
	}
	runtime.SetFinalizer(result, func(me *RowsEnumerator[T]) { me.Rows.Close() })

	return result, nil
}

// QueryxP returns IEnumerable[*T], which is faster when T is a large struct,
// because it avoids copying the struct in LINQ operations.
func QueryxP[T any](db sqlx.Queryer, query string, args ...any) (glinq.IEnumerable[*T], error) {
	rows, err := db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	result := &RowsEnumerator[*T]{
		Rows:    rows,
		pointer: true,
		current: new(T),
	}
	runtime.SetFinalizer(result, func(me *RowsEnumerator[*T]) { me.Rows.Close() })

	return result, nil
}

// IEnumerable[T]

func (me *RowsEnumerator[T]) GetEnumerator() glinq.IEnumerator[T] {
	return me
}

func (me *RowsEnumerator[T]) Count() int {
	return -1
}

func (me *RowsEnumerator[T]) Any() bool {
	if me.any {
		return true
	}
	if me.peeked {
		return me.any
	}
	me.peeked = true
	me.any = me.Rows.Next()
	return me.any
}

// IEnumerator[T]

func (me *RowsEnumerator[T]) MoveNext() (err error) {
	if !me.peeked {
		me.peeked = true
		any := me.Rows.Next()
		me.any = me.any || any
		if !any {
			return io.EOF
		}
	}
	if !me.any {
		return io.EOF
	}

	me.peeked = false
	if me.pointer {
		err = me.Rows.StructScan(me.current)
	} else {
		err = me.Rows.StructScan(&me.current)
	}
	return
}

func (me *RowsEnumerator[T]) Current() (result T) {
	return me.current
}
