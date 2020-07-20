package main

import "fmt"

type person struct {
	name string
	age  int
	string
	int
}

type user struct {
	like     string
	userinfo person
}

/**
字段匿名:匿名字段是指用类型作为索引
*/
func case1() {
	person := person{
		name:   "kaka",
		age:    24,
		string: "liu",
		int:    23,
	}
	fmt.Println(person)
}

/**
结构体嵌套
*/
func case2() {
	u := user{
		userinfo: person{
			name:   "kaka",
			age:    24,
			string: "liu",
			int:    23,
		},
		like: "code",
	}
	fmt.Println(u)
	fmt.Printf("user=%#v\n", u)
	fmt.Printf("user=%+v\n", u)
	fmt.Printf("user=%v\n", u)
	fmt.Printf("user=%T\n", u)
	fmt.Printf("user=%p\n", &u)
}

func main() {
	case1()
	case2()
}
