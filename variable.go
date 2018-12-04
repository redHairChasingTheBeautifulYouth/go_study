package main

import (
	"errors"
	"fmt"
)

/**
https://github.com/astaxie/build-web-application-with-golang教程
*/

/**
大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。
大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。
*/

/**
定义一个名称为“variableName”，类型为"type"的变量
*/
var variableName float32 = 3.2

/**
定义多个变量
*/
var vname1, vname2, vname3 int = 1, 2, 3

/**
可以直接忽略类型声明
*/
var a1, a2, a3 = 2, "ss", 4

/**
常量声明
*/
const name = "name"

/**
Boolean
*/
var flag bool = false

/**
Go内置有一个error类型，专门用来处理错误信息，Go的package里面还专门有一个包errors来处理错误：
*/
var err = errors.New("emit macho dwarf: elf header corrupted")

func main() {

	/**
	只能用在函数内部
	定义三个变量，它们分别初始化为相应的值
	vname1为v1，vname2为v2，vname3为v3
	编译器会根据初始化的值自动推导出相应的类型
	*/
	b1, b2, b3 := "sss", 3.3, "ss"
	b1 = "sssss"
	b2 = 4.4
	b3 = "sssssss"

	/**
	数组声明
	*/
	var arr [10]int
	a := [3]int{1, 2, 3}

	/**
	slice动态数组,和声明array一样，只是少了长度,切片
	lice可以从一个数组或一个已经存在的slice中再次声明。slice通过array[i:j]来获取，其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i
	*/
	var fslice []byte

	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	fslice = ar[2:5]

	/**
	声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	*/
	var numbers map[string]int
	// 另一种map的声明方式
	numbers = make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3

	fmt.Println(len(numbers))

	// 初始化一个字典
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println(csharpRating)
	} else {
		fmt.Println("没有值")
	}
	// 删除key为C的元素
	delete(rating, "C")

	/**
	make、new操作
	内建函数new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，
	即一个*T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：
	new返回指针。
	make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，而不是*T。本质来讲，
	导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。例如，一个slice，
	是一个包含指向数据（内部array）的指针、长度和容量的三项描述符；在这些项目被初始化之前，
	slice为nil。对于slice、map和channel来说，make初始化了内部的数据结构，填充适当的值。
	make返回初始化后的（非零）值。
	*/
}

func changeTo(name string, sn string) (sex int8, gnder string) {
	return 1, "2"
}
