package goTask2

import "fmt"

// 面向对象
// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Println("员工编号：", e.EmployeeID)
	fmt.Println("员工姓名：", e.Name)
	fmt.Println("员工年龄：", e.Age)
}

func RunPrintInfo() {
	employee := Employee{Person{"张三", 18}, "111111"}
	employee.PrintInfo()
}
