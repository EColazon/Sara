package chgoroutines

import (
	//"chgoroutines"
	//. "chgoroutines"
	"fmt"
	"time"
	
)

func SelectMain() {
	/*
	ch1 := make(chan int)
	ch2 := make(chan int)
	//ch := make(chan string)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)
	fmt.Println("Done---> ")

	

	time.Sleep(1e9)
	fmt.Println("END---> 01")
	*/
	ChMain()
	//go SendData(ch)
	//go GetData(ch)
	
	PrintSomething()
	time.Sleep(1e9)
	fmt.Println("END---> 02")
}

func pump1(ch chan int)  {

	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	 
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {

	for {
		select {
		case v := <- ch1:
		fmt.Printf("Received on channel 1 ---> %d\n", v)
		case v := <- ch2:
		fmt.Printf("Received on channel 2 ---> %d\n", v)
		}
	}
}