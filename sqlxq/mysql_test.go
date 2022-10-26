package sqlxq

import (
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/szmcdull/glinq"
)

type acts_trades struct {
	Exchange string
	Symbol   string
	Trade_id string
	Time     int64
	Price    float64
	Amount   float64
	Side     string
}

var (
	connstr = `szb02:*****@tcp(47.115.143.171:3306)/szb02?charset=utf8&loc=Asia%2FShanghai&parseTime=true`
)

func Test2(t *testing.T) {
	db, err := sqlx.Open(`mysql`, connstr)
	if err != nil {
		t.Fatal(err)
	}
	rows, err := Queryx[acts_trades](db, `SELECT * FROM acts_trades`)
	if err != nil {
		t.Fatal(err)
	}
	q := glinq.Select(
		glinq.Where(rows,
			func(x acts_trades) bool {
				return x.Side == `sell`
			}),
		func(x acts_trades) string {
			return time.Unix(x.Time, 0).Format(time.RFC3339)
		})
	sl, err := glinq.ToSlice(q)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf(`%+v`, sl)
}

func TestPointer(t *testing.T) {
	db, err := sqlx.Open(`mysql`, connstr)
	if err != nil {
		t.Fatal(err)
	}
	rows, err := QueryxP[acts_trades](db, `SELECT * FROM acts_trades`)
	if err != nil {
		t.Fatal(err)
	}
	q := glinq.Select(
		glinq.Where(rows,
			func(x *acts_trades) bool {
				return x.Side == `sell`
			}),
		func(x *acts_trades) string {
			return time.Unix(x.Time, 0).Format(time.RFC3339)
		})
	sl, err := glinq.ToSlice(q)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf(`%+v`, sl)
}
