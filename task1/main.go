package main

import (
	"fmt"

	task "github.com/Haley01114/init_project/task1/pkg"
)

func main() {
	fmt.Println("===== 任务1-1. 两数之和：", task.TwoSum([]int{2, 7, 11, 15}, 9))

	fmt.Println("===== 任务1-9. 回文数：", task.IsPalindrome(121))

	fmt.Println("===== 任务1-14. 最长公共前缀：", task.LongestCommonPrefix([]string{"flower", "flow", "flight"}))

	fmt.Println("===== 任务1-20. 有效的括号：", task.IsValid("()[]{}"))

	fmt.Println("===== 任务1-26. 删除有序数组中的重复项：", task.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	fmt.Println("===== 任务1-56. 合并区间：", task.Merge([][]int{{1, 4}, {2, 5}, {6, 8}, {7, 9}}))

	fmt.Println("===== 任务1-66. 加一：", task.PlusOne([]int{4, 3, 2, 1}))

	fmt.Println("===== 任务1-136. 只出现一次的数字：", task.SingleNumber([]int{2, 2, 1}))
}
