package main

import "fmt"

type hotdog int

var x hotdog
var y int

func main() {
	fmt.Printf(`
Value: %v
Type: %T`, x, x)

	x = 42
	fmt.Printf(`
Value: %v
Type: %T`, x, x)

	y = int(x)
	fmt.Printf(`
Value: %v
Type: %T`, y, y)
}
