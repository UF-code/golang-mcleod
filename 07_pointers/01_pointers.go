package main

import "fmt"

func main() {
	a := 42
	fmt.Println("Value of a", a)
	fmt.Println("Location of a", &a)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", &a)
	fmt.Println(*&a)

	// var b *int = &a
	b := &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 42
	fmt.Println(a)
}
