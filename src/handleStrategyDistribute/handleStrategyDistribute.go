package handleStrategyDistribute

import (
	"fmt"
	"time"
	Shared "handleShared"
	Redis "handleRedis"
)

func HandleStrategyDistributeManage() {
	fmt.Println("---> HandleStrategyDistributeManage.")

	for {
		handleStrategyDistribute()
		time.Sleep(1 * time.Second)
	}
}

func handleStrategyDistribute() {
	fmt.Println("---> HandleStrategyDistributeManage.")

	// 特殊策略和节假日策略不执行阶段调光，阶段开关标志
	flagStepMask := 0	
	jsonStepMask := make(map[string]interface{})

	// 回路常开标志
	flagLoopAlwaysON := 0
	// jsonLoopAlwaysON := make(map[string]interface{})

	// 获取当前时间
	timeNowYear 	:= time.Now().Year()
	timeNowMonth	:= time.Now().Month()
	timeNowDay		:= time.Now().Day()
	timeNowWeek	  	:= time.Now().Weekday()
	timeNowHour 	:= time.Now().Hour()
	timeNowMinute 	:= time.Now().Minute()
	timeNowSecond	:= time.Now().Second()

	// 获取回路数量
	loopQuantityBase := Redis.HandleRedisJsonGet(Shared.WDQuantityLoop)
	loopQuantity := loopQuantityBase.(int)

	// loopQuantity == syspara.l_num
	// WDNumSpecial == syspara.ss_num
	for i := 0; i < loopQuantity; i++ {
		for j := 0; j < Shared.WDNumSpecial[i]; j++ {
			if Shared.WDDataStgSpecial[i*150+j*10] == timeNowYear &&  Shared.WDDataStgSpecial[i*150+j*10+1] == int(timeNowMonth) && Shared.WDDataStgSpecial[i*150+j*10+2] == timeNowDay {
				flagStrategySH[i] = 0x01

				timerRONSpecial[i*3] = 0x00
				timerRONSpecial[i*3+1] = Shared.WDDataStgSpecial[i*150+j*10+4]
				timerRONSpecial[i*3+2] = Shared.WDDataStgSpecial[i*150+j*10+3]
				timerROFFSpecial[i*3] = 0x00
				timerROFFSpecial[i*3+1] = Shared.WDDataStgSpecial[i*150+j*10+9]
				timerROFFSpecial[i*3+2] = Shared.WDDataStgSpecial[i*150+j*10+8]
				break
				
			}
		}
	}

	// 节假日策略
	for i := 0; i < loopQuantity; i++ {
		for j := 0; j < Shared.WDNumHoliday[i]; j++ {
			if Shared.WDDataStgHoliday[i*150+j*10] == timeNowYear && Shared.WDDataStgHoliday[i*150+j*10+1] == int(timeNowMonth) && Shared.WDDataStgHoliday[i*150+j*10+2] == timeNowDay {
				if flagStrategySH[i] == 0 { // 如果今天特殊策略起作用,则不进行节假日策略动作
					flagStrategySH[i] = 0x02

					timerRONHoliday[i*3] = 0x00
					timerRONHoliday[i*3+1] = Shared.WDDataStgHoliday[i*150+j*10+4]
					timerRONHoliday[i*3+2] = Shared.WDDataStgHoliday[i*150+j*10+3]
					timerROFFHoliday[i*3] = 0x00
					timerROFFHoliday[i*3+1] = Shared.WDDataStgHoliday[i*150+j*10+9]
					timerROFFHoliday[i*3+2] = Shared.WDDataStgHoliday[i*150+j*10+8]
					break

				}
			}
		}
	}



	for i := 0; i < loopQuantity; i++ {
		if flagStrategySH[i] == 0x01 || flagStrategySH[i] == 0x02 {
			flagStepMask = 0x55
			jsonStepMask[Shared.WDFlagMaskStep] = flagStepMask
			Redis.HandleRedisJsonInsert(Shared.WDFlagMaskStep, jsonStepMask)			
		}
	}

	// 获取回路常开标志
	flagLoopAlwaysONBase := Redis.HandleRedisJsonGet(Shared.WDFlagLoopAlwaysON)
	flagLoopAlwaysON = flagLoopAlwaysONBase.(int)
	for i := 0; i < loopQuantity; i++ {
		if flagStrategySH[i] == 0x01 {
			if flagLoopAlwaysON != 0x55 {
				// strategy_special()
				fmt.Println("---> strategy_special")
				handleStrategySpecial(i, timeNowHour, timeNowMinute, timeNowSecond, int(timeNowWeek))
			}
			// 定时开关时间
			Shared.WDTimerRON[i*3] = timerRONSpecial[i*3]
			Shared.WDTimerRON[i*3+1] = timerRONSpecial[i*3+1]
			Shared.WDTimerRON[i*3+2] = timerRONSpecial[i*3+2]
			Shared.WDTimerROFF[i*3] = timerROFFSpecial[i*3]
			Shared.WDTimerROFF[i*3+1] = timerROFFSpecial[i*3+1]
			Shared.WDTimerROFF[i*3+2] = timerROFFSpecial[i*3+2]
		} else if flagStrategySH[i] == 0x02 {
			if flagLoopAlwaysON != 0x55 {
				// strategy_holiday()
				fmt.Println("---> strategy_special")
				handleStrategyHoliday(i, timeNowHour, timeNowMinute, timeNowSecond, int(timeNowWeek))
			}
			// 定时开关时间
			Shared.WDTimerRON[i*3] = timerRONHoliday[i*3]
			Shared.WDTimerRON[i*3+1] = timerRONHoliday[i*3+1]
			Shared.WDTimerRON[i*3+2] = timerRONHoliday[i*3+2]
			Shared.WDTimerROFF[i*3] = timerROFFHoliday[i*3]
			Shared.WDTimerROFF[i*3+1] = timerROFFHoliday[i*3+1]
			Shared.WDTimerROFF[i*3+2] = timerROFFHoliday[i*3+2]
		} else {
			if Shared.WDTypeStrategyON[i] == 0x01 {
				// get_JW_onoff_time(i)
				handleGetTimeLatiLongiONOFF(i)
				Shared.WDTimerRON[i*3] = timerOffsetON[0]
				Shared.WDTimerRON[i*3+1] = timerOffsetON[1]
				Shared.WDTimerRON[i*3+2] = timerOffsetON[2]
			} else if Shared.WDTypeStrategyON[i] == 0x02 {
				Shared.WDTimerRON[i*3] = Shared.WDTimerRONTime[i*3]
				Shared.WDTimerRON[i*3+1] = Shared.WDTimerRONTime[i*3+1]
				Shared.WDTimerRON[i*3+2] = Shared.WDTimerRONTime[i*3+2]
			} else if Shared.WDTypeStrategyON[i] == 0x03 {
				// get_week_onoff_time(i, systime.week)
				handleGetTimeWeekONOFF(i, int(timeNowWeek))
				Shared.WDTimerRON[i*3] = timerOffsetON[0]
				Shared.WDTimerRON[i*3+1] = timerOffsetON[1]
				Shared.WDTimerRON[i*3+2] = timerOffsetON[2]
			} else {
				// get_JW_onoff_time(i)
				handleGetTimeLatiLongiONOFF(i)
				Shared.WDTimerRON[i*3] = timerOffsetON[0]
				Shared.WDTimerRON[i*3+1] = timerOffsetON[1]
				Shared.WDTimerRON[i*3+2] = timerOffsetON[2]
			}

			// 关
			if Shared.WDTypeStrategyOFF[i] == 0x01 {
				// get_JW_onoff_time(i)
				handleGetTimeLatiLongiONOFF(i)
				Shared.WDTimerROFF[i*3] = timerOffsetOFF[0]
				Shared.WDTimerROFF[i*3+1] = timerOffsetOFF[1]
				Shared.WDTimerROFF[i*3+2] = timerOffsetOFF[2]

			} else if Shared.WDTypeStrategyOFF[i] == 0x02 {
				Shared.WDTimerROFF[i*3] = Shared.WDTimerROFFTime[i*3]
				Shared.WDTimerROFF[i*3+1] = Shared.WDTimerROFFTime[i*3+1]
				Shared.WDTimerROFF[i*3+1] = Shared.WDTimerROFFTime[i*3+2]
			} else if Shared.WDTypeStrategyOFF[i] == 0x03 {
				// get_week_onoff_time(i, systime.week)
				handleGetTimeWeekONOFF(i, int(timeNowWeek))
				Shared.WDTimerROFF[i*3] = timerOffsetOFF[0]
				Shared.WDTimerROFF[i*3+1] = timerOffsetOFF[1]
				Shared.WDTimerROFF[i*3+2] = timerOffsetOFF[2]
			} else {
				// get_JW_onoff_time(i)
				handleGetTimeLatiLongiONOFF(i)
				Shared.WDTimerROFF[i*3] = timerOffsetOFF[0]
				Shared.WDTimerROFF[i*3+1] = timerOffsetOFF[1]
				Shared.WDTimerROFF[i*3+2] = timerOffsetOFF[2]

			}

			if flagLoopAlwaysON != 0x55 {
				// time_compare_relay()
				fmt.Println("---> handleTimeCompareRelay(i,1,1,hour,minute,second,week)")
				fmt.Println("---> handleTimeCompareRelay(i,0,1,hour,minute,second,week)")
				handleTimeCompareRelay(i, 1, 1, timeNowHour, timeNowMinute, timeNowSecond, int(timeNowWeek))
				handleTimeCompareRelay(i, 0, 1, timeNowHour, timeNowMinute, timeNowSecond, int(timeNowWeek))
			}
		}
		
	}

}