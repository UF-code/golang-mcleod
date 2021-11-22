package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)

	x = true

	fmt.Println(x)

	a := 7
	b := 42
	fmt.Printf(`%v == %v is %t`, a, b, a == b)

}
