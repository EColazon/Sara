package handleAlarmUpload

/*
流程:
	初始化系统报警相关变量
*/
import (
	"fmt"
)
const (
	CMD33HEAD = 0x33
	CMD33TAIL = 0x99
	CMD2FHEAD00 = 0x2F
	CMD2FHEAD01 = 0x43
	CMD2FHEAD02 = 0x2F
	CMD2FTAIL = 0xCC

)
var (
	// 切片类型报警缓存(noUse)
	AlarmBuff = make([]uint, 12)
	
	// 解析异常标识
	FlagParsing33 = 0

	// E1&E2&E5&E6&E7&E8报警标志
	FlagAlarmE1			= make([]int, 6)
	FlagAlarmE2			= make([]int, 6)
	FlagAlarmE5			= make([]int, 6)
	FlagAlarmE6			= make([]int, 6)
	FlagAlarmE7			= make([]int, 1024)
	FlagAlarmE8			= make([]int, 1024)

	// E1&E2&E5&E6&E7&E8报警内容
	ValueAlarmE1			= make([]int, 6)
	ValueAlarmE2			= make([]int, 6)
	ValueAlarmE5			= make([]int, 6)
	ValueAlarmE6			= make([]int, 6)
	ValueAlarmE7			= make([]int, 1024)
	ValueAlarmE8			= make([]int, 1024)

	// 数据库读写异常报警标志
	FlagAlarmD0			= make([]int, 1)		// DBEleC
	FlagAlarmD1			= make([]int, 1)		// DBLamp
	FlagAlarmD2			= make([]int, 1)		// DBNode
	FlagAlarmD3			= make([]int, 1)		// DBOneKeys
	FlagAlarmD4			= make([]int, 1)		// DBTenKeys
	FlagAlarmD5			= make([]int, 1)		// DBTimer
	FlagAlarmD6			= make([]int, 1)
	FlagAlarmD7			= make([]int, 1)
	FlagAlarmD8			= make([]int, 1)
	FlagAlarmD9			= make([]int, 1)
	FlagAlarmDA			= make([]int, 1)
	FlagAlarmDB			= make([]int, 1)
	FlagAlarmDC			= make([]int, 1)
	FlagAlarmDD			= make([]int, 1)
	FlagAlarmDE			= make([]int, 1)
	FlagAlarmDF			= make([]int, 1)

	// 数据库读写异常报警内容
	ValueAlarmD0			= make([]int, 1)
	ValueAlarmD1			= make([]int, 1)
	ValueAlarmD2			= make([]int, 1)
	ValueAlarmD3			= make([]int, 1)
	ValueAlarmD4			= make([]int, 1)
	ValueAlarmD5			= make([]int, 1)
	ValueAlarmD6			= make([]int, 1)
	ValueAlarmD7			= make([]int, 1)
	ValueAlarmD8			= make([]int, 1)
	ValueAlarmD9			= make([]int, 1)
	ValueAlarmDA			= make([]int, 1)
	ValueAlarmDB			= make([]int, 1)
	ValueAlarmDC			= make([]int, 1)
	ValueAlarmDD			= make([]int, 1)
	ValueAlarmDE			= make([]int, 1)
	ValueAlarmDF			= make([]int, 1)

	// 命令解析异常报警
	FlagAlarmC0			= make([]int, 1)		// 33 Fail SR
	FlagAlarmC1			= make([]int, 1)		// 2F Fail SR
	FlagAlarmC2			= make([]int, 1)		// 33 Fail RR
	FlagAlarmC3			= make([]int, 1)		// 
	FlagAlarmC4			= make([]int, 1)		// 
	FlagAlarmC5			= make([]int, 1) 		// 

	ValueAlarmC0			= make([]int, 1)
	ValueAlarmC1			= make([]int, 1)
	ValueAlarmC2			= make([]int, 1)
	ValueAlarmC3			= make([]int, 1)
	ValueAlarmC4			= make([]int, 1)
	ValueAlarmC5			= make([]int, 1)

	// 清空日志异常报警
	FlagDBLog10			= make([]int, 1)
	FlagDBLog11			= make([]int, 1)
	FlagDBLog12			= make([]int, 1)
	FlagDBLog13			= make([]int, 1)
	FlagDBLog14			= make([]int, 1)
	FlagDBLog15			= make([]int, 1)
	FlagDBLog16			= make([]int, 1)
	FlagDBLog17			= make([]int, 1)
	FlagDBLog18			= make([]int, 1)
	FlagDBLog19			= make([]int, 1)
	FlagDBLog1A			= make([]int, 1)

	ValueDBLog10			= make([]int, 1)
	ValueDBLog11			= make([]int, 1)
	ValueDBLog12			= make([]int, 1)
	ValueDBLog13			= make([]int, 1)
	ValueDBLog14			= make([]int, 1)
	ValueDBLog15			= make([]int, 1)
	ValueDBLog16			= make([]int, 1)
	ValueDBLog17			= make([]int, 1)
	ValueDBLog18			= make([]int, 1)
	ValueDBLog19			= make([]int, 1)
	ValueDBLog1A			= make([]int, 1)
)

func init() {
	fmt.Println("---> handleAlarmUploadInit.")
}

