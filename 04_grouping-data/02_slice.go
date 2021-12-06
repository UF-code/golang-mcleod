package main

import "fmt"

func main() {
	// x := type{values} // composit literal
	x := []int{4, 5, 7, 8, 42}

	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}
}