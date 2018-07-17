package handleCmdsManagesSR

import (
	"fmt"
	"time"
)

// 实现命令重传机制
/*
流程:
	1.使用map:MapSnum记录命令序列号snum
	2.发送cmd到zigebee串口计时开始
	3.监测timeTick范围(0.5s/loop) and flagRealBack
		3.1snumBack in MapSnum and timeTick in 0-5 and flagRealBack is true:
			delete snum and flagRealBack false
		3.2snumBack in MapSnum and timeTick in 5-10 and flagRealBack is true:
			delete snum and flagRealBack false and timeTick = 0
		3.3snumBack not in MapSnum and timeTick in 5-10 and flagRealBack is not true:
			DoFuncSendOnce
		3.4snumBack not in MapSnum and timeTick > 10 and flagRealBack is not true:
			delete snum and DoFuncAlarmTimeOut and timeTick = 0
		
*/

// map:MapSnum用于记录序列号snum,最大长度1024
// FlagRealBack用于标识协调器返回命令
var(
	// SliceSnum = make([]int, 1024)
	MapSnum = make(map[int]int, 1024)
	FlagRealBack = 0
	SnumBack = 1
)

// 实现1个接口s
// HandleCmdRetransmissionDistributer:监听channel
// HandleCmdRetransmissionDistributer:DoZigBeeTasks

type CmdsZigbeeDistributer interface {
	HandleCmdRetransmissionDistributer()
}

// channel数据
type CmdZigbeeChannel struct {
	id		int
	data 	[]int
	snum 	int
}
// 协程用于监测channel数据;分发不同命令
func (cmdZigbee CmdZigbeeChannel)HandleCmdRetransmissionDistributer() {
	fmt.Println("---> CmdGeter Start.")
	go func() {
		for{
			select {
			case buffZigbee := <- ChCmdZigbeeSend:
				fmt.Println("---> missok", buffZigbee)
				cmdZigbee.id, _ = buffZigbee["id"].(int)
				cmdZigbee.data, _ = buffZigbee["data"].([]int)
				cmdZigbee.snum, _ = buffZigbee["snum"].(int)
				fmt.Println("---> missData", cmdZigbee)
				// 记录序列号
				MapSnum[cmdZigbee.snum] = cmdZigbee.snum
				// 命令下发处理
				go HandleCmdDistributerRetransmission(cmdZigbee.data, cmdZigbee.snum)
				// switch cmdZigbee.id {
				// case 72053:
				// 	fmt.Println("---> miss in 72053", cmdZigbee.snum)
				// 	go HandleCmdDistributerRetransmission(cmdZigbee.snum)
				// case 0x02:
				// 	fmt.Println("---> miss in 02", cmdZigbee.snum)
				// 	go HandleCmdDistributerRetransmission(cmdZigbee.snum)
				// }

			}
		}
	}()
}

// 实现超时重传
func HandleCmdDistributerRetransmission(data []int, snum int) {

	// zigbee 命令下发
	HandleCmdSendOnce(data) // 换为ZigBee相关发送接口
	fmt.Println("---> Strat HandleRetransmission.", len(MapSnum), MapSnum)
	// 计数器
	timeTick := 0

	for {
		_, ok := MapSnum[SnumBack];
		if  ok && timeTick < 5 {
			fmt.Println("---> SnumBack is Ok 01.", timeTick)
			// 数据返回正常
			delete(MapSnum, SnumBack)
			break
		} else if ok && timeTick >= 5 && timeTick < 10{
			fmt.Println("---> SnumBack is Ok 02.", timeTick, SnumBack)
			// 数据返回正常
			delete(MapSnum, SnumBack)
			timeTick = 0
			fmt.Println("---> Delete HandleRetransmission.", len(MapSnum), MapSnum)
			break
		} else if !ok && timeTick >=5 && timeTick < 10 {
			// 数据返回超时
			HandleCmdSendOnce(data)
			fmt.Println("---> SnumBack is NOT OK and ReTry.", timeTick)
		} else if !ok && timeTick >= 15 {
			// 第二次命令下发无数据返回
			delete(MapSnum, SnumBack)
			timeTick = 0
			// 异常报警
			handleAlarmTimeout(snum)
			fmt.Println("---> SnumBack is NOT Ok.", timeTick)
			break
			
		}
		timeTick += 1
		time.Sleep(1*time.Second)

	}
	fmt.Println("---> OutLoop.")






}


func HanldeCmdZigbeeLoading(id int, sliceCommand []int, snum int) {
	// 将数据(id&data&snum)发送到channel
	MapCmdZigbee["id"] = id
	MapCmdZigbee["data"] = sliceCommand
	MapCmdZigbee["snum"] = snum
	ChCmdZigbeeSend <- MapCmdZigbee
}

// 调用ZigBee接口下发相关命令
func HandleCmdSendOnce(data []int) {
	fmt.Println("---> Strat handleCmdSendOnce.")
	for i := 0; i < 3; i++{
		MapCmdZigbee["id"] = i
		MapCmdZigbee["data"] = data
		MapCmdZigbee["snum"] = i
		ChCmdZigbeeSend <- MapCmdZigbee
		fmt.Println("---> 01")
		time.Sleep(5*time.Second)
	}
}

func handleAlarmTimeout(snum int) {
	fmt.Println("---> Strat handleAlarmTimeout.")

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0

	sliceCommand := make([]int, 16)
	sliceCommandHead := make([]int, 5)
	sliceCommand33 := make([]int, 11)
	length := len(sliceCommand33) + 1

	sliceCommand33[0] = 0x33
	sliceCommand33[1] = 0x01
	sliceCommand33[2] = 0x02
	sliceCommand33[3] = 0xFF
	sliceCommand33[4] = 0x04
	sliceCommand33[5] = 0xFF
	sliceCommand33[6] = 0xFF
	sliceCommand33[7] = 0x01
	sliceCommand33[8] = snum

	for i := 0; i < 9; i++ {
		checkSum ^= sliceCommand33[i]
	}

	sliceCommand33[9] = checkSum
	sliceCommand33[10] = 0x99

	sliceCommandHead[0] = (length>>24) & 0xFF
	sliceCommandHead[1] = (length>>16) & 0xFF
	sliceCommandHead[2] = (length>>8) & 0xFF
	sliceCommandHead[3] = length & 0xFF
	sliceCommandHead[5] = snum

	for i := 0; i < 5; i++ {
		sliceCommand[i] = sliceCommandHead[i]
	}
	for i := 0; i < 11; i++ {
		sliceCommand[i+5] = sliceCommand33[i]
	}

	// TODO TCP.send



}