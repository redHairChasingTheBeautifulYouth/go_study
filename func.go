package main

import "fmt"

/**
我们知道，变量在内存中是存放于一定地址上的，修改变量实际是修改变量地址处的内存。
只有add1函数知道x变量所在的地址，才能修改x变量的值。所以我们需要将x所在地址&x传入函数，
并将函数的参数的类型由int改为*int，即改为指针类型，才能在函数中修改x变量的值。
此时参数仍然是按copy传递的，只是copy的是一个指针
*/
//func main(){
//	x := 3
//	fmt.Println("x=",x)
//	x1 := add(&x)
//	fmt.Println("x+1=",x1)
//	fmt.Println("x=",x)
//}
//func add (a *int) int {
//	*a = *a+1
//	return *a
//}

/**
Go语言中有种不错的设计，即延迟（defer）语句，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题。
*/
//func ReadWrite() bool {
//	file.Open("file")
//	defer file.Close()
//}

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd) // 函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)
}

/**
Panic和Recove
Go没有像Java那样的异常机制，它不能抛出异常，而是使用了panic和recover机制。一定要记住，你应当把它作为最后的手段来使用，也就是说，你的代码中应当没有，或者很少有panic的东西。这是个强大的工具，请明智地使用它。那么，我们应该如何使用它呢
是一个内建函数，可以中断原有的控制流程，进入一个panic状态中。当函数F调用panic，函数F的执行被中断，但是F中的延迟函数会正常执行，然后F返回到调用它的地方。在调用的地方，F的行为就像调用了panic。这一过程继续向上，直到发生panic的goroutine中所有调用的函数返回，此时程序退出。panic可以直接调用panic产生。也可以由运行时错误产生，例如访问越界的数组。
是一个内建的函数，可以让进入panic状态的goroutine恢复过来。recover仅在延迟函数中有效。在正常的执行过程中，调用recover会返回nil，并且没有其它任何效果。如果当前的goroutine陷入panic状态，调用recover可以捕获到panic的输入值，并且恢复正常的执行。
*/
