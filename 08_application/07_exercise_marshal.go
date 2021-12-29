package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Username string `json:"Username"`
	Name     string
	Age      int
}

func main() {

	u1 := user{
		Username: "uf-code",
		Name:     "uf",
		Age:      99,
	}

	u2 := user{
		Username: "Kate",
		Name:     "Katarina",
		Age:      9999,
	}

	users := []user{u1, u2}

	fmt.Println(users)

	// Marshaling slice
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
