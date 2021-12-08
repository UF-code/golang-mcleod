package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Hellooooooo, James"}

	z := [][]string{x, y}

	for i := 0; i < len(z); i++ {
		for j, v := range z[i] {
			fmt.Printf(`Index of Array:%d Index of Item: %d Value of Item: %v
`, i, j, v)
		}
	}
}
