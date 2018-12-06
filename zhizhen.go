package main

import "fmt"

/**
指向类型 T 的指针用 *T 表示,Go 不支持指针运算
*/
func main() {
	//& 操作符用来获取一个变量的地址。在上面的程序中，第 9 行我们将 b 的地址赋给 a（a 的类型为 *int）
	b := 344
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	//指针的空值为 nil
	var c *int
	fmt.Println("c is", c)

	//通过指针访问值,解引用指针的意思是通过指针访问被指向的值。指针 a 的解引用表示为：*a
	fmt.Println(a)
	fmt.Println(*a)

	//不要传递指向数组的指针给函数，而是使用切片
	d := [3]int{23, 34, 45}
	modify(d[:])

	//结构体是值类型
	t := Test{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", t.firstName)
	fmt.Println("Age:", t.age)

	//方法指针接收者和值接收者
	//以指针为接收者也是可以的。两者的区别在于，以指针作为接收者时，方法内部对其的修改对于调用者是可见的，但是以值作为接收者却不是。
	//
	//如果一个method的receiver是*T,你可以在一个T类型的实例变量V上面调用这个method，而不需要&V去调用这个method
	//
	//类似的
	//
	//如果一个method的receiver是T，你可以在一个*T类型的变量P上面调用这个method，而不需要 *P去调用这个method

	/**
	一般来讲，指针接收者可用于对接收者的修改应该对调用者可以见的场合。

	指针接收者也可用于拷贝结构体代价较大的场合。考虑一个包含较多字段的结构体，若使用值作为接收者则必须拷贝整个结构体，这样的代价很大。这种情况下使用指针接收者将避免结构体的拷贝，而仅仅是指向结构体指针的拷贝。

	其他情况下可以使用值接收者。
	*/
	//当一个方法有一个值接收者时，它可以接受值和指针接收者。
	//一个接收者为指针的方法可以接受值接收者和指针接收者。
}

func modify(sls []int) {
	sls[0] = 90
}

type Test struct {
	firstName, lastName string
	age, salary         int
}
