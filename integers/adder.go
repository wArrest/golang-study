package integers

//为测试的运行编写最少量的代码并检查失败测试的输出
//当你有多个相同类型的参数（在我们的例子中是两个整数），可以将它缩短为 (x，y int) 而不是 (x int, y int)。
//我们在 上一节 中学习了 命名返回值，但没有在这里使用。它通常应该在结果的含义在上下文中不明确时使用，在我们的例子中，Add 函数将参数相加是非常明显的。

//Add takes two integers and returns the sum of them
func Adder(x, y int) int {
	return x + y
}
