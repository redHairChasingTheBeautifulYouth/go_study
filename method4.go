package main

import "fmt"

/**
可以定义非结构体类型的方法，不过这里需要注意一点。为了定义某个类型的方法，接收者类型的定义与方法的定义必须在同一个包中
*/
//使用别名,否则报错
type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}
