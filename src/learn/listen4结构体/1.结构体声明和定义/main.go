package main

import "fmt"

type Person struct {
	name string
	age  int
	like []string
}

/**
第一种定义方式
*/
func case1() {
	p := Person{}
	p.name = "kaka"
	p.age = 24
	p.like = []string{"php", "go"}
	//{kaka 24 [php go]}
	fmt.Println(p)
}

/**
第二种定义方式
*/
func case2() {
	var person Person = Person{
		name: "kaka",
		age:  24,
		like: []string{"php1", "go"},
	}
	//{kaka 24 [php1 go]}
	fmt.Println(person)
}

/**
第三种方式
*/
func case3() {
	fmt.Println(Person{})
}

func main() {
	case1()
	case2()
	case3()
}
