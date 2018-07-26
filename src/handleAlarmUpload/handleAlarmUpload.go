package handleAlarmUpload

import (
	"fmt"
	"time"

	sharePara "handleShared"
)

func HandleModuleAlarm() {
	// 声明变量
	// 报警发送间隔:默认60s/once
	alarmTimeInterval := 60
	// 计数:1s/once
	alarmTimeCount 	  := 0
	index := 0

	for {
		// 计数自增
		alarmTimeCount += 1
		time.Sleep(1 * time.Second)

		if alarmTimeCount >= alarmTimeInterval {
			fmt.Println("---> alarmTimeCount: ", alarmTimeCount)
			// 报警计数清零
			alarmTimeCount = 0
			// 检测单回路数量报警

			// 电压报警
			for index = 0; index < sharePara.WDValueLoop; index++ {
				if FlagAlarmE1[index] == 1 {
					// 清零报警标志
					FlagAlarmE1[index] = 0
					handleAlarmBuffUpload2Server(index+1, 0xE1, ValueAlarmE1[index]) // 电压放大100倍
				}
			}
			// 电流报警
			for index = 0; index < sharePara.WDValueLoop; index++ {
				if FlagAlarmE2[index] == 1 {
					// 清零报警标志
					FlagAlarmE2[index] = 0
					handleAlarmBuffUpload2Server(index+1, 0xE2, ValueAlarmE2[index]) // 电流放大100倍
				}
			}
			// 意外灭灯报警
			for index = 0; index < sharePara.WDValueLoop; index++ {
				if FlagAlarmE5[index] == 1 {
					// 清零报警标志
					FlagAlarmE5[index] = 0
					handleAlarmBuffUpload2Server(index+1, 0xE5, ValueAlarmE5[index]) // 电压放大100倍
				}
			}
			// 意外亮灯报警
			for index = 0; index < sharePara.WDValueLoop; index++ {
				if FlagAlarmE6[index] == 1 {
					// 清零报警标志
					FlagAlarmE6[index] = 0
					handleAlarmBuffUpload2Server(index+1, 0xE6, ValueAlarmE6[index]) // 电压放大100倍
				}
			}
			// 单灯失恋报警
			for index = 0; index < sharePara.WDValueLamp+1; index++ {
				if FlagAlarmE7[index] == 1 {
					// TODO新增E7报警标志判断
					// 清零报警标志
					FlagAlarmE7[index] = 0
					// TODO新增sl_eamask!=0x55判断
					handleAlarmBuffUpload2Server(index+1, 0xE7, ValueAlarmE7[index]) // 电压放大100倍
				}
			}
			// 单灯异常报警
			for index = 0; index < sharePara.WDValueLamp+1; index++ {
				if FlagAlarmE8[index] == 1 {
					// TODO新增E8报警标志判断
					// 清零报警标志
					FlagAlarmE8[index] = 0
					// TODO新增sl_eamask!=0x55判断
					handleAlarmBuffUpload2ServerE8(index+1, 0xE8, ValueAlarmE8[index]) // 电压放大100倍
				}
			}

			//数据库读写异常报警
			// DBElec
			if FlagAlarmD0[0] == 1 {
				// 清零报警标志
				FlagAlarmD0[0] = 0
				// 上传报警信心
				handleAlarmBuffUpload2Server(1, 0xD0, ValueAlarmD0[0])
			}
			// DBLamp
			if FlagAlarmD1[0] == 1 {
				// 清零报警标志
				FlagAlarmD1[0] = 0
				// 上传报警信心
				handleAlarmBuffUpload2Server(1, 0xD0, ValueAlarmD1[0])
			}
			// DBNode
			if FlagAlarmD2[0] == 1 {
				// 清零报警标志
				FlagAlarmD2[0] = 0
				// 上传报警信心
				handleAlarmBuffUpload2Server(1, 0xD0, ValueAlarmD2[0])
			}
			// DBOneKeys
			if FlagAlarmD3[0] == 1 {
				// 清零报警标志
				FlagAlarmD3[0] = 0
				// 上传报警信心
				handleAlarmBuffUpload2Server(1, 0xD0, ValueAlarmD3[0])
			}
			// DBTenKeys
			if FlagAlarmD4[0] == 1 {
				// 清零报警标志
				FlagAlarmD4[0] = 0
				// 上传报警信心
				handleAlarmBuffUpload2Server(1, 0xD0, ValueAlarmD4[0])
			}
			// DBTimer
			if FlagAlarmD5[0] == 1 {
				// 清零报警标志
				FlagAlarmD5[0] = 0
				// 上传报警信心
				handleAlarmBuffUpload2Server(1, 0xD0, ValueAlarmD5[0])
			}
		}

	}
}