package handleCmdsManages

/*
流程: 
	1.CmdGeter监听ChCmd
	2.根据不同命令id做不同动作
		2.1DoRedisMessages
		2.2DoStateCallback
		2.3DoSender
	3.DoZigbeeTasks
		3.xDoTasks...

Author:mengfei.wu@foxmail.com
---------start:2018.07.20---------
*/

import (
	"fmt"
	"time"
	// "reflect"
	DBLog "handleMySql"
	Sheard "handleShared"

)

// 实现接口
// HandleCmdBackGeter:监听channel
// HandleCmdSender:DoZigBeeTasks

type CmdsBackDistributer interface {
	HandleCmdBackGeter()
}

// channel数据
type CmdBackChannel struct {
	id		int
	data 	[]int
	snum 	int
}

// 返回命令
// 协程用于监测channel数据;分发不同命令
func (cmd CmdBackChannel)HandleCmdBackGeter() {
	fmt.Println("---> CmdGeter Start.")
	go func() {
		for{
			select {

			case buff33Back := <- ChCmd33Back:
				lampBackCommand := make([]int, 21)
				fmt.Println("---> buff33Back: ", len(ChCmd33Back), cap(ChCmd33Back), buff33Back, time.Now())
				cmd.id, _ = buff33Back["id"].(int)
				cmd.data, _ = buff33Back["data"].([]int)
				cmd.snum, _ = buff33Back["snum"].(int)
				// switch buff33Back["id"] {
				switch cmd.id {
				case 73500: // 设置开始时间.关闭时间.亮度
					lampBackCommand[13] = 0x02
					lampBackCommand[14] = 0xC1

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73500")
				case 73501: // 设置关联单灯&单灯恢复时间
					lampBackCommand[13] = 0x02
					lampBackCommand[14] = 0xC2

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73501")
				case 73502: // 召测节点基本信息 

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73502")
				case 73503: // 节点巡检

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73503")
				case 73504: // 分组命令
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x00
					lampBackCommand[15] = 0x92
					lampBackCommand[16] = 0x00
					lampBackCommand[17] = 0x00
					lampBackCommand[18] = 0x00
					lampBackCommand[19] = cmd.data[8]

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73504")
				case 73505: // 主辅灯软件上对调
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x94

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73505")
				case 73506: // 单灯报警开,目前命令不起作用
					lampBackCommand[13] = 0xE2
					lampBackCommand[14] = 0xC1

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73506")
				case 73507: // 单灯报警关,目前命令不起作用
					lampBackCommand[13] = 0xE2
					lampBackCommand[14] = 0xC2

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73507")
				case 73508: // 报警时间设置，目前命令不起作用
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x95

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73508")
				case 73509: // 电压上下限设置
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xA2

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73509")
				case 73510: // 电流上下限设置
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xA3

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73510")
				case 73511: // 功率上下限设置
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xA4

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73511")
				case 73512: // 功率因素上下限设置
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xA5

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73512")
				case 73513: // 清能量记录
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xA6

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73513")
				case 73514: // 命令不起作用
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xA7

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73514")
				case 73515: // 人体感应组设置
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x96

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73515")
				case 73516: // 人体感应延时时间设置
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x98

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73516")
				case 73517: // 命令不起作用
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x99

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73517")
				case 73518: // 命令不起作用
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x9A

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73518")
				case 73519: // 单灯编号采用硬件读或软件读，命令带有编号信息
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x9B

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73519")
				case 73520: // 设置主灯电压比例
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xA8

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73520")
				case 73521: // 设置辅灯电压比例
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xA9

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73521")
				case 73522: // 设置主灯电流比例
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xAA

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73522")
				case 73523: // 设置辅灯电流比例
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xAB

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73523")
				case 73524: // 设置主辅灯是否响应定时开关
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x9E

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73524")
				case 73525: // 设置panid
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xB2

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73525")
				case 73526: // 命令不起作用
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xB3

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73526")
				case 73527: // 设置开始充电和充电结束电压
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xB4

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73527")
				case 73528: // 设置终止放电和开始放电电压
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xB5

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73528")
				case 73529: // 设置关灯和开灯电压
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xB6

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73529")
				case 73530: // 调整主灯电流
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xB7

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73500")
				case 73531: // 调整辅灯电流
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xB8

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73531")
				case 73532: // 开关节点调光功能
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x15

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73532")
				case 73533: // 单灯开闪烁功能
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0xBF

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73533")
				case 73534: // 查询stc时间
					lampBackCommand[13] = cmd.data[8]
					lampBackCommand[14] = cmd.data[9]
					lampBackCommand[15] = cmd.data[10]
					lampBackCommand[16] = cmd.data[11]
					lampBackCommand[17] = cmd.data[12]
					lampBackCommand[18] = cmd.data[13]

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73534")
				case 73535: // 手持机单灯总开包括回路

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73535")
				case 73536: // 手持机单灯总关包括回路

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73536")
				case 73537: // 手持机主辅灯总开

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73537")
				case 73538: // 手持机主辅灯总关

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73538")
				case 73539: // 单灯巡检

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73539")
				case 73540: // 主辅灯总开
					lampBackCommand[13] = 0xC0
					lampBackCommand[14] = 0xC1

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73540")
				case 73541: // 主辅灯总关
					lampBackCommand[13] = 0xC0
					lampBackCommand[14] = 0xC2

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73541")
				case 73542: // 主灯开
					lampBackCommand[13] = 0xC1
					lampBackCommand[14] = 0xC1

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73542")
				case 73543: // 主灯关
					lampBackCommand[13] = 0xC1
					lampBackCommand[14] = 0xC2

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73543")
				case 73544: // 辅灯开 
					lampBackCommand[13] = 0xC2
					lampBackCommand[14] = 0xC1

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73544")
				case 73545: // 辅灯关
					lampBackCommand[13] = 0xC2
					lampBackCommand[14] = 0xC2

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73545")
				case 73546: // 单灯召测

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73546")
				case 73547: // 主辅调光
					lampBackCommand[13] = 0xC0
					lampBackCommand[14] = 0xC1

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73547")
				case 73548: // 主灯调光
					lampBackCommand[13] = 0xC1
					lampBackCommand[14] = 0xC1

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73548")
				case 73549: // 辅灯调光
					lampBackCommand[13] = 0xC2
					lampBackCommand[14] = 0xC1

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73549")
				case 73550: // 返回单灯程序版本
					lampBackCommand[12] = 0x00
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x00
					lampBackCommand[15] = 0x00
					lampBackCommand[16] = 0x00
					lampBackCommand[17] = cmd.data[8]
					lampBackCommand[18] = cmd.data[9]

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73550")
				case 73551: // 复位设备
					lampBackCommand[12] = 0x00
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x00

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73551")
				case 73552: // 擦除NV，重新加入网络
					lampBackCommand[12] = 0x00
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x00

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73552")
				case 73553: // 发命令让单灯自校验
					lampBackCommand[12] = 0x00
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x00

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73553")
				case 73554: // 返回panid
					lampBackCommand[12] = 0x00
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x00
					lampBackCommand[15] = 0x00
					lampBackCommand[16] = 0x00
					lampBackCommand[17] = cmd.data[8]
					lampBackCommand[18] = cmd.data[9]

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73554")
				case 73555: // 命令不起作用
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x00

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73555")
				case 73556: // 读取协调器ieee地址

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73556")
				case 73557: // 获取锂电池电压

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73557")
				case 73558: // 获取温度
					
					// 温度采样成功
					Sheard.WDFlagTempreratureBack = 1

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73558")
				case 73559:

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73559")
				case 73560:

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73560")
				case 73561:

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73561")
				case 73562:
					lampBackCommand[12] = 0xFF
					lampBackCommand[13] = 0x00
					lampBackCommand[14] = 0x00
					lampBackCommand[15] = 0x00
					lampBackCommand[16] = 0x00
					lampBackCommand[17] = cmd.data[8]
					lampBackCommand[18] = cmd.data[9]

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73562")
				case 73563:
					lampBackCommand[12] = 0x00
					lampBackCommand[15] = 0x00
					lampBackCommand[16] = 0x00
					lampBackCommand[17] = cmd.data[8]
					lampBackCommand[18] = cmd.data[9]

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73563")
				case 73564: // 窨井盖报警
					lampBackCommand[0] = 0x2F
					lampBackCommand[1] = 0x2F
					lampBackCommand[2] = 0x2F
					lampBackCommand[3] = 0x01
					lampBackCommand[4] = cmd.data[15]
					lampBackCommand[5] = cmd.data[14]
					lampBackCommand[6] = cmd.data[13]
					lampBackCommand[7] = cmd.data[12]
					lampBackCommand[8] = cmd.data[11]
					lampBackCommand[9] = cmd.data[10]
					lampBackCommand[10] = cmd.data[9]
					lampBackCommand[11] = cmd.data[8]
					lampBackCommand[12] = 0x00
					lampBackCommand[13] = cmd.data[16]
					lampBackCommand[14] = 0xEB
					lampBackCommand[15] = cmd.data[21]
					lampBackCommand[16] = cmd.data[22]

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73564")
				case 73565:

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73565")
				case 73566:

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73566")
				case 73567:

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73567")
				case 73568:

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73568")
				case 73569:

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK02)
					fmt.Println("---> id 73569")
				}
			}
		}
	}()
}