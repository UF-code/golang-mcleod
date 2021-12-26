package main

import "fmt"

func main() {
	fi := foo()
	fmt.Println(fi)

	bi, bs := bar()
	fmt.Println(bi, bs)
}

func foo() int {
	return 42
}
func bar() (int, string) {
	return 42, "Whats up"
}
