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


type GZigbeeNodeActer interface {
	GZB72053InductAndTrigeTimeActer()
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
	sliceCommand[4] = TYPENODE
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

}
