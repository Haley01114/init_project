package pkg

import (
	"fmt"
	"time"
)

// Goroutine
// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。

func method1(num int) {
	for i := 1; i <= num; i++ {
		if i%2 != 0 {
			fmt.Printf("	1到%d的奇数：%d\n", num, i)
		}
	}
}

func method2(num int) {
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			fmt.Printf("	2到%d的偶数：%d\n", num, i)
		}
	}
}

func RunMethod() {
	go method1(10)
	go method2(10)
	time.Sleep(time.Second * 1)
}
