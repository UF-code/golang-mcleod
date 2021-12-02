package main

import "fmt"

func main() {
	for i, c := range []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		fmt.Printf(`Index: %v Value: %v ASCII: %v UTF-8: %U
`, i, string(c), c, c)
	}

}
