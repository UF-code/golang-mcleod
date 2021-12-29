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
	s := `[{"Username":"uf-code","Name":"uf","Age":99},{"Username":"Kate","Name":"Katarina","Age":9999}]`
	bs := []byte(s)

	var users []user
	err := json.Unmarshal(bs, &users)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range users {
		fmt.Println("---User Number", i)
		fmt.Println(v.Username, v.Name, v.Age)
	}

}
