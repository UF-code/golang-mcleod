package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did something wrong and here's the error:", err)
	}
}
