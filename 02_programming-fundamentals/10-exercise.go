package main

import "fmt"

func main() {
	x := 42
	fmt.Printf(`Decimal: %d  Binary: %b  Hexadecimal: %#x
`, x, x, x)

	y := x >> 1
	fmt.Printf(`Decimal %d  Binary: %b  Hexadecimal: %#x
`, y, y, y)

	z := x << 1
	fmt.Printf(`Decimal: %d Binary: %b  Hexadecimal: %#x
`, z, z, z)
}
