package main

import "fmt"

type hotdog int

var x hotdog

func main() {
	fmt.Printf(`Value: %v
Type: %T
`, x, x)

	x = 42
	fmt.Printf(`New Value: %d
Type: %T
`, x, x)
}
