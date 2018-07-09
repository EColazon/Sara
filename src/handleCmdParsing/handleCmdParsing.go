package handleCmdParsing

/* 72000:2F命令起始编号
		 7代表月
		 2代表原2F格式
		 000代表三位编号,最大999
   73000:33命令起始编号
		 7代表月
		 3代表原33格式
		 000代表三位编号,最大999

Author:mengfei.wu@foxmail.com
---------2018.07.09---------
*/

import (
	"fmt"
	"reflect"
)

const (
	CMD33HEAD = 0x33
	CMD33TAIL = 0x99
	CMD2FHEAD00 = 0x2F
	CMD2FHEAD01 = 0x43
	CMD2FHEAD02 = 0x2F
	CMD2FTAIL = 0xCC

)
// 声明全局缓冲通道用于命令解析到命令分发间通信
var ChCmd = make(chan map[string]interface{}, 1024)
// 声明map格式用于拼组数据
var MapCmd = make(map[string]interface{})

func Cmd33Parsing() {
	
	sliceCmd := []int{0x33, 0x00, 0x01, 0x03, 0x01, 0x60, 0x06, 0x01, 0x00, 0x57, 0x99}
	//fmt.Println("---> typeOf(sliceCmd)---> ",reflect.TypeOf(sliceCmd))

	for index, value := range sliceCmd {

		if index == 0 {
			fmt.Println(reflect.TypeOf(sliceCmd[index]), sliceCmd[index])
			fmt.Println(reflect.TypeOf(CMD33HEAD), CMD33HEAD)
			fmt.Println("Hello", index, value)
			if sliceCmd[index] == CMD33HEAD {
				fmt.Println("Head is here.")
				lengthData := sliceCmd[index+7]
				fmt.Println("lengthData--> ", lengthData)
				snumData := sliceCmd[index+1]
				fmt.Println("snumData---> ", snumData)
				if sliceCmd[lengthData+9] == CMD33TAIL {
					fmt.Println("Tail is here.")
				check := 0
					for i := 0; i < lengthData+8; i++ {
						check ^= sliceCmd[i]
					}
					fmt.Println("check---> ", check)
					if check == sliceCmd[lengthData+8] {
						fmt.Println("Check is ok")
						commandCmd := (sliceCmd[5]<<8) | sliceCmd[6]
						fmt.Println("commandCmd ---> ", commandCmd)

						switch commandCmd {
						case 0x6002:
							MapCmd["id"] = 73000
							MapCmd["data"] = sliceCmd
							ChCmd <- MapCmd
							fmt.Println("---> 6002 is here.")
						case 0x6004:
							MapCmd["id"] = 73001
							MapCmd["data"] = sliceCmd
							ChCmd <- MapCmd
							fmt.Println("6004 is here.")
						case 0x6005:
							MapCmd["id"] = 73002
							MapCmd["data"] = sliceCmd
							ChCmd <- MapCmd
							fmt.Println("6005 is here.")
						case 0x6006:
							MapCmd["id"] = 73003
							MapCmd["data"] = sliceCmd
							ChCmd <- MapCmd
							fmt.Println("6006 is here.")
						case 0x6007:
							MapCmd["id"] = 73004
							MapCmd["data"] = sliceCmd
							ChCmd <- MapCmd
							fmt.Println("6007 is here.")
						case 0x6019:
							MapCmd["id"] = 73005
							MapCmd["data"] = sliceCmd
							ChCmd <- MapCmd
							fmt.Println("6019 is here.")
						case 0x6020:
							MapCmd["id"] = 73006
							MapCmd["data"] = sliceCmd
							ChCmd <- MapCmd
							fmt.Println("6020 is here.")
						case 0x6021:
							MapCmd["id"] = 73007
							MapCmd["data"] = sliceCmd
							ChCmd <- MapCmd
							fmt.Println("6021 is here.")
						case 0x2600:
							MapCmd["id"] = 73008
							MapCmd["data"] = sliceCmd
							ChCmd <- MapCmd
							fmt.Println("2600 is here.")
						case 0x2400:
							MapCmd["id"] = 73009
							MapCmd["data"] = sliceCmd
							ChCmd <- MapCmd
							fmt.Println("2400 is here.")
						default:
							fmt.Println("33 not fetched..")
							return
						}
					} else {
						fmt.Println("Check33 is error.")
						break
					}
				} else {
					fmt.Println("Tail33 is error.")
					break
				}
			} else {
				fmt.Println("Head33 is error.")
				break
			}
		} else {
			fmt.Println("Do 33 Nothing.")
			break
		}

	}

	fmt.Println("Hello cmdparsing")

}

func Cmd2FParsing() {
	// 序列号放在首位
	
	sliceCmd := []int{0x01, 0x2F, 0x43, 0x2F, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,0x0A, 0xC0, 0xB4, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xCC}

	for index, value := range sliceCmd {

		if index == 0 {
			fmt.Println(reflect.TypeOf(sliceCmd[index]), sliceCmd[index])
			fmt.Println(reflect.TypeOf(CMD2FHEAD00), CMD2FHEAD00)
			fmt.Println("Hello", len(sliceCmd), index, value)

			if sliceCmd[index+1] == CMD2FHEAD00 && sliceCmd[index+2] == CMD2FHEAD01 && sliceCmd[index+3] == CMD2FHEAD02 && sliceCmd[index+21] == CMD2FTAIL {
				fmt.Println("Head is here.")			
				// 目的识别码 objIdent
				objIdent := sliceCmd[index+4]

				switch objIdent {
					case 0xA2:
						// 返回8字节IEEE地址
						MapCmd["id"] = 72000
						MapCmd["data"] = sliceCmd
						ChCmd <- MapCmd
						fmt.Println("TO DO ---> 0xA2.")
					
					case 0xA3:
						// 返回程序版本号
						MapCmd["id"] = 72001
						MapCmd["data"] = sliceCmd
						ChCmd <- MapCmd
						fmt.Println("TO DO ---> 0xA3.")

					case 0xA4:
						// 复位设备-RTU
						MapCmd["id"] = 72002
						MapCmd["data"] = sliceCmd
						ChCmd <- MapCmd
						fmt.Println("TO DO ---> 0xA4.")
					
					case 0xAD:
						// 复位设备-协调器
						MapCmd["id"] = 72003
						MapCmd["data"] = sliceCmd
						ChCmd <- MapCmd
						fmt.Println("TO DO ---> 0xAD.")
					
					case 0xA5:
						// 读取eeprom数据-RTU
						MapCmd["id"] = 72004
						MapCmd["data"] = sliceCmd
						ChCmd <- MapCmd
						fmt.Println("TO DO ---> 0xA5.")
					
					case 0xA6:
						// RTU供电方式查询
						MapCmd["id"] = 72005
						MapCmd["data"] = sliceCmd
						ChCmd <- MapCmd
						fmt.Println("TO DO ---> 0xA6.")
					
					case 0x05:
						// 其他RTU命令处理-RTU
						fmt.Println("TO DO ---> 0x05.")
						cmdRtuParsing(sliceCmd)
					
					case 0x01, 0x02, 0x03:
						// another
						fmt.Println("TO DO ---> Another.")
						cmdZigbeeParsing(sliceCmd)
					
					default:
						fmt.Println("objIdent not fetched.")
						break

				}
				
			} else {
					fmt.Println("Head2f is error.")
					break
			}
		} else {
			fmt.Println("Do 2f Nothing.")
			break
		}

	}

	fmt.Println("Hello cmdparsing")

}

func cmdRtuParsing(buff []int) {
	// RTU相关命令处理 Another

	sliceCmdRTU := buff
	fmt.Println("sliceCmdRTU---> ", sliceCmdRTU)

	check := 0
	for i := 4; i < 20; i++ {
		check ^= sliceCmdRTU[i]
	}
	fmt.Println("CheckRTU---> ", check)

	if check == sliceCmdRTU[20] {
		// 标识功能动作
		proIdent := sliceCmdRTU[13]

		switch proIdent {
		case 0x0A:
			fmt.Println("Process ---> 0x0A")
			parsingProcess0A(sliceCmdRTU)
		case 0x8A:
			fmt.Println("Process ---> 0x8A")
			parsingProcess8A(sliceCmdRTU)
		case 0x8B:
			fmt.Println("Process ---> 0x8B")
			parsingProcess8B(sliceCmdRTU)
		case 0x8C:
			fmt.Println("Process ---> 0x8C")
			parsingProcess8C(sliceCmdRTU)
		case 0x8D:
			fmt.Println("Process ---> 0x8D")
			parsingProcess8D(sliceCmdRTU)
		case 0x8E:
			fmt.Println("Process ---> 0x8E")
			parsingProcess8E(sliceCmdRTU)
		case 0x8F:
			fmt.Println("Process ---> 0x8F")
			parsingProcess8F(sliceCmdRTU)
		case 0x90:
			fmt.Println("Process ---> 0x90")
			parsingProcess90(sliceCmdRTU)
		case 0x93:
			fmt.Println("Process ---> 0x93")
		default:
			fmt.Println("Process ---> not fetched in sliceCmdRTU.")

		}
	}

}

func cmdZigbeeParsing(buff []int) {
	// Zigbee相关命令处理 01-02-03
	sliceCmdZig := buff

	fmt.Println("sliceCmdZig---> ", sliceCmdZig)

	if sliceCmdZig[12] == 0x70 {
		fmt.Println("sliceCmdZig---> inter 0x70")
		parsingProcess70(sliceCmdZig)
	} else {
		fmt.Println("sliceCmdZig---> inter to parsingZigDeeper.")

		cmdZigbeeParsingDeeper(sliceCmdZig)
	}

}

func cmdZigbeeParsingDeeper(buff []int) {
	// zigbee命令深层解析
	sliceCmdZigDeep := buff

	fmt.Println("sliceCmdZigDeep---> ", sliceCmdZigDeep)

	if ((sliceCmdZigDeep[11]<<8)|(sliceCmdZigDeep[12])>1000) {
		fmt.Println("OverSliceMax ---> zigDeeper")
	} else {
		zig13Ident := sliceCmdZigDeep[13]
		zig14Ident := sliceCmdZigDeep[14]
		zig15Ident := sliceCmdZigDeep[15]
		
		switch zig13Ident {
		case 0x0A:
			fmt.Println("zig13Parsing ---> 0x0A")
			switch zig14Ident {
			case 0x00:
				fmt.Println("zig14Parsing ---> 0x00")
				switch zig15Ident {
				case 0x20:
					MapCmd["id"] = 72053
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x0020")
				case 0x92:
					MapCmd["id"] = 72054
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x0092")
				case 0x94:
					MapCmd["id"] = 72055
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x0094")
				case 0x95:
					MapCmd["id"] = 72056
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x0095")
				case 0x96:
					MapCmd["id"] = 72057
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x0096")
				case 0x98:
					MapCmd["id"] = 72058
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x0098")
				case 0x99:
					MapCmd["id"] = 72059
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x0099")
				case 0x9A:
					MapCmd["id"] = 72060
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x009A")
				case 0x9D:
					MapCmd["id"] = 72061
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x009D")
				case 0x9E:
					MapCmd["id"] = 72062
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x009E")
				case 0x9F:
					MapCmd["id"] = 72063
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x009F")
				case 0xAC:
					MapCmd["id"] = 72064
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00AC")
				case 0xAD:
					MapCmd["id"] = 72065
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00AD")
				case 0xAE:
					MapCmd["id"] = 72066
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00AE")
				case 0xAF:
					MapCmd["id"] = 72067
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00AF")
				case 0xB1:
					MapCmd["id"] = 72068
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00B1")
				case 0xB2:
					MapCmd["id"] = 72069
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00B2")
				case 0xB3:
					MapCmd["id"] = 72070
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00B3")
				case 0xB4:
					MapCmd["id"] = 72071
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00B4")
				case 0xB5:
					MapCmd["id"] = 72072
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00B5")
				case 0xB6:
					MapCmd["id"] = 72073
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00B6")
				case 0xB8:
					MapCmd["id"] = 72074
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00B8")
				case 0xB9:
					MapCmd["id"] = 72075
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00B9")
				case 0xAA:
					MapCmd["id"] = 72076
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00AA")
				case 0xAB:
					MapCmd["id"] = 72077
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00AB")
				case 0xFA:
					MapCmd["id"] = 72078
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00FA")
				case 0xFB:
					MapCmd["id"] = 72079
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00FB")
				case 0xFC:
					MapCmd["id"] = 72080
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00FC")
				case 0xBA:
					MapCmd["id"] = 72081
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00BA")
				case 0xBB:
					MapCmd["id"] = 72082
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00BB")
				case 0xBC:
					MapCmd["id"] = 72083
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00BC")
				case 0xBD:
					MapCmd["id"] = 72084
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00BD")
				case 0xBE:
					MapCmd["id"] = 72085
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00BE")
				case 0xBF:
					MapCmd["id"] = 72086
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x00BF")
				case 0x80:
					MapCmd["id"] = 72087
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0x0080")
				default:
					fmt.Println("zig15Parsing 00---> not fetched")
			
				}
			case 0xC0:
				fmt.Println("zig14Parsing ---> 0xC0")
				switch zig15Ident {
				case 0xB4:
					MapCmd["id"] = 72088
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0xC0B4")
				case 0xB6:
					MapCmd["id"] = 72089
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0xC0B6")
				case 0xC1:
					MapCmd["id"] = 72090
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0xC0C1")
				case 0xC2:
					MapCmd["id"] = 72091
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0xC0C2")
				case 0xBE:
					MapCmd["id"] = 72092
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0xC0BE")
				case 0xBF:
					MapCmd["id"] = 72093
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0xC0BF")
				default:
					fmt.Println("zig15Parsing C0---> not fetched")
				}
			case 0xC1:
				fmt.Println("zig14Parsing ---> 0xC1")
				switch zig15Ident {
				case 0xC1:
					MapCmd["id"] = 72094
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0xC1C1")
				case 0xC2:
					MapCmd["id"] = 72095
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0xC1C2")
				default:
					fmt.Println("zig15Parsing C1---> not fetched")
				}
			case 0xC2:
				fmt.Println("zig14Parsing ---> 0xC2")
				switch zig15Ident {
				case 0xC1:
					MapCmd["id"] = 72096
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0xC2C1")
				case 0xC2:
					MapCmd["id"] = 72097
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing ---> 0xC2C2")
				default:
					fmt.Println("zig15Parsing C2---> not fetched")
				}
			case 0xE2:
				MapCmd["id"] = 72098
				MapCmd["data"] = sliceCmdZigDeep
				ChCmd <- MapCmd
				fmt.Println("zig14Parsing ---> 0xE2")
			case 0x01:
				MapCmd["id"] = 72099
				MapCmd["data"] = sliceCmdZigDeep
				ChCmd <- MapCmd
				fmt.Println("zig14Parsing ---> 0x01")
				if zig15Ident == 0x15 {
					MapCmd["id"] = 72100
					MapCmd["data"] = sliceCmdZigDeep
					ChCmd <- MapCmd
					fmt.Println("zig15Parsing 01---> 0x0115")
				} else {
					fmt.Println("zig15Parsing 0115---> not fetched")
				}
			default:
				fmt.Println("zig15Parsing 00---> not fetched")
			}
		case 0x8E:
			fmt.Println("zig13Parsing ---> 0x8E")
			if zig14Ident == 0xC0 && zig15Ident == 0xC1 {
				MapCmd["id"] = 72101
				MapCmd["data"] = sliceCmdZigDeep
				ChCmd <- MapCmd
				fmt.Println("zig15Parsing 8E---> 0xC0C1")
			} else if zig14Ident == 0xC1 && zig15Ident == 0xC1 {
				MapCmd["id"] = 72102
				MapCmd["data"] = sliceCmdZigDeep
				ChCmd <- MapCmd
				fmt.Println("zig15Parsing 8E---> 0xC1C1")
			} else if zig14Ident == 0xC2 && zig15Ident == 0xC1 {
				MapCmd["id"] = 72103
				MapCmd["data"] = sliceCmdZigDeep
				ChCmd <- MapCmd
				fmt.Println("zig15Parsing 8E---> 0xC2C1")
			} else {
				fmt.Println("zig15Parsing 8E---> not fetched")
			}
		case 0xFC:
			fmt.Println("zig13Parsing ---> 0xFC")
			if zig14Ident == 0x00 && zig15Ident == 0x00 {
				MapCmd["id"] = 72104
				MapCmd["data"] = sliceCmdZigDeep
				ChCmd <- MapCmd
				fmt.Println("zig15Parsing FC---> 0x0000")
			} else {
				fmt.Println("zig15Parsing FC---> not fetched")
			}
		case 0xFD:
			fmt.Println("zig13Parsing ---> 0xFD")
			if zig14Ident == 0x00 && zig15Ident == 0x00 {
				MapCmd["id"] = 72105
				MapCmd["data"] = sliceCmdZigDeep
				ChCmd <- MapCmd
				fmt.Println("zig15Parsing FD---> 0x0000")
			} else {
				fmt.Println("zig15Parsing FD---> not fetched")
			}
		case 0xFE:
			fmt.Println("zig13Parsing ---> 0xFE")
			if zig14Ident == 0x00 && zig15Ident == 0x00 {
				MapCmd["id"] = 72106
				MapCmd["data"] = sliceCmdZigDeep
				ChCmd <- MapCmd
				fmt.Println("zig15Parsing FE---> 0x0000")
			} else {
				fmt.Println("zig15Parsing FE---> not fetched")
			}
		default:
			fmt.Println("zig15Parsing 0A---> not fetched")
		}
	}


}
//RTU相关命令处理-deepMore
func parsingProcess0A(buff []int) {
	sliceCmdPrc := buff
	fmt.Println("sliceCmdZig---> 0x0A", sliceCmdPrc)

	if sliceCmdPrc[14] == 0x00 {
		prcIdent := sliceCmdPrc[15]

		switch prcIdent {
		case 0x30:
			MapCmd["id"] = 72006
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x30", prcIdent)
		case 0x31:
			MapCmd["id"] = 72007
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x31", prcIdent)
		case 0x32:
			MapCmd["id"] = 72008
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x30", prcIdent)
		case 0x62:
			MapCmd["id"] = 72008
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x31", prcIdent)
		case 0x63:
			MapCmd["id"] = 72009
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x30", prcIdent)
		case 0x71:
			MapCmd["id"] = 72010
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x31", prcIdent)
		case 0x72:
			MapCmd["id"] = 72011
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x30", prcIdent)
		case 0x81:
			MapCmd["id"] = 72012
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x31", prcIdent)
		case 0x82:
			MapCmd["id"] = 72013
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x30", prcIdent)
		case 0x88:
			MapCmd["id"] = 72014
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x31", prcIdent)
		case 0x89:
			MapCmd["id"] = 72015
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x30", prcIdent)
		case 0x8B:
			MapCmd["id"] = 72016
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("--> sliceCmdPrc00---> 0x31", prcIdent)
		case 0x8C:
			MapCmd["id"] = 72017
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x30", prcIdent)
		case 0x92:
			MapCmd["id"] = 72018
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x31", prcIdent)
		case 0x95:
			MapCmd["id"] = 72019
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x30", prcIdent)
		case 0x96:
			MapCmd["id"] = 72020
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x31", prcIdent)
		case 0x97:
			MapCmd["id"] = 72021
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x30", prcIdent)
		case 0x98:
			MapCmd["id"] = 72022
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x31", prcIdent)
		case 0x9A:
			MapCmd["id"] = 72023
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("--> sliceCmdPrc00---> 0x30", prcIdent)
		case 0x9B:
			MapCmd["id"] = 72024
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("---> sliceCmdPrc00---> 0x31", prcIdent)
		case 0x9C:
			MapCmd["id"] = 72025
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0x9D:
			MapCmd["id"] = 72026
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0x9F:
			MapCmd["id"] = 72027
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0xA0:
			MapCmd["id"] = 72028
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0xA1:
			MapCmd["id"] = 72029
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0xA2:
			MapCmd["id"] = 72030
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0xA3:
			MapCmd["id"] = 72031
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0xA4:
			MapCmd["id"] = 72032
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0xA5:
			MapCmd["id"] = 72033
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0xA6:
			MapCmd["id"] = 72034
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0xF7:
			MapCmd["id"] = 72035
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0xF8:
			MapCmd["id"] = 72036
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0xF9:
			MapCmd["id"] = 72037
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0xFA:
			MapCmd["id"] = 72038
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0xFB:
			MapCmd["id"] = 72039
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0xFC:
			MapCmd["id"] = 72040
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0xFD:
			MapCmd["id"] = 72041
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0xFE:
			MapCmd["id"] = 72042
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0xFF:
			MapCmd["id"] = 72043
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		default:
			fmt.Println("sliceCmdPrc00---> not fetched.", prcIdent)

		}
	} else if sliceCmdPrc[14] == 0xC0 {
		prcIdent := sliceCmdPrc[15]

		switch prcIdent {
		case 0xB4:
			MapCmd["id"] = 72107
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrcC0---> 0x30", prcIdent)
		case 0xB5:
			MapCmd["id"] = 72108
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrcC0---> 0x31", prcIdent)
		case 0xC1:
			MapCmd["id"] = 72109
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrcC0---> 0x30", prcIdent)
		case 0xC2:
			MapCmd["id"] = 72110
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrcC0---> 0x31", prcIdent)
		default:
			fmt.Println("sliceCmdPrcC0---> not fetched.", prcIdent)

		}
	} else if sliceCmdPrc[14] >= 0xC1 && sliceCmdPrc[14] <= 0xC8 {
		prcIdent := sliceCmdPrc[15]

		switch prcIdent {
		case 0x91:
			MapCmd["id"] = 72111
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrcC0---> 0x30", prcIdent)
		case 0x94:
			MapCmd["id"] = 72112
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrcC0---> 0x31", prcIdent)
		case 0xC1:
			MapCmd["id"] = 72113
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrcC0---> 0x30", prcIdent)
		case 0xC2:
			MapCmd["id"] = 72114
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrcC0---> 0x31", prcIdent)
		default:
			fmt.Println("sliceCmdPrcC1-C8---> not fetched.", prcIdent)

		}
	} else if sliceCmdPrc[14] == 0xCE {
		if sliceCmdPrc[15] ==0xB2 {
			MapCmd["id"] = 72115
			MapCmd["data"] = sliceCmdPrc
			ChCmd <- MapCmd
			fmt.Println("sliceCmdPrcCE---> 0x31", sliceCmdPrc[15])
		} else {
			fmt.Println("sliceCmdPrcCE---> not fetched.", sliceCmdPrc[15])
		}
	} else {
		fmt.Println("sliceCmdPrc0A---> not fetched.")
	}

}

func parsingProcess8A(buff []int) {
	sliceCmdPrc := buff
	fmt.Println("sliceCmdZig---> 0x8A", sliceCmdPrc)

	if sliceCmdPrc[14] == 0xC0 && sliceCmdPrc[15] == 0xC1 {
		MapCmd["id"] = 72044
		MapCmd["data"] = sliceCmdPrc
		ChCmd <- MapCmd
		fmt.Println("sliceCmdZig---> 0x8A-C0C1")
	} else if sliceCmdPrc[14] == 0xC0 && sliceCmdPrc[15] == 0xC2 {
		MapCmd["id"] = 72045
		MapCmd["data"] = sliceCmdPrc
		ChCmd <- MapCmd
		fmt.Println("sliceCmdZig---> 0x8A-C0C2")
	} else {
		fmt.Println("sliceCmdPrc8A---> not fetched.")
	}
}

func parsingProcess8B(buff []int) {
	sliceCmdPrc := buff
	fmt.Println("sliceCmdZig---> 0x8B", sliceCmdPrc)

	if sliceCmdPrc[14] == 0xC0 && sliceCmdPrc[15] == 0xC1 {
		MapCmd["id"] = 72046
		MapCmd["data"] = sliceCmdPrc
		ChCmd <- MapCmd
		fmt.Println("sliceCmdZig---> 0x8B-C0C1")
	} else {
		fmt.Println("sliceCmdPrc8B---> not fetched.")
	}
}

func parsingProcess8C(buff []int) {
	sliceCmdPrc := buff
	fmt.Println("sliceCmdZig---> 0x8C", sliceCmdPrc)

	if sliceCmdPrc[14] == 0xC0 && sliceCmdPrc[15] == 0xC1 {
		MapCmd["id"] = 72047
		MapCmd["data"] = sliceCmdPrc
		ChCmd <- MapCmd
		fmt.Println("sliceCmdZig---> 0x8C-C0C1")
	} else {
		fmt.Println("sliceCmdPrc8C---> not fetched.")
	}
}

func parsingProcess8D(buff []int) {
	sliceCmdPrc := buff
	fmt.Println("sliceCmdZig---> 0x8D", sliceCmdPrc)

	if sliceCmdPrc[14] == 0xC0 && sliceCmdPrc[15] == 0xC1 {
		MapCmd["id"] = 72048
		MapCmd["data"] = sliceCmdPrc
		ChCmd <- MapCmd
		fmt.Println("sliceCmdZig---> 0x8D-C0C1")
	} else {
		fmt.Println("sliceCmdPrc8D---> not fetched.")
	}
}

func parsingProcess8E(buff []int) {
	sliceCmdPrc := buff
	fmt.Println("sliceCmdZig---> 0x8E", sliceCmdPrc)

	if sliceCmdPrc[14] == 0xC0 && sliceCmdPrc[15] == 0xC1 {
		MapCmd["id"] = 72049
		MapCmd["data"] = sliceCmdPrc
		ChCmd <- MapCmd
		fmt.Println("sliceCmdZig---> 0x8E-C0C1")
	} else {
		fmt.Println("sliceCmdPrc8E---> not fetched.")
	}
}

func parsingProcess8F(buff []int) {
	sliceCmdPrc := buff
	fmt.Println("sliceCmdZig---> 0x8F", sliceCmdPrc)

	if sliceCmdPrc[14] == 0xC0 && sliceCmdPrc[15] == 0xC1 {
		MapCmd["id"] = 72050
		MapCmd["data"] = sliceCmdPrc
		ChCmd <- MapCmd
		fmt.Println("sliceCmdZig---> 0x8F-C0C1")
	} else {
		fmt.Println("sliceCmdPrc8F---> not fetched.")
	}
}

func parsingProcess90(buff []int) {
	sliceCmdPrc := buff
	fmt.Println("sliceCmdZig---> 0x90", sliceCmdPrc)

	if sliceCmdPrc[14] == 0xC0 && sliceCmdPrc[15] == 0xC1 {
		MapCmd["id"] = 72051
		MapCmd["data"] = sliceCmdPrc
		ChCmd <- MapCmd
		fmt.Println("sliceCmdZig---> 0x90-C0C1")
	} else {
		fmt.Println("sliceCmdPrc90---> not fetched.")
	}
}

func parsingProcess93(buff []int) {
	sliceCmdPrc := buff
	fmt.Println("sliceCmdZig---> 0x93", sliceCmdPrc)
	if sliceCmdPrc[14] == 0xC0 {
		fmt.Println("sliceCmdZig---> 0x93-C0")
		prcIdent := sliceCmdPrc[15]
		switch prcIdent {
		case 0xC1:
			fmt.Println("sliceCmdPrcC1---> 0x30", prcIdent)
		case 0xC2:
			fmt.Println("sliceCmdPrcC2---> 0x31", prcIdent)
		case 0xC3:
			fmt.Println("sliceCmdPrcC3---> 0x30", prcIdent)
		case 0xC4:
			fmt.Println("sliceCmdPrcC4---> 0x31", prcIdent)
		case 0xC5:
			fmt.Println("sliceCmdPrcC5---> 0x31", prcIdent)
		default:
			fmt.Println("sliceCmdPrc93---> not fetched.", prcIdent)
		}	
	} else {
		fmt.Println("sliceCmdPrc93---> not fetched.")
	}
}

func parsingProcess70(buff []int) {
	sliceCmdPrc := buff
	fmt.Println("sliceCmdZig---> 0x70", sliceCmdPrc)

	if sliceCmdPrc[14] == 0xC0 && sliceCmdPrc[15] >= 0xC1 && sliceCmdPrc[15] <= 0xC4{
		MapCmd["id"] = 72052
		MapCmd["data"] = sliceCmdPrc
		ChCmd <- MapCmd
		fmt.Println("sliceCmdZig70---> C0C1")
	} else {
		fmt.Println("sliceCmdPrc70---> not fetched.")
	}
}

