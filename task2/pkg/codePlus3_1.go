package pkg

import (
	"fmt"
	"math"
)

// 面向对象
// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。

// 定义：Shape接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义：矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// 实现Shape接口——矩形面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 实现Shape接口——矩形周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 定义：圆结构体
type Circle struct {
	Radius float64
}

// 实现Shape接口——圆面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 实现Shape接口——圆周长
func (c Circle) Perimeter() float64 {
	return math.Pi * (2 * c.Radius)
}

func CalMetod() {
	var shape Shape

	shape = Rectangle{Width: 6, Height: 4}
	fmt.Println("	矩形面积6*4=", shape.Area())
	fmt.Println("	矩形周长2*(6+4)=", shape.Perimeter())

	shape = Circle{Radius: 5}
	fmt.Println("	圆形面积pi*5*5=", shape.Area())
	fmt.Println("	圆形周长pi*2*5=", shape.Perimeter())
}
