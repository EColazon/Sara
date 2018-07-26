package handleShared

import (
	"bytes"
)

// 切片2字符串
func Slice2String(sliceBuff []int) string {
	var str string
	content := bytes.Buffer{}

	for i := 0; i < len(sliceBuff); i++ {
		str = DecimalToAny(sliceBuff[i], 16, 8)
		content.WriteString(str)
	}


	strFinal := content.String()

	return strFinal
}
// 进制转换
func DecimalToAny(num, n, count int) string {

	num2char := "0123456789abcdef"
 
	new_num_str := ""
	var remainder int
	var remainder_string string
	for num != 0 {
		remainder = num % n
		remainder_string = string(num2char[remainder])
		new_num_str = remainder_string + new_num_str //注意顺序
		num = num / n
	}
	// length := len(new_num_str)
	// if length < count { //如果小于8位
	// 	for i := 1; i <= count-length; i++ {
	// 		new_num_str = "0" + new_num_str
	// 	}
	// } else {
	// 	return "ERROR"
	// }

	// 特殊处理个位数字
	if new_num_str == "0" {
		new_num_str = "00"
	} else if new_num_str == "1" {
		new_num_str = "01"
	} else if new_num_str == "2" {
		new_num_str = "02"
	} else if new_num_str == "3" {
		new_num_str = "03"
	} else if new_num_str == "4" {
		new_num_str = "04"
	} else if new_num_str == "5" {
		new_num_str = "05"
	} else if new_num_str == "6" {
		new_num_str = "06"
	} else if new_num_str == "7" {
		new_num_str = "07"
	} else if new_num_str == "8" {
		new_num_str = "08"
	} else if new_num_str == "9" {
		new_num_str = "09"
	}
	return new_num_str
}
