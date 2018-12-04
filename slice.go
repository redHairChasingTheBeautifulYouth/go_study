package main

import "fmt"

//func main()  {
//	a := [5]int{98,90,98}
//	var b []int = a[1:3]
//	fmt.Println(b)
//	//切片本身不包含任何数据。它仅仅是底层数组的一个上层表示。对切片进行的任何修改都将反映在底层数组中。
//	for i := range b{
//		b[i]++
//	}
//}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}

}

/**
可以看到切片包含长度、容量、以及一个指向首元素的指针。当将一个切片作为参数传递给一个函数时，虽然是值传递，
但是指针始终指向同一个数组。因此将切片作为参数传给函数时，函数对该切片的修改在函数外部也可以看到。
*/
/*func main() {

	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside

}*/

/**
内存优化
切片保留对底层数组的引用。只要切片存在于内存中，数组就不能被垃圾回收。这在内存管理方便可能是值得关注的。假设我们有一个非常大的数组，而我们只需要处理它的一小部分，
为此我们创建这个数组的一个切片，并处理这个切片。这里要注意的事情是，数组仍然存在于内存中，因为切片正在引用它。
*/
func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}
func main() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}
