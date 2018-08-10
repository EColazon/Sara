package handleRelayCheck

import (
	"fmt"
	"time"
	Shread "handleShared"
	Redis "handleRedis"
)

func HandleRelayCheckManage() {
	fmt.Println("---> HandleRelayCheckManage.")

	for {
		// 获取继电器状态改变值
		relayStateChange := Redis.HandleRedisJsonGet(Shread.WDStateChangeRelay)
		HandleDoRelay(relayStateChange)
		time.Sleep(1 * time.Second)
	}
}

func HandleDoRelay(state int) {
	kvJson := make(map[string]interface{})
	// 获取继电器状态改变标志
	relayStateChange := Redis.HandleRedisJsonGet(Shread.WDStateChangeRelay)
	if relayStateChange == 0x55 {
		relayStateChange = 0
		state = ~state
		// 存储继电器状态
		kvJson[Sheard.WDStateChangeRelay] = relayStateChange
		Redis.HandleRedisJsonInsert(Sheard.WDStateChangeRelay, kvJson)
		// 状态写入I2C-1
		Shread.HandleSharedExecCSoI2C1Write(Shread.WDMCPADDR20, Shread.WDMCP20PINMODE, 0x00)
		Shread.HandleSharedExecCSoI2C1Write(Shread.WDMCPADDR20, Shread.WDMCP20PINVALUE, state)

	}

}