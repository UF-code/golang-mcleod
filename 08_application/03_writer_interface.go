package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hey there!")
	fmt.Fprintln(os.Stdout, "Hey there!")
	io.WriteString(os.Stdout, "Hey there!")
}
