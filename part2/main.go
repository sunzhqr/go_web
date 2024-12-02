package main

import (
	"encoding/json"
	"fmt"
)

var Data = `
	{
		"id":1,
		"age":19,
		"badge": [
			{
				"leetcode": "yes",
				 "db" : "yes"
			}, 
			{
				"leetcode": "no", 
				"db" : "yes"
			}
		]
	}
`

type DataStruct struct {
	Id    uint    `json:"id"`
	Age   uint    `json:"age"`
	Badge []Badge `json:"badge"`
	Brad  bool
}

type Badge struct {
	Leetcode string `json:"leetcode"`
	Db       string `json:"db"`
}

type User struct {
	Id   int    `json:"identification"`
	Name string `json:"username"`
}

func main() {
	user1 := User{10, "Sanzhar"}
	output, _ := json.MarshalIndent(user1, "", "    ")
	fmt.Println(string(output))

	var data DataStruct
	err := json.Unmarshal([]byte(Data), &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
