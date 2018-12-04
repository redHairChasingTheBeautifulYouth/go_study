package main

import "fmt"

/**
如果human里面有一个字段叫做phone，而student也有一个字段叫做phone，那么该怎么办呢？
o里面很简单的解决了这个问题，最外层的优先访问，也就是当你通过student.phone访问的时候，是访问student里面的字段，而不是human里面的字段。

这样就允许我们去重载通过匿名字段继承的一些字段，当然如果我们想访问重载后对应匿名类型里面的字段，可以通过匿名字段名来访问。请看下面的例子
*/
type Human1 struct {
	name  string
	age   int
	phone string // Human类型拥有的字段
}

type Employee struct {
	Human1     // 匿名字段Human
	speciality string
	phone      string // 雇员的phone字段
}

func main() {
	Bob := Employee{Human1{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	// 如果我们要访问Human的phone字段
	fmt.Println("Bob's personal phone is:", Bob.Human1.phone)
}
