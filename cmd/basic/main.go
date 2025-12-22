package main

import (
	"fmt"
	"time"

	task1 "github.com/Haley01114/init_project/pkg/basic/goTask1"
	task2 "github.com/Haley01114/init_project/pkg/basic/goTask2"
)

func main() {
	fmt.Println("===== 任务1-1. 两数之和：", task1.TwoSum([]int{2, 7, 11, 15}, 9))

	fmt.Println("===== 任务1-9. 回文数：", task1.IsPalindrome(121))

	fmt.Println("===== 任务1-14. 最长公共前缀：", task1.LongestCommonPrefix([]string{"flower", "flow", "flight"}))

	fmt.Println("===== 任务1-20. 有效的括号：", task1.IsValid("()[]{}"))

	fmt.Println("===== 任务1-26. 删除有序数组中的重复项：", task1.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	fmt.Println("===== 任务1-56. 合并区间：", task1.Merge([][]int{{1, 4}, {2, 5}, {6, 8}, {7, 9}}))

	fmt.Println("===== 任务1-66. 加一：", task1.PlusOne([]int{4, 3, 2, 1}))

	fmt.Println("===== 任务1-136. 只出现一次的数字：", task1.SingleNumber([]int{2, 2, 1}))

	num := 0
	fmt.Println("===== 任务2-1-1 将指针指向的值加10：", task2.AddPointNum(&num))

	fmt.Println("===== 任务2-1-2 通过整数切片的指针，将切片中的每个元素乘以2：", task2.DoubleSlice(&[]int{3, 5}))

	fmt.Println("===== 任务2-2-1 使用两个协程打印从1到10的奇数、偶数：")
	task2.RunMethod()

	fmt.Println("===== 任务2-2-2 使用协程并发执行这些任务，同时统计每个任务的执行时间：")
	task2.RunTask()
	time.Sleep(time.Second * 5)

	fmt.Println("===== 任务2-3-1 使用结构体实现接口，求取面积、周长：")
	task2.CalMetod()

	fmt.Println("===== 任务2-3-2 使用嵌套结构体，输出学生信息：")
	task2.RunPrintInfo()

	fmt.Println("===== 任务2-4-1 使用两个协程实现分别从Channel发送、接收数据：")
	task2.RunChannel()

	fmt.Println("===== 任务2-4-2 使用两个协程实现分别从带有缓冲的Channel中发送、接收数据：")
	task2.RunChannel2()

	fmt.Println("===== 任务2-5-1 使用10个协程操作一个被sync.Mutex保护的计数器：")
	task2.RunCalMeThod()

	fmt.Println("===== 任务2-5-2 使用10个协程操作一个原子计数器：")
	task2.RunCalAdd()
}
