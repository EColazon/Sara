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
---------start:2018.07.10---------
*/

import (
	"fmt"
	"time"
	"reflect"

	"handleShared"
	DBLog "handleMySql"
	Sheard "handleShared"
	Redis "handleRedis"
)

// 实现两个接口s
// HandleCmdGeter:监听channel
// HandleCmdSender:DoZigBeeTasks

type CmdsDistributer interface {
	HandleCmdGeter()
	HandleCmdSender()
}

// channel数据
type CmdChannel struct {
	id		int
	data 	[]int
	snum 	int
}

// flagLogDB 00
// flagLogDB 01
// flagLogDB 02
// flagLogDB 03


// 下发命令
// 协程用于监测channel数据;分发不同命令
func (cmd CmdChannel)HandleCmdGeter() {
	fmt.Println("---> CmdGeter Start.")
	kvJson := make(map[string]interface{})
	go func() {
		for{
			select {

			// fmt.Println("---> buff33: ", len(ChCmd33), cap(ChCmd33), buff33)
			// if buff33["id"] == 73003 {
			// 	fmt.Println("---> id 73003")

			// }
			// case buff2f := <- ChCmd2F:
			// 	fmt.Println("---> buff2f: ", len(ChCmd2F), cap(ChCmd2F), buff2f)
			// 	if buff2f["id"] == 72088 {
			// 		fmt.Println("---> id 72001")
			// 	}
			case buff33 := <- ChCmd33:
				fmt.Println("---> buff33: ", len(ChCmd33), cap(ChCmd33), buff33, time.Now())
				cmd.id, _ = buff33["id"].(int)
				cmd.data, _ = buff33["data"].([]int)
				cmd.snum, _ = buff33["snum"].(int)
				// switch buff33["id"] {
				switch cmd.id {
				case 73000:
					// DoTimeSaveLatLongitude()
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 73000")
				
				case 73001:
					// DoTimeSaveHoliday()
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 73001")
				
				case 73002:
					// DoTimeSaveSpecial()
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 73002")
				
				case 73003:
					// DoTimeSavePwmStage()
					// 将interface{}类型转为[]int类型(通过断言)
					// value, _ := buff33["data"].([]int)
					// snum, _ := buff33["snum"].(int)
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 73003", reflect.TypeOf(cmd.data), cmd.data[:3], cmd.snum)
				
				case 73004:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 73004")
				
				case 73005:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 73005")
				
				case 73006:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 73006")
				
				case 73007:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 73007")
				
				case 73008:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 73008")
				
				case 73009:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 73009")
				default:
					fmt.Println("---> id 33 not fetched..")
					return
				}
			case buff2f := <- ChCmd2F:
				fmt.Println("---> buff2f: ", len(ChCmd2F), cap(ChCmd2F), buff2f, time.Now())
				cmd.id, _ = buff2f["id"].(int)
				cmd.data, _ = buff2f["data"].([]int)
				cmd.snum, _ = buff2f["snum"].(int)
				// switch buff2f["id"] {
				switch cmd.id {
				case 72000: // 返回8字节IEEE地址
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72000")
				case 72001: // 返回程序版本号
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72001")
				case 72002: // 复位设备-RTU
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72002")
				case 72003: // 复位设备-协调器
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72003")
				case 72004: // 读取eeprom数据-RTU
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72004")
				case 72005: // RTU供电方式查询
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72005")
				case 72006: //设置RTU配置:IP
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72006")
				case 72007: //设置RTU配置:Port
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72007")
				case 72008: //设置RTU意外亮灭灯报警阈值syspara.I_oc
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					
					levelTopUnusualSwitch := cmd.data[18] <<8 | cmd.data[19]
					kvJson[Sheard.WDLevelTopUnusualSwitch] = levelTopUnusualSwitch
					Redis.HandleRedisJsonInsert(Sheard.WDLevelTopUnusualSwitch, kvJson)
					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72008")
				case 72009: //RTU经纬度开关状态
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72009")
				case 72010: //设置经度
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					WDTimeLatitude00 := cmd.data[16]
					WDTimeLatitude01 := cmd.data[17]
					WDTimeLatitude02 := cmd.data[18]
					WDTimeLatitude03 := cmd.data[19]
					kvJson[Sheard.WDTimeLatitude00] = WDTimeLatitude00
					Redis.HandleRedisJsonInsert(Sheard.WDTimeLatitude00, kvJson)
					kvJson[Sheard.WDTimeLatitude01] = WDTimeLatitude01
					Redis.HandleRedisJsonInsert(Sheard.WDTimeLatitude01, kvJson)
					kvJson[Sheard.WDTimeLatitude02] = WDTimeLatitude02
					Redis.HandleRedisJsonInsert(Sheard.WDTimeLatitude02, kvJson)
					kvJson[Sheard.WDTimeLatitude03] = WDTimeLatitude03
					Redis.HandleRedisJsonInsert(Sheard.WDTimeLatitude03, kvJson)
					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72010")
				case 72011: //设置纬度
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					WDTimeLongitude00 := cmd.data[16]
					WDTimeLongitude01 := cmd.data[17]
					WDTimeLongitude02 := cmd.data[18]
					WDTimeLongitude03 := cmd.data[19]
					kvJson[Sheard.WDTimeLongitude00] = WDTimeLongitude00
					Redis.HandleRedisJsonInsert(Sheard.WDTimeLongitude00, kvJson)
					kvJson[Sheard.WDTimeLongitude01] = WDTimeLongitude01
					Redis.HandleRedisJsonInsert(Sheard.WDTimeLongitude01, kvJson)
					kvJson[Sheard.WDTimeLongitude02] = WDTimeLongitude02
					Redis.HandleRedisJsonInsert(Sheard.WDTimeLongitude02, kvJson)
					kvJson[Sheard.WDTimeLongitude03] = WDTimeLongitude03
					Redis.HandleRedisJsonInsert(Sheard.WDTimeLongitude03, kvJson)
					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72011")
				case 72012: //设置RTU时间(秒/分/时/星期)
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72012")
				case 72013: //设置RTU时间(日/月/年#)
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72013")
				case 72014: //查询RTU时间(秒/分/时/星期) 
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72014")
				case 72015: //查询RTU时间(日/月/年)
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72015")
				case 72016: //设置RTU电参数采样间隔时间
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72016")
				case 72017: //设置外接电流互感器比例
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					
					// 外接电流互感器比例:syspara.I_t
					ratioTransformerI := cmd.data[18] << 8 | cmd.data[19]
					kvJson[Sheard.WDRatioTransformer] = ratioTransformerI
					Redis.HandleRedisJsonInsert(Sheard.WDRatioTransformer, kvJson)
					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72017")
				case 72018: //设置回路数量
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					loopQuantity := cmd.data[19]
					// 回路数量最大为6路
					if loopQuantity > 6 {
						loopQuantity = 6
					}
					// 存储回路数量
					kvJson[Sheard.WDQuantityLoop] = loopQuantity
					Redis.HandleRedisJsonInsert(Sheard.WDQuantityLoop, kvJson)
					// 计算回路状态
					temp := 0
					loopState := 0 // 回路状态
					relayStateChange := 0  // 继电器状态改变
					for i := 0; i < loopQuantity; i++ {
						// temp |= (1 << i)
						temp |= (i >> 1) // 什么鬼操作
					}
					loopState &= temp
					// 存储回路状态
					kvJson[Sheard.WDStateLoop] = loopState
					Redis.HandleRedisJsonInsert(Sheard.WDStateLoop, kvJson) 
					relayStateChange = 0x55
					// 存储继电器状态
					kvJson[Sheard.WDStateChangeRelay] = relayStateChange
					Redis.HandleRedisJsonInsert(Sheard.WDStateChangeRelay, kvJson)
					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72018")
				case 72019: //设置报警屏蔽
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72019")
				case 72020:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72020")
				case 72021: //设置单灯数量
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72021")
				case 72022: //按位开继电器(总开)
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72022")
				case 72023: //查询RTU温度值(##buff[17]buff[18])
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72023")
				case 72024: //查询RTU门的状态
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72024")
				case 72025: //保留
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72025")
				case 72026: //设置单灯组数
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72026")
				case 72027: //保留
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72027")
				case 72028: //保留
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72028")
				case 72029: //查询锂电池电平值
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72029")
				case 72030: //保留
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72030")
				case 72031:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72031")
				case 72032:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72032")
				case 72033:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72033")
				case 72034:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72034")
				case 72035:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72035")
				case 72036:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72036")
				case 72037:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72037")
				case 72038:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72038")
				case 72039:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72039")
				case 72040:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72040")
				case 72041:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72041")
				case 72042:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72042")
				case 72043:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72043")
				case 72044: // 设置节点功能定时开(秒/分/时/PWM)
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72044")
				case 72045: // 设置节点功能定时关(秒/分/时/PWM)
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72045")
				case 72046: // 设置单灯报警屏蔽(#/#/D1/D2)D1为0x55屏蔽单灯失去连接报警,为0不屏蔽
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72046")
				case 72047: // 设置单灯巡检时间(#/#/DH/DL)范围500ms~20000ms
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72047")
				case 72048: // 保留
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72048")
				case 72049: // 设置回路常开模式(###DD)DD为0x55回路常开,影响单灯报警的处理方式
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72049")
				case 72050: // 设置辅灯不接灯(#/#/Num/DD)Num单灯编号,DD为0x55辅灯不接灯
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72050")
				case 72051: // 保留
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72051")
				case 72052: // 主灯定时开(秒/分/时/#)第11字节代表组号
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72052")
				case 72053: //设置节点控制器的感应恢复时间和触发间隔时间
					GZigbeeNode.GZB72053InductAndTrigeTimeActer(cmd.data[11]<<8|cmd.data[12], cmd.data[16]<<8|cmd.data[17], cmd.data[18]<<8|cmd.data[19], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72053")
					// ZigbeeActer Start.
				case 72054: // 设置单灯组号
					GZigbeeNode.GZB72054SetLampGroupActer(cmd.data[11]<<8|cmd.data[12], cmd.data[19], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72054")
				case 72055: // 单灯主辅互换
					GZigbeeNode.GZB72055SetLampRelayChangeActer(cmd.data[11]<<8|cmd.data[12], cmd.data[19], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72055")
				case 72056: // 保留
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)
					fmt.Println("---> id 72056")
				case 72057: // 设置单灯与节点关联
					GZigbeeNode.GZB72057SetLampAndNodeRelatedActer(cmd.data[11]<<8|cmd.data[12], cmd.data[18], cmd.data[19], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72057")
				case 72058: // 设置单灯在节点触发后的感应恢复时间
					GZigbeeNode.GZB72058SetLampInductRecoverTimeActer(cmd.data[11]<<8|cmd.data[12], cmd.data[18]<<8|cmd.data[19], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72058")
				case 72059: // 设置命令单灯自校验
					GZigbeeNode.GZB72059SetLampSelfCheckActer(cmd.data[11]<<8|cmd.data[12], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72059")
				case 72060: // 固定单灯拨码
					GZigbeeNode.GZB72060SetLampFixedKeyActer(cmd.data[11]<<8|cmd.data[12], cmd.data[19], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72060")
				case 72061: // 查询返回单灯ieee地址
					GZigbeeNode.GZB72061QueryLampIEEEActer(cmd.data[11]<<8|cmd.data[12], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72061")
				case 72062: // 单灯常关
					GZigbeeNode.GZB72062SetLampAlwaysCloseActer(cmd.data[11]<<8|cmd.data[12], cmd.data[18]<<8|cmd.data[19], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72062")
				case 72063: // 单灯内部开主灯
					GZigbeeNode.GZB72063SetLampInterOpenActer(cmd.data[11]<<8|cmd.data[12], LAMPADV, cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72063")
				case 72064: // 单灯内部关主灯
					GZigbeeNode.GZB72064SetLampInterCloseActer(cmd.data[11]<<8|cmd.data[12], LAMPADV, cmd.data[4], cmd.data[12], cmd.snum)
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72064")
				case 72065: // 单灯内部开辅灯
					GZigbeeNode.GZB72063SetLampInterOpenActer(cmd.data[11]<<8|cmd.data[12], LAMPAUX, cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72065")
				case 72066: // 单灯内部关辅灯
					GZigbeeNode.GZB72064SetLampInterCloseActer(cmd.data[11]<<8|cmd.data[12], LAMPAUX, cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72066")
				case 72067: // 单灯内部开主辅灯
					GZigbeeNode.GZB72063SetLampInterOpenActer(cmd.data[11]<<8|cmd.data[12], LAMPADU, cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72067")
				case 72068: // 单灯内部关主辅灯
					GZigbeeNode.GZB72064SetLampInterCloseActer(cmd.data[11]<<8|cmd.data[12], LAMPADU, cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72068")
				case 72069: // 单灯电量保存
					GZigbeeNode.GZB72069SetLampSaveElecsActer(cmd.data[11]<<8|cmd.data[12], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72069")
				case 72070: // 设置单灯电压报警上下限
					GZigbeeNode.GZB72070SetLampAlarmLimitVActer(cmd.data[11]<<8|cmd.data[12], cmd.data[16]<<8|cmd.data[17], cmd.data[18]<<8|cmd.data[19], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72070")
				case 72071: // 设置单灯电流报警上下限
					GZigbeeNode.GZB72071SetLampAlarmLimitIActer(cmd.data[11]<<8|cmd.data[12], cmd.data[16]<<8|cmd.data[17], cmd.data[18]<<8|cmd.data[19], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72071")
				case 72072: // 设置单灯功率报警上下限
					GZigbeeNode.GZB72072SetLampAlarmLimitPActer(cmd.data[11]<<8|cmd.data[12], cmd.data[16]<<8|cmd.data[17], cmd.data[18]<<8|cmd.data[19], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72072")
				case 72073: // 设置单灯功率因素报警上下限
					GZigbeeNode.GZB72073SetLampAlarmLimitPFActer(cmd.data[11]<<8|cmd.data[12], cmd.data[16]<<8|cmd.data[17], cmd.data[18]<<8|cmd.data[19], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72073")
				case 72074: // 清除单灯能量数据
					GZigbeeNode.GZB72074SetLampAnergyClearActer(cmd.data[11]<<8|cmd.data[12], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72074")
				case 72075: // 设置单灯电压放大倍数主灯
					GZigbeeNode.GZB72075SetLampAmplifyVActer(cmd.data[11]<<8|cmd.data[12], cmd.data[19], cmd.data[18], LAMPADV, cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72075")
				case 72076: // 设置单灯电压放大倍数辅灯
					GZigbeeNode.GZB72075SetLampAmplifyVActer(cmd.data[11]<<8|cmd.data[12], cmd.data[19], cmd.data[18], LAMPAUX, cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72076")
				case 72077: // 设置单灯电流放大倍数主灯
					GZigbeeNode.GZB72077SetLampAmplifyIActer(cmd.data[11]<<8|cmd.data[12], cmd.data[19], cmd.data[18], LAMPADV, cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72077")
				case 72078: // 设置单灯电流放大倍数辅灯
					GZigbeeNode.GZB72077SetLampAmplifyIActer(cmd.data[11]<<8|cmd.data[12], cmd.data[19], cmd.data[18], LAMPAUX, cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72078")
				case 72079: // rn8209有效值offset和有功offset校正
					GZigbeeNode.GZB72079SetLampRN8209OffsetActer(cmd.data[11]<<8|cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72079")
				case 72080: // rn8209参数设置 
					GZigbeeNode.GZB72080SetLampRN8209ParasActer(cmd.data[11]<<8|cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72080")
				case 72081: // rn8209比例系数设置
					GZigbeeNode.GZB72081SetLampRN8209RatioActer(cmd.data[11]<<8|cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72081")
				case 72082: // 单灯与rtu时间同步命令
					// GZigbeeNode.GZB72082SetLampRtuSyncTimeActer()
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72082")
				case 72083: // 开关时间同步?
					// GZigbeeNode.GZB72083SetLampSwitchSyncTimeActer()
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72083")
				case 72084: // 查询锂电池电平
					GZigbeeNode.GZB72084QueryBatteryVoltActer(cmd.data[11]<<8|cmd.data[12], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72084")
				case 72085: // 查询温度
					GZigbeeNode.GZB72085QueryTemperatureActer(cmd.data[11]<<8|cmd.data[12], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72085")
				case 72086: // 擦除NV，重新加入网络
					GZigbeeNode.GZB72086SetClearNVAdderActer(cmd.data[11]<<8|cmd.data[12], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72086")
				case 72087: // 单灯开闪烁功能
					GZigbeeNode.GZB72087SetLampTwinkleActer(cmd.data[11]<<8|cmd.data[12], cmd.data[19], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72087")
				case 72088: // 单灯召测命令
					// value, _ := buff2f["data"].([]int)
					// snum, _ := buff2f["snum"].(int)
					// handleShared.HandleSharedCmdOk(22, value[:8], snum)
					GZigbeeNode.GZB72088QueryLampCalledDataActer(cmd.data[11]<<8|cmd.data[12], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72088")
					fmt.Println("---> 72088 amd: ", cmd.data, cmd.data[0], cmd.snum)
				case 72089: // 单灯召测命令
					GZigbeeNode.GZB72088QueryLampCalledDataActer(cmd.data[11]<<8|cmd.data[12], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72089")
				case 72090: // 单灯手动开主辅灯
					GZigbeeNode.GZB72090SetLampOpenByHandsActer(cmd.data[11]<<8|cmd.data[12], LAMPADU, cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72090")
				case 72091: // 单灯手动关主辅灯
					GZigbeeNode.GZB72091SetLampCloseByHandsActer(cmd.data[11]<<8|cmd.data[12], LAMPADU, cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72091")
				case 72092: // 巡检manny命令 
					GZigbeeNode.GZB72092QueryLampDetecDataMannyActer(cmd.data[11]<<8|cmd.data[12], cmd.data[18], cmd.data[19], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72092")
				case 72093: // 巡检内部命令 
					GZigbeeNode.GZB72093QueryLampDetecDataInterActer(cmd.data[11]<<8|cmd.data[12], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72093")
				case 72094: // 单灯手动开主灯
					GZigbeeNode.GZB72090SetLampOpenByHandsActer(cmd.data[11]<<8|cmd.data[12], LAMPADV, cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72094")
				case 72095: // 单灯手动关主灯
					GZigbeeNode.GZB72091SetLampCloseByHandsActer(cmd.data[11]<<8|cmd.data[12], LAMPADV, cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72095")
				case 72096: // 单灯手动开辅灯
					GZigbeeNode.GZB72090SetLampOpenByHandsActer(cmd.data[11]<<8|cmd.data[12], LAMPAUX, cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72096")
				case 72097: // 单灯手动关辅灯
					GZigbeeNode.GZB72091SetLampCloseByHandsActer(cmd.data[11]<<8|cmd.data[12], LAMPAUX, cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72097")
				case 72098: // 保留
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72098")
				case 72099: // 保留
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72099")
				case 72100: // 节点开关命令
					GZigbeeNode.GZB72100SetNodeSwitchActer(cmd.data[11]<<8|cmd.data[12], cmd.data[19], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72100")
				case 72101: // 单灯调光主辅灯
					GZigbeeNode.GZB72101SetLampDimmer(cmd.data[11]<<8|cmd.data[12], LAMPADU, cmd.data[19], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72101")
				case 72102: // 单灯调光主灯
					GZigbeeNode.GZB72101SetLampDimmer(cmd.data[11]<<8|cmd.data[12], LAMPADV, cmd.data[19], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72102")
				case 72103: // 单灯调光辅灯
					GZigbeeNode.GZB72101SetLampDimmer(cmd.data[11]<<8|cmd.data[12], LAMPAUX, cmd.data[19], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72103")
				case 72104: // 返回单灯panid
					GZigbeeNode.GZB72104QueryLampPanidActer(cmd.data[11]<<8|cmd.data[12], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72104")
				case 72105: // 返回单灯程序版本
					GZigbeeNode.GZB72105QueryLampProgramVersionActer(cmd.data[11]<<8|cmd.data[12], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72105")
				case 72106: // 重启单灯
					GZigbeeNode.GZB72106SetLampResetActer(cmd.data[11]<<8|cmd.data[12], cmd.data[4], cmd.data[12], cmd.snum)
					// handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72106")

					// ZigbeeActer End.
				case 72107: //查询RTU召测
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72107")
				case 72108: //液晶召测(保留)
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72108")
				case 72109: //设置RTU继电器全开
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72109")
				case 72110: //设置RTU继电器全关
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72110")
				case 72111: //设置RTU回路电流上下限
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72111")
				case 72112: //设置RTU回路电压上下限
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72112")
				case 72113: //设置RTU继电器1-6开
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72113")
				case 72114: //设置RTU继电器1-6关
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72114")
				case 72115: //获取历史电参量?
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72115")
				case 72116: // 查询stc时间
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72116")
				case 72117:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72117")
				case 72118:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72118")
				case 72119:
					handleShared.HandleSharedCmdOk(22, cmd.data[:8], cmd.snum)

					// 日志记录
					content := Sheard.Slice2String(cmd.data)
					DBLog.HandleDBLogInsert(cmd.id, content, DBNameOK00)
					fmt.Println("---> id 72119")
				default:
					fmt.Println("---> id 2F not fetched..")
					return
				}
			
			
			}
			// time.Sleep(1*time.Second)
		}
	}()
}

func HandleCmdSender() {
	fmt.Println("---> CmdSender Start.")
}