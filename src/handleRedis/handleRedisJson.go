package handleRedis

import (
	"fmt"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

//判断键是否存在,1删除
//插入键值信息
//return:true&false
func HandleRedisJsonInsert(key string, kvJson map[string]interface{}) interface{}{
	imap := kvJson
	value, _ := json.Marshal(imap)
	valueGet, err := redis.Bytes(RConn.Do("GET", key))
	if err != nil {
		fmt.Println("---> GET ERROR:", err)
		
	} else {
		fmt.Println("---> valueGet: ", valueGet)
		fmt.Println("---> value ok.")
		RConn.Do("DEL", key)
		fmt.Println("---> DEL key.")
	}
	

	ok, err := RConn.Do("SETNX", key, value)
	if err != nil {
		fmt.Println("---> SETNX error: ", err)
		return false
	}

	if ok == int64(1) {
		fmt.Println("---> Insert Sucessed.")
	}
	return true
	
}

//通过键获取键对应的信息
//return:value
func HandleRedisJsonGet(key string) interface{} {
	var imapGet map[string]interface{}

	valueGet, err := redis.Bytes(RConn.Do("GET", key))
	if err != nil {
		fmt.Println("---> GET ERROR:", err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println("---> Unmashal error: ", err)
		return nil
	}

	fmt.Println("---> key:value ",key, imapGet[key])
	fmt.Println("---> value", imapGet)

	return imapGet[key]
}
//删除键极其对应的信息
//return:true&false
func HandleRedisJsonDel(key string) interface{}{

	valueDel, err := redis.Bool(RConn.Do("DEL", key))
	if err != nil {
		fmt.Println("---> Del ERROR:", err)
		return nil
	}
	fmt.Println("---> Del valueDel: ", valueDel)

	return valueDel

}

