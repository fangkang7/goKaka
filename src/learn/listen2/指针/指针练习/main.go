package main

import "fmt"

/**
练习一:获取一个变量的地址
*/
func getAddr() {
	var a int
	a = 10
	fmt.Printf("a addr is %p\n", &a)
}

func case1(a *int) {
	*a = 100
}

func main() {
	getAddr()

	b := 10
	case1(&b)
	fmt.Println(b)
}
