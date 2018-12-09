package main

import "fmt"

func main() {
	name := "Hello World"
	printBytes(name)
	for index, rune := range name {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

/**
访问字符串中的字节
*/
func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printBytes1(s string) {
	for index, s := range s {
		fmt.Printf("%c starts at byte %d\n", s, index)
	}
}
