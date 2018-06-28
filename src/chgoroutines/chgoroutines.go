package chgoroutines

import (
	"fmt"
	"time"
)

func ChMain() {

	ch := make(chan string)

	go SendData(ch)
	go GetData(ch)

	

	time.Sleep(1e9)
}

func SendData(ch chan string) {
	ch <- "aa1"
	ch <- "aa2"
	ch <- "aa3"
	ch <- "aa4"
	ch <- "aa5"
	ch <- "aa6"
}

func GetData(ch chan string) {
	var input string
	for {
		input = <- ch
		fmt.Printf("---> %s\n", input)
	}

}

func PrintSomething() {
	fmt.Println("I'm In PrintSomething---> ")
}