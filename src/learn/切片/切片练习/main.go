package main

import (
	"fmt"
	"sort"
)

/**
练习一
这个练习检测对切片长度和容量的认识
*/
func case1() {
	var sa = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		sa = append(sa, i)
	}
	// [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	fmt.Println(sa)
}

/**
练习二
使用sort包来对切片进行排序
*/
func case2() {
	var section []int = []int{1, 5, 8, 7, 6}
	sort.Ints(section)
	fmt.Println(section)

	var a []string = []string{"a", "d", "b", "v"}
	sort.Strings(a)
	fmt.Println(a)
}

/**
打印出A到Z的字母
后边实现一次性打印出[A-Za-z]
*/
func all() {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf("%c", i)
	}
}

func main() {
	//case1()
	//case2()
	all()
}
