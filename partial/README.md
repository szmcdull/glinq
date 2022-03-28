# Partial function argument binders

## Usage:

`Bind[x_y][r[z]](Func, ArgToBind)` where:
- x - number of args to Func
- y - the yth number of arg of Func to bind. y <= x
- z - number of return values of Func

Currently supports binding only 1 argument, to a function of maximum 4 arguments and 2 return values.

For example:

```go
func Add(a, b int) int {
    return a + b
}

func main() {
    AddTo5 := Bind2_1r(Add, 5)  // binds 5 to the 1st argument of Add(). Add() has 2 arguments and a return value
    fmt.Println(AddTo5(10))     // prints 15
}
```

Will consider implementing more flexible forms like Bind__xxr, where every _ is an arguments to bind and every x to bypass.