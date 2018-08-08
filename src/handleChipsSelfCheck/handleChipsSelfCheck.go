package handleChipsSelfCheckManages

import (
	"fmt"
	"time"
	"math"

	Sheard "handleShared"
	Cmd "handleCmdsManages"
)

var (
	timeCountTemperature 		= 60*1
	timeCountTemperatureBack 	= 70
	timeCountLiVolt 			= 60*1
	timeCountLatLong 			= 60*60
	timeCountClock				= 60*1

	timeAbsSysPCF 				= 60*3 // 系统时间和时钟芯片绝对值差默认3分钟
)

func HandleChipsSelfCheckManages() {
	fmt.Println("---> HandleChipsSelfCheckManages.")

	pinErr 					:= 40 // 芯片异常40pin

	
	checkTimeCount 			:= 0
	checkCountEeprom 		:= 0 // EEPROM检测计数
	checkCountPCF8563 		:= 0 // PCF8563检测计数
	checkCountRN8209 		:= 0 // RN8209检测计数
	checkCountZigbee 		:= 0 // zigbeeIEEE地址读取检测计数
	checkCountClock 		:= 0 // 时钟校时检测计数
	checkCountTemperature 	:= 0 // 温度采样计数
	checkCountTemperatureBack 	:= 0 // 温度采样返回计数
	checkCountLiVolt 		:= 0 // 锂电池采样计数
	checkCountLatLong 		:= 0 // 经纬度时间更新计数

	checkFlagEeprom 		:= 1 // EEPROM检测标志
	checkFlagPCF8563 		:= 1 // PCF8563检测标志
	checkFlagMCP23008		:= 1 // MCP23008检测标志
	checkFlagRN8209 		:= 1 // RN8209检测标志
	checkFlagZigbee 		:= 1 // zigbeeIEEE地址读取检测标志
	checkFlagClock 			:= 1 // 时钟校时标志
	checkFlagTemperature 	:= 1 // 温度采样标志
	checkFlagLiVolt 		:= 1 // 锂电池采样标志
	checkFlagLatLong 		:= 1 // 经纬度时间更新标志

	errFlagEeprom 		:= 0 // EEPROM检测异常标志
	errFlagPCF8563 		:= 0 // PCF8563检测异常标志
	errFlagMCP23008		:= 0 // MCP23008检测异常标志
	errFlag01RN8209 	:= 0 // RN8209检测异常标志
	errFlag02RN8209 	:= 0 // RN8209检测异常标志
	errFlag03RN8209 	:= 0 // RN8209检测异常标志
	errFlagZigbee 		:= 0 // zigbeeIEEE地址读取检测异常标志
	errFlagClock 		:= 0 // 时钟校时异常标志
	errFlagTemperature 	:= 0 // 温度采样异常标志
	errFlagLiVolt 		:= 0 // 锂电池采样异常标志
	errFlagLatLong 		:= 0 // 经纬度时间更新异常标志



	for {
		// 计数自增
		checkTimeCount += 1
		checkCountTemperature += 1 // 温度采样
		checkCountTemperatureBack += 1 // 温度采样返回
		checkCountLiVolt += 1 // 锂电池采样
		checkCountClock += 1 // 系统时间和时钟芯片时间检测
		time.Sleep(1 * time.Second)

		

		// EEPROM检测计数
		if (checkFlagEeprom == 1) {
			checkFlagEeprom = 0

			// 写数据
			Sheard.HandleSharedExecCSoI2C0Write(Sheard.DEEPADDR57, Sheard.WDEEP57CHECK001, 0x55)
			// 读数据
			readData := Sheard.HandleSharedExecCSoI2C0Read(Sheard.DEEPADDR57, Sheard.WDEEP57CHECK001)

			if (readData != 0x55) {
				// EEPROM异常
				errFlagEeprom = 1
				// TODO socketToServer
			}

		} else if (checkFlagPCF8563 == 1) {
			checkFlagPCF8563 = 0

			// 写数据
			Sheard.HandleSharedExecCSoI2C0Write(Sheard.WDPCFADDR51, Sheard.WDPCF51CHECK001, 0xd9)
			// 读数据
			readData := Sheard.HandleSharedExecCSoI2C0Read(Sheard.WDPCFADDR51, Sheard.WDPCF51CHECK001)

			if (readData != 0xd9) {
				// PCF8563异常
				errFlagPCF8563 = 1
				// TODO socketToServer

			}

		} else if (checkFlagMCP23008 == 1) {
			checkFlagMCP23008 = 0

			// 写数据
			Sheard.HandleSharedExecCSoI2C1Write(Sheard.WDMCPADDR20, Sheard.WDMCP20CHECK001, 0x55)
			// 读数据
			readData := Sheard.HandleSharedExecCSoI2C1Read(Sheard.WDMCPADDR20, Sheard.WDMCP20CHECK001)
			if (readData != 0x55) {
				// MCP23008异常
				errFlagMCP23008 = 1
				// TODO socketToServer

			}
		} else if (checkFlagRN8209 == 1) {
			checkFlagRN8209 = 0
			readData01 := Sheard.HandleSharedExecCSoRN8209ReadFromID(1)
			if readData01[0] != 0x82 || readData01[1] != 0x09 {
				errFlag01RN8209 = 1
				// TODO socketToServer
			}
			readData02 := Sheard.HandleSharedExecCSoRN8209ReadFromID(2)
			if readData02[0] != 0x82 || readData02[1] != 0x09 {
				errFlag02RN8209 = 1
				// TODO socketToServer
			}
			readData03 := Sheard.HandleSharedExecCSoRN8209ReadFromID(3)
			if readData03[0] != 0x82 || readData03[1] != 0x09 {
				errFlag03RN8209 = 1
				// TODO socketToServer
			}

		} else if checkFlagZigbee == 1 {
			checkFlagZigbee = 0
			// 获取ZigBee IEEEADDR
			sliceGetZigbeeIEEE := []int{0xf1,0x2f,0x43,0x2f,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0a,0x00,0x9d,0x00,0x00,0x00,0x00,0x34,0xcc}
			Cmd.CmdZigbeeParsingDeeper(sliceGetZigbeeIEEE, sliceGetZigbeeIEEE[0])

		} else if checkCountTemperature >= timeCountTemperature {
			checkCountTemperature = 0
			// 获取温度采样
			sliceGetTemperature := []int{0xf1,0x2f,0x43,0x2f,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0a,0x00,0xbd,0x00,0x00,0x00,0x00,0x34,0xcc}
			Cmd.CmdZigbeeParsingDeeper(sliceGetTemperature, sliceGetTemperature[0])

		} else if checkCountTemperatureBack >= timeCountTemperatureBack {
			checkCountTemperatureBack = 0
			// 温度采样是否返回
			if Sheard.WDFlagTempreratureBack == 1 {
				Sheard.WDFlagTempreratureBack = 0
				Sheard.WDFlagNoTempreratureCeiling = 0
				fmt.Print("---> OK WDFlagTempreratureBack")
			}
			if Sheard.WDFlagTempreratureBack == 0 {
				Sheard.WDFlagNoTempreratureCeiling += 1
				fmt.Print("---> NOT OK WDFlagTempreratureBack")
			}
			

		} else if checkCountLiVolt >= timeCountLiVolt {
			checkCountLiVolt = 0
			// 获取锂电池电平采样
			sliceGetLiVolt := []int{0xf1,0x2f,0x43,0x2f,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0a,0x00,0xbc,0x00,0x00,0x00,0x00,0x34,0xcc}
			Cmd.CmdZigbeeParsingDeeper(sliceGetLiVolt, sliceGetLiVolt[0])
		} else if checkCountLatLong >= timeCountLatLong {
			checkCountLatLong = 0
			// 更新经纬度时间
			// TODO
		} else if checkCountClock >= timeCountClock {
			checkCountClock = 0
			// 时间计数初始化
			countSysTime := 0
			countPCFTime := 0

			timeSys := make([]int, 3)
			timePCF := make([]int, 7)

			sliceTimeAlarm := []int{0x2f, 0x2f, 0x2f, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0A,0xC0, 0xEF, 0x00, 0x00, 0x00, 0x00, 0x00, 0xCC, 0xFF}

			timeSys[0] = time.Now().Second()
			timeSys[1] = time.Now().Minute()
			timeSys[2] = time.Now().Hour()
			timePCF = Sheard.HandleSharedExecCSoPCFRead() // 秒、分、时、天、星期、月、年

			countSysTime = timeSys[0] + timeSys[1]*60 + timeSys[2]*60*60
			countPCFTime = timePCF[0] = timePCF[1]*60 + timePCF{2}*60*60

			// 系统时间和时钟芯片绝对值差
			if math.Abs(countSysTime - countPCFTime) >= timeAbsSysPCF {
				fmt.Println("---> ERR timeAbsSysPCF.")
				Sheard.HandleSharedCmdOk(22, sliceTimeAlarm[12:19], sliceTimeAlarm[21])

			}
		
		}

	}
}