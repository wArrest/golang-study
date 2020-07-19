package struct_funciton_interface

import "math"

//这种定义 interface 的方式与大部分其他编程语言不同。通常接口定义需要这样的代码 My type Foo implements interface Bar。
//但是在我们的例子里，
//Rectangle 有一个返回值类型为 float64 的方法 Area，所以它满足接口 Shape
//Circle 有一个返回值类型为 float64 的方法 Area，所以它满足接口 Shape
//string 没有这种方法，所以它不满足这个接口
//等等
//声明接口，这样我们可以定义适合不同参数类型的函数（参数多态）
type Shape interface {
	Area() float64
}

type Rectangle struct {
	X float64
	Y float64
}
type Circle struct {
	Radius float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.X + r.Y) * 2
}

func (r Rectangle) Area() float64 {
	return r.X * r.Y
}

func (r Circle) Area() float64 {
	return math.Pi * r.Radius * r.Radius
}

//请注意我们的辅助函数是怎样实现不需要关心参数是矩形，圆形还是三角形的。通过声明一个接口，辅助函数能从具体类型解耦而只关心方法本身需要做的工作

