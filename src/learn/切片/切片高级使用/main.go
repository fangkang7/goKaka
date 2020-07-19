package main

import "fmt"

/**
把一个切片放到另一个切片里
*/
func case1() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	a = append(a, 3, 2, 1)
	a = append(a, b...)
	fmt.Println(a)
}

/**
利用切片实现数组所有元素之和
*/
func case2(array []int) int {
	var sum int
	for value := range array {
		sum += value
	}
	return sum
}

/**
测试切片是引用传递
*/

func test(a []int) {
	a[0] = 1000
}

func case3() {
	//a := [4]int{1, 2, 3, 4}
	//b := a[:]
	a := []int{1, 2, 3, 4, 5}
	test(a)
	// [1000 2 3 4]
	fmt.Println(a)
}

func main() {
	//case1()
	//array := [6]int{1, 2, 3, 4, 5, 6}
	//b := array[:]
	//sum := case2(b)
	//fmt.Println(sum)
	case3()
}
