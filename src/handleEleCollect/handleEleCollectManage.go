package handleEleCollect

import (
	"fmt"
	"time"

	Redis "handleRedis"
	Shared "handleShared"
)

func HandleEleCollectManage() {
	fmt.Println("---> HandleEleCollectManage.")
	// 初始化变量
	jsonModePower := make(map[string]interface{})
	// modePower := Redis.HandleRedisJsonGet(Shared.WDSysPowerMode)
	sliceEnergyIntC1 := make([]int, 4)
	sliceEnergyIntC2 := make([]int, 4)
	sliceEnergyIntC3 := make([]int, 4)
	sliceEnergyIntC4 := make([]int, 4)
	sliceEnergyIntC5 := make([]int, 4)
	sliceEnergyIntC6 := make([]int, 4)
	energyIntC1 := 0
	energyIntC2 := 0
	energyIntC3 := 0
	energyIntC4 := 0
	energyIntC5 := 0
	energyIntC6 := 0

	// 计数相关
	countEleDataGet 	:= 0
	countEleCheckReg 	:= 0
	countEleSaveW 		:= 0
	countEleCheckAlarm	:= 0

	for {
		countEleDataGet 		+= 1
		countEleCheckReg 		+= 1
		countEleSaveW 			+= 1
		countEleCheckAlarm 		+= 1
		time.Sleep(1 * time.Second)
		handleCheckFlagClear()
		
		if countEleDataGet > 5 {
			countEleDataGet = 0
			handleGetDataEle()
		} else if countEleCheckReg > 60 {
			countEleCheckReg = 0
			// Shared.HandleSharedExecCSoGpioRN8209RegCheck(1)
			// Shared.HandleSharedExecCSoGpioRN8209RegCheck(2)
			// Shared.HandleSharedExecCSoGpioRN8209RegCheck(3)
		} else if countEleSaveW >= 60 * 30 {
			countEleSaveW = 0
			// 获取供电模式
			// modePower = Redis.HandleRedisJsonGet(Shared.WDSysPowerMode)
			modePower := Redis.HandleRedisJsonGet(Shared.WDSysPowerMode)
			if modePower == Shared.WDPOWERMODEAC {
				fmt.Println("---> AC.")
			}
			
			// WriteProcess
			// C1
			sliceEnergyIntC1[0] = (int(valueEleC1["energy"]*10000) << 24) & 0xff			
			sliceEnergyIntC1[1] = (int(valueEleC1["energy"]*10000) << 16) & 0xff			
			sliceEnergyIntC1[2] = (int(valueEleC1["energy"]*10000) << 8) & 0xff			
			sliceEnergyIntC1[3] = (int(valueEleC1["energy"]*10000)) & 0xff
			
			energyIntC1 = sliceEnergyIntC1[0]<<24|sliceEnergyIntC1[1]<<16|sliceEnergyIntC1[2]<<8|sliceEnergyIntC1[3]	
			jsonModePower[Shared.WDEleEnergyC1] = energyIntC1
			Redis.HandleRedisJsonInsert(Shared.WDEleEnergyC1, jsonModePower)

			// C2
			sliceEnergyIntC2[0] = (int(valueEleC2["energy"]*10000) << 24) & 0xff			
			sliceEnergyIntC2[1] = (int(valueEleC2["energy"]*10000) << 16) & 0xff			
			sliceEnergyIntC2[2] = (int(valueEleC2["energy"]*10000) << 8) & 0xff			
			sliceEnergyIntC2[3] = (int(valueEleC2["energy"]*10000)) & 0xff
			
			energyIntC2 = sliceEnergyIntC2[0]<<24|sliceEnergyIntC2[1]<<16|sliceEnergyIntC2[2]<<8|sliceEnergyIntC2[3]	
			jsonModePower[Shared.WDEleEnergyC2] = energyIntC2
			Redis.HandleRedisJsonInsert(Shared.WDEleEnergyC2, jsonModePower)

			// C3
			sliceEnergyIntC3[0] = (int(valueEleC3["energy"]*10000) << 24) & 0xff			
			sliceEnergyIntC3[1] = (int(valueEleC3["energy"]*10000) << 16) & 0xff			
			sliceEnergyIntC3[2] = (int(valueEleC3["energy"]*10000) << 8) & 0xff			
			sliceEnergyIntC3[3] = (int(valueEleC3["energy"]*10000)) & 0xff
			
			energyIntC3 = sliceEnergyIntC3[0]<<24|sliceEnergyIntC3[1]<<16|sliceEnergyIntC3[2]<<8|sliceEnergyIntC3[3]	
			jsonModePower[Shared.WDEleEnergyC3] = energyIntC3
			Redis.HandleRedisJsonInsert(Shared.WDEleEnergyC3, jsonModePower)

			// C4
			sliceEnergyIntC4[0] = (int(valueEleC4["energy"]*10000) << 24) & 0xff			
			sliceEnergyIntC4[1] = (int(valueEleC4["energy"]*10000) << 16) & 0xff			
			sliceEnergyIntC4[2] = (int(valueEleC4["energy"]*10000) << 8) & 0xff			
			sliceEnergyIntC4[3] = (int(valueEleC4["energy"]*10000)) & 0xff
			
			energyIntC4 = sliceEnergyIntC4[0]<<24|sliceEnergyIntC4[1]<<16|sliceEnergyIntC4[2]<<8|sliceEnergyIntC4[3]	
			jsonModePower[Shared.WDEleEnergyC4] = energyIntC4
			Redis.HandleRedisJsonInsert(Shared.WDEleEnergyC4, jsonModePower)

			// C5
			sliceEnergyIntC5[0] = (int(valueEleC5["energy"]*10000) << 24) & 0xff			
			sliceEnergyIntC5[1] = (int(valueEleC5["energy"]*10000) << 16) & 0xff			
			sliceEnergyIntC5[2] = (int(valueEleC5["energy"]*10000) << 8) & 0xff			
			sliceEnergyIntC5[3] = (int(valueEleC5["energy"]*10000)) & 0xff
			
			energyIntC5 = sliceEnergyIntC5[0]<<24|sliceEnergyIntC5[1]<<16|sliceEnergyIntC5[2]<<8|sliceEnergyIntC5[3]	
			jsonModePower[Shared.WDEleEnergyC5] = energyIntC5
			Redis.HandleRedisJsonInsert(Shared.WDEleEnergyC5, jsonModePower)

			// C6
			sliceEnergyIntC6[0] = (int(valueEleC6["energy"]*10000) << 24) & 0xff			
			sliceEnergyIntC6[1] = (int(valueEleC6["energy"]*10000) << 16) & 0xff			
			sliceEnergyIntC6[2] = (int(valueEleC6["energy"]*10000) << 8) & 0xff			
			sliceEnergyIntC6[3] = (int(valueEleC6["energy"]*10000)) & 0xff
			
			energyIntC6 = sliceEnergyIntC6[0]<<24|sliceEnergyIntC6[1]<<16|sliceEnergyIntC6[2]<<8|sliceEnergyIntC6[3]	
			jsonModePower[Shared.WDEleEnergyC6] = energyIntC6
			Redis.HandleRedisJsonInsert(Shared.WDEleEnergyC6, jsonModePower)

		

		} else if countEleCheckAlarm > 5 {
			countEleCheckAlarm = 0
			handleCheckAlarmForElec()
		}

	}

}