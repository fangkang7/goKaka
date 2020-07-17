package main

import "fmt"

/**
闭包案例一：传过的来的参数就是d  x初始值为0
*/
func Adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func main() {
	f := Adder()
	fmt.Println(f(10))
	fmt.Println(f(20))
	fmt.Println(f(300))
	f = Adder()
	fmt.Println(f(1))
}
