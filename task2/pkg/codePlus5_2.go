package pkg

import (
	"fmt"
	"sync/atomic"
	"time"
)

// 锁机制
// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

var count int64 = 0

func add() {
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&count, 1)
	}
}

func RunCalAdd() {
	for i := 1; i <= 10; i++ {
		go add()
	}
	time.Sleep(time.Second * 5)
	fmt.Println("期望结果：", 10000)
	fmt.Println("最后结果：", count)
}
