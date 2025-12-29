package pkg

import (
	"fmt"
	"time"
)

// Channel
// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。

func sendOnly(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("	发送：", i)
	}
	close(ch)
}
func receiveOnly(ch <-chan int) {
	for value := range ch {
		fmt.Println("		接收：", value)
	}
}
func RunChannel() {
	ch := make(chan int)
	go sendOnly(ch)
	go receiveOnly(ch)
	time.Sleep(time.Second * 2)
}
