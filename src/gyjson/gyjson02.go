package gyjson

import (
	"fmt"
	"log"
	"encoding/json"
	"reflect"
)

func GYJson02() {
	instore := make(map[string]interface{})
	instore["name"] = "Gopher"
	instore["title"] = "programmer"
	instore["contract"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	data, err := json.MarshalIndent(instore, "", "    ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println("---> ", reflect.TypeOf(data))
	fmt.Println(string(data))
}
