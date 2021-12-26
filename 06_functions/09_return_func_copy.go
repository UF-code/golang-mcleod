package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)

	// x is function integer
	// x ==>> return func() int{ return 42 }
	x := bar()
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// x1 is integer
	// x1 ==>> return 42
	x1 := x()
	fmt.Println(x1)
	fmt.Printf("%T\n", x1)

	// bar()() = return 42
	z := bar()()
	fmt.Println(z)
	fmt.Printf("%T", z)

}

func foo() string {
	s := "Hello World"
	return s
}

func bar() func() int {
	return func() int {
		return 42
	}
}
