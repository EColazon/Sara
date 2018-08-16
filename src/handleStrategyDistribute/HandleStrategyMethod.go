package handleStrategyDistribute

import (
	"fmt"
	Shared "handleShared"
	Redis "handleRedis"
)

func handleTimeCompareRelay(numR, a, count, hour, minute, second, week int) {
	fmt.Println("---> handleTimeCompareRelay.")
	sON := make([]int, 3)
	sOFF := make([]int, 3)

	relayOpenAllBit := 0
	jsonRelayOpenAllBit := make(map[string]interface{})
	relayCloseAllBit := 0
	jsonRelayCloseAllBit := make(map[string]interface{})

	// 回路状态
	stateLoop := 0
	// jsonStateLoop := make(map[string]interface{})

	// 定时开关标志
	flagLight := 0
	jsonFlagLight := make(map[string]interface{})

	if week == 0 {
		week += 7
	}

	if count == 0 {
		sON[0] = timerRONSpecial[numR*3]
		sON[1] = timerRONSpecial[numR*3+1]
		sON[2] = timerRONSpecial[numR*3+2]
		sOFF[0] = timerROFFSpecial[numR*3]
		sOFF[1] = timerROFFSpecial[numR*3+1]
		sOFF[2] = timerROFFSpecial[numR*3+2]
	} else if count == 1 {
		sON[0] = Shared.WDTimerRON[numR*3]
		sON[1] = Shared.WDTimerRON[numR*3+1]
		sON[2] = Shared.WDTimerRON[numR*3+2]
		sOFF[0] = Shared.WDTimerROFF[numR*3]
		sOFF[1] = Shared.WDTimerROFF[numR*3+1]
		sOFF[2] = Shared.WDTimerROFF[numR*3+2]
	} else if count == 2 {
		sON[0] = timerRONHoliday[numR*3]
		sON[1] = timerRONHoliday[numR*3+1]
		sON[2] = timerRONHoliday[numR*3+2]
		sOFF[0] = timerROFFHoliday[numR*3]
		sOFF[1] = timerROFFHoliday[numR*3+1]
		sOFF[2] = timerROFFHoliday[numR*3+2]
	} else if count == 3 {
		sON[0] = Shared.WDTimerRONWeek[numR*3+24*(week-1)]
		sON[1] = Shared.WDTimerRONWeek[numR*3+1+24*(week-1)]
		sON[2] = Shared.WDTimerRONWeek[numR*3+2+24*(week-1)]
		sOFF[0] = Shared.WDTimerROFFWeek[numR*3+24*(week-1)]
		sOFF[1] = Shared.WDTimerROFFWeek[numR*3+1+24*(week-1)]
		sOFF[2] = Shared.WDTimerROFFWeek[numR*3+2+24*(week-1)]
	}

	// hahaha...
	sON[0] = (sON[0]+(0x1F&Shared.WDAddrBoard[7]))%60
	sOFF[0] = (sOFF[0]+(0x1F&Shared.WDAddrBoard[7]))%60

	if sON[0] == sOFF[0] && sON[1] == sOFF[1] && sON[2] == sOFF[2] {
		fmt.Println("---> son==sff")
	} else {
		if a == 1 {
			if hour == sON[2] && minute == sON[1] && second == sON[0] {
				// 继电器开-改变标志-记录标志
				Shared.HandlerelayOpen(numR >> 1)
				flagLight |= (numR >> 1)
				jsonFlagLight[Shared.WDFlagLight] = flagLight
				Redis.HandleRedisJsonInsert(Shared.WDFlagLight, jsonFlagLight)
				fmt.Println("---> relayOpen001.")
			} else {
				if hour == sOFF[2] && minute == sOFF[1] && second == sOFF[0] {
					// 继电器开-改变标志-记录标志
					Shared.HandlerelayClose(numR >> 1)
					flagLight &= ^(numR >> 1)
					jsonFlagLight[Shared.WDFlagLight] = flagLight
					Redis.HandleRedisJsonInsert(Shared.WDFlagLight, jsonFlagLight)
					fmt.Println("---> relayClose001.")
				}
			}

			countNow := hour * 3600 + minute * 60 + second
			countOpen := sON[2] * 3600 + sON[1] * 60 + sON[0]
			countClose := sOFF[2] * 3600 + sOFF[1] * 60 + sOFF[0]

			// 获取RelayOpenAllBit
			relayOpenAllBitBase := Redis.HandleRedisJsonGet(Shared.WDRelayOpenAllBit)
			relayOpenAllBit = relayOpenAllBitBase.(int)
			relayCloseAllBitBase := Redis.HandleRedisJsonGet(Shared.WDRelayCloseAllBit)
			relayCloseAllBit = relayCloseAllBitBase.(int)
			stateLoopBase := Redis.HandleRedisJsonGet(Shared.WDStateLoop)
			stateLoop = stateLoopBase.(int)
			
			if countNow >= countOpen && countNow - countOpen < 120 {
				// 开灯前两分钟,清手动开关标志
				if relayOpenAllBit&(numR >> 1) > 0 {
					relayOpenAllBit &= ^(numR >> 1)
					// 存储relayOpenAllBit
					jsonRelayOpenAllBit[Shared.WDRelayOpenAllBit] = relayOpenAllBit
					Redis.HandleRedisJsonInsert(Shared.WDRelayOpenAllBit, jsonRelayOpenAllBit)
				}
				
				if relayCloseAllBit&(numR >> 1) > 0 {
					relayCloseAllBit &= ^(numR >> 1)
					// 存储relayCloseAllBit
					jsonRelayCloseAllBit[Shared.WDRelayCloseAllBit] = relayCloseAllBit
					Redis.HandleRedisJsonInsert(Shared.WDRelayCloseAllBit, jsonRelayCloseAllBit)
				}
			}

			if countNow >= countClose && countNow - countClose < 120 {
				// 关灯前两分钟,清手动开关标志
				if relayOpenAllBit&(numR >> 1) > 0 {
					relayOpenAllBit &= ^(numR >> 1)
					// 存储relayOpenAllBit
					jsonRelayOpenAllBit[Shared.WDRelayOpenAllBit] = relayOpenAllBit
					Redis.HandleRedisJsonInsert(Shared.WDRelayOpenAllBit, jsonRelayOpenAllBit)
				}
				
				if relayCloseAllBit&(numR >> 1) > 0 {
					relayCloseAllBit &= ^(numR >> 1)
					// 存储relayCloseAllBit
					jsonRelayCloseAllBit[Shared.WDRelayCloseAllBit] = relayCloseAllBit
					Redis.HandleRedisJsonInsert(Shared.WDRelayCloseAllBit, jsonRelayCloseAllBit)
				}
			}

			// 开灯比关灯晚
			if countOpen > countClose {
				if countNow < countOpen && countNow >= countClose {
					// 关灯时间
					if relayOpenAllBit & (numR >> 1) != 0|| (stateLoop & (numR >> 1)) == 0 {
						fmt.Println("---> timeCloseLamp.")
					} else {
						Shared.HandlerelayClose(numR >> 1)
					}
					flagLightBase := Redis.HandleRedisJsonGet(Shared.WDFlagLight)
					flagLight = flagLightBase.(int)
					if flagLight & (numR >> 1) > 0 {
						flagLight &= ^(numR >> 1)
						jsonFlagLight[Shared.WDFlagLight] = flagLight
						Redis.HandleRedisJsonInsert(Shared.WDFlagLight, jsonFlagLight)
					}

				} else {
					// 开灯时间
					if relayCloseAllBit & (numR >> 1) != 0 || (stateLoop & (numR >> 1)) != 0 {
						// 手动开过或者继电器状态是对的
						fmt.Println("---> timeOpenLamp.")
					} else {
						Shared.HandlerelayOpen(numR >> 1)
					}

					if stateLoop & (numR >> 1) == 0 {
						stateLoop |= (numR >> 1)
						jsonFlagLight[Shared.WDFlagLight] = flagLight
						Redis.HandleRedisJsonInsert(Shared.WDFlagLight, jsonFlagLight)
					}
				}
			} else if countOpen < countClose {
				// 开灯比关灯早
				if countNow >= countOpen && countNow < countClose {
					// 开灯时间
					if relayCloseAllBit & (numR >> 1) == 1 || stateLoop & (numR >> 1) == 1 {
						// 手动关过或者继电器是对的
						fmt.Println("---> timeOpenLamp01")
					} else {
						Shared.HandlerelayOpen(numR >> 1)
					}

					if flagLight & (numR >> 1) == 0 {
						flagLight |= (numR >> 1)
						jsonFlagLight[Shared.WDFlagLight] = flagLight
						Redis.HandleRedisJsonInsert(Shared.WDFlagLight, jsonFlagLight)
					
					}
				} else {
					// 关灯时间
					if relayOpenAllBit & (numR >> 1) == 1 || stateLoop & (numR >> 1) == 0 {
						// 手动开过或者继电器是对的
						fmt.Println("---> timeCloseLamp01")
					} else {
						Shared.HandlerelayClose(numR >> 1)
					}

					if flagLight & (numR >> 1) > 0 {
						flagLight &= ^(numR >> 1)
						jsonFlagLight[Shared.WDFlagLight] = flagLight
						Redis.HandleRedisJsonInsert(Shared.WDFlagLight, jsonFlagLight)
					}
				}
			}
		}
	}
}

func handleStrategySpecial(num, hour, minute, second, week int) {
	fmt.Println("---> handleStrategySpecial.")

	if timerRONSpecial[num*3] == timerROFFSpecial[num*3] && timerRONSpecial[num*3+1] == timerROFFSpecial[num*3+1] && timerRONSpecial[num*3+2] == timerROFFSpecial[num*3+2] {
		fmt.Println("---> handleStrategySpecial-DoNothing.")
	} else {
		handleTimeCompareRelay(num, 1, 0, hour, minute, second, week)
		handleTimeCompareRelay(num, 0, 0, hour, minute, second, week)
	}
}

func handleStrategyHoliday(num, hour, minute, second, week int) {
	fmt.Println("---> handleStrategyHoliday.")
	if timerRONHoliday[num*3] == timerROFFHoliday[num*3] && timerRONHoliday[num*3+1] == timerROFFHoliday[num*3+1] && timerRONHoliday[num*3+2] == timerROFFHoliday[num*3+2] {
		fmt.Println("---> handleStrategyHoliday-DoNothing.")
	} else {
		handleTimeCompareRelay(num, 1, 2, hour, minute, second, week)
		handleTimeCompareRelay(num, 0, 2, hour, minute, second, week)
	}
}

func handleGetTimeWeekONOFF(num, week int) {
	if week == 0 {
		week += 7
	}

	timerOffsetON[0] = Shared.WDTimerRONWeek[num*21+3*(week-1)]
	timerOffsetON[1] = Shared.WDTimerRONWeek[num*21+3*(week-1)+1]
	timerOffsetON[2] = Shared.WDTimerRONWeek[num*21+3*(week-1)+2]
	timerOffsetOFF[0] = Shared.WDTimerROFFWeek[num*21+3*(week-1)]
	timerOffsetOFF[1] = Shared.WDTimerROFFWeek[num*21+3*(week-1)+1]
	timerOffsetOFF[2] = Shared.WDTimerROFFWeek[num*21+3*(week-1)+2]
}

func handleGetTimeLatiLongiONOFF(num int) {
	latiLongiON := make([]int, 3)
	latiLongiOFF := make([]int, 3)

	flagTempON := 0
	flagTempOFF := 0
	tempT := 0

	latiLongiON[0] = Shared.WDTimerLatiLongiON[0]
	latiLongiON[1] = Shared.WDTimerLatiLongiON[1]
	latiLongiON[2] = Shared.WDTimerLatiLongiON[2]
	latiLongiOFF[0] = Shared.WDTimerLatiLongiOFF[0]
	latiLongiOFF[1] = Shared.WDTimerLatiLongiOFF[1]
	latiLongiOFF[2] = Shared.WDTimerLatiLongiOFF[2]

	// typeLatiLongi := 0
	// typeLatiLongiBase := Redis.HandleRedisJsonGet(Shared.WDTypeLatiLongi)
	// typeLatiLongi = typeLatiLongiBase.(int)
	// if typeLatiLongi != 0x55

	// // 经纬度开关正负偏移标志
	flagROffsetON := 0
	flagROffsetONBase := Redis.HandleRedisJsonGet(Shared.WDFlagROffsetON)
	flagROffsetON = flagROffsetONBase.(int)
	flagROffsetOFF := 0
	flagROffsetOFFBase := Redis.HandleRedisJsonGet(Shared.WDFlagROffsetOFF)
	flagROffsetOFF = flagROffsetOFFBase.(int)


	if flagROffsetON & (num >> 1) > 0 {
		flagTempON = 1
	} else {
		flagTempON = 0
	}

	if flagROffsetOFF & (num >> 1) >= 0 {
		flagTempOFF = 1
	} else {
		flagTempOFF = 0
	}

	if flagTempON == 1 {
		Shared.WDTimerLatiLongiON[0] += Shared.WDTimeLatiLongiROffsetON[num]%60
		if Shared.WDTimerLatiLongiON[0] >= 60 {
			Shared.WDTimerLatiLongiON[1] += 1
			Shared.WDTimerLatiLongiON[0] -= 60
			
			if Shared.WDTimerLatiLongiON[1] >= 60 {
				Shared.WDTimerLatiLongiON[2] += 1
				Shared.WDTimerLatiLongiON[1] -= 60
				
				
				if Shared.WDTimerLatiLongiON[2] >= 60 {
					Shared.WDTimerLatiLongiON[2] = 0
				}
			}
		}
		
		Shared.WDTimerLatiLongiON[1] += (Shared.WDTimerLatiLongiON[num]%3600) / 60
		if Shared.WDTimerLatiLongiON[1] >= 60 {
			Shared.WDTimerLatiLongiON[2] += 1
			Shared.WDTimerLatiLongiON[1] -= 60
			if Shared.WDTimerLatiLongiON[2] >= 24 {
				Shared.WDTimerLatiLongiON[2] = 0
			}
		}
		Shared.WDTimerLatiLongiON[2] += (Shared.WDTimerLatiLongiON[num]/3600) / 60
		if Shared.WDTimerLatiLongiON[2] >= 24 {
			Shared.WDTimerLatiLongiON[2] = -24
		}
		
	} else {
		tempT = Shared.WDTimeLatiLongiROffsetON[num] % 60
		if Shared.WDTimerLatiLongiON[0] < tempT {
			Shared.WDTimerLatiLongiON[0] += 60 - tempT
			Shared.WDTimerLatiLongiON[1] -= 1
			if Shared.WDTimerLatiLongiON[1] < 0 {
				Shared.WDTimerLatiLongiON[1] = 59
				Shared.WDTimerLatiLongiON[2] -= 1
				if Shared.WDTimerLatiLongiON[2] < 0 {
					Shared.WDTimerLatiLongiON[2] = 23
				}
			}
		} else {
			Shared.WDTimerLatiLongiON[0] -= tempT
		}

		tempT = Shared.WDTimeLatiLongiROffsetON[num] % 3600 /60

		if Shared.WDTimerLatiLongiON[1] < tempT {
			Shared.WDTimerLatiLongiON[1] += 60 - tempT
			Shared.WDTimerLatiLongiON[2] -= 1

			if Shared.WDTimerLatiLongiON[2] < 0 {
				Shared.WDTimerLatiLongiON[2] = 23
			}

		} else {
			Shared.WDTimerLatiLongiON[1] -= tempT
		}

		tempT = Shared.WDTimeLatiLongiROffsetON[num] / 3600
		if Shared.WDTimerLatiLongiON[2] < tempT {
			Shared.WDTimerLatiLongiON[2] += 24 - tempT
		} else {
			Shared.WDTimerLatiLongiON[2] -= tempT
		}
	}

	// 
	if flagTempOFF == 1 {
		Shared.WDTimerLatiLongiOFF[0] += Shared.WDTimeLatiLongiROffsetOFF[num]%60
		if Shared.WDTimerLatiLongiOFF[0] >= 60 {
			Shared.WDTimerLatiLongiOFF[1] += 1
			Shared.WDTimerLatiLongiOFF[0] -= 60
			
			if Shared.WDTimerLatiLongiOFF[1] >= 60 {
				Shared.WDTimerLatiLongiOFF[2] += 1
				Shared.WDTimerLatiLongiOFF[1] -= 60
				
				
				if Shared.WDTimerLatiLongiOFF[2] >= 60 {
					Shared.WDTimerLatiLongiOFF[2] = 0
				}
			}
		}
		
		Shared.WDTimerLatiLongiOFF[1] += (Shared.WDTimerLatiLongiOFF[num]%3600) / 60
		if Shared.WDTimerLatiLongiOFF[1] >= 60 {
			Shared.WDTimerLatiLongiOFF[2] += 1
			Shared.WDTimerLatiLongiOFF[1] -= 60
			if Shared.WDTimerLatiLongiOFF[2] >= 24 {
				Shared.WDTimerLatiLongiOFF[2] = 0
			}
		}
		Shared.WDTimerLatiLongiOFF[2] += (Shared.WDTimerLatiLongiOFF[num]/3600) / 60
		if Shared.WDTimerLatiLongiOFF[2] >= 24 {
			Shared.WDTimerLatiLongiOFF[2] = -24
		}
		
	} else {
		tempT = Shared.WDTimeLatiLongiROffsetOFF[num] % 60
		if Shared.WDTimerLatiLongiOFF[0] < tempT {
			Shared.WDTimerLatiLongiOFF[0] += 60 - tempT
			Shared.WDTimerLatiLongiOFF[1] -= 1
			if Shared.WDTimerLatiLongiOFF[1] < 0 {
				Shared.WDTimerLatiLongiOFF[1] = 59
				Shared.WDTimerLatiLongiOFF[2] -= 1
				if Shared.WDTimerLatiLongiOFF[2] < 0 {
					Shared.WDTimerLatiLongiOFF[2] = 23
				}
			}
		} else {
			Shared.WDTimerLatiLongiOFF[0] -= tempT
		}

		tempT = Shared.WDTimeLatiLongiROffsetOFF[num] % 3600 /60

		if Shared.WDTimerLatiLongiOFF[1] < tempT {
			Shared.WDTimerLatiLongiOFF[1] += 60 - tempT
			Shared.WDTimerLatiLongiOFF[2] -= 1

			if Shared.WDTimerLatiLongiOFF[2] < 0 {
				Shared.WDTimerLatiLongiOFF[2] = 23
			}

		} else {
			Shared.WDTimerLatiLongiOFF[1] -= tempT
		}

		tempT = Shared.WDTimeLatiLongiROffsetOFF[num] / 3600
		if Shared.WDTimerLatiLongiOFF[2] < tempT {
			Shared.WDTimerLatiLongiOFF[2] += 24 - tempT
		} else {
			Shared.WDTimerLatiLongiOFF[2] -= tempT
		}
	}


	timerOffsetON[0] = Shared.WDTimerLatiLongiON[0]
	timerOffsetON[1] = Shared.WDTimerLatiLongiON[1]
	timerOffsetON[2] = Shared.WDTimerLatiLongiON[1]
	timerOffsetOFF[0] = Shared.WDTimerLatiLongiOFF[0]
	timerOffsetOFF[1] = Shared.WDTimerLatiLongiOFF[1]
	timerOffsetOFF[2] = Shared.WDTimerLatiLongiOFF[2]

	Shared.WDTimerLatiLongiON[0] = latiLongiON[0]
	Shared.WDTimerLatiLongiON[1] = latiLongiON[1]
	Shared.WDTimerLatiLongiON[2] = latiLongiON[2]
	Shared.WDTimerLatiLongiOFF[0] = latiLongiOFF[0]
	Shared.WDTimerLatiLongiOFF[1] = latiLongiOFF[1]
	Shared.WDTimerLatiLongiOFF[2] = latiLongiOFF[2]
}