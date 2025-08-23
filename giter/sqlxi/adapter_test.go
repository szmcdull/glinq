package sqlxi

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/szmcdull/glinq/giter"
)

func ExampleQueryx() {
	type User struct {
		ID   int
		Name string
		Age  int
	}
	var db *sqlx.DB

	err := giter.TryEach(
		Queryx[User](db, "SELECT * FROM users WHERE age > ?", 18),
		func(user User) bool {
			fmt.Println(user.Name)
			return true
		})
	if err != nil {
		println(err.Error())
	}
}
