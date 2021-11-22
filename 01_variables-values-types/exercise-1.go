package main

import "fmt"

func main() {
	x := 42
	y := `James Bond`
	z := true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf(`Integer Value: %v 
String Value: %v 
Boolean Value: %v`, x, y, z)

}
