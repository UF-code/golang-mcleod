package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f)

	x := bar()
	fmt.Printf("%T \n", x)

	fmt.Println(x())
}

func foo() string {
	s := "Hello World!"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}
