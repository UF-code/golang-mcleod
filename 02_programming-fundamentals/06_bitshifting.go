package main

import "fmt"

const (
	_ = iota
	// kb = 1024
	kb = 1 << (iota * 10)
	// mb = 1024 * kb
	mb = 1 << (iota * 10)
	// gb = 1024 * mb
	gb = 1 << (iota * 10)
)

func main() {

	fmt.Printf(`%d			%b
`, kb, kb)
	fmt.Printf(`%d			%b
`, mb, mb)
	fmt.Printf(`%d		%b
`, gb, gb)

	//
	//
	// 	x := 4

	// 	fmt.Printf(`%d		%b
	// `, x, x)

	// 	y := x << 1
	// 	fmt.Printf(`%d		%b
	// `, y, y)
}
