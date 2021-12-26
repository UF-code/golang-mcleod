package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("Hey function expression")
	}

	f()

	f1 := func(x int) {
		fmt.Println("Value of X is ", x)
	}

	f1(12)

}
