package gyjson

import (
	"fmt"
	"log"
	"encoding/json"
	"reflect"
)

var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}	
}`

func GYJson01() {
	var instore map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &instore)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println("Name:", instore["name"])
	fmt.Println("Title:", instore["title"])
	fmt.Println("Contract")
	fmt.Println("H:", instore["contact"].(map[string]interface{})["home"])
	fmt.Println("C:", instore["contact"].(map[string]interface{})["cell"])
	fmt.Println("---> ", reflect.TypeOf(instore["contact"].(map[string]interface{})["cell"]))
}