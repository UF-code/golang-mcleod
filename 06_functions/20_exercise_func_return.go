package main

import "fmt"

func main() {
	r2 := return2()
	fmt.Printf("Value of r2: %v Type of r2: %T \n", r2, r2)

	r1 := r2()
	fmt.Printf("Value of r1: %v Type of r1: %T \n", r1, r1)

}

func return2() func() int {
	return func() int {
		return 452
	}
}
