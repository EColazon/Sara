package handleShared

import (
	"fmt"
	"crypto/md5"
	"hash/crc32"
	"reflect"
)

func ExecCRC32(data string) {
	checkData := []byte(data)
	result := crc32.ChecksumIEEE(checkData)
	fmt.Println("---> CRC32Result: ", reflect.TypeOf(result), result)
}

func ExecMD5(data string) {
	checkData := []byte(data)
	result := md5.Sum(checkData)
	fmt.Println("---> MD5Result: ", reflect.TypeOf(result), result)
}