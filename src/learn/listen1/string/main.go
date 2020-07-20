package main

import (
	"fmt"
)

func add(a int,b int) (int,int) {
	return a + b,a - b
}

const (
	a = iota
	b
	c
	d = 8
	e
	f = iota
	g
)

func main() {
	i, i2 := add(3, 2)
	fmt.Println(i)
	fmt.Println(i2)
	//a := 3
	fmt.Println(a)

	fmt.Println("常量输出的值")
	fmt.Println(a,b,c,d,e,f,g)
}
