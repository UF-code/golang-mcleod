package main

import "fmt"

func main() {
	s := `What's up "DUDE"`
	fmt.Printf(`Value: %v Type: %T
`, s, s)

	//
	// ascii codes
	byte_s := []byte(s)
	fmt.Printf(`Value: %v Type: %T
`, byte_s, byte_s)

	for index, char := range s {
		// fmt.Println(i, " => ", string(char))

		fmt.Printf(`Index: %d  Value: %v  ASCII: %v  UTF-8: %U  Hex: %#x
`, index, string(char), char, char, char)
	}
}
