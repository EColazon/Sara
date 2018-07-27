package handleAlarmUpload

/*
流程:
	系统报警相关方法

Author:
	mengfei.wu@foxmail.com
*/
import (
	"fmt"
)

// 解析报警缓存
func HandleAlarmBuffParsing(alarmBuff []int) {
	fmt.Println("---> handleAlarmBuffParsing.")

	indexEBase := 0

	// 初始化数据长度
	lengthAlarmBuff := len(alarmBuff)
	// 判断数据长度
	if lengthAlarmBuff != 0x0C {
		fmt.Println("---> Error:lengthAlarmBuff.")
		FlagParsing33 = 1 // 解析异常
	} else {
		fmt.Println("---> Right:lengthAlarmBuff.")
		if alarmBuff[0] == CMD33HEAD {
			// 获取有效数据长度
			lengthValue := alarmBuff[4]<<8 | alarmBuff[5]
			if alarmBuff[lengthValue+7] == CMD33TAIL {
				FlagParsing33 = 0 // 解析正常
				// 获取控制码
				commandCode := alarmBuff[2]<<8 | alarmBuff[3]
				switch commandCode {
				case 0x1000: // RTU报警
					if alarmBuff[7] == 0xE1 {
						indexEBase = alarmBuff[6] - 1
						FlagAlarmE1[indexEBase] = 1
						ValueAlarmE1[indexEBase] = alarmBuff[8] << 24 | alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xE1.")
					} else if alarmBuff[7] == 0xE2 {
						indexEBase = alarmBuff[6] - 1
						FlagAlarmE2[indexEBase] = 1
						ValueAlarmE2[indexEBase] = alarmBuff[8] << 24 | alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xE2.")
					} else if alarmBuff[7] == 0xE5 {
						indexEBase = alarmBuff[6] - 1
						FlagAlarmE5[indexEBase] = 1
						ValueAlarmE5[indexEBase] = alarmBuff[8] << 24 | alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xE5.")
					} else if alarmBuff[7] == 0xE6 {
						indexEBase = alarmBuff[6] - 1
						FlagAlarmE6[indexEBase] = 1
						ValueAlarmE6[indexEBase] = alarmBuff[8] << 24 | alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xE6.")
					}
				case 0x1001: // 单灯报警
					if alarmBuff[7] == 0xE7 {
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmE7[indexEBase] = 1
						ValueAlarmE7[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xE7.")
					} else if alarmBuff[7] == 0xE8 {
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmE8[indexEBase] = 1
						ValueAlarmE7[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xE8.")
					}
				case 0x1002: // 数据库读写异常报警
					if alarmBuff[7] == 0xD0 { // DBElec
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmD0[indexEBase] = 1
						ValueAlarmD0[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xD0.")
					} else if alarmBuff[7] == 0xD1 { // DBLamp
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmD1[indexEBase] = 1
						ValueAlarmD1[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xD1.")
					} else if alarmBuff[7] == 0xD2 { // DBNode
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmD2[indexEBase] = 1
						ValueAlarmD2[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xD2.")
					} else if alarmBuff[7] == 0xD3 { // DBOneKeys
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmD3[indexEBase] = 1
						ValueAlarmD3[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xD3.")
					} else if alarmBuff[7] == 0xD4 { // DBTenKeys
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmD4[indexEBase] = 1
						ValueAlarmD4[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xD4.")
					} else if alarmBuff[7] == 0xD5 { // DBTimer
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmD5[indexEBase] = 1
						ValueAlarmD5[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xD5.")
					} else if alarmBuff[7] == 0xD6 {
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmD6[indexEBase] = 1
						ValueAlarmD6[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xD6.")
					} else if alarmBuff[7] == 0xD7 {
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmD7[indexEBase] = 1
						ValueAlarmD7[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xD7.")
					} else if alarmBuff[7] == 0xD8 {
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmD8[indexEBase] = 1
						ValueAlarmD8[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xD8.")
					} else if alarmBuff[7] == 0xD9 {
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmD9[indexEBase] = 1
						ValueAlarmD9[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xD9.")
					} else if alarmBuff[7] == 0xDA {
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmDA[indexEBase] = 1
						ValueAlarmDA[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xDA.")
					} else if alarmBuff[7] == 0xDB {
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmDB[indexEBase] = 1
						ValueAlarmDB[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xDB.")
					} else if alarmBuff[7] == 0xDC {
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmDC[indexEBase] = 1
						ValueAlarmDC[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xDC.")
					} else if alarmBuff[7] == 0xDD {
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmDD[indexEBase] = 1
						ValueAlarmDD[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xDD.")
					} else if alarmBuff[7] == 0xDE {
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmDE[indexEBase] = 1
						ValueAlarmDE[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xDE.")
					} else if alarmBuff[7] == 0xDF {
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmDF[indexEBase] = 1
						ValueAlarmDF[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xDF.")
					}
				case 0x1003: // 命令解析异常报警
					if alarmBuff[7] == 0xC0 { // 33 Fail SR
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmC0[indexEBase] = 1
						ValueAlarmC0[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xC0.")
					} else if alarmBuff[7] == 0xC1 { // 2F Fail SR
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmC1[indexEBase] = 1
						ValueAlarmC1[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xC1.")
					} else if alarmBuff[7] == 0xC2 { // 33 Fail RR
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmC2[indexEBase] = 1
						ValueAlarmC2[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xD2.")
					} else if alarmBuff[7] == 0xD3 { // 
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmC3[indexEBase] = 1
						ValueAlarmC3[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xC3.")
					} else if alarmBuff[7] == 0xC4 { // 
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmC4[indexEBase] = 1
						ValueAlarmC4[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xC4.")
					} else if alarmBuff[7] == 0xC5 { // 
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagAlarmC5[indexEBase] = 1
						ValueAlarmC5[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> Alarm 0xC5.")
					}
				case 0x1004: // 清空日志异常报警
					if alarmBuff[7] == 0x10 { // dbloghearta0
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagDBLog10[indexEBase] = 1
						ValueDBLog10[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> DBLog 0x10.")
					} else if alarmBuff[7] == 0x11 { // dbloghearta1
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagDBLog11[indexEBase] = 1
						ValueDBLog11[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> DBLog 0x11.")
					} else if alarmBuff[7] == 0x12 { // dblogokd0
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagDBLog12[indexEBase] = 1
						ValueDBLog12[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> DBLog 0x12.")
					} else if alarmBuff[7] == 0x13 { // dblogokd1
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagDBLog13[indexEBase] = 1
						ValueDBLog13[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> DBLog 0x13.")
					} else if alarmBuff[7] == 0x14 { // dblogokd2
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagDBLog14[indexEBase] = 1
						ValueDBLog14[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> DBLog 0x14.")
					} else if alarmBuff[7] == 0x15 { // dblogokd3
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagDBLog15[indexEBase] = 1
						ValueDBLog15[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> DBLog 0x15.")
					} else if alarmBuff[7] == 0x16 { // dblogokd4
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagDBLog16[indexEBase] = 1
						ValueDBLog16[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> DBLog 0x16.")
					} else if alarmBuff[7] == 0x17 { // dblogcmderr
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagDBLog17[indexEBase] = 1
						ValueDBLog17[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> DBLog 0x17.")
					} else if alarmBuff[7] == 0x18 { // dblogsysb
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagDBLog18[indexEBase] = 1
						ValueDBLog18[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> DBLog 0x18.")
					} else if alarmBuff[7] == 0x19 { // dblogmodulec
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagDBLog19[indexEBase] = 1
						ValueDBLog19[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> DBLog 0x19.")
					} else if alarmBuff[7] == 0x1A { // dblogalarme
						indexEBase = alarmBuff[6] << 8 | alarmBuff[7]
						FlagDBLog1A[indexEBase] = 1
						ValueDBLog1A[indexEBase] = alarmBuff[9] << 16 | alarmBuff[10] << 8 | alarmBuff[11]
						fmt.Println("---> DBLog 0x1A.")
					} 
					
				default:
					fmt.Println("---> Alarm DoNothing.")
				}
			} else {
				fmt.Println("---> Error:lengthAlarmBuff-Tail.")
				FlagParsing33 = 1 // 解析异常
			}
		} else {
			fmt.Println("---> Error:lengthAlarmBuff-Head.")
			FlagParsing33 = 1 // 解析异常
		}
	}

}

// 回调报警缓存
func handleAlarmBuffRollback() {
	fmt.Println("---> handleAlarmBuffRollback.")

}

// 报警信息发送到服务器
// 上传报警缓存
func handleAlarmBuffUpload2Server(alarmChannel, alarmType, alarmValue int) {
	// 声明26字节报警缓存
	var alarmBuff = make([]int, 26)

	// 声明1字节校验和
	var checkSum int = 0
	// E8报警数据上传内容和其他不同
	if alarmType == 0xE8 {
		handleAlarmBuffUpload2ServerE8(alarmChannel, alarmType, alarmValue)
	} else {
		//初始化头部相关
		alarmBuff[0], alarmBuff[1], alarmBuff[2], alarmBuff[3] = 0x00, 0x00, 0x00, 0x16
		alarmBuff[4], alarmBuff[5], alarmBuff[6], alarmBuff[7], alarmBuff[8] = 0xFF, 0x2F, 0x2F, 0x2F, 0x05

		// 拼组panid地址相关
		for i:= 0; i < 8; i++ {
			alarmBuff[i+9] = i // 暂时填充
		}

		// 拼组报警类型
		switch alarmType {
		case 0xE1, 0xE2, 0xE5, 0xE6:
			alarmBuff[18] = 0xC0 + alarmChannel
		default:
			alarmBuff[18] = alarmChannel
		}

		alarmBuff[19] = alarmType
		alarmBuff[20] = (int(alarmValue)>>24) & 0xFF
		alarmBuff[21] = (int(alarmValue)>>16) & 0xFF
		alarmBuff[22] = (int(alarmValue)>>8) & 0xFF
		alarmBuff[23] = int(alarmValue) & 0xFF
		
		// 计算校验和
		for i := 0; i < 16; i++ {
			checkSum ^= alarmBuff[i+8]
		}

		alarmBuff[24], alarmBuff[25] = checkSum, 0xCC

		// TODO TCP.Send

	}

	
} 

func handleAlarmBuffUpload2ServerE8(alarmChannel, alarmType, alarmValue int) {
	// 声明14字节单灯数据缓存
	var lampBuff = make([]int, 14)
	// 声明(26+14)字节报警缓存
	var alarmBuff = make([]int, 40)

	// 声明1字节校验和
	var checkSum int = 0

	// 获取14字节单灯数据
	for i := 0; i < 14; i++ {
		lampBuff[i] = i // 模拟填充数据
	}


	//初始化头部相关
	alarmBuff[0], alarmBuff[1], alarmBuff[2], alarmBuff[3] = 0x00, 0x00, 0x00, 0x24
	alarmBuff[4], alarmBuff[5], alarmBuff[6], alarmBuff[7], alarmBuff[8] = 0xFF, 0x2F, 0x2F, 0x2F, 0x05

	// 拼组panid地址相关
	for i:= 0; i < 8; i++ {
		alarmBuff[i+9] = i // 暂时填充
	}

	// 拼组报警类型
	alarmBuff[18] = alarmChannel
	alarmBuff[19] = alarmType
	alarmBuff[20] = (int(alarmValue)>>24) & 0xFF
	alarmBuff[21] = (int(alarmValue)>>16) & 0xFF
	alarmBuff[22] = (int(alarmValue)>>8) & 0xFF
	alarmBuff[23] = int(alarmValue) & 0xFF
	
	// 计算校验和
	for i := 0; i < 40; i++ {
		checkSum ^= alarmBuff[i+8]
	}

	alarmBuff[38], alarmBuff[39] = checkSum, 0xCC

	// TODO TCP.Send
}

// 异常上传到服务器
/*
	***数据库
	***其他

*/
// 全局可调用
func HandleAlarmBuffUpload2ServerAbnormity(identNum int) {
	fmt.Println("---> Strat HandleAlarmBuffUpload2ServerAbnormity.")

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0

	sliceCommand := make([]int, 16)
	sliceCommandHead := make([]int, 5)
	sliceCommand33 := make([]int, 11)
	length := len(sliceCommand33) + 1

	sliceCommand33[0] = 0x33
	sliceCommand33[1] = 0xFE
	sliceCommand33[2] = 0x02
	sliceCommand33[3] = 0xFF
	sliceCommand33[4] = 0x5
	sliceCommand33[5] = 0xFE
	sliceCommand33[6] = 0xFE
	sliceCommand33[7] = 0x01
	sliceCommand33[8] = identNum

	for i := 0; i < 9; i++ {
		checkSum ^= sliceCommand33[i]
	}

	sliceCommand33[9] = checkSum
	sliceCommand33[10] = 0x99

	sliceCommandHead[0] = (length>>24) & 0xFF
	sliceCommandHead[1] = (length>>16) & 0xFF
	sliceCommandHead[2] = (length>>8) & 0xFF
	sliceCommandHead[3] = length & 0xFF
	sliceCommandHead[4] = 0xFE

	for i := 0; i < 5; i++ {
		sliceCommand[i] = sliceCommandHead[i]
	}
	for i := 0; i < 11; i++ {
		sliceCommand[i+5] = sliceCommand33[i]
	}

	fmt.Println("---> Alarm Send.")
	return
	// TODO TCP.send
}
