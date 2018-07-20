package handleCmdsManagesSR

import (
	"fmt"
	"reflect"
)

// ZigBee返回数据解析处理
func Cmd33ParsingBack() {
	
	sliceCmd33Back := []int{0x33, 0xFF, 0x01, 0x03, 0x01, 0x60, 0x06, 0x01, 0x00, 0xA8, 0x99}
	//fmt.Println("---> typeOf(sliceCmd33Back)---> ",reflect.TypeOf(sliceCmd33Back))

	for index, value := range sliceCmd33Back {

		if index == 0 {
			fmt.Println(reflect.TypeOf(sliceCmd33Back[index]), sliceCmd33Back[index])
			fmt.Println(reflect.TypeOf(CMD33HEAD), CMD33HEAD)
			fmt.Println("---> 33Back Hello", index, value)
			if sliceCmd33Back[index] == CMD33HEAD { // 头部
				fmt.Println("---> Head is here.")
				lengthData := sliceCmd33Back[index+7]
				fmt.Println("---> lengthData--> ", lengthData)
				snumData := sliceCmd33Back[index+1]
				fmt.Println("---> snumData---> ", snumData)
				if sliceCmd33Back[lengthData+9] == CMD33TAIL { // 尾部
					fmt.Println("---> Tail is here.")
					check := 0 // 校验码
					for i := 0; i < lengthData+8; i++ {
						check ^= sliceCmd33Back[i]
					}
					fmt.Println("---> check---> ", check)
					if check == sliceCmd33Back[lengthData+8] { // 校验
						fmt.Println("---> Check is ok")
						commandCmd := (sliceCmd33Back[5]<<8) | sliceCmd33Back[6] // 命令码
						fmt.Println("---> commandCmd ---> ", commandCmd)

						switch commandCmd {
						case 0x6019: // 设置开始时间.关闭时间.亮度
							MapCmd33Back["id"] = 73500
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73500 is here.")
						case 0x6020: // 设置关联单灯&单灯恢复时间
							MapCmd33Back["id"] = 73501
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73501 is here.")
						case 0x6021: // 召测节点基本信息 
							MapCmd33Back["id"] = 73502
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73502 is here.")
						case 0x6022: // 节点巡检
							MapCmd33Back["id"] = 73503
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73503 is here.")
						case 0x2001: // 分组命令
							MapCmd33Back["id"] = 73504
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73504 is here.")
						case 0x2002: // 主辅灯软件上对调
							MapCmd33Back["id"] = 73505
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73505 is here.")
						case 0x2003: // 单灯报警开,目前命令不起作用
							MapCmd33Back["id"] = 73506
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73506 is here.")
						case 0x2004: // 单灯报警关,目前命令不起作用
							MapCmd33Back["id"] = 73507
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73507 is here.")
						case 0x2005: // 报警时间设置，目前命令不起作用
							MapCmd33Back["id"] = 73508
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73508 is here.")
						case 0x2006: // 电压上下限设置
							MapCmd33Back["id"] = 73509
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73509 is here.")
						case 0x2007: // 电流上下限设置
							MapCmd33Back["id"] = 73510
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73510 is here.")
						case 0x2008: // 功率上下限设置
							MapCmd33Back["id"] = 73511
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73511 is here.")
						case 0x2009: // 功率因素上下限设置
							MapCmd33Back["id"] = 73512
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73512 is here.")
						case 0x200A: // 清能量记录
							MapCmd33Back["id"] = 73513
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73513 is here.")
						case 0x200B: // 命令不起作用
							MapCmd33Back["id"] = 73514
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73514 is here.")
						case 0x200C: // 人体感应组设置
							MapCmd33Back["id"] = 73515
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73515 is here.")
						case 0x200D: // 人体感应延时时间设置
							MapCmd33Back["id"] = 73516
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73516 is here.")
						// case 0x200E: // 命令不起作用
						// 	MapCmd33Back["id"] = 73517
						// 	MapCmd33Back["data"] = sliceCmd33Back
						// 	MapCmd33Back["snum"] = snumData
						// 	ChCmd33Back <- MapCmd33Back
						// 	fmt.Println("---> 73517 is here.")
						case 0x200F: // 命令不起作用
							MapCmd33Back["id"] = 73518
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73518 is here.")
						case 0x2010: // 单灯编号采用硬件读或软件读，命令带有编号信息
							MapCmd33Back["id"] = 73519
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73519 is here.")
						case 0x2011: // 设置主灯电压比例
							MapCmd33Back["id"] = 73520
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73520 is here.")
						case 0x2012: // 设置辅灯电压比例
							MapCmd33Back["id"] = 73521
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73521 is here.")
						case 0x2013: // 设置主灯电流比例
							MapCmd33Back["id"] = 73522
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73522 is here.")
						case 0x2014: // 设置辅灯电流比例
							MapCmd33Back["id"] = 73523
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73523 is here.")
						case 0x2015: // 设置主辅灯是否响应定时开关
							MapCmd33Back["id"] = 73524
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73524 is here.")
						case 0x2016: // 设置panid
							MapCmd33Back["id"] = 73525
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73525 is here.")
						case 0x2017: // 命令不起作用
							MapCmd33Back["id"] = 73526
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73526 is here.")
						case 0x2018: // 设置开始充电和充电结束电压
							MapCmd33Back["id"] = 73527
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73527 is here.")
						case 0x2019: // 设置终止放电和开始放电电压
							MapCmd33Back["id"] = 73528
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73528 is here.")
						case 0x201A: // 设置关灯和开灯电压
							MapCmd33Back["id"] = 73529
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73529 is here.")
						case 0x201B: // 调整主灯电流
							MapCmd33Back["id"] = 73530
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73530 is here.")
						case 0x201C: // 调整辅灯电流
							MapCmd33Back["id"] = 73531
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73531 is here.")
						case 0x2021: // 开关节点调光功能
							MapCmd33Back["id"] = 73532
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73532 is here.")
						case 0x2023: // 单灯开闪烁功能
							MapCmd33Back["id"] = 73533
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73533 is here.")
						case 0x2025: // 查询stc时间
							MapCmd33Back["id"] = 73534
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73534 is here.")
						case 0x1201: // 手持机单灯总开包括回路
							MapCmd33Back["id"] = 73535
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73535 is here.")
						case 0x1202: // 手持机单灯总关包括回路
							MapCmd33Back["id"] = 73536
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73536 is here.")
						case 0x1203: // 手持机主辅灯总开
							MapCmd33Back["id"] = 73537
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73537 is here.")
						case 0x1204: // 手持机主辅灯总关
							MapCmd33Back["id"] = 73538
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73538 is here.")
						case 0x2102: // 单灯巡检
							MapCmd33Back["id"] = 73539
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73539 is here.")
						case 0x1003: // 主辅灯总开
							MapCmd33Back["id"] = 73540
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73540 is here.")
						case 0x1004: // 主辅灯总关
							MapCmd33Back["id"] = 73541
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73541 is here.")
						case 0x1005: // 主灯开
							MapCmd33Back["id"] = 73542
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73542 is here.")
						case 0x1006: // 主灯关
							MapCmd33Back["id"] = 73543
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73543 is here.")
						case 0x1007: // 辅灯开 
							MapCmd33Back["id"] = 73544
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73544 is here.")
						case 0x1008: // 辅灯关
							MapCmd33Back["id"] = 73545
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73545 is here.")
						case 0x1009: // 单灯召测
							MapCmd33Back["id"] = 73546
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73546 is here.")
						case 0x100A: // 主辅调光
							MapCmd33Back["id"] = 73547
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73547 is here.")
						case 0x100B: // 主灯调光
							MapCmd33Back["id"] = 73548
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73548 is here.")
						case 0x100C: // 辅灯调光
							MapCmd33Back["id"] = 73549
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73549 is here.")
						case 0x100D: // 返回单灯程序版本
							MapCmd33Back["id"] = 73550
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73550 is here.")
						case 0x100E: // 复位设备
							MapCmd33Back["id"] = 73551
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73551 is here.")
						case 0x1020: // 擦除NV，重新加入网络
							MapCmd33Back["id"] = 73552
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73552 is here.")
						case 0x200E: // 发命令让单灯自校验
							MapCmd33Back["id"] = 73553
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73553 is here.")
						case 0x100F: // 返回panid
							MapCmd33Back["id"] = 73554
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73554 is here.")
						case 0x1010: // 命令不起作用
							MapCmd33Back["id"] = 73555
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73555 is here.")
						case 0x1011: // 读取协调器ieee地址
							MapCmd33Back["id"] = 73556
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73556 is here.")
						case 0x1018: // 获取锂电池电压
							MapCmd33Back["id"] = 73557
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73557 is here.")
						case 0x1019: // 获取温度
							MapCmd33Back["id"] = 73558
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73558 is here.")
						case 0x3001: // ?
							MapCmd33Back["id"] = 73559
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73559 is here.")
						case 0x3002: // 
							MapCmd33Back["id"] = 73560
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73560 is here.")
						case 0x2030: // 
							MapCmd33Back["id"] = 73561
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73561 is here.")
						case 0x1000: // 
							MapCmd33Back["id"] = 73562
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73562 is here.")
						case 0x1001: // 
							MapCmd33Back["id"] = 73563
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73563 is here.")
						case 0x3004: // 窨井盖报警
							MapCmd33Back["id"] = 73564
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73564 is here.")
						case 0x0997: // 
							MapCmd33Back["id"] = 73565
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73565 is here.")
						case 0x2400: // 
							MapCmd33Back["id"] = 73566
							MapCmd33Back["data"] = sliceCmd33Back
							MapCmd33Back["snum"] = snumData
							ChCmd33Back <- MapCmd33Back
							fmt.Println("---> 73566 is here.")
						default:
							fmt.Println("---> Another Cmds.")
						}
					}
				}
			}
		}
	}
}
