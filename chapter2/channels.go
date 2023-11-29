package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func Sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}
func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Println("Before sleep")
	Sleep(3)
	fmt.Println("After sleep")

	//Буферизованный канал
	cbuf := make(chan int, 1)
	fmt.Scanln(&input)
}
