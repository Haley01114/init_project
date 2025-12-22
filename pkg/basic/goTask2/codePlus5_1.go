package goTask2

import (
	"fmt"
	"sync"
	"time"
)

// 锁机制
// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

var num int = 0
var mu sync.Mutex

func addMethod() {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		num = num + 1
		mu.Unlock()
	}
}

func RunCalMeThod() {
	for i := 1; i <= 10; i++ {
		go addMethod()
	}
	time.Sleep(time.Second * 5)
	fmt.Println("期望结果：", 10000)
	fmt.Println("最后结果：", num)
}
