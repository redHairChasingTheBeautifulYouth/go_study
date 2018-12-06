package main

import "fmt"

func main() {
	myMap := make(map[string]int)
	if myMap == nil {
		fmt.Println("error")
	} else {
		fmt.Println(myMap)
	}

	youMap := map[string]int{
		"s":  12,
		"ss": 123,
	}
	fmt.Println(youMap)

	_, ok := youMap["sss"]

	if ok {
		fmt.Println("true")
	}
	//map是引用类型
	delete(youMap, "s")

}
