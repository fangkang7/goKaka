package main

import "fmt"

type person struct {
	name string
	age  int
}

func case1() {
	var user *person
	fmt.Println(user)

	var user0 *person = &person{}
	user0.name = "kaka"
	user0.age = 24
	fmt.Println(user0)
}

func main() {
	case1()
}
