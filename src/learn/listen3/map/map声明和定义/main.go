package main

import "fmt"

/**
定义方式一
*/
func case1() {
	kakaMap := make(map[string]string)
	kakaMap["age"] = "24"
	kakaMap["name"] = "kaka"
	fmt.Println(kakaMap)
}

/**
定义方式二
*/
func case2() {
	var kakaMap map[string]string
	kakaMap = make(map[string]string)
	kakaMap["name"] = "kaka"
	fmt.Println(kakaMap)
}

/**
定义方式三
*/
func case3() {
	kakaMap := make(map[string]string)
	kakaMap["name"] = "kaka"
	fmt.Println(kakaMap)
}

func main() {
	case1()
	case2()
	case3()
}
