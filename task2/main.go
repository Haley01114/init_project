package main

import (
	"fmt"
	"time"

	task "github.com/Haley01114/init_project/task2/pkg"
)

func main() {

	num := 0
	fmt.Println("===== 任务2-1-1 将指针指向的值加10：", task.AddPointNum(&num))

	fmt.Println("===== 任务2-1-2 通过整数切片的指针，将切片中的每个元素乘以2：", task.DoubleSlice(&[]int{3, 5}))

	fmt.Println("===== 任务2-2-1 使用两个协程打印从1到10的奇数、偶数：")
	task.RunMethod()

	fmt.Println("===== 任务2-2-2 使用协程并发执行这些任务，同时统计每个任务的执行时间：")
	task.RunTask()
	time.Sleep(time.Second * 5)

	fmt.Println("===== 任务2-3-1 使用结构体实现接口，求取面积、周长：")
	task.CalMetod()

	fmt.Println("===== 任务2-3-2 使用嵌套结构体，输出学生信息：")
	task.RunPrintInfo()

	fmt.Println("===== 任务2-4-1 使用两个协程实现分别从Channel发送、接收数据：")
	task.RunChannel()

	fmt.Println("===== 任务2-4-2 使用两个协程实现分别从带有缓冲的Channel中发送、接收数据：")
	task.RunChannel2()

	fmt.Println("===== 任务2-5-1 使用10个协程操作一个被sync.Mutex保护的计数器：")
	task.RunCalMeThod()

	fmt.Println("===== 任务2-5-2 使用10个协程操作一个原子计数器：")
	task.RunCalAdd()
}
