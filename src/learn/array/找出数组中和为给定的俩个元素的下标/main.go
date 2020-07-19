package main

import "fmt"

/**
解法一
*/
func calc(array [5]int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == 8 {
				fmt.Printf("(%d,%d)", i, j)
			} else {
				continue
			}
		}
	}
}

/**
解法二
*/
func calc1(array [5]int, target int) {
	for i := 0; i < len(array); i++ {
		other := target - array[i]
		for j := i + 1; j < len(array); j++ {
			if other == array[j] {
				fmt.Printf("(%d,%d)", i, j)
			}
		}
	}
}

/**
比如 数组[1,3,5,8,7] 找出俩个元素之和等于8的下标  分别为(0,4) (1,2)
*/
func main() {
	var array [5]int = [5]int{1, 3, 5, 8, 7}
	calc(array)
	calc1(array, 8)
}
