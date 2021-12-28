package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"` // for mapping
	Last  string
	Age   int
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)
	fmt.Println(s)
	fmt.Printf("%T \n", s)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)

	// people := []person{}
	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range people {
		fmt.Println("---- Person Number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
