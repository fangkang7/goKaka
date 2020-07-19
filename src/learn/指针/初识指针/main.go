package main

import "fmt"

/**
初识指针类型
*/
func case1() {
	var a int = 20
	b := &a
	//a type of 0xc00000a0b8
	fmt.Printf("a type of %p\n", &a)
	//b type of 0xc00000a0b8
	fmt.Printf("b type of %p\n", b)
	//20
	fmt.Println(*b)
}

func edit(a *[3]int) {
	(*a)[0] = 100
}

func case2() {
	var b [3]int = [3]int{1, 2, 3}
	edit(&b)
	fmt.Println(b)
}

func case3() {
	var a int = 1000
	var b *int = &a
	var c *int = b
	fmt.Println(&a, b, c)
	*b = 100
	fmt.Println(*b)
	fmt.Println(a, *b, *c)
}

func main() {
	//case1()
	//case2()
	case3()
}
