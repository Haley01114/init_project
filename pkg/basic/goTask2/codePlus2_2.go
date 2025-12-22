package goTask2

import (
	"fmt"
	"time"
)

// Goroutine
// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。

func task1() {
	time.Sleep(time.Second * 1)
	fmt.Println("	task1 执行完成！")
}

func task2() {
	time.Sleep(time.Second * 3)
	fmt.Println("	task2 执行完成！")
}

func task3() {
	time.Sleep(time.Second * 2)
	fmt.Println("	task3 执行完成！")
}

func RunTask() {
	go task1()
	go task2()
	go task3()
}
