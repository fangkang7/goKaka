package main

import (
	"fmt"
	"strings"
	"time"
)

/**
返回值为int 的闭包
*/
func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

/**
返回值为string的闭包
*/
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		// 监测字符串中是否存在.bmp .jpg
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

/**
双返回值的闭包
*/
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func forGo() {
	for i := 0; i < 5; i++ {
		go func(index int) {
			fmt.Println(index)
		}(i)
	}
	time.Sleep(time.Second)
}

func main() {

	tmp := add(10)
	fmt.Println(tmp(1), tmp(2))
	tmp1 := add(10)
	fmt.Println(tmp1(100), tmp1(200))

	func1 := makeSuffixFunc(".bmp")
	func2 := makeSuffixFunc(".jpg")
	fmt.Println(func1("test"))
	fmt.Println(func2("test"))

	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))
	fmt.Println(f1(7), f2(8))

	forGo()
}
