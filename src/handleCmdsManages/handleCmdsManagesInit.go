package handleCmdsManages

import (
	"fmt"
)
const (
	CMD33HEAD = 0x33
	CMD33TAIL = 0x99
	CMD2FHEAD00 = 0x2F
	CMD2FHEAD01 = 0x43
	CMD2FHEAD02 = 0x2F
	CMD2FTAIL = 0xCC

)
// 声明全局缓冲通道用于命令解析到命令分发间通信
var ChCmd2F = make(chan map[string]interface{}, 1024)
var ChCmd33 = make(chan map[string]interface{}, 1024)

// zigbee返回数据解析channel
var ChCmd33Back = make(chan map[string]interface{}, 1024)

// 声明全局缓冲channel用于下发zigbee命令
var ChCmdZigbeeSend = make(chan map[string]interface{}, 1024)


// 声明eixt通道用于阻塞channel
var ChExit = make(chan int)
// 声明map格式用于拼组数据
var MapCmd2f = make(map[string]interface{})
var MapCmd33 = make(map[string]interface{})

// zigbee返回数据
var MapCmd33Back = make(map[string]interface{})

// zigbee下发数据
var MapCmdZigbee = make(map[string]interface{})

func init() {
	fmt.Println("---> handleCmdsManagesInit.")
}