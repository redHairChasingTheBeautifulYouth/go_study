package main

import "fmt"

/**
数组是值类型不是引用类型，这意味着当数组变量被赋值时，将会获得原数组（译者注：也就是等号右面的数组）的拷贝。新数组中元素的改变不会影响原数组中元素的值。
*/

func main() {
	a := [...]string{"aa", "bb"}
	b := a
	fmt.Println(b)
}
