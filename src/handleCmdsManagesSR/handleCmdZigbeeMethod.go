package handleCmdsManagesSR

import (
	"fmt"
	"reflect"
)

// 这里是。。。这个应该怎么说?
const (
	TYPELAMP 		= 0x01		// 单灯类型
	TYPENODE		= 0x02		// 节点类型
	LAMPADV			= 0x01		// 主灯Advocete
	LAMPAUX			= 0x02		// 辅灯Auxiliary
	LAMPADU			= 0x03		// 主辅灯
)

// 单灯结构内容
// BT:Both&SN:Single
type Lamp struct{
	LampNum			int				//单灯编号1
	LampNumGroup 	int				//单灯组号1

	LAdvV			int				//主灯电压,未放大2
	LAdvI			int				//主灯电流,放大1000倍2
	LAdvP			int				//主灯功率,放大10倍2
	LAdvPF			int 			//主灯功率因素1
	LAuxV			int				//辅灯电压,未放大2
	LAuxI			int				//辅灯电流,放大1000倍2
	LAuxP			int				//辅灯功率,放大10倍2
	LAuxPF			int 			//辅灯功率因素1

	LAdvPwm			int				//主灯PWM调光0-100,值为0关灯1
	LAuxPwm			int				//辅灯PWM调光0-100,值为0关灯1
	LStateRelayBT	int				//单灯两路继电器开关状态,默认0x03,1
									//bit0代表主灯,1开0关&bit1代表辅灯,1开0关
	LModeTX			int				//保留1
	LModeRX			int 			//值为0x55表示单灯拨码固定1
	LTimeAlarm		int				//保留2
	LRelayChange	int				//值为0x55表示单灯主辅灯互换1
	LFlagEX			int				//保留1
	LStateAlarm		int				//报警指示2
									//1111111111111111
									//0主灯电压超下限&1主灯电压超上限&2主灯功率超下限&3主灯功率超上限&
									//4主灯功率因素超下限&5主灯功率因素超上限&6主灯继电器黏连&7未定义
									//8辅灯电压超下限&9辅灯电压超上限&10辅灯功率超下限&11辅灯功率超上限&
									//12辅灯功率因素超下限&13辅灯功率因素超上限&14辅灯继电器黏连&15未定义
	//以上共24字节
	ZBNetAddr		int				//Zigbee网络短地址2
	LAdvPower		int				//主灯能量4
	LAuxPower		int				//辅灯能量4
	LampHigherV		int				//电压上限2
	LampLowerV		int				//电压下限2
	LampHigherI		int				//电流上限2
	LampLowerI		int				//电流下限2
	LampHigherP		int				//功率上限2
	LampLowerP		int				//功率下限2
	LampHigherPF	int				//功率因素上限1
	LampLowerPF		int				//功率因素下限1
	LFlagSetNum		int				//保留1
	LFlagSetAdu		int				//保留:主辅灯?
	LChecksum		int				//校验和1
	//以上54字节
	//以下不上传

	LNAdvSet			int				//保留:主灯?2
	LNAuxSet			int				//保留:辅灯?2
	LNStateRelayBT		int				//单灯两路继电器开关状态,默认0x03,1
	LNCountDisconn		int				//单灯失联检测时计数,巡检中累加,在0x1009和0x2012返回中清零1
	LNFlagDataback		int				//保留1
	LNAdvPwm			int				//主灯PWM调光0-100,值为0关灯1
	LNAuxPwm			int				//辅灯PWM调光0-100,值为0关灯1
	LNStateByHand		int 			//单灯手动开关标志,值为0x55在巡检时不去纠正开关不一致1
	LNNumOutput			int				//单灯输出路数,值为2两路输出,值为1一路输出1
										//在报警时判断要不要报辅灯报警
	LNDetecOK			int				//标识巡检时是否被巡检到
}

// id:标识ZigBeeNode对应编号,0-1023
// lamp:代表单灯数据内容
type ZigebeeNode struct {
	id			int
	lamp		Lamp
}

// 声明[]ZigebeeNode类型
type GZBNode []ZigebeeNode

// 声明ZigBeeNode数据缓存,默认1024
// G:全局
// var GZigbeeNode []ZigebeeNode
// 相当于面向对象的类(GZigbeeNode)
var GZigbeeNode GZBNode

// Methods:51
type GZigbeeNodeActer interface {
	GZB72053InductAndTrigeTimeActer()			// 设置节点控制器的感应恢复时间和触发间隔时间
	GZB72054SetLampGroupActer()					// 设置单灯组号
	GZB72055SetLampRelayChangeActer()			// 单灯主辅互换
	GZB72057SetLampAndNodeRelatedActer()		// 设置单灯与节点关联
	GZB72058SetLampInductRecoverTimeActer()  	// 设置单灯在节点触发后的感应恢复时间
	GZB72059SetLampSelfCheckActer()				// 设置命令单灯自校验
	GZB72060SetLampFixedKeyActer()				// 固定单灯拨码
	GZB72061QueryLampIEEEActer()				// 查询返回单灯ieee地址
	GZB72062SetLampAlwaysCloseActer()			// 单灯常关
	GZB72063SetLampInterOpenActer()   	  		// 单灯内部开主灯
	GZB72064SetLampInterCloseActer()   	  		// 单灯内部关主灯
	// GZB72065SetLampInterAuxOpenActer()   	// 单灯内部开辅灯
	// GZB72066SetLampInterAuxCloseActer()   	// 单灯内部关辅灯
	// GZB72067SetLampInterAduOpenActer()   	// 单灯内部开主辅灯
	// GZB72068SetLampInterAduCloseActer()   	// 单灯内部关主辅灯
	GZB72069SetLampSaveElecsActer()				// 单灯电量保存
	GZB72070SetLampAlarmLimitVActer()			// 设置单灯电压报警上下限
	GZB72071SetLampAlarmLimitIActer()			// 设置单灯电流报警上下限
	GZB72072SetLampAlarmLimitPActer()			// 设置单灯功率报警上下限
	GZB72073SetLampAlarmLimitPFActer()			// 设置单灯功率因素报警上下限
	GZB72074SetLampAnergyClearActer()			// 清除单灯能量数据
	GZB72075SetLampAmplifyVActer()				// 设置单灯电压放大倍数主灯
	// GZB72076SetLampAmplifyVActer()			// 设置单灯电压放大倍数辅灯
	GZB72077SetLampAmplifyIActer()				// 设置单灯电流放大倍数主灯
	// GZB72078SetLampmplifyIActer()			// 设置单灯电流放大倍数辅灯
	GZB72079SetLampRN8209OffsetActer()			// rn8209有效值offset和有功offset校正
	GZB72080SetLampRN8209ParasActer()			// rn8209参数设置
	GZB72081SetLampRN8209RatioActer()			// rn8209比例系数设置
	GZB72082SetLampRtuSyncTimeActer()			// 单灯与rtu时间同步命令
	GZB72083SetLampSwitchSyncTimeActer()		// 开关时间同步?
	GZB72084QueryBatteryVoltActer()				// 查询锂电池电平
	GZB72085QueryTemperatureActer()				// 查询温度
	GZB72086SetClearNVAdderActer()				// 擦除NV，重新加入网络
	GZB72087SetLampTwinkleActer()				// 单灯开闪烁功能
	GZB72088QueryLampCalledDataActer()			// 单灯召测命令
	// GZB72089QueryLampCalledDataNextActer()	// 单灯召测命令
	GZB72090SetLampOpenByHandsActer()			// 单灯手动开主辅灯
	GZB72091SetLampCloseByHandsActer()			// 单灯手动关主辅灯
	GZB72092QueryLampDetecDataMannyActer()		// 巡检manny命令
	GZB72093QueryLampDetecDataInterActer()		// 巡检内部命令	// GZB72094SetLampAdvOpenByHandsActer()		// 单灯手动开主灯
	// GZB72095SetLampAdvOpenByHandsActer()		// 单灯手动关主灯
	// GZB72096SetLampAuxOpenByHandsActer()		// 单灯手动开辅灯
	// GZB72097SetLampAuxOpenByHandsActer()		// 单灯手动关辅灯
	GZB72100SetNodeSwitchActer()				// 节点开关命令
	GZB7210SetLampDimmer()						// 单灯调光主灯
	// GZB72102SetLampAdvDimmerActer()			// 单灯调光辅灯
	// GZB72103SetLampAuxDimmerActer()			// 单灯调光主辅灯
	GZB72104QueryLampPanidActer()				// 返回单灯panid
	GZB72105QueryLampProgramVersionActer()		// 返回单灯程序版本
	GZB72106SetLampResetActer()					// 重启单灯
	GZB72116QueryLampTimeSTC()					// 查询STC时间



}
func init() {
	// var zigbeeNode ZigebeeNode
	// fmt.Println("---> 01 ", zigbeeNode)
	// fmt.Println("---> 02 ", reflect.TypeOf(zigbeeNode))

	var node ZigebeeNode
	// var GZigbeeNode []ZigebeeNode

	for i := 0; i < 3; i++ {
		node.id = i
		// node.lamp.LampNum = i
		// node.lamp.LAdvPwm = i*100
		// node.lamp.LAdvPower = i*i
		// node.lamp.LampHigherV = i+100
		node.lamp.LampNum 			= 1
		node.lamp.LampNumGroup 		= 1
		node.lamp.LAdvV 			= 0
		node.lamp.LAdvI 			= 0
		node.lamp.LAdvP 			= 0
		node.lamp.LAdvPF 			= 0
		node.lamp.LAuxV 			= 0
		node.lamp.LAuxI 			= 0
		node.lamp.LAuxP 			= 0
		node.lamp.LAuxPF 			= 0

		node.lamp.LAdvPwm 			= 100
		node.lamp.LAuxPwm 			= 100
		node.lamp.LStateRelayBT 	= 0x03
		node.lamp.LModeTX 			= 0
		node.lamp.LModeRX 			= 0
		node.lamp.LTimeAlarm 		= 0
		node.lamp.LRelayChange 		= 0
		node.lamp.LFlagEX 			= 0
		node.lamp.LStateAlarm 		= 0

		node.lamp.ZBNetAddr 		= 0
		node.lamp.LAdvPower 		= 0
		node.lamp.LAuxPower 		= 0
		node.lamp.LampHigherV 		= 0
		node.lamp.LampLowerV 		= 0 
		node.lamp.LampHigherI 		= 0
		node.lamp.LampLowerI 		= 0
		node.lamp.LampHigherP 		= 0
		node.lamp.LampLowerP 		= 0
		node.lamp.LampHigherPF 		= 0 
		node.lamp.LampLowerPF 		= 0
		node.lamp.LFlagSetNum 		= 0
		node.lamp.LFlagSetAdu 		= 0
		node.lamp.LChecksum 		= 0
		//以上54字节
		//以下不上传
		node.lamp.LNAdvSet 			= 0
		node.lamp.LNAuxSet 			= 0
		node.lamp.LNStateRelayBT 	= 0x03
		node.lamp.LNCountDisconn 	= 0
		node.lamp.LNFlagDataback 	= 0
		node.lamp.LNAdvPwm 			= 100
		node.lamp.LNAuxPwm 			= 100
		node.lamp.LNStateByHand 	= 0
		node.lamp.LNNumOutput 		= 2
		node.lamp.LNDetecOK 		= 0

		// 追加操作
		GZigbeeNode = append(GZigbeeNode, node)

	}
	fmt.Println("---> GZigbeeNode: ", reflect.TypeOf(GZigbeeNode), reflect.ValueOf(GZigbeeNode))
	fmt.Println("--->01 ", GZigbeeNode)
	fmt.Println("--->02 ", GZigbeeNode[0].id, GZigbeeNode[1].id)
	fmt.Println("--->03 ", GZigbeeNode[2].lamp, GZigbeeNode[2].lamp.LampHigherV)

	GZigbeeNode.GZB72053InductAndTrigeTimeActer(1, 1, 2, 3, 4, 5)
	
}

//设置节点控制器的感应恢复时间和触发间隔时间
func (gzbNode GZBNode)GZB72053InductAndTrigeTimeActer(lampNum, buffA, buffB, modeTX, groupNum, snum int) {
	fmt.Println("---> Acter 72053", gzbNode[2].id)
	gzbNode[2].lamp.LampHigherV = 1000
	fmt.Println("---> ActerVVV ", GZigbeeNode[2].lamp, GZigbeeNode[2].lamp.LampHigherV)

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPENODE | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x23
	sliceCommand[7] = 0x04
	sliceCommand[8] = (buffA>>8) & 0xFF
	sliceCommand[9] = buffA & 0xFF
	sliceCommand[10] = (buffB>>8) & 0xFF
	sliceCommand[11] = buffB & 0xFF

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72053 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72053, sliceCommand, snum)

	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// 设置单灯组号
func (gzbNode GZBNode)GZB72054SetLampGroupActer(lampNum, gNum, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x01
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = gNum

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72054 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72054, sliceCommand, snum)

	// TODO Uart.Send()

	// TODO CMD.StateOK
}

// 单灯主辅互换
func (gzbNode GZBNode)GZB72055SetLampRelayChangeActer(lampNum, stateOk, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x02
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = stateOk

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72055 sliceCommandOK.", sliceCommand)


	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72055, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK
}

// 设置单灯与节点关联
func (gzbNode GZBNode)GZB72057SetLampAndNodeRelatedActer(lampNum, numMax, numMin, modeTX, groupNum, snum int) {
	
	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x0C
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = numMax
	sliceCommand[11] = numMin

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72057 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72057, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// 设置单灯在节点触发后的感应恢复时间
func (gzbNode GZBNode)GZB72058SetLampInductRecoverTimeActer(lampNum, timeRecover, modeTX, groupNum, snum int) {
	
	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x0D
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = (timeRecover>>8)&0xFF
	sliceCommand[11] = timeRecover&0xFF

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72058 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72058, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// 设置命令单灯自校验
func (gzbNode GZBNode)GZB72059SetLampSelfCheckActer(lampNum, modeTX, groupNum, snum int) {
	
	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x0E
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72059 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72059, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

//  固定单灯拨码
func (gzbNode GZBNode)GZB72060SetLampFixedKeyActer(lampNum, stateOk, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x0F
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = stateOk

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72060 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72060, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK
}

// 查询返回单灯ieee地址
func (gzbNode GZBNode)GZB72061QueryLampIEEEActer(lampNum, modeTX, groupNum, snum int) {
	
	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x10
	sliceCommand[6] = 0x11
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72061 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72061, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// 单灯常关
func (gzbNode GZBNode)GZB72062SetLampAlwaysCloseActer(lampNum, stateOk, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x15
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = (stateOk>>8)&0xFF
	sliceCommand[11] = stateOk&0xFF

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72062 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72062, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK
}

//单灯内部开主灯
// stateRelay:标识不同类型(主灯&辅灯&主辅灯)
func (gzbNode GZBNode)GZB72063SetLampInterOpenActer(lampNum, stateRelay, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)

	switch stateRelay {
	case 0x01: // 开主灯
		sliceCommand[5] = 0x10
		sliceCommand[6] = 0x12
		if modeTX == 0x01 {
			// 开一个灯
			fmt.Println("---> modeTx0x01.")
		} else if modeTX == 0x02 {
			// 开所有灯
			fmt.Println("---> modeTx0x02.")
		} else {
			// 开一组灯
			fmt.Println("---> modeTx0x03.")
		}
	case 0x02: // 开辅灯
		sliceCommand[5] = 0x10
		sliceCommand[6] = 0x14
		if modeTX == 0x01 {
			// 开一个灯
			fmt.Println("---> modeTx0x01.")
		} else if modeTX == 0x02 {
			// 开所有灯
			fmt.Println("---> modeTx0x02.")
		} else {
			// 开一组灯
			fmt.Println("---> modeTx0x03.")
		}
	case 0x03: // 开主辅灯
		sliceCommand[5] = 0x10
		sliceCommand[6] = 0x16
		if modeTX == 0x01 {
			// 开一个灯
			fmt.Println("---> modeTx0x01.")
		} else if modeTX == 0x02 {
			// 开所有灯
			fmt.Println("---> modeTx0x02.")
		} else {
			// 开一组灯
			fmt.Println("---> modeTx0x03.")
		}
	}
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72063 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72063, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK
}

// 单灯内部关主灯
// stateRelay:标识不同类型(主灯&辅灯&主辅灯)
func (gzbNode GZBNode)GZB72064SetLampInterCloseActer(lampNum, stateRelay, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)

	switch stateRelay {
	case 0x01: // 关主灯
		sliceCommand[5] = 0x10
		sliceCommand[6] = 0x13
		if modeTX == 0x01 {
			// 关一个灯
			fmt.Println("---> modeTx0x01.")
		} else if modeTX == 0x02 {
			// 关所有灯
			fmt.Println("---> modeTx0x02.")
		} else {
			// 关一组灯
			fmt.Println("---> modeTx0x03.")
		}
	case 0x02: // 关辅灯
		sliceCommand[5] = 0x10
		sliceCommand[6] = 0x15
		if modeTX == 0x01 {
			// 关一个灯
			fmt.Println("---> modeTx0x01.")
		} else if modeTX == 0x02 {
			// 关所有灯
			fmt.Println("---> modeTx0x02.")
		} else {
			// 关一组灯
			fmt.Println("---> modeTx0x03.")
		}
	case 0x03: // 关主辅灯
		sliceCommand[5] = 0x10
		sliceCommand[6] = 0x17
		if modeTX == 0x01 {
			// 开一个灯
			fmt.Println("---> modeTx0x01.")
		} else if modeTX == 0x02 {
			// 开所有灯
			fmt.Println("---> modeTx0x02.")
		} else {
			// 开一组灯
			fmt.Println("---> modeTx0x03.")
		}
	}
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72064 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72064, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK
}

// 单灯电量保存
func (gzbNode GZBNode)GZB72069SetLampSaveElecsActer(lampNum, modeTX, groupNum, snum int) {
	
	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x27
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72069 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72069, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// 设置单灯电压报警上下限
func (gzbNode GZBNode)GZB72070SetLampAlarmLimitVActer(lampNum, numMax, numMin, modeTX, groupNum, snum int) {
	
	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x06
	sliceCommand[7] = 0x04
	sliceCommand[8] = (numMax>>8)&0xFF
	sliceCommand[9] = numMax&0xFF
	sliceCommand[10] = (numMin>>8)&0xFF
	sliceCommand[11] = numMin&0xFF

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72070 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72070, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// 设置单灯电流报警上下限
func (gzbNode GZBNode)GZB72071SetLampAlarmLimitIActer(lampNum, numMax, numMin, modeTX, groupNum, snum int) {
	
	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x07
	sliceCommand[7] = 0x04
	sliceCommand[8] = (numMax>>8)&0xFF
	sliceCommand[9] = numMax&0xFF
	sliceCommand[10] = (numMin>>8)&0xFF
	sliceCommand[11] = numMin&0xFF

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72071 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72071, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// 设置单灯功率报警上下限
func (gzbNode GZBNode)GZB72072SetLampAlarmLimitPActer(lampNum, numMax, numMin, modeTX, groupNum, snum int) {
	
	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x08
	sliceCommand[7] = 0x04
	sliceCommand[8] = (numMax>>8)&0xFF
	sliceCommand[9] = numMax&0xFF
	sliceCommand[10] = (numMin>>8)&0xFF
	sliceCommand[11] = numMin&0xFF

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72072 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72072, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// 设置单灯功率因素报警上下限
func (gzbNode GZBNode)GZB72073SetLampAlarmLimitPFActer(lampNum, numMax, numMin, modeTX, groupNum, snum int) {
	
	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x09
	sliceCommand[7] = 0x04
	sliceCommand[8] = (numMax>>8)&0xFF
	sliceCommand[9] = numMax&0xFF
	sliceCommand[10] = (numMin>>8)&0xFF
	sliceCommand[11] = numMin&0xFF

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72073 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72073, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// 清除单灯能量数据
func (gzbNode GZBNode)GZB72074SetLampAnergyClearActer(lampNum, modeTX, groupNum, snum int) {
	
	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x0A
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72074 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72074, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// 设置单灯电压放大倍数
// stateRelay:标识不同类型(主灯&辅灯)
func (gzbNode GZBNode)GZB72075SetLampAmplifyVActer(lampNum, stateOk, valueAmpli, stateRelay, modeTX, groupNum, snum int) {
	
	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPENODE | ((lampNum>>8)<<4)
	switch stateRelay {
	case 0x01:
		sliceCommand[5] = 0x20
		sliceCommand[6] = 0x11
	case 0x02:
		sliceCommand[5] = 0x20
		sliceCommand[6] = 0x12
	default:
		sliceCommand[5] = 0x20
		sliceCommand[6] = 0x11
	}
	
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = valueAmpli
	sliceCommand[11] = stateOk

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72075 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72075, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}


// 设置单灯电流放大倍数
// stateRelay:标识不同类型(主灯&辅灯)
func (gzbNode GZBNode)GZB72077SetLampAmplifyIActer(lampNum, stateOk, valueAmpli, stateRelay, modeTX, groupNum, snum int) {
	
	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPENODE | ((lampNum>>8)<<4)
	switch stateRelay {
	case 0x01:
		sliceCommand[5] = 0x20
		sliceCommand[6] = 0x13
	case 0x02:
		sliceCommand[5] = 0x20
		sliceCommand[6] = 0x14
	default:
		sliceCommand[5] = 0x20
		sliceCommand[6] = 0x13
	}
	
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = valueAmpli
	sliceCommand[11] = stateOk

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72077 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72077, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// rn8209有效值offset和有功offset校正
func (gzbNode GZBNode)GZB72079SetLampRN8209OffsetActer(lampNum, snum int) {
	
	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum
	sliceCommand[2] = lampNum & 0xFF
	sliceCommand[3] = 0x01
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x09
	sliceCommand[6] = 0x96
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72079 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72079, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// rn8209参数设置
func (gzbNode GZBNode)GZB72080SetLampRN8209ParasActer(lampNum, snum int) {
	
	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum
	sliceCommand[2] = lampNum & 0xFF
	sliceCommand[3] = 0x01
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x09
	sliceCommand[6] = 0x97
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72080 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72080, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// rn8209比例系数设置
func (gzbNode GZBNode)GZB72081SetLampRN8209RatioActer(lampNum, snum int) {
	
	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum
	sliceCommand[2] = lampNum & 0xFF
	sliceCommand[3] = 0x01
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x09
	sliceCommand[6] = 0x98
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72081 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72081, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// 单灯与rtu时间同步命令
func (gzbNode GZBNode)GZB72082SetLampRtuSyncTimeActer(lampNum, timeSec, timeMin, timeH, timeD, timeW, timeMon, timeY, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 17)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x1E
	sliceCommand[7] = 0x07
	sliceCommand[8] = timeSec
	sliceCommand[9] = timeMin
	sliceCommand[10] = timeH
	sliceCommand[11] = timeD
	sliceCommand[12] = timeW
	sliceCommand[13] = timeMon
	sliceCommand[14] = timeY%100

	for i := 0; i < 15; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[15] = checkSum
	sliceCommand[16] = 0x99

	fmt.Println("---> 72082 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72082, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK
}

// 开关时间同步?
func (gzbNode GZBNode)GZB72083SetLampSwitchSyncTimeActer(lampNum, timeOnSec, timeOnMin, timeOnH, timeOffSec, timeOffMin, timeOffH, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 16)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x20
	sliceCommand[7] = 0x06
	sliceCommand[8] = timeOnSec
	sliceCommand[9] = timeOnMin
	sliceCommand[10] = timeOnH
	sliceCommand[11] = timeOffSec
	sliceCommand[12] = timeOffMin
	sliceCommand[13] = timeOffH

	for i := 0; i < 14; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[14] = checkSum
	sliceCommand[15] = 0x99

	fmt.Println("---> 72083 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72083, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK
}

// 查询锂电池电平
func (gzbNode GZBNode)GZB72084QueryBatteryVoltActer(lampNum, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x10
	sliceCommand[6] = 0x18
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72084 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72084, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK
}

// 查询温度
func (gzbNode GZBNode)GZB72085QueryTemperatureActer(lampNum, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x10
	sliceCommand[6] = 0x19
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72085 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72085, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK
}

// 擦除NV，重新加入网络
func (gzbNode GZBNode)GZB72086SetClearNVAdderActer(lampNum, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x10
	sliceCommand[6] = 0x20
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72086 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72086, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK
}

// 单灯开闪烁功能
func (gzbNode GZBNode)GZB72087SetLampTwinkleActer(lampNum, stateOk, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x23
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = stateOk

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72087 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72087, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK
}

// 单灯召测命令
func (gzbNode GZBNode)GZB72088QueryLampCalledDataActer(lampNum, modeTX, groupNum, snum int) {
	fmt.Println("---> 72088 sliceCommandOK.", modeTX)

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01: // 单召
		sliceCommand[2] = lampNum & 0xFF
		sliceCommand[3] = modeTX
		sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
		sliceCommand[5] = 0x10
		sliceCommand[6] = 0x09
		sliceCommand[7] = 0x04
		sliceCommand[8] = 0x00
		sliceCommand[9] = 0x00
		sliceCommand[10] = 0x00
		sliceCommand[11] = 0x00

		for i := 0; i < 12; i++ {
			checkSum ^= sliceCommand[i]
		}
		sliceCommand[12] = checkSum
		sliceCommand[13] = 0x99

		fmt.Println("---> 72088 sliceCommandOK.", sliceCommand)

		// 将数据(id&data&snum)发送到channel
		HanldeCmdZigbeeLoading(72088, sliceCommand, snum)
		// TODO Uart.Send()

		// TODO CMD.StateOK

	case 0x02: //召测全部
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	// sliceCommand[3] = modeTX
	// sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	// sliceCommand[5] = 0x20
	// sliceCommand[6] = 0x25
	// sliceCommand[7] = 0x04
	// sliceCommand[8] = 0x00
	// sliceCommand[9] = 0x00
	// sliceCommand[10] = 0x00
	// sliceCommand[11] = 0x00

	// for i := 0; i < 12; i++ {
	// 	checkSum ^= sliceCommand[i]
	// }
	// sliceCommand[12] = checkSum
	// sliceCommand[13] = 0x99

	// fmt.Println("---> 72053 sliceCommandOK.", sliceCommand)

	// TODO Uart.Send()

	// TODO CMD.StateOK
}

//单灯手动开主灯
// stateRelay:标识不同类型(主灯&辅灯&主辅灯)
func (gzbNode GZBNode)GZB72090SetLampOpenByHandsActer(lampNum, stateRelay, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)

	switch stateRelay {
	case 0x01: // 开主灯
		sliceCommand[5] = 0x10
		sliceCommand[6] = 0x05
		if modeTX == 0x01 {
			// 开一个灯
			fmt.Println("---> modeTx0x01.")
		} else if modeTX == 0x02 {
			// 开所有灯
			fmt.Println("---> modeTx0x02.")
		} else {
			// 开一组灯
			fmt.Println("---> modeTx0x03.")
		}
	case 0x02: // 开辅灯
		sliceCommand[5] = 0x10
		sliceCommand[6] = 0x07
		if modeTX == 0x01 {
			// 开一个灯
			fmt.Println("---> modeTx0x01.")
		} else if modeTX == 0x02 {
			// 开所有灯
			fmt.Println("---> modeTx0x02.")
		} else {
			// 开一组灯
			fmt.Println("---> modeTx0x03.")
		}
	case 0x03: // 开主辅灯
		sliceCommand[5] = 0x10
		sliceCommand[6] = 0x03
		if modeTX == 0x01 {
			// 开一个灯
			fmt.Println("---> modeTx0x01.")
		} else if modeTX == 0x02 {
			// 开所有灯
			fmt.Println("---> modeTx0x02.")
		} else {
			// 开一组灯
			fmt.Println("---> modeTx0x03.")
		}
	}
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72090 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72090, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK
}

// 单灯手动关主灯
// stateRelay:标识不同类型(主灯&辅灯&主辅灯)
func (gzbNode GZBNode)GZB72091SetLampCloseByHandsActer(lampNum, stateRelay, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)

	switch stateRelay {
	case 0x01: // 关主灯
		sliceCommand[5] = 0x10
		sliceCommand[6] = 0x06
		if modeTX == 0x01 {
			// 关一个灯
			fmt.Println("---> modeTx0x01.")
		} else if modeTX == 0x02 {
			// 关所有灯
			fmt.Println("---> modeTx0x02.")
		} else {
			// 关一组灯
			fmt.Println("---> modeTx0x03.")
		}
	case 0x02: // 关辅灯
		sliceCommand[5] = 0x10
		sliceCommand[6] = 0x08
		if modeTX == 0x01 {
			// 关一个灯
			fmt.Println("---> modeTx0x01.")
		} else if modeTX == 0x02 {
			// 关所有灯
			fmt.Println("---> modeTx0x02.")
		} else {
			// 关一组灯
			fmt.Println("---> modeTx0x03.")
		}
	case 0x03: // 关主辅灯
		sliceCommand[5] = 0x10
		sliceCommand[6] = 0x04
		if modeTX == 0x01 {
			// 开一个灯
			fmt.Println("---> modeTx0x01.")
		} else if modeTX == 0x02 {
			// 开所有灯
			fmt.Println("---> modeTx0x02.")
		} else {
			// 开一组灯
			fmt.Println("---> modeTx0x03.")
		}
	}
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72091 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72091, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK
}

// 巡检manny命令
func (gzbNode GZBNode)GZB72092QueryLampDetecDataMannyActer(lampNum, numMax, numMin, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x21
	sliceCommand[6] = 0x04
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = numMin
	sliceCommand[11] = numMax

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72092 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72092, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// 巡检内部命令
func (gzbNode GZBNode)GZB72093QueryLampDetecDataInterActer(lampNum, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x21
	sliceCommand[6] = 0x02
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72093 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72093, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// 节点开关命令
func (gzbNode GZBNode)GZB72100SetNodeSwitchActer(lampNum, stateOk, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 16)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum
	sliceCommand[2] = lampNum & 0xFF
	sliceCommand[3] = 0x02
	sliceCommand[4] = TYPENODE | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x21
	sliceCommand[7] = 0x06
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = stateOk
	switch stateOk {
	case 0x65:
		sliceCommand[12] = 0x00
		sliceCommand[13] = stateOk
	default:
		sliceCommand[12] = 0x01
		sliceCommand[13] = stateOk

	}
	for i := 0; i < 14; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[14] = checkSum
	sliceCommand[15] = 0x99

	fmt.Println("---> 72100 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72100, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}


//单灯调光
// stateRelay:标识不同类型(主灯&辅灯&主辅灯)
func (gzbNode GZBNode)GZB72101SetLampDimmer(lampNum, stateRelay, lampPwm, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)

	switch stateRelay {
	case 0x01: // 主灯调光
		sliceCommand[5] = 0x10
		sliceCommand[6] = 0x0B
		if modeTX == 0x01 {
			// 一个
			fmt.Println("---> modeTx0x01.")
		} else if modeTX == 0x02 {
			// 所有
			fmt.Println("---> modeTx0x02.")
		} else {
			// 一组
			fmt.Println("---> modeTx0x03.")
		}
	case 0x02: // 辅灯调光
		sliceCommand[5] = 0x10
		sliceCommand[6] = 0x0C
		if modeTX == 0x01 {
			// 一个
			fmt.Println("---> modeTx0x01.")
		} else if modeTX == 0x02 {
			// 所有
			fmt.Println("---> modeTx0x02.")
		} else {
			// 一组
			fmt.Println("---> modeTx0x03.")
		}
	case 0x03: // 主辅灯调光
		sliceCommand[5] = 0x10
		sliceCommand[6] = 0x0A
		if modeTX == 0x01 {
			// 开一个灯
			fmt.Println("---> modeTx0x01.")
		} else if modeTX == 0x02 {
			// 开所有灯
			fmt.Println("---> modeTx0x02.")
		} else {
			// 开一组灯
			fmt.Println("---> modeTx0x03.")
		}
	}
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = lampPwm

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72101 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72101, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK
}

// 返回单灯panid
func (gzbNode GZBNode)GZB72104QueryLampPanidActer(lampNum, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x10
	sliceCommand[6] = 0x0F
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72104 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72104, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// 返回单灯程序版本
func (gzbNode GZBNode)GZB72105QueryLampProgramVersionActer(lampNum, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x10
	sliceCommand[6] = 0x0D
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72105 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72105, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// 重启单灯
func (gzbNode GZBNode)GZB72106SetLampResetActer(lampNum, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x10
	sliceCommand[6] = 0x0E
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72106 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72106, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK

}

// 查询STC时间
func (gzbNode GZBNode)GZB72116QueryLampTimeSTC(lampNum, modeTX, groupNum, snum int) {

	// 初始化局部变量:校验和&命令缓存
	checkSum := 0
	sliceCommand := make([]int, 14)

	// Start
	sliceCommand[0] = 0x33
	sliceCommand[1] = snum

	switch modeTX {
	case 0x01:
		sliceCommand[2] = lampNum & 0xFF
	case 0x02:
		sliceCommand[2] = lampNum & 0xFF
	default:
		sliceCommand[2] = groupNum
	}

	sliceCommand[3] = modeTX
	sliceCommand[4] = TYPELAMP | ((lampNum>>8)<<4)
	sliceCommand[5] = 0x20
	sliceCommand[6] = 0x25
	sliceCommand[7] = 0x04
	sliceCommand[8] = 0x00
	sliceCommand[9] = 0x00
	sliceCommand[10] = 0x00
	sliceCommand[11] = 0x00

	for i := 0; i < 12; i++ {
		checkSum ^= sliceCommand[i]
	}
	sliceCommand[12] = checkSum
	sliceCommand[13] = 0x99

	fmt.Println("---> 72116 sliceCommandOK.", sliceCommand)

	// 将数据(id&data&snum)发送到channel
	HanldeCmdZigbeeLoading(72116, sliceCommand, snum)
	// TODO Uart.Send()

	// TODO CMD.StateOK
}