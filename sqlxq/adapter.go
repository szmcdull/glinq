package sqlxq

import (
	"io"
	"runtime"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/szmcdull/glinq"
)

type (
	RowsEnumerator[T any] struct {
		*sqlx.Rows
		peeked  bool
		any     bool
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

func (me *RowsEnumerator[T]) MoveNext() error {
	if !me.peeked {
		me.peeked = true
		any := me.Rows.Next()
		me.any = me.any || any
	}
	if !me.any {
		return io.EOF
	}

	me.peeked = false
	err := me.Rows.StructScan(&me.current)
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), `Rows are closed`) {
		return io.EOF
	}
	return err
}

func (me *RowsEnumerator[T]) Current() (result T) {
	return me.current
}
