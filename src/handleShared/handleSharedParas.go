package handleShared

/*
内容:
	全局共享变量用于其他模块使用
	WD:设置参数
	FQ:查询参数
Author:mengfei.wu@foxmail.com
*/
var (
	FQRTUVERSION 				= 0xFFFFFFFF 	//RTU程序版本编号
	FQLevelSignal				= 0			 	//3G/4G信号强度
	FQStateLatLongitude 		= 0				//RTU经纬度开关状态 
	FQTemperature				= 0				//查询RTU温度值
	FQStateDoor					= 0				//查询RTU门的状态
	FQValueBattery				= 0 			//查询锂电池电平值

	WDLevelTopUnusualSwitch		= 0 			//RTU意外亮灭灯报警阈值
	WDValueLatitude 			= 0				//设置经度
	WDValueLongitude 			= 0 			//设置纬度
	WDIntervalTimeElec 			= 0				//设置RTU电参数采样间隔时间
	WDRatioTransformer			= 0 			//设置外接电流互感器比例
	WDValueLoop					= 0				//设置回路数量
	WDAlarmShielding 			= 0				//设置报警屏蔽
	WDValueLamp					= 0				//设置单灯数量
	WDRelayOpenAllBit			= 0				//按位开继电器(总开)
	WDGroupLamp					= 0				//设置单灯组数
	WDRelayOpenAll				= 0 			//设置RTU继电器全开
	WDRelayCloseAll				= 0 			//设置RTU继电器全关
	WDCurrentLimitLoop0X		= 0 			//设置RTU回路电流上下限
	WDVoltLimitLoop0X			= 0				//设置RTU回路电压上下限
	WDRelayOpen					= 0 			//设置RTU继电器1-6开
	WDRelayClose				= 0				//设置RTU继电器1-6关
	WDLampAlarmShielding		= 0				// 设置单灯报警屏蔽
	WDDetcTimeLamp				= 0				// 设置单灯巡检时间
	WDLoopAlwaysOn				= 0 			// 设置回路常开模式


	WDFlagTempreratureBack		= 0 			// 温度采样返回标志
	WDFlagNoTempreratureCeiling = 0				// 温度采样未返回数据标志
	
	// Redis相关
	WDTimeLatitude00			= "Latitude00"	// 经度
	WDTimeLatitude01			= "Latitude01"	// 经度
	WDTimeLatitude02			= "Latitude02"	// 经度
	WDTimeLatitude03			= "Latitude03"	// 经度
	WDTimeLongitude00			= "Longitude00"	// 纬度
	WDTimeLongitude01			= "Longitude01"	// 纬度
	WDTimeLongitude02			= "Longitude02"	// 纬度
	WDTimeLongitude03			= "Longitude03"	// 纬度
	
)

const (
	WDMCPADDR20					= 0x20 // MCP23008页地址
	WDPCFADDR51					= 0x51  // PCF8563页地址0x51
	WDEEPADDR57					= 0x57  // eeprom页地址0x57

	WDMCP20CHECK001				= 1 // MCP23008: 0x20-1
	WDEEP57CHECK001				= 1 // eeprom: 0x57-1
	WDPCF51CHECK001				= 1 // PCF8563: 0x51-1
)