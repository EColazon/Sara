package handleLogsManages

/*
流程：
	1.每1s获取一次时间
	2.拿当前时间与期望时间比对
	3.条件成立执行相应动作
	4.删除日志失败执行相应动作
*/
import (
	"fmt"
	"time"

	MySql "handleMySql"
	alarmMethod "handleAlarmUpload"
)

var (
	AlarmBuffDBDelLog = []int{0x33, 0x01, 0x10, 0x04, 0x00, 0x06, 0x00, 0x01, 0xD0, 0x00, 0x00, 0x00, 0x32, 0x99}
)

const (
	// 时间相关
	DelWeek			= 	"Friday"
	DelHour 		= 	7
	DelMinute		= 	30
	DelSecond00		= 	0
	DelSecond01		= 	1

	// 数据库表相关
	DBTableHeart00		= "dbloghearta0"
	DBTableHeart01		= "dbloghearta1"

	DBTableCmd00		= "dblogokd0"
	DBTableCmd01		= "dblogokd1"
	DBTableCmd02		= "dblogokd2"
	DBTableCmd03		= "dblogokd3"
	DBTableCmd04		= "dblogokd4"
	DBTableCmd05		= "dblogcmderr"

	// 系统参数相关
	DBTableSys00		= "dblogsysb"

	// 异常报警相关
	DBTableModule00		= "dblogmodulec"

	DBTableAlarm00		= "dblogalarme"
)
func HandleLogsManages() {

	fmt.Println("---> start HandleLogsManages.")

	logTimeCount := 0


	for {
		// 计数自增
		logTimeCount += 1
		time.Sleep(1 * time.Second)

		// 获取当前时间
		timeNowWeek	  	:= time.Now().Weekday().String()
		timeNowHour 	:= time.Now().Hour()
		timeNowMinute 	:= time.Now().Minute()
		timeNowSecond	:= time.Now().Second()
		
		fmt.Println("---> TimeNow: ",timeNowWeek, timeNowHour, timeNowMinute)

		// 每周五07:30清日志数据
		if timeNowWeek == DelWeek && timeNowHour == DelHour && timeNowMinute == DelMinute {
			if timeNowSecond == DelSecond00 || timeNowSecond == DelSecond01 {
				fmt.Println("---> It's Time To Delete Something.")

				// Do Del
				// 心跳包日志
				okDel00 := MySql.HandleDBLogDelete(7, DBTableHeart00)
				okDel01 := MySql.HandleDBLogDelete(7, DBTableHeart01)
				
				// 命令日志
				okDel02 := MySql.HandleDBLogDelete(7, DBTableCmd00)
				okDel03 := MySql.HandleDBLogDelete(7, DBTableCmd01)
				okDel04 := MySql.HandleDBLogDelete(7, DBTableCmd02)
				okDel05 := MySql.HandleDBLogDelete(7, DBTableCmd03)
				okDel06 := MySql.HandleDBLogDelete(7, DBTableCmd04)
				okDel07 := MySql.HandleDBLogDelete(7, DBTableCmd05)

				// 系统参数日志
				okDel08 := MySql.HandleDBLogDelete(7, DBTableSys00)

				// 程序异常日志
				okDel09 := MySql.HandleDBLogDelete(7, DBTableModule00)

				// 设备异常日志
				okDel0A := MySql.HandleDBLogDelete(7, DBTableAlarm00)

				// 删除日志异常
				if !okDel00 {
					AlarmBuffDBDelLog[8] = 0x10
					alarmMethod.HandleAlarmBuffParsing(AlarmBuffDBDelLog)
					fmt.Println("---> DBLog Del Fail DBTableHeart00")
				} else if !okDel01 {
					AlarmBuffDBDelLog[8] = 0x11
					alarmMethod.HandleAlarmBuffParsing(AlarmBuffDBDelLog)
					fmt.Println("---> DBLog Del Fail DBTableHeart01")
				} else if !okDel02 {
					AlarmBuffDBDelLog[8] = 0x12
					alarmMethod.HandleAlarmBuffParsing(AlarmBuffDBDelLog)
					fmt.Println("---> DBLog Del Fail DBTableCmd00")
				} else if !okDel03 {
					AlarmBuffDBDelLog[8] = 0x13
					alarmMethod.HandleAlarmBuffParsing(AlarmBuffDBDelLog)
					fmt.Println("---> DBLog Del Fail DBTableCmd01")
				} else if !okDel04 {
					AlarmBuffDBDelLog[8] = 0x14
					alarmMethod.HandleAlarmBuffParsing(AlarmBuffDBDelLog)
					fmt.Println("---> DBLog Del Fail DBTableCmd02")
				} else if !okDel05 {
					AlarmBuffDBDelLog[8] = 0x15
					alarmMethod.HandleAlarmBuffParsing(AlarmBuffDBDelLog)
					fmt.Println("---> DBLog Del Fail DBTableCmd03")
				} else if !okDel06 {
					AlarmBuffDBDelLog[8] = 0x16
					alarmMethod.HandleAlarmBuffParsing(AlarmBuffDBDelLog)
					fmt.Println("---> DBLog Del Fail DBTableCmd04")
				} else if !okDel07 {
					AlarmBuffDBDelLog[8] = 0x17
					alarmMethod.HandleAlarmBuffParsing(AlarmBuffDBDelLog)
					fmt.Println("---> DBLog Del Fail DBTableCmd05")
				} else if !okDel08 {
					AlarmBuffDBDelLog[8] = 0x18
					alarmMethod.HandleAlarmBuffParsing(AlarmBuffDBDelLog)
					fmt.Println("---> DBLog Del Fail DBTableSys00")
				} else if !okDel09 {
					AlarmBuffDBDelLog[8] = 0x19
					alarmMethod.HandleAlarmBuffParsing(AlarmBuffDBDelLog)
					fmt.Println("---> DBLog Del Fail DBTableModule00")
				} else if !okDel0A {
					AlarmBuffDBDelLog[8] = 0x1A
					alarmMethod.HandleAlarmBuffParsing(AlarmBuffDBDelLog)
					fmt.Println("---> DBLog Del Fail DBTableAlarm00")
				}
			}
		}

		

		break
	}

}