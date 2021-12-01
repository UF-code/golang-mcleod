package main

import "fmt"

func main() {

	//
	// for  init; condition; post {}
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	//
	// for index, item := range string, list
	for i, s := range "Hello World" {
		fmt.Println(i, string(s))
	}
	for i, n := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0} {
		fmt.Println(i, n)
	}

	//
	// nested for loop
	for i := 0; i <= 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
		}
	}

	//
	// while loop
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}

	//
	// infinite loop
	y := 1
	for {
		if y > 9 {
			break
		}
		fmt.Println(y)
		y++
	}

}
