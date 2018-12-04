package main

import "fmt"

/**
	method重写
上面的例子中，如果Employee想要实现自己的SayHi,怎么办？简单，和匿名字段冲突一样的道理，我们可以在Employee上面定义一个method，重写了匿名字段的方法。请看下面的例子
*/
type Human3 struct {
	name  string
	age   int
	phone string
}

type Student3 struct {
	Human3 //匿名字段
	school string
}

type Employee3 struct {
	Human3  //匿名字段
	company string
}

//Human定义method
func (h *Human3) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Employee的method重写Human的method
func (e *Employee3) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func main() {
	mark := Student3{Human3{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee3{Human3{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}
