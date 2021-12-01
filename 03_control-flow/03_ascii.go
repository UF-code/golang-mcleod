package main

import "fmt"

func main() {
	for i, c := range "Ugur Firat" {
		fmt.Printf(`Index: %v Ascii: %v Value: %v
`, i, c, string(c))
	}
}
