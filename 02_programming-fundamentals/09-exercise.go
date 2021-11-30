package main

import "fmt"

const (
	x int = 20
	y     = 20

	z string = `Hello`
	s        = `Hello2`
)

func main() {
	fmt.Printf(`x: %d  y: %d
`, x, y)

	fmt.Printf(`z: %v s: %v
`, z, s)
}
