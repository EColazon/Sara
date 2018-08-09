package handleShared

import (
	"fmt"
	// "math"
	// Redis "handleRedis"
)

func HandleUpdateTimeLatiLongi() {
	fmt.Println("---> HandleUpdateTimeLatiLongi")
	// 获取PCF时间
	// timePCF := HandleSharedExecCSoPCFRead()
	// 获取经度
	/*
	WDTimeLatitude00 := Redis.HandleRedisJsonGet(WDTimeLatitude00) 
	WDTimeLatitude01 := Redis.HandleRedisJsonGet(WDTimeLatitude01) 
	WDTimeLatitude02 := Redis.HandleRedisJsonGet(WDTimeLatitude02) 
	WDTimeLatitude03 := Redis.HandleRedisJsonGet(WDTimeLatitude03)
	
	//获取纬度
	WDTimeLongitude00 := Redis.HandleRedisJsonGet(WDTimeLongitude00)
	WDTimeLongitude01 := Redis.HandleRedisJsonGet(WDTimeLongitude01)
	WDTimeLongitude02 := Redis.HandleRedisJsonGet(WDTimeLongitude02)
	WDTimeLongitude03 := Redis.HandleRedisJsonGet(WDTimeLongitude03)
    
	tDay := (timePCF[5] - 1)*30 + timePCF[3]
	WDTimeLatitude := WDTimeLatitude02 + WDTimeLatitude01 / 60.0 + WDTimeLatitude00 / 3600.0
	WDTimeLongitude := WDTimeLongitude02 + WDTimeLongitude01 / 60.0 + WDTimeLongitude00 / 3600.0
	tDayNum := WDTimeLatitude03

	handleUpDownMeasure(WDTimeLatitude, WDTimeLongitude, tDay, tDayNum, 1)
	handleUpDownTransfrom()
	handleUpDownMeasure(WDTimeLatitude, WDTimeLongitude, tDay, tDayNum, 0)
	*/
}

func handleUpDownMeasure(timeLatitude float32, timeLongitude float32, tDay int, tDayNum int, flag int) int {
	/*
	#经纬度超出范围时，开灯6:00，关灯17:00;
    #东西半球坐标范围-180到180；负代表西半球，正代表东北球；
	#南北半球的坐标范围是-90到90；负代表南半球，正代表北半球；
	*/
	if timeLatitude > 180 || timeLatitude < (-180) {
		if flag == 0 {
			return 6
		} else {
			return 17
		}
	}
	if timeLongitude > 90 || timeLongitude < (-90) {
		if flag == 0 {
			return 6
		} else {
			return 17
		}
	}

	// 经纬度在正范围内，通过公式计算
	/*
	tDayNum = 8
	m1 := math.Cos(360*(tDay + 9) / 365 * 3.14 / 180) * (-23.4)
	m2 := -math.Tan(m1 * 3.14 / 180) * math.Tan(timeLongitude * 3.14 / 180)
	m3 := 24 * (180 + tDayNum * 15 - timeLatitude - math.Acos(m2) * 180 / 3.14) /360
	m4 := 24 * (1 + (tDayNum * 15 - timeLatitude) / 180) - m3
	
	if flag == 0 {
		return m3
	} else {
		return m4
	}
	*/
	return 32730
}

func handleUpDownTransfrom() {
	fmt.Println("---> handleUpDownTransfrom.")
}