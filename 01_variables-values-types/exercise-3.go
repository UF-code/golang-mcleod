package main

import "fmt"

var x int = 42
var y string = `James Bond`
var z bool = true

func main() {
	s := fmt.Sprintf(`Integer Value: %d
String Value: %s
Boolean Value: %t`, x, y, z)

	fmt.Println(s)
}
