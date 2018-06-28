package timefunc

import (
	"fmt"
	"time"
)

func TimeFunc(f func()) {

	funcStart := time.Now()
	f()
	funcEnd := time.Now()
	funcDelta := funcEnd.Sub(funcStart)
	fmt.Printf("%p took this amount of time: %s\n",f, funcDelta)
}
