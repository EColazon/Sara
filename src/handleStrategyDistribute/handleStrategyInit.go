package handleStrategyDistribute

var (
	// 特殊策略开关时间
	timerRONSpecial 		= make([]int, 24)
	timerROFFSpecial 		= make([]int, 24)
	// 节假日策略
	timerRONHoliday 		= make([]int, 24)
	timerROFFHoliday 		= make([]int, 24)
	// 初始化策略标志
	flagStrategySH 			= make([]int, 8)

	// timer_offset_on
	timerOffsetON			= make([]int, 3)
	// timer_offset_off
	timerOffsetOFF			= make([]int, 3)
)