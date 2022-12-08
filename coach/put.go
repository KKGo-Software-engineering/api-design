package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	v := &User{}
	fmt.Printf("%#v\n", &v)

	// from client
	// PUT /users/1
	put := []byte(`{
			"name": "Nong",
			"age": 22
		}`)

	var u User
	u.ID = 1

	err := json.Unmarshal(put, &u)

	// insert to database
	fmt.Printf("% #v\n", u)
	fmt.Println(err)
}
