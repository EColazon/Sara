package handleShared

import (
	"fmt"
)
var (
	headBuff = []int{0x00, 0x00, 0x00, 0x00, 0x00}
	originalBuff = []int{0x2f, 0x2f, 0x2f, 0x05,0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,0x0A, 0xC0, 0xB4,0x00, 0x00, 0x00, 0x00,0x00, 0xCC}
	uploadBuff = []int{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
)
func HandleSharedCmdOk(length int, data []int, serNum int) interface{}{

	if length < 0 || len(data) != 0x08 || serNum > 255 || serNum < 1 {
		fmt.Println("---> HandleSharedCmdOk paramsLen err", length, len(data), serNum)
		return nil
	}
	fmt.Println("---> handleShared : uploadBuff-len ", len(uploadBuff))

	for i := 0; i < 8; i++ {
		originalBuff[4+i] = 0x99
	}
	
	fmt.Println("---> handleShared : originalBuff-data ", originalBuff)
	originalBuff[12] = 0x00 //OK
	for index, value := range data {
		originalBuff[13+index] = value
	}
	fmt.Println("---> handleShared : originalBuff-data ", originalBuff)
	//尾部校验
	checkCode := 0
	for i := 3; i < 19; i++ {
		checkCode ^= originalBuff[i]
	}
	fmt.Println("---> handleShared : originalBuff-checkCode ", checkCode)
	originalBuff[19] = checkCode
	originalBuff[20] = 0xCC
	fmt.Println("---> handleShared : originalBuff-data-ok1 ", originalBuff)

	if length==0x00 {
		fmt.Println("---> handleShared : parameters-length-error ", length)
	} else {
		headBuff[0] = (length>>24)&0xff
		headBuff[1] = (length>>16)&0xff
		headBuff[2] = (length>>8)&0xff
		headBuff[3] = length&0xff
		headBuff[4] = serNum
	}
	fmt.Println("---> handleShared : originalBuff-data-ok2 ", headBuff, originalBuff)
	for i := 0; i < 5; i++ {
		uploadBuff[i] = headBuff[i]
	}
	for i := 0; i < 21; i++ {
		uploadBuff[i+5] = originalBuff[i]
	}
	fmt.Println("---> handleShared : uploadBuff-data-ok ", len(uploadBuff), uploadBuff)


	return true
}

func HandleSharedCmdError(length int, data []int, serNum int) interface{}{
	if length < 0 || len(data) != 0x08 || serNum > 255 || serNum < 1 {
		fmt.Println("---> HandleSharedCmdOk paramsLen err", length, len(data), serNum)
		return nil
	}
	fmt.Println("---> handleShared : uploadBuff-len ", len(uploadBuff))

	for i := 0; i < 8; i++ {
		originalBuff[4+i] = 0x99
	}
	
	fmt.Println("---> handleShared : originalBuff-data ", originalBuff)
	originalBuff[12] = 0xFF //ERROR
	for index, value := range data {
		originalBuff[13+index] = value
	}
	fmt.Println("---> handleShared : originalBuff-data ", originalBuff)
	//尾部校验
	checkCode := 0
	for i := 3; i < 19; i++ {
		checkCode ^= originalBuff[i]
	}
	fmt.Println("---> handleShared : originalBuff-checkCode ", checkCode)
	originalBuff[19] = checkCode
	originalBuff[20] = 0xCC
	fmt.Println("---> handleShared : originalBuff-data-ok1 ", originalBuff)

	if length==0x00 {
		fmt.Println("---> handleShared : parameters-length-error ", length)
	} else {
		headBuff[0] = (length>>24)&0xff
		headBuff[1] = (length>>16)&0xff
		headBuff[2] = (length>>8)&0xff
		headBuff[3] = length&0xff
		headBuff[4] = serNum
	}
	fmt.Println("---> handleShared : originalBuff-data-ok2 ", headBuff, originalBuff)
	for i := 0; i < 5; i++ {
		uploadBuff[i] = headBuff[i]
	}
	for i := 0; i < 21; i++ {
		uploadBuff[i+5] = originalBuff[i]
	}
	fmt.Println("---> handleShared : uploadBuff-data-ok ", len(uploadBuff), uploadBuff)


	return true
}

