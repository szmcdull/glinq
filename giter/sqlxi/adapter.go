package sqlxi

import (
	"iter"

	"github.com/jmoiron/sqlx"
)

func Queryx[T any](db sqlx.Queryer, query string, args ...any) iter.Seq[T] {
	rows, err := db.Queryx(query, args...)
	if err != nil {
		panic(err)
	}
	return func(yield func(T) bool) {
		for rows.Next() {
			var current T
			if err := rows.StructScan(&current); err != nil {
				panic(err)
			}
			if !yield(current) {
				break
			}
		}
	}
}

// QueryxP returns IEnumerable[*T], which is faster when T is a large struct,
// because it avoids copying the struct in LINQ operations.
func QueryxP[T any](db sqlx.Queryer, query string, args ...any) iter.Seq[*T] {
	rows, err := db.Queryx(query, args...)
	if err != nil {
		panic(err)
	}
	return func(yield func(*T) bool) {
		for rows.Next() {
			var current T
			if err := rows.StructScan(&current); err != nil {
				panic(err)
			}
			if !yield(&current) {
				break
			}
		}
	}
}
