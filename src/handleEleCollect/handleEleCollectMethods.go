package handleEleCollect

import (
	"fmt"
	"time"
	"math"
	// Shread "handleShared"
	// Redis "handleRedis"
	AlarmUpload "handleAlarmUpload"
)

const (
	timeClearHour	= 12
	timeClearMinute = 0
)

var (
	valueEleC1 = map[string]float64{"volt":0,"current":0,"power":0,"energy":0,"pf":0,"ki":0,"ku":0,"kp":0,"current_limit_max":0,"current_limit_min":0,"volt_limit_max":0,"volt_limit_min":0}
	valueEleC2 = map[string]float64{"volt":0,"current":0,"power":0,"energy":0,"pf":0,"ki":0,"ku":0,"kp":0,"current_limit_max":0,"current_limit_min":0,"volt_limit_max":0,"volt_limit_min":0}
	valueEleC3 = map[string]float64{"volt":0,"current":0,"power":0,"energy":0,"pf":0,"ki":0,"ku":0,"kp":0,"current_limit_max":0,"current_limit_min":0,"volt_limit_max":0,"volt_limit_min":0}
	valueEleC4 = map[string]float64{"volt":0,"current":0,"power":0,"energy":0,"pf":0,"ki":0,"ku":0,"kp":0,"current_limit_max":0,"current_limit_min":0,"volt_limit_max":0,"volt_limit_min":0}
	valueEleC5 = map[string]float64{"volt":0,"current":0,"power":0,"energy":0,"pf":0,"ki":0,"ku":0,"kp":0,"current_limit_max":0,"current_limit_min":0,"volt_limit_max":0,"volt_limit_min":0}
	valueEleC6 = map[string]float64{"volt":0,"current":0,"power":0,"energy":0,"pf":0,"ki":0,"ku":0,"kp":0,"current_limit_max":0,"current_limit_min":0,"volt_limit_max":0,"volt_limit_min":0}

	countRTUE1 = make([]int, 6) // 电压
	countRTUE2 = make([]int, 6) // 电流
	countRTUE5 = make([]int, 6) // 意外亮灯
	countRTUE6 = make([]int, 6) // 意外灭灯

	flagRTUE1 = make([]int, 6)
	flagRTUE2 = make([]int, 6)
	flagRTUE5 = make([]int, 6)
	flagRTUE6 = make([]int, 6)

	flagE1V1  = 0
	flagE1V2  = 0
	flagE1V3  = 0
	flagE1V4  = 0
	flagE1V5  = 0
	flagE1V6  = 0

	flagE2I1  = 0 
	flagE2I2  = 0 
	flagE2I3  = 0 
	flagE2I4  = 0 
	flagE2I5  = 0 
	flagE2I6  = 0 

)

func handleInitEle() {
	fmt.Println("---> handleInitEle.")

	readData1 := make([]int, 4)
	readData2 := make([]int, 4)
	readData3 := make([]int, 4)
	readData4 := make([]int, 4)
	readData5 := make([]int, 4)
	readData6 := make([]int, 4)
	// 电流比例系数
	// 舍弃hex_to_f&&f_to_hex转换方法
	// 改为原始数据*10000倍,然后取整数，再转换为四个int
	/* eg: 0.1023
	 * 		0.1023*10000 = 1023
	 *  	Int00 = (1023>>24)&0xff = 0
	 *		iInt01= (1023>>16)&0xff = 0
	 *		iInt02= (1023>>8)&0xff = 3
	 *		iInt03= 1023&0xff = 255
	 *	
	 * 		1023 = Int00<<24|Int01<<16|Int02<<8|Int03
	 * 		0.1023 = 1023 / 10000
	 */
	// KIA
	for i := 0; i < 4; i++ {
		// readData1[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KIA1BIT0+i)
		// readData2[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KIA2BIT0+i)
		// readData3[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KIA3BIT0+i)
		// readData4[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KIA4BIT0+i)
		// readData5[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KIA5BIT0+i)
		// readData6[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KIA6BIT0+i)
		readData1[i] = i
		readData2[i] = i
		readData3[i] = i
		readData4[i] = i
		readData5[i] = i
		readData6[i] = i
	}
	valueEleC1["ki"] = float64(readData1[0]<<24|readData1[1]<<16|readData1[2]<<8|readData1[3]) / 10000
	if math.IsNaN(valueEleC1["ki"]) == true {
		valueEleC1["ki"] = 0.0
	}
	valueEleC2["ki"] = float64(readData2[0]<<24|readData2[1]<<16|readData2[2]<<8|readData2[3]) / 10000
	if math.IsNaN(valueEleC2["ki"]) == true {
		valueEleC2["ki"] = 0.0
	}
	valueEleC3["ki"] = float64(readData3[0]<<24|readData3[1]<<16|readData3[2]<<8|readData3[3]) / 10000
	if math.IsNaN(valueEleC3["ki"]) == true {
		valueEleC3["ki"] = 0.0
	}
	valueEleC4["ki"] = float64(readData4[0]<<24|readData4[1]<<16|readData4[2]<<8|readData4[3]) / 10000
	if math.IsNaN(valueEleC4["ki"]) == true {
		valueEleC4["ki"] = 0.0
	}
	valueEleC5["ki"] = float64(readData5[0]<<24|readData5[1]<<16|readData5[2]<<8|readData5[3]) / 10000
	if math.IsNaN(valueEleC5["ki"]) == true {
		valueEleC5["ki"] = 0.0
	}
	valueEleC6["ki"] = float64(readData6[0]<<24|readData6[1]<<16|readData6[2]<<8|readData6[3]) / 10000
	if math.IsNaN(valueEleC6["ki"]) == true {
		valueEleC6["ki"] = 0.0
	}

	// KUA
	for i := 0; i < 4; i++ {
		// readData1[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KUA1BIT0+i)
		// readData2[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KUA2BIT0+i)
		// readData3[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KUA3BIT0+i)
		// readData4[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KUA4BIT0+i)
		// readData5[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KUA5BIT0+i)
		// readData6[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KUA6BIT0+i)
		readData1[i] = i
		readData2[i] = i
		readData3[i] = i
		readData4[i] = i
		readData5[i] = i
		readData6[i] = i
	}
	valueEleC1["ku"] = float64(readData1[0]<<24|readData1[1]<<16|readData1[2]<<8|readData1[3]) / 10000
	if math.IsNaN(valueEleC1["ku"]) == true {
		valueEleC1["ku"] = 0.0
	}
	valueEleC2["ku"] = float64(readData2[0]<<24|readData2[1]<<16|readData2[2]<<8|readData2[3]) / 10000
	if math.IsNaN(valueEleC2["ku"]) == true {
		valueEleC2["ku"] = 0.0
	}
	valueEleC3["ku"] = float64(readData3[0]<<24|readData3[1]<<16|readData3[2]<<8|readData3[3]) / 10000
	if math.IsNaN(valueEleC3["ku"]) == true {
		valueEleC3["ku"] = 0.0
	}
	valueEleC4["ku"] = float64(readData4[0]<<24|readData4[1]<<16|readData4[2]<<8|readData4[3]) / 10000
	if math.IsNaN(valueEleC4["ku"]) == true {
		valueEleC4["ku"] = 0.0
	}
	valueEleC5["ku"] = float64(readData5[0]<<24|readData5[1]<<16|readData5[2]<<8|readData5[3]) / 10000
	if math.IsNaN(valueEleC5["ku"]) == true {
		valueEleC5["ku"] = 0.0
	}
	valueEleC6["ku"] = float64(readData6[0]<<24|readData6[1]<<16|readData6[2]<<8|readData6[3]) / 10000
	if math.IsNaN(valueEleC6["ku"]) == true {
		valueEleC6["ku"] = 0.0
	}

	// KPA
	for i := 0; i < 4; i++ {
		// readData1[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KPA1BIT0+i)
		// readData2[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KPA2BIT0+i)
		// readData3[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KPA3BIT0+i)
		// readData4[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KPA4BIT0+i)
		// readData5[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KPA5BIT0+i)
		// readData6[i] = Shread.HandleSharedExecCSoI2C0Read(Shread.WDRN1ADDR54, Shread.WDRN1KPA6BIT0+i)
		readData1[i] = i
		readData2[i] = i
		readData3[i] = i
		readData4[i] = i
		readData5[i] = i
		readData6[i] = i
	}
	valueEleC1["kp"] = float64(readData1[0]<<24|readData1[1]<<16|readData1[2]<<8|readData1[3]) / 10000
	if math.IsNaN(valueEleC1["kp"]) == true {
		valueEleC1["kp"] = 0.0
	}
	valueEleC2["kp"] = float64(readData2[0]<<24|readData2[1]<<16|readData2[2]<<8|readData2[3]) / 10000
	if math.IsNaN(valueEleC2["kp"]) == true {
		valueEleC2["kp"] = 0.0
	}
	valueEleC3["kp"] = float64(readData3[0]<<24|readData3[1]<<16|readData3[2]<<8|readData3[3]) / 10000
	if math.IsNaN(valueEleC3["kp"]) == true {
		valueEleC3["kp"] = 0.0
	}
	valueEleC4["kp"] = float64(readData4[0]<<24|readData4[1]<<16|readData4[2]<<8|readData4[3]) / 10000
	if math.IsNaN(valueEleC4["kp"]) == true {
		valueEleC4["kp"] = 0.0
	}
	valueEleC5["kp"] = float64(readData5[0]<<24|readData5[1]<<16|readData5[2]<<8|readData5[3]) / 10000
	if math.IsNaN(valueEleC5["kp"]) == true {
		valueEleC5["kp"] = 0.0
	}
	valueEleC6["kp"] = float64(readData6[0]<<24|readData6[1]<<16|readData6[2]<<8|readData6[3]) / 10000
	if math.IsNaN(valueEleC6["kp"]) == true {
		valueEleC6["kp"] = 0.0
	}
}

func handleCheckFlagClear() {
	// 获取时间
	timeNowHour 	:= time.Now().Hour()
	timeNowMinute 	:= time.Now().Minute()

	if timeNowHour == timeClearHour && timeNowMinute == timeClearMinute {
		flagE1V1  = 0
		flagE1V2  = 0
		flagE1V3  = 0
		flagE1V4  = 0
		flagE1V5  = 0
		flagE1V6  = 0

		flagE2I1  = 0 
		flagE2I2  = 0 
		flagE2I3  = 0 
		flagE2I4  = 0 
		flagE2I5  = 0 
		flagE2I6  = 0 
	}
}



func handleGetDataEle() {
	fmt.Println("---> handleGetDataEle.")
	// sliceParms := make([]int, 4)
	// sliceVarms := make([]int, 3)
	// sliceIarms := make([]int, 3)
	// sliceWarms := make([]int, 3)

	// // jsonCEnergy := make(map[string]interface{})
	// energyC1 := 0
	// energyC2 := 0
	// energyC3 := 0
	// energyC4 := 0
	// energyC5 := 0
	// energyC6 := 0

	// sliceParms = Shread.HandleSharedExecCSoGpioRN8209GetRegParams(1)
	// valueEleC1["power"] = float64(sliceParms[0]<<24|sliceParms[1]<<16|sliceParms[2]<<8|sliceParms[3]) * valueEleC1["kp"]
	// sliceParms = Shread.HandleSharedExecCSoGpioRN8209GetRegParams(2)
	// valueEleC2["power"] = float64(sliceParms[0]<<24|sliceParms[1]<<16|sliceParms[2]<<8|sliceParms[3]) * valueEleC2["kp"]
	// sliceParms = Shread.HandleSharedExecCSoGpioRN8209GetRegParams(3)
	// valueEleC3["power"] = float64(sliceParms[0]<<24|sliceParms[1]<<16|sliceParms[2]<<8|sliceParms[3]) * valueEleC3["kp"]
	// sliceParms = Shread.HandleSharedExecCSoGpioRN8209GetRegParams(4)
	// valueEleC4["power"] = float64(sliceParms[0]<<24|sliceParms[1]<<16|sliceParms[2]<<8|sliceParms[3]) * valueEleC4["kp"]
	// sliceParms = Shread.HandleSharedExecCSoGpioRN8209GetRegParams(5)
	// valueEleC5["power"] = float64(sliceParms[0]<<24|sliceParms[1]<<16|sliceParms[2]<<8|sliceParms[3]) * valueEleC5["kp"]
	// sliceParms = Shread.HandleSharedExecCSoGpioRN8209GetRegParams(6)
	// valueEleC6["power"] = float64(sliceParms[0]<<24|sliceParms[1]<<16|sliceParms[2]<<8|sliceParms[3]) * valueEleC6["kp"]

	// sliceVarms = Shread.HandleSharedExecCSoGpioRN8209GetRegVarms(1)
	// valueEleC1["volt"] = float64(sliceVarms[0]<<16|sliceVarms[1]<<8|sliceVarms[2]) * valueEleC1["ku"]
	// sliceVarms = Shread.HandleSharedExecCSoGpioRN8209GetRegVarms(2)
	// valueEleC2["volt"] = float64(sliceVarms[0]<<16|sliceVarms[1]<<8|sliceVarms[2]) * valueEleC2["ku"]
	// sliceVarms = Shread.HandleSharedExecCSoGpioRN8209GetRegVarms(3)
	// valueEleC3["volt"] = float64(sliceVarms[0]<<16|sliceVarms[1]<<8|sliceVarms[2]) * valueEleC3["ku"]
	// sliceVarms = Shread.HandleSharedExecCSoGpioRN8209GetRegVarms(4)
	// valueEleC4["volt"] = float64(sliceVarms[0]<<16|sliceVarms[1]<<8|sliceVarms[2]) * valueEleC4["ku"]
	// sliceVarms = Shread.HandleSharedExecCSoGpioRN8209GetRegVarms(5)
	// valueEleC5["volt"] = float64(sliceVarms[0]<<16|sliceVarms[1]<<8|sliceVarms[2]) * valueEleC5["ku"]
	// sliceVarms = Shread.HandleSharedExecCSoGpioRN8209GetRegVarms(6)
	// valueEleC6["volt"] = float64(sliceVarms[0]<<16|sliceVarms[1]<<8|sliceVarms[2]) * valueEleC6["ku"]

	// sliceIarms = Shread.HandleSharedExecCSoGpioRN8209GetRegIarms(1)
	// valueEleC1["current"] = float64(sliceIarms[0]<<16|sliceIarms[1]<<8|sliceIarms[2]) * valueEleC1["ki"]
	// if valueEleC1["current"] < 0.5 && valueEleC1["power"] < 0.1 {
	// 	valueEleC1["current"] = 0
	// }
	// sliceIarms = Shread.HandleSharedExecCSoGpioRN8209GetRegIarms(2)
	// valueEleC2["current"] = float64(sliceIarms[0]<<16|sliceIarms[1]<<8|sliceIarms[2]) * valueEleC2["ki"]
	// if valueEleC2["current"] < 0.5 && valueEleC2["power"] < 0.1 {
	// 	valueEleC2["current"] = 0
	// }
	// sliceIarms = Shread.HandleSharedExecCSoGpioRN8209GetRegIarms(3)
	// valueEleC3["current"] = float64(sliceIarms[0]<<16|sliceIarms[1]<<8|sliceIarms[2]) * valueEleC3["ki"]
	// if valueEleC3["current"] < 0.5 && valueEleC3["power"] < 0.1 {
	// 	valueEleC3["current"] = 0
	// }
	// sliceIarms = Shread.HandleSharedExecCSoGpioRN8209GetRegIarms(4)
	// valueEleC4["current"] = float64(sliceIarms[0]<<16|sliceIarms[1]<<8|sliceIarms[2]) * valueEleC4["ki"]
	// if valueEleC4["current"] < 0.5 && valueEleC4["power"] < 0.1 {
	// 	valueEleC4["current"] = 0
	// }
	// sliceIarms = Shread.HandleSharedExecCSoGpioRN8209GetRegIarms(5)
	// valueEleC5["current"] = float64(sliceIarms[0]<<16|sliceIarms[1]<<8|sliceIarms[2]) * valueEleC5["ki"]
	// if valueEleC5["current"] < 0.5 && valueEleC5["power"] < 0.1 {
	// 	valueEleC5["current"] = 0
	// }
	// sliceIarms = Shread.HandleSharedExecCSoGpioRN8209GetRegIarms(6)
	// valueEleC6["current"] = float64(sliceIarms[0]<<16|sliceIarms[1]<<8|sliceIarms[2]) * valueEleC6["ki"]
	// if valueEleC6["current"] < 0.5 && valueEleC6["power"] < 0.1 {
	// 	valueEleC6["current"] = 0
	// }

	// // 获取energyC*
	// energyC1 = Redis.HandleRedisJsonGet(Shread.WDEleEnergyC1)
	// energyC2 = Redis.HandleRedisJsonGet(Shread.WDEleEnergyC2)
	// energyC3 = Redis.HandleRedisJsonGet(Shread.WDEleEnergyC3)
	// energyC4 = Redis.HandleRedisJsonGet(Shread.WDEleEnergyC4)
	// energyC5 = Redis.HandleRedisJsonGet(Shread.WDEleEnergyC5)
	// energyC6 = Redis.HandleRedisJsonGet(Shread.WDEleEnergyC6)
	// sliceWarms = Shread.HandleSharedExecCSoGpioGetRegWarms(1)
	// valueEleC1["energy"] = float64(sliceWarms[0]<<16|sliceWarms[1]<<8|sliceWarms[2]) / 3200.0 + float64(energyC1)
	// sliceWarms = Shread.HandleSharedExecCSoGpioGetRegWarms(2)
	// valueEleC2["energy"] = float64(sliceWarms[0]<<16|sliceWarms[1]<<8|sliceWarms[2]) / 3200.0 + float64(energyC2)
	// sliceWarms = Shread.HandleSharedExecCSoGpioGetRegWarms(3)
	// valueEleC3["energy"] = float64(sliceWarms[0]<<16|sliceWarms[1]<<8|sliceWarms[2]) / 3200.0 + float64(energyC3)
	// sliceWarms = Shread.HandleSharedExecCSoGpioGetRegWarms(4)
	// valueEleC4["energy"] = float64(sliceWarms[0]<<16|sliceWarms[1]<<8|sliceWarms[2]) / 3200.0 + float64(energyC4)
	// sliceWarms = Shread.HandleSharedExecCSoGpioGetRegWarms(5)
	// valueEleC5["energy"] = float64(sliceWarms[0]<<16|sliceWarms[1]<<8|sliceWarms[2]) / 3200.0 + float64(energyC5)
	// sliceWarms = Shread.HandleSharedExecCSoGpioGetRegWarms(6)
	// valueEleC6["energy"] = float64(sliceWarms[0]<<16|sliceWarms[1]<<8|sliceWarms[2]) / 3200.0 + float64(energyC6)

	if valueEleC1["current"] == 0 || valueEleC1["volt"] == 0 {
		valueEleC1["pf"] = 0
	} else {
		valueEleC1["pf"] = valueEleC1["power"] / (valueEleC1["current"] * valueEleC1["volt"])
		if valueEleC1["pf"] > 1 {
			valueEleC1["pf"] = 0
		}
	}

	if valueEleC2["current"] == 0 || valueEleC2["volt"] == 0 {
		valueEleC2["pf"] = 0
	} else {
		valueEleC2["pf"] = valueEleC2["power"] / (valueEleC2["current"] * valueEleC2["volt"])
		if valueEleC2["pf"] > 1 {
			valueEleC2["pf"] = 0
		}
	}
	
	if valueEleC3["current"] == 0 || valueEleC3["volt"] == 0 {
		valueEleC3["pf"] = 0
	} else {
		valueEleC3["pf"] = valueEleC3["power"] / (valueEleC3["current"] * valueEleC3["volt"])
		if valueEleC3["pf"] > 1 {
			valueEleC3["pf"] = 0
		}
	}

	if valueEleC4["current"] == 0 || valueEleC4["volt"] == 0 {
		valueEleC4["pf"] = 0
	} else {
		valueEleC4["pf"] = valueEleC4["power"] / (valueEleC4["current"] * valueEleC4["volt"])
		if valueEleC4["pf"] > 1 {
			valueEleC4["pf"] = 0
		}
	}

	if valueEleC5["current"] == 0 || valueEleC5["volt"] == 0 {
		valueEleC5["pf"] = 0
	} else {
		valueEleC5["pf"] = valueEleC5["power"] / (valueEleC5["current"] * valueEleC5["volt"])
		if valueEleC5["pf"] > 1 {
			valueEleC5["pf"] = 0
		}
	}

	if valueEleC6["current"] == 0 || valueEleC6["volt"] == 0 {
		valueEleC6["pf"] = 0
	} else {
		valueEleC6["pf"] = valueEleC6["power"] / (valueEleC6["current"] * valueEleC6["volt"])
		if valueEleC6["pf"] > 1 {
			valueEleC6["pf"] = 0
		}
	}
}
////






////

func handleClearEnergyRN8209(numRN int) {

	fmt.Println("---> handleClearEnergyRN8209.")
}


func handleCheckAlarmForElec() {

	fmt.Println("---> handleCheckAlarmForElec.")

	// 初始化
	alarmBuff := []int{0x33,0x01,0x10,0x00,0x00,0x06,0x01,0xE1,0x00,0x00,0x00,0x00,0x32,0x99}
	countAlarmDect := 5

	// 第1路电压
	if valueEleC1["volt"] > valueEleC1["volt_limit_max"] || valueEleC1["volt"] < valueEleC1["volt_limit_min"] {
		countRTUE1[0] += 1
		if countRTUE1[0] >= countAlarmDect {
			countRTUE1[0] = 0
			flagRTUE1[0] = 1
			if flagE1V1 <= 3 {
				flagE1V1 += 1
				alarmBuff[6] = 1
				alarmBuff[7] = 0xE1
				alarmBuff[8] = ((int(valueEleC1["volt"]*100))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC1["volt"]*100))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC1["volt"]*100))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC1["volt"]*100))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		}
	} else {
		countRTUE1[0] = 0
		flagRTUE1[0] = 0
		flagE1V1 = 0
	}

	// 第2路电压
	if valueEleC2["volt"] > valueEleC2["volt_limit_max"] || valueEleC2["volt"] < valueEleC2["volt_limit_min"] {
		countRTUE1[1] += 1
		if countRTUE1[1] >= countAlarmDect {
			countRTUE1[1] = 0
			flagRTUE1[1] = 1
			if flagE1V2 <= 3 {
				flagE1V2 += 1
				alarmBuff[6] = 1
				alarmBuff[7] = 0xE1
				alarmBuff[8] = ((int(valueEleC2["volt"]*100))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC2["volt"]*100))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC2["volt"]*100))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC2["volt"]*100))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		}
	} else {
		countRTUE1[1] = 0
		flagRTUE1[1] = 0
		flagE1V2 = 0
	}

	// 第3路电压
	if valueEleC3["volt"] > valueEleC3["volt_limit_max"] || valueEleC3["volt"] < valueEleC3["volt_limit_min"] {
		countRTUE1[2] += 1
		if countRTUE1[2] >= countAlarmDect {
			countRTUE1[2] = 0
			flagRTUE1[2] = 1
			if flagE1V3 <= 3 {
				flagE1V3 += 1
				alarmBuff[6] = 1
				alarmBuff[7] = 0xE1
				alarmBuff[8] = ((int(valueEleC3["volt"]*100))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC3["volt"]*100))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC3["volt"]*100))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC3["volt"]*100))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		}
	} else {
		countRTUE1[2] = 0
		flagRTUE1[2] = 0
		flagE1V3 = 0
	}

	// 第4路电压
	if valueEleC4["volt"] > valueEleC4["volt_limit_max"] || valueEleC4["volt"] < valueEleC4["volt_limit_min"] {
		countRTUE1[3] += 1
		if countRTUE1[3] >= countAlarmDect {
			countRTUE1[3] = 0
			flagRTUE1[3] = 1
			if flagE1V4 <= 3 {
				flagE1V4 += 1
				alarmBuff[6] = 1
				alarmBuff[7] = 0xE1
				alarmBuff[8] = ((int(valueEleC4["volt"]*100))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC4["volt"]*100))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC4["volt"]*100))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC4["volt"]*100))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		}
	} else {
		countRTUE1[3] = 0
		flagRTUE1[3] = 0
		flagE1V4 = 0
	}

	// 第5路电压
	if valueEleC5["volt"] > valueEleC5["volt_limit_max"] || valueEleC5["volt"] < valueEleC5["volt_limit_min"] {
		countRTUE1[4] += 1
		if countRTUE1[4] >= countAlarmDect {
			countRTUE1[4] = 0
			flagRTUE1[4] = 1
			if flagE1V5 <= 3 {
				flagE1V5 += 1
				alarmBuff[6] = 1
				alarmBuff[7] = 0xE1
				alarmBuff[8] = ((int(valueEleC5["volt"]*100))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC5["volt"]*100))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC5["volt"]*100))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC5["volt"]*100))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		}
	} else {
		countRTUE1[4] = 0
		flagRTUE1[4] = 0
		flagE1V5 = 0
	}

	// 第6路电压
	if valueEleC6["volt"] > valueEleC6["volt_limit_max"] || valueEleC6["volt"] < valueEleC6["volt_limit_min"] {
		countRTUE1[5] += 1
		if countRTUE1[5] >= countAlarmDect {
			countRTUE1[5] = 0
			flagRTUE1[5] = 1
			if flagE1V6 <= 3 {
				flagE1V6 += 1
				alarmBuff[6] = 1
				alarmBuff[7] = 0xE1
				alarmBuff[8] = ((int(valueEleC6["volt"]*100))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC6["volt"]*100))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC6["volt"]*100))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC6["volt"]*100))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		}
	} else {
		countRTUE1[5] = 0
		flagRTUE1[5] = 0
		flagE1V6 = 0
	}

	// 获取回路状态
	// loopState := Redis.HandleRedisJsonGet(Shared.WDStateLoop) // syspara.r_value
	loopState := 0

	// 获取外接电流互感器比例
	// ratioTransformerI := Redis.HandleRedisJsonGet(Shared.WDRatioTransformer)
	ratioTransformerI := 0

	// 获取RTU意外亮灭灯报警阈值
	// levelTopUnusualSwitch := Redis.HandleRedisJsonGet(Shared.WDLevelTopUnusualSwitch)
	levelTopUnusualSwitch := 0

	// 获取定时开状态
	// flagLampState := Redis.HandleRedisJsonGet(Shared.WDFlagLampState)
	flagLampState := 0

	// conditionLoopState := loopState & 0x01 // 回路关状态
	// if loopState & 0x01 {
	// 第1路电流
	if loopState & 0x01 == 0 { // 回路关状态
		countRTUE2[0] = 0
		flagRTUE2[0] = 0
	} else if valueEleC1["current"] * float64(ratioTransformerI) > valueEleC1["current_limit_max"] / 10.0 || valueEleC1["current"] * float64(ratioTransformerI) < valueEleC1["current_limit_min"] / 10.0 {
		// 电流上下限放大10倍
		countRTUE2[0] += 1
		if countRTUE2[0] >= countAlarmDect {
			countRTUE2[0] = 0
			flagRTUE2[0] = 1
			if flagE2I1 <= 3 {
				flagE2I1 += 1
				alarmBuff[6] = 1
				alarmBuff[7] = 0xE2
				alarmBuff[8] = ((int(valueEleC1["current"]*100))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC1["current"]*100))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC1["current"]*100))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC1["current"]*100))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		}
	} else {
		countRTUE2[0] = 0
		flagRTUE2[0] = 0
		flagE2I1 = 0
	}

	if valueEleC1["current"] * float64(ratioTransformerI) <= float64(levelTopUnusualSwitch) / 10.0 {
		// 意外亮灭灯电流阈值放大10倍
		if flagLampState & 0x01 != 0 {
			countRTUE5[0] += 1
			if countRTUE5[0] >= countAlarmDect {
				countRTUE5[0] = 0
				flagRTUE5[0] = 1
				alarmBuff[6] = 1
				alarmBuff[7] = 0xE5
				alarmBuff[8] = ((int(valueEleC1["current"]*100))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC1["current"]*100))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC1["current"]*100))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC1["current"]*100))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		} else {
			countRTUE5[0] = 0
			flagRTUE5[0] = 0
		}
		// 小于阈值就不会出现意外亮灯的情况
		countRTUE6[0] = 0
		flagRTUE6[0] = 0
	} else {
		if flagLampState & 0x01 == 0 {
			// 定时关，检测电流大于阈值，意外亮灯
			countRTUE6[0] += 1
			if countRTUE6[0] >= countAlarmDect {
				countRTUE6[0] =0
				flagRTUE6[0] = 1
				alarmBuff[6] = 1
				alarmBuff[7] = 0xE6
				alarmBuff[8] = ((int(valueEleC1["current"]*100))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC1["current"]*100))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC1["current"]*100))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC1["current"]*100))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		} else {
			countRTUE6[0] = 0
			flagRTUE6[0] = 0
		}
		// 大于阈值就不会出现意外灭灯的情况
		countRTUE5[0] = 0
		flagRTUE5[0] = 0
	} 

	// 第2路电流
	if loopState & 0x02 == 0 { // 回路关状态
		countRTUE2[1] = 0
		flagRTUE2[1] = 0
	} else if valueEleC2["current"] * float64(ratioTransformerI) > valueEleC2["current_limit_max"] / 10.0 || valueEleC2["current"] * float64(ratioTransformerI) < valueEleC2["current_limit_min"] / 10.0 {
		// 电流上下限放大10倍
		countRTUE2[1] += 1
		if countRTUE2[1] >= countAlarmDect {
			countRTUE2[1] = 0
			flagRTUE2[1] = 1
			if flagE2I2 <= 3 {
				flagE2I2 += 1
				alarmBuff[6] = 2
				alarmBuff[7] = 0xE2
				alarmBuff[8] = ((int(valueEleC2["current"]*100))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC2["current"]*100))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC2["current"]*100))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC2["current"]*100))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		}
	} else {
		countRTUE2[1] = 0
		flagRTUE2[1] = 0
		flagE2I2 = 0
	}

	if valueEleC2["current"] * float64(ratioTransformerI) <= float64(levelTopUnusualSwitch) / 10.0 {
		// 意外亮灭灯电流阈值放大10倍
		if flagLampState & 0x02 != 0 {
			countRTUE5[1] += 1
			if countRTUE5[1] >= countAlarmDect {
				countRTUE5[1] = 0
				flagRTUE5[1] = 1
				alarmBuff[6] = 2
				alarmBuff[7] = 0xE5
				alarmBuff[8] = ((int(valueEleC2["current"]*100))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC2["current"]*100))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC2["current"]*100))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC2["current"]*100))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		} else {
			countRTUE5[1] = 0
			flagRTUE5[1] = 0
		}
		// 小于阈值就不会出现意外亮灯的情况
		countRTUE6[1] = 0
		flagRTUE6[1] = 0
	} else {
		if flagLampState & 0x02 == 0 {
			// 定时关，检测电流大于阈值，意外亮灯
			countRTUE6[1] += 1
			if countRTUE6[1] >= countAlarmDect {
				countRTUE6[1] =0
				flagRTUE6[1] = 1
				alarmBuff[6] = 2
				alarmBuff[7] = 0xE6
				alarmBuff[8] = ((int(valueEleC2["current"]*100))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC2["current"]*100))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC2["current"]*100))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC2["current"]*100))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		} else {
			countRTUE6[1] = 0
			flagRTUE6[1] = 0
		}
		// 大于阈值就不会出现意外灭灯的情况
		countRTUE5[1] = 0
		flagRTUE5[1] = 0
	} 

	// 第3路电流
	if loopState & 0x04 == 0 { // 回路关状态
		countRTUE2[2] = 0
		flagRTUE2[2] = 0
	} else if valueEleC3["current"] * float64(ratioTransformerI) > valueEleC3["current_limit_max"] / 10.0 || valueEleC3["current"] * float64(ratioTransformerI) < valueEleC3["current_limit_min"] / 10.0 {
		// 电流上下限放大10倍
		countRTUE2[2] += 1
		if countRTUE2[2] >= countAlarmDect {
			countRTUE2[2] = 0
			flagRTUE2[2] = 1
			if flagE2I3 <= 3 {
				flagE2I3 += 1
				alarmBuff[6] = 3
				alarmBuff[7] = 0xE2
				alarmBuff[8] = ((int(valueEleC3["current"]*1000))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC3["current"]*1000))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC3["current"]*1000))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC3["current"]*1000))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		}
	} else {
		countRTUE2[2] = 0
		flagRTUE2[2] = 0
		flagE2I3 = 0
	}

	if valueEleC3["current"] * float64(ratioTransformerI) <= float64(levelTopUnusualSwitch) / 10.0 {
		// 意外亮灭灯电流阈值放大10倍
		if flagLampState & 0x04 != 0 {
			countRTUE5[2] += 1
			if countRTUE5[2] >= countAlarmDect {
				countRTUE5[2] = 0
				flagRTUE5[2] = 1
				alarmBuff[6] = 3
				alarmBuff[7] = 0xE5
				alarmBuff[8] = ((int(valueEleC3["current"]*1000))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC3["current"]*1000))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC3["current"]*1000))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC3["current"]*1000))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		} else {
			countRTUE5[2] = 0
			flagRTUE5[2] = 0
		}
		// 小于阈值就不会出现意外亮灯的情况
		countRTUE6[2] = 0
		flagRTUE6[2] = 0
	} else {
		if flagLampState & 0x04 == 0 {
			// 定时关，检测电流大于阈值，意外亮灯
			countRTUE6[2] += 1
			if countRTUE6[2] >= countAlarmDect {
				countRTUE6[2] =0
				flagRTUE6[2] = 1
				alarmBuff[6] = 3
				alarmBuff[7] = 0xE6
				alarmBuff[8] = ((int(valueEleC3["current"]*1000))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC3["current"]*1000))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC3["current"]*1000))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC3["current"]*1000))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		} else {
			countRTUE6[2] = 0
			flagRTUE6[2] = 0
		}
		// 大于阈值就不会出现意外灭灯的情况
		countRTUE5[2] = 0
		flagRTUE5[2] = 0
	} 

	// 第4路电流
	if loopState & 0x08 == 0 { // 回路关状态
		countRTUE2[3] = 0
		flagRTUE2[3] = 0
	} else if valueEleC4["current"] * float64(ratioTransformerI) > valueEleC4["current_limit_max"] / 10.0 || valueEleC4["current"] * float64(ratioTransformerI) < valueEleC4["current_limit_min"] / 10.0 {
		// 电流上下限放大10倍
		countRTUE2[3] += 1
		if countRTUE2[3] >= countAlarmDect {
			countRTUE2[3] = 0
			flagRTUE2[3] = 1
			if flagE2I4 <= 3 {
				flagE2I4 += 1
				alarmBuff[6] = 4
				alarmBuff[7] = 0xE2
				alarmBuff[8] = ((int(valueEleC4["current"]*1000))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC4["current"]*1000))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC4["current"]*1000))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC4["current"]*1000))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		}
	} else {
		countRTUE2[3] = 0
		flagRTUE2[3] = 0
		flagE2I4 = 0
	}

	if valueEleC4["current"] * float64(ratioTransformerI) <= float64(levelTopUnusualSwitch) / 10.0 {
		// 意外亮灭灯电流阈值放大10倍
		if flagLampState & 0x08 != 0 {
			countRTUE5[3] += 1
			if countRTUE5[3] >= countAlarmDect {
				countRTUE5[3] = 0
				flagRTUE5[3] = 1
				alarmBuff[6] = 4
				alarmBuff[7] = 0xE5
				alarmBuff[8] = ((int(valueEleC4["current"]*1000))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC4["current"]*1000))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC4["current"]*1000))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC4["current"]*1000))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		} else {
			countRTUE5[3] = 0
			flagRTUE5[3] = 0
		}
		// 小于阈值就不会出现意外亮灯的情况
		countRTUE6[3] = 0
		flagRTUE6[3] = 0
	} else {
		if flagLampState & 0x08 == 0 {
			// 定时关，检测电流大于阈值，意外亮灯
			countRTUE6[3] += 1
			if countRTUE6[3] >= countAlarmDect {
				countRTUE6[3] =0
				flagRTUE6[3] = 1
				alarmBuff[6] = 4
				alarmBuff[7] = 0xE6
				alarmBuff[8] = ((int(valueEleC4["current"]*1000))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC4["current"]*1000))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC4["current"]*1000))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC4["current"]*1000))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		} else {
			countRTUE6[3] = 0
			flagRTUE6[3] = 0
		}
		// 大于阈值就不会出现意外灭灯的情况
		countRTUE5[3] = 0
		flagRTUE5[3] = 0
	} 

	// 第5路电流
	if loopState & 0x10 == 0 { // 回路关状态
		countRTUE2[4] = 0
		flagRTUE2[4] = 0
	} else if valueEleC5["current"] * float64(ratioTransformerI) > valueEleC5["current_limit_max"] / 10.0 || valueEleC5["current"] * float64(ratioTransformerI) < valueEleC5["current_limit_min"] / 10.0 {
		// 电流上下限放大10倍
		countRTUE2[4] += 1
		if countRTUE2[4] >= countAlarmDect {
			countRTUE2[4] = 0
			flagRTUE2[4] = 1
			if flagE2I5 <= 3 {
				flagE2I5 += 1
				alarmBuff[6] = 5
				alarmBuff[7] = 0xE2
				alarmBuff[8] = ((int(valueEleC5["current"]*1000))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC5["current"]*1000))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC5["current"]*1000))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC5["current"]*1000))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		}
	} else {
		countRTUE2[4] = 0
		flagRTUE2[4] = 0
		flagE2I5 = 0
	}

	if valueEleC5["current"] * float64(ratioTransformerI) <= float64(levelTopUnusualSwitch) / 10.0 {
		// 意外亮灭灯电流阈值放大10倍
		if flagLampState & 0x10 != 0 {
			countRTUE5[4] += 1
			if countRTUE5[4] >= countAlarmDect {
				countRTUE5[4] = 0
				flagRTUE5[4] = 1
				alarmBuff[6] = 5
				alarmBuff[7] = 0xE5
				alarmBuff[8] = ((int(valueEleC5["current"]*1000))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC5["current"]*1000))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC5["current"]*1000))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC5["current"]*1000))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		} else {
			countRTUE5[4] = 0
			flagRTUE5[4] = 0
		}
		// 小于阈值就不会出现意外亮灯的情况
		countRTUE6[4] = 0
		flagRTUE6[4] = 0
	} else {
		if flagLampState & 0x10 == 0 {
			// 定时关，检测电流大于阈值，意外亮灯
			countRTUE6[4] += 1
			if countRTUE6[4] >= countAlarmDect {
				countRTUE6[4] =0
				flagRTUE6[4] = 1
				alarmBuff[6] = 5
				alarmBuff[7] = 0xE6
				alarmBuff[8] = ((int(valueEleC5["current"]*1000))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC5["current"]*1000))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC5["current"]*1000))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC5["current"]*1000))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		} else {
			countRTUE6[4] = 0
			flagRTUE6[4] = 0
		}
		// 大于阈值就不会出现意外灭灯的情况
		countRTUE5[4] = 0
		flagRTUE5[4] = 0
	} 

	// 第6路电流
	if loopState & 0x20 == 0 { // 回路关状态
		countRTUE2[5] = 0
		flagRTUE2[5] = 0
	} else if valueEleC6["current"] * float64(ratioTransformerI) > valueEleC6["current_limit_max"] / 10.0 || valueEleC6["current"] * float64(ratioTransformerI) < valueEleC6["current_limit_min"] / 10.0 {
		// 电流上下限放大10倍
		countRTUE2[5] += 1
		if countRTUE2[5] >= countAlarmDect {
			countRTUE2[5] = 0
			flagRTUE2[5] = 1
			if flagE2I6 <= 3 {
				flagE2I6 += 1
				alarmBuff[6] = 6
				alarmBuff[7] = 0xE2
				alarmBuff[8] = ((int(valueEleC6["current"]*1000))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC6["current"]*1000))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC6["current"]*1000))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC6["current"]*1000))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		}
	} else {
		countRTUE2[5] = 0
		flagRTUE2[5] = 0
		flagE2I6 = 0
	}

	if valueEleC6["current"] * float64(ratioTransformerI) <= float64(levelTopUnusualSwitch) / 10.0 {
		// 意外亮灭灯电流阈值放大10倍
		if flagLampState & 0x20 != 0 {
			countRTUE5[5] += 1
			if countRTUE5[5] >= countAlarmDect {
				countRTUE5[5] = 0
				flagRTUE5[5] = 1
				alarmBuff[6] = 6
				alarmBuff[7] = 0xE5
				alarmBuff[8] = ((int(valueEleC6["current"]*1000))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC6["current"]*1000))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC6["current"]*1000))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC6["current"]*1000))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		} else {
			countRTUE5[5] = 0
			flagRTUE5[5] = 0
		}
		// 小于阈值就不会出现意外亮灯的情况
		countRTUE6[5] = 0
		flagRTUE6[5] = 0
	} else {
		if flagLampState & 0x20 == 0 {
			// 定时关，检测电流大于阈值，意外亮灯
			countRTUE6[5] += 1
			if countRTUE6[5] >= countAlarmDect {
				countRTUE6[5] =0
				flagRTUE6[5] = 1
				alarmBuff[6] = 6
				alarmBuff[7] = 0xE6
				alarmBuff[8] = ((int(valueEleC6["current"]*1000))>>24) & 0xFF
				alarmBuff[9] = ((int(valueEleC6["current"]*1000))>>16) & 0xFF
				alarmBuff[10] = ((int(valueEleC6["current"]*1000))>>8) & 0xFF
				alarmBuff[11] = ((int(valueEleC6["current"]*1000))) & 0xFF
				AlarmUpload.HandleAlarmBuffParsing(alarmBuff)
			}
		} else {
			countRTUE6[5] = 0
			flagRTUE6[5] = 0
		}
		// 大于阈值就不会出现意外灭灯的情况
		countRTUE5[5] = 0
		flagRTUE5[5] = 0
	} 
}

