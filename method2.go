package main

import "fmt"

/**
那么你也会发现Go的一个神奇之处，method也是可以继承的。如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method。让我们来看下面这个例子
*/

type Human2 struct {
	name  string
	age   int
	phone string
}

type Student2 struct {
	Human2 //匿名字段
	school string
}

type Employee2 struct {
	Human2  //匿名字段
	company string
}

//在human上面定义了一个method
func (h *Human2) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
	mark := Student2{Human2{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee2{Human2{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}
