package main

import "fmt"

type person struct {
	name string
	age  int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

//
//func main(){
//	//var tom person
//	tom := person{"tom" ,25}
//	var paul person
//	paul.name ,paul.age = "paul" ,256
//
//	tb_older ,tb_diff := Older(tom ,paul)
//	fmt.Printf("Of %s and %s, %s is older by %d years\n",
//		tom.name, tb_older, paul.name, tb_diff)
//
//}

/**
匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。
*/

type Student struct {
	person
	sex string
}

func main() {
	stu := Student{person{"ss", 23}, "ssss"}
	fmt.Println("His name is ", stu.name)
	//student还能访问person这个字段作为字段名
	stu.person.age = 1
}
