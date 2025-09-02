package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name": "John Wick", "Age": 27}`

	var data User

	var err = json.Unmarshal([]byte(jsonString), &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age  :", data.Age)

	var object = []User{{"John Wick", 27}, {"Ethan Hunt", 32}}
	jsonData, err := json.Marshal(object)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonString2 := string(jsonData)
	fmt.Println(jsonString2)
}
