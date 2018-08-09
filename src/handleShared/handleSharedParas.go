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
	WDQuantityLoop				= "WDQuantityLoop"				//设置回路数量
	WDStateLoop					= "WDStateLoop"					//回路状态
	WDStateChangeRelay			= "WDStateChangeRelay"			//继电器状态改变
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

	WDEleEnergyC1				= "EleEnergyC1" // 回路1能量
	WDEleEnergyC2				= "EleEnergyC2" // 回路2能量
	WDEleEnergyC3				= "EleEnergyC3" // 回路3能量
	WDEleEnergyC4				= "EleEnergyC4" // 回路4能量
	WDEleEnergyC5				= "EleEnergyC5" // 回路5能量
	WDEleEnergyC6				= "EleEnergyC6" // 回路6能量

	// *andleEleCollectManage
	WDSysPowerMode				= "SysPowerMode" //供电方式

	// 回路相关
	
	
)

const (

	WDPOWERMODEAC				= 0x01	// AC供电
	WDPOWERMODELI				= 0x02	// LI供电

	WDMCPADDR20					= 0x20 // MCP23008页地址
	WDPCFADDR51					= 0x51  // PCF8563页地址0x51
	WDRN1ADDR54					= 0x54 // RN8209页地址 使用AT24c08
	WDEEPADDR57					= 0x57  // eeprom页地址0x57

	// 0x20
	WDMCP20CHECK001				= 1 // MCP23008: 0x20-1
	WDMCP20PINMODE				= 2 // PINMODE
	WDMCP20PINVALUE				= 3 // PINVALUE

	// 0x51
	WDPCF51CHECK001				= 1 // PCF8563: 0x51-1

	// 0x54
	WDRN1KIA1BIT0				= 179
	WDRN1KIA1BIT1				= 180
	WDRN1KIA1BIT2				= 181
	WDRN1KIA1BIT3				= 182
	WDRN1KIA2BIT0				= 183
	WDRN1KIA2BIT1				= 184
	WDRN1KIA2BIT2				= 185
	WDRN1KIA2BIT3				= 186
	WDRN1KIA3BIT0				= 187
	WDRN1KIA3BIT1				= 188
	WDRN1KIA3BIT2				= 189
	WDRN1KIA3BIT3				= 190
	WDRN1KIA6BIT0				= 191
	WDRN1KIA6BIT1				= 192
	WDRN1KIA6BIT2				= 193
	WDRN1KIA6BIT3				= 194
	WDRN1KIA5BIT0				= 195
	WDRN1KIA5BIT1				= 196
	WDRN1KIA5BIT2				= 197
	WDRN1KIA5BIT3				= 198
	WDRN1KIA4BIT0				= 199
	WDRN1KIA4BIT1				= 200
	WDRN1KIA4BIT2				= 201
	WDRN1KIA4BIT3				= 202

	WDRN1KUA1BIT0				= 203
	WDRN1KUA1BIT1				= 204
	WDRN1KUA1BIT2				= 205
	WDRN1KUA1BIT3				= 206
	WDRN1KUA2BIT0				= 207
	WDRN1KUA2BIT1				= 208
	WDRN1KUA2BIT2				= 209
	WDRN1KUA2BIT3				= 210
	WDRN1KUA3BIT0				= 211
	WDRN1KUA3BIT1				= 212
	WDRN1KUA3BIT2				= 213
	WDRN1KUA3BIT3				= 214
	WDRN1KUA6BIT0				= 215
	WDRN1KUA6BIT1				= 216
	WDRN1KUA6BIT2				= 217
	WDRN1KUA6BIT3				= 218
	WDRN1KUA5BIT0				= 219
	WDRN1KUA5BIT1				= 220
	WDRN1KUA5BIT2				= 221
	WDRN1KUA5BIT3				= 222
	WDRN1KUA4BIT0				= 223
	WDRN1KUA4BIT1				= 224
	WDRN1KUA4BIT2				= 225
	WDRN1KUA4BIT3				= 226

	WDRN1KPA1BIT0				= 227
	WDRN1KPA1BIT1				= 228
	WDRN1KPA1BIT2				= 229
	WDRN1KPA1BIT3				= 230
	WDRN1KPA2BIT0				= 231
	WDRN1KPA2BIT1				= 232
	WDRN1KPA2BIT2				= 233
	WDRN1KPA2BIT3				= 234
	WDRN1KPA3BIT0				= 235
	WDRN1KPA3BIT1				= 236
	WDRN1KPA3BIT2				= 237
	WDRN1KPA3BIT3				= 238
	WDRN1KPA6BIT0				= 239
	WDRN1KPA6BIT1				= 240
	WDRN1KPA6BIT2				= 241
	WDRN1KPA6BIT3				= 242
	WDRN1KPA5BIT0				= 243
	WDRN1KPA5BIT1				= 244
	WDRN1KPA5BIT2				= 245
	WDRN1KPA5BIT3				= 246
	WDRN1KPA4BIT0				= 247
	WDRN1KPA4BIT1				= 248
	WDRN1KPA4BIT2				= 249
	WDRN1KPA4BIT3				= 250



	// 0x57
	WDEEP57CHECK001				= 1 // eeprom: 0x57-1

	
)