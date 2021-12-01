package main

import "fmt"

func main() {
	if true {
		fmt.Println(`01`)
	}
	if false {
		fmt.Println(`Not Gonna Show This One`)
	}
	if !false {
		fmt.Println(`03`)
	}
	if !true {
		fmt.Println(`Not Gonna Show This One`)
	}

	if x := 42; x == 41 {
		fmt.Println(`01`)
	} else {
		fmt.Println(x)
	}
	// fmt.Println(x) that doesn't work cause is out of scope

	fmt.Println("Here's a Statement")

}
