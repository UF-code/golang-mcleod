package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr)

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	fmt.Println(arr)
	fmt.Printf(`%T`, arr)
}
