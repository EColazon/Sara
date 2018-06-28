package handleCmdParsing

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
							fmt.Println("6002 is here.")
						case 0x6004:
							fmt.Println("6004 is here.")
						case 0x6005:
							fmt.Println("6005 is here.")
						case 0x6006:
							fmt.Println("6006 is here.")
						case 0x6007:
							fmt.Println("6007 is here.")
						case 0x6019:
							fmt.Println("6019 is here.")
						case 0x6020:
							fmt.Println("6020 is here.")
						case 0x6021:
							fmt.Println("6021 is here.")
						case 0x2600:
							fmt.Println("2600 is here.")
						case 0x2400:
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
						fmt.Println("TO DO ---> 0xA2.")
					
					case 0xA3:
						// 返回程序版本号
						fmt.Println("TO DO ---> 0xA3.")

					case 0xA4:
						// 复位设备-RTU
						fmt.Println("TO DO ---> 0xA4.")
					
					case 0xAD:
						// 复位设备-协调器
						fmt.Println("TO DO ---> 0xAD.")
					
					case 0xA5:
						// 读取eeprom数据-RTU
						fmt.Println("TO DO ---> 0xA5.")
					
					case 0xA6:
						// RTU供电方式查询
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
					fmt.Println("zig15Parsing ---> 0x0020")
				case 0x92:
					fmt.Println("zig15Parsing ---> 0x0092")
				case 0x94:
					fmt.Println("zig15Parsing ---> 0x0094")
				case 0x95:
					fmt.Println("zig15Parsing ---> 0x0095")
				case 0x96:
					fmt.Println("zig15Parsing ---> 0x0096")
				case 0x98:
					fmt.Println("zig15Parsing ---> 0x0098")
				case 0x99:
					fmt.Println("zig15Parsing ---> 0x0099")
				case 0x9A:
					fmt.Println("zig15Parsing ---> 0x009A")
				case 0x9D:
					fmt.Println("zig15Parsing ---> 0x009D")
				case 0x9E:
					fmt.Println("zig15Parsing ---> 0x009E")
				case 0x9F:
					fmt.Println("zig15Parsing ---> 0x009F")
				case 0xAC:
					fmt.Println("zig15Parsing ---> 0x00AC")
				case 0xAD:
					fmt.Println("zig15Parsing ---> 0x00AD")
				case 0xAE:
					fmt.Println("zig15Parsing ---> 0x00AE")
				case 0xAF:
					fmt.Println("zig15Parsing ---> 0x00AF")
				case 0xB1:
					fmt.Println("zig15Parsing ---> 0x00B1")
				case 0xB2:
					fmt.Println("zig15Parsing ---> 0x00B2")
				case 0xB3:
					fmt.Println("zig15Parsing ---> 0x00B3")
				case 0xB4:
					fmt.Println("zig15Parsing ---> 0x00B4")
				case 0xB5:
					fmt.Println("zig15Parsing ---> 0x00B5")
				case 0xB6:
					fmt.Println("zig15Parsing ---> 0x00B6")
				case 0xB8:
					fmt.Println("zig15Parsing ---> 0x00B8")
				case 0xB9:
					fmt.Println("zig15Parsing ---> 0x00B9")
				case 0xAA:
					fmt.Println("zig15Parsing ---> 0x00AA")
				case 0xAB:
					fmt.Println("zig15Parsing ---> 0x00AB")
				case 0xFA:
					fmt.Println("zig15Parsing ---> 0x00FA")
				case 0xFB:
					fmt.Println("zig15Parsing ---> 0x00FB")
				case 0xFC:
					fmt.Println("zig15Parsing ---> 0x00FC")
				case 0xBA:
					fmt.Println("zig15Parsing ---> 0x00BA")
				case 0xBB:
					fmt.Println("zig15Parsing ---> 0x00BB")
				case 0xBC:
					fmt.Println("zig15Parsing ---> 0x00BC")
				case 0xBD:
					fmt.Println("zig15Parsing ---> 0x00BD")
				case 0xBE:
					fmt.Println("zig15Parsing ---> 0x00BE")
				case 0xBF:
					fmt.Println("zig15Parsing ---> 0x00BF")
				case 0x80:
					fmt.Println("zig15Parsing ---> 0x0080")
				default:
					fmt.Println("zig15Parsing 00---> not fetched")
			
				}
			case 0xC0:
				fmt.Println("zig14Parsing ---> 0xC0")
				switch zig15Ident {
				case 0xB4:
					fmt.Println("zig15Parsing ---> 0xC0B4")
				case 0xB6:
					fmt.Println("zig15Parsing ---> 0xC0B6")
				case 0xC1:
					fmt.Println("zig15Parsing ---> 0xC0C1")
				case 0xC2:
					fmt.Println("zig15Parsing ---> 0xC0C2")
				case 0xBE:
					fmt.Println("zig15Parsing ---> 0xC0BE")
				case 0xBF:
					fmt.Println("zig15Parsing ---> 0xC0BF")
				default:
					fmt.Println("zig15Parsing C0---> not fetched")
				}
			case 0xC1:
				fmt.Println("zig14Parsing ---> 0xC1")
				switch zig15Ident {
				case 0xC1:
					fmt.Println("zig15Parsing ---> 0xC1C1")
				case 0xC2:
					fmt.Println("zig15Parsing ---> 0xC1C2")
				default:
					fmt.Println("zig15Parsing C1---> not fetched")
				}
			case 0xC2:
				fmt.Println("zig14Parsing ---> 0xC2")
				switch zig15Ident {
				case 0xC1:
					fmt.Println("zig15Parsing ---> 0xC2C1")
				case 0xC2:
					fmt.Println("zig15Parsing ---> 0xC2C2")
				default:
					fmt.Println("zig15Parsing C2---> not fetched")
				}
			case 0xE2:
				fmt.Println("zig14Parsing ---> 0xE2")
			case 0x01:
				fmt.Println("zig14Parsing ---> 0x01")
				if zig15Ident == 0x15 {
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
				fmt.Println("zig15Parsing 8E---> 0xC0C1")
			} else if zig14Ident == 0xC1 && zig15Ident == 0xC1 {
				fmt.Println("zig15Parsing 8E---> 0xC1C1")
			} else if zig14Ident == 0xC2 && zig15Ident == 0xC1 {
				fmt.Println("zig15Parsing 8E---> 0xC2C1")
			} else {
				fmt.Println("zig15Parsing 8E---> not fetched")
			}
		case 0xFC:
			fmt.Println("zig13Parsing ---> 0xFC")
			if zig14Ident == 0x00 && zig15Ident == 0x00 {
				fmt.Println("zig15Parsing FC---> 0x0000")
			} else {
				fmt.Println("zig15Parsing FC---> not fetched")
			}
		case 0xFD:
			fmt.Println("zig13Parsing ---> 0xFD")
			if zig14Ident == 0x00 && zig15Ident == 0x00 {
				fmt.Println("zig15Parsing FD---> 0x0000")
			} else {
				fmt.Println("zig15Parsing FD---> not fetched")
			}
		case 0xFE:
			fmt.Println("zig13Parsing ---> 0xFE")
			if zig14Ident == 0x00 && zig15Ident == 0x00 {
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
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0x31:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0x32:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0x62:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0x63:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0x71:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0x72:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0x81:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0x82:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0x88:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0x89:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0x8B:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0x8C:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0x92:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0x95:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0x96:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0x97:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0x98:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0x9A:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0x9B:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0x9C:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0x9D:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0x9F:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0xA0:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0xA1:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0xA2:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0xA3:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0xA4:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0xA5:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0xA6:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0xF7:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0xF8:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0xF9:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0xFA:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0xFB:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0xFC:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0xFD:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		case 0xFE:
			fmt.Println("sliceCmdPrc00---> 0x31", prcIdent)
		case 0xFF:
			fmt.Println("sliceCmdPrc00---> 0x30", prcIdent)
		default:
			fmt.Println("sliceCmdPrc00---> not fetched.", prcIdent)

		}
	} else if sliceCmdPrc[14] == 0xC0 {
		prcIdent := sliceCmdPrc[15]

		switch prcIdent {
		case 0xB4:
			fmt.Println("sliceCmdPrcC0---> 0x30", prcIdent)
		case 0xB5:
			fmt.Println("sliceCmdPrcC0---> 0x31", prcIdent)
		case 0xC1:
			fmt.Println("sliceCmdPrcC0---> 0x30", prcIdent)
		case 0xC2:
			fmt.Println("sliceCmdPrcC0---> 0x31", prcIdent)
		default:
			fmt.Println("sliceCmdPrcC0---> not fetched.", prcIdent)

		}
	} else if sliceCmdPrc[14] >= 0xC1 && sliceCmdPrc[14] <= 0xC8 {
		prcIdent := sliceCmdPrc[15]

		switch prcIdent {
		case 0x91:
			fmt.Println("sliceCmdPrcC0---> 0x30", prcIdent)
		case 0x94:
			fmt.Println("sliceCmdPrcC0---> 0x31", prcIdent)
		case 0xC1:
			fmt.Println("sliceCmdPrcC0---> 0x30", prcIdent)
		case 0xC2:
			fmt.Println("sliceCmdPrcC0---> 0x31", prcIdent)
		default:
			fmt.Println("sliceCmdPrcC1-C8---> not fetched.", prcIdent)

		}
	} else if sliceCmdPrc[14] == 0xCE {
		if sliceCmdPrc[15] ==0xB2 {
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
		fmt.Println("sliceCmdZig---> 0x8A-C0C1")
	} else if sliceCmdPrc[14] == 0xC0 && sliceCmdPrc[15] == 0xC2 {
		fmt.Println("sliceCmdZig---> 0x8A-C0C2")
	} else {
		fmt.Println("sliceCmdPrc8A---> not fetched.")
	}
}

func parsingProcess8B(buff []int) {
	sliceCmdPrc := buff
	fmt.Println("sliceCmdZig---> 0x8B", sliceCmdPrc)

	if sliceCmdPrc[14] == 0xC0 && sliceCmdPrc[15] == 0xC1 {
		fmt.Println("sliceCmdZig---> 0x8B-C0C1")
	} else {
		fmt.Println("sliceCmdPrc8B---> not fetched.")
	}
}

func parsingProcess8C(buff []int) {
	sliceCmdPrc := buff
	fmt.Println("sliceCmdZig---> 0x8C", sliceCmdPrc)

	if sliceCmdPrc[14] == 0xC0 && sliceCmdPrc[15] == 0xC1 {
		fmt.Println("sliceCmdZig---> 0x8C-C0C1")
	} else {
		fmt.Println("sliceCmdPrc8C---> not fetched.")
	}
}

func parsingProcess8D(buff []int) {
	sliceCmdPrc := buff
	fmt.Println("sliceCmdZig---> 0x8D", sliceCmdPrc)

	if sliceCmdPrc[14] == 0xC0 && sliceCmdPrc[15] == 0xC1 {
		fmt.Println("sliceCmdZig---> 0x8D-C0C1")
	} else {
		fmt.Println("sliceCmdPrc8D---> not fetched.")
	}
}

func parsingProcess8E(buff []int) {
	sliceCmdPrc := buff
	fmt.Println("sliceCmdZig---> 0x8E", sliceCmdPrc)

	if sliceCmdPrc[14] == 0xC0 && sliceCmdPrc[15] == 0xC1 {
		fmt.Println("sliceCmdZig---> 0x8E-C0C1")
	} else {
		fmt.Println("sliceCmdPrc8E---> not fetched.")
	}
}

func parsingProcess8F(buff []int) {
	sliceCmdPrc := buff
	fmt.Println("sliceCmdZig---> 0x8F", sliceCmdPrc)

	if sliceCmdPrc[14] == 0xC0 && sliceCmdPrc[15] == 0xC1 {
		fmt.Println("sliceCmdZig---> 0x8F-C0C1")
	} else {
		fmt.Println("sliceCmdPrc8F---> not fetched.")
	}
}

func parsingProcess90(buff []int) {
	sliceCmdPrc := buff
	fmt.Println("sliceCmdZig---> 0x90", sliceCmdPrc)

	if sliceCmdPrc[14] == 0xC0 && sliceCmdPrc[15] == 0xC1 {
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
		fmt.Println("sliceCmdZig70---> C0C1")
	} else {
		fmt.Println("sliceCmdPrc70---> not fetched.")
	}
}

