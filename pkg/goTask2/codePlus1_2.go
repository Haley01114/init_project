package goTask2

//指针
//题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。

func DoubleSlice(pointArr *[]int) []int {
	if pointArr == nil {
		return nil
	}
	for i := range *pointArr {
		(*pointArr)[i] *= 2
	}
	return *pointArr
}
