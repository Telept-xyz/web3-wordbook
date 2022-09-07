package main

import (
	"fmt"
	"os"
)

// @title web3-wordbook validator
// @version 1.0
// @description verify the word uniqueness
func main() {
	
	// verify the word list to make sure no duplicate word
	var wordIndex map[string]interface{}
	dat, err := os.ReadFile("../words.json")
    check(err)
    fmt.Print(string(dat))
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		panic(err)
	}







	var wordBooks map[string]interface{}

	dat, err := os.ReadFile("/tmp/dat")
    check(err)
    fmt.Print(string(dat))
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("text =", data["text"])
	fmt.Println("number =", data["number"])
	fmt.Println("floats =", data["floats"])
	fmt.Println("innermap =", data["innermap"])
	
	innermap, ok := data["innermap"].(map[string]interface{})
	if !ok {
		panic("inner map is not a map!")
	}
	fmt.Println("innermap.foo =", innermap["foo"])
	fmt.Println("innermap.bar =", innermap["bar"])
	
	fmt.Println("The whole map:", data)

}