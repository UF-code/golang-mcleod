package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("Hey")
	}

	f()

	f1 := func(x int) {
		fmt.Println("Value", x)
	}

	f1(42)
}
