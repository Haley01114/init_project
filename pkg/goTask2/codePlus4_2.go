package goTask2

import (
	"fmt"
	"time"
)

// Channel
// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。

func send(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Println("	生产通道：", i)
	}
	close(ch)
}
func receive(ch <-chan int) {
	for value := range ch {
		fmt.Println("		消费通道：", value)
	}
}
func RunChannel2() {
	ch := make(chan int, 5)
	go send(ch)
	go receive(ch)
	time.Sleep(time.Second * 5)
}
