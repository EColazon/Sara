package handleRedis

import (
	"testing"
)

//单元测试
func TestHandleRedisJsonInsert(t *testing.T) {
	//初始化键值信息
	keyString := "1"
	kvString := make(map[string]interface{})
	kvInt := make(map[string]interface{})

	kvString["1"] = "1string"
	kvInt["2"] = 2
	//测试开始
	//value:stirng
	if HandleRedisJsonInsert(keyString, kvString) != true{		
		t.Log("---> testInsert valueString err.")
	}


	//value:int
	if HandleRedisJsonInsert(keyString, kvInt) != true{			
		t.Log("---> testInsert valueString err.")
	}

}

func TestHandleRedisJsonGet(t *testing.T) {
	//初始化键值信息
	keyString := "1"
	//Get
	if HandleRedisJsonGet(keyString) != nil {
		t.Log("---> testGet err.")
	}
}

func TestHandleRedisJsonDel(t *testing.T) {
	//初始化键值信息
	keyString := "1"
	//Del
	if HandleRedisJsonDel(keyString) == nil  {
		t.Log("---> testDel err.")
	}
}
//性能测试
func BenchmarkHandleRedisJsonInsert(b *testing.B) {
	//初始化键值信息
	keyString := "1"
	kvString := make(map[string]interface{})
	kvInt := make(map[string]interface{})

	kvString["1"] = "1string"
	kvInt["2"] = 2
	for i := 0; i < b.N; i++ {
		_ = HandleRedisJsonInsert(keyString, kvString)
	}

	for i := 0; i < b.N; i++ {
		_ = HandleRedisJsonInsert(keyString, kvInt)
	}


}

func BenchmarkHandleRedisJsonGet(b *testing.B) {
	//初始化键值信息
	keyString := "1"
	for i := 0; i < b.N; i++ {
		_ = HandleRedisJsonGet(keyString)
	}
}

func BenchmarkHandleRedisJsonDel(b *testing.B) {
	//初始化键值信息
	keyString := "1"
	for i := 0; i < b.N; i++ {
		_ = HandleRedisJsonDel(keyString)
	}
}