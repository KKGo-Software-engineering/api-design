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
	// query from database
	u := User{
		ID:   1,
		Name: "Aorjoa",
	}

	fmt.Printf("from db: % #v\n", u)

	// from client
	patch := []byte(`{
		"age": 19
	}`)

	err := json.Unmarshal(patch, &u)

	fmt.Printf("patched: % #v\n", u)
	fmt.Println(err)
}
