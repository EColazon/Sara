package gyjson

import (
	"fmt"
	"log"
	"encoding/json"
	"reflect"
	"time"
)
func ChanJson() {
	chJson := make(chan map[string]interface{}, 1024)
	mapJson := make(map[string]interface{})
	go func() {
		for i := 0; i < 10; i++ {
			chRecv := <- chJson
			fmt.Println("---> chRecv: ", chRecv["id"], chRecv["data"])
	}
		
	}()
	for i := 0; i < 10; i++ {
			mapJson["id"] = i
			mapJson["data"] = i*i
			chJson <- mapJson
	}
	fmt.Println("---> Send Done.")
	// select {
	// case msg := <- chJson:
	// 	fmt.Println("---> ", msg)
	// default:
	// 	time.Sleep(2*time.Second)
	// }
	
	time.Sleep(2*time.Second)
	
}

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
