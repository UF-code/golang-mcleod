package main

import "fmt"

// const a = 42
// const b = 42.78
// const c = "James Bond"

const (
	a int     = 42
	b float64 = 42.89
	c         = "James Bond"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
