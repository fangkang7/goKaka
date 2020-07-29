package main

import "fmt"

/**
对于fmt.Print() 这类的演示就不写了  没有什么意义  在学习中多多使用就可以了
*/

func case1() {
	fmt.Print("请输入一个值")
	var a int
	var b int
	fmt.Scan(&a)
	c := a + b
	fmt.Print(c)
}

func case2() {
	var a int
	var b string
	var c float32
	//fmt.Scanf("%d %s %f", &a, &b, &c)
	fmt.Scan(&a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f", a, b, c)
}

func main() {
	//case1()
	case2()
}
