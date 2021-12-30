package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here is the error %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "need more wine with some lemon ",
	}

	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran - ", e, "\n", e.(customErr).info)
}
