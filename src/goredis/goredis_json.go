package goredis

import (
	"fmt"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

func GoRedisJson() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("---> Connect to redis error", err)
		return
	}
	defer conn.Close()

	key := "profile"
	imap := map[string]string{"username": "Sara", "addr": "Zhengzhou"}
	value, _ := json.Marshal(imap)

	n, err := conn.Do("SETNX", key, value)
	if err != nil {
		fmt.Println("---> SETNX error: ", err)
	}

	if n == int64(1) {
		fmt.Println("---> Sucessed.")
	}

	var imapGet map[string]string

	valueGet, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		fmt.Println("---> GET ERROR:", err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println("---> Unmashal error: ", err)
	}

	fmt.Println("---> username: ", imapGet["username"])
	fmt.Println("---> addr", imapGet["addr"])
}