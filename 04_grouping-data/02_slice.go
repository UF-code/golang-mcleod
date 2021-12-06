package main

import "fmt"

func main() {
	// x := type{values} // composit literal
	x := []int{4, 5, 7, 8, 42}

	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x)

	//
	// for
	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	//
	// slicing a slice
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	//
	// append
	x = append(x, 88, 99, 00)
	fmt.Println(x)

	y := []int{234, 456, 678, 890}
	fmt.Println(y)
	x = append(x, y...)
	fmt.Println(x)

	//
	// delete
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
