package main

import "fmt"

func sort(array [5]int) [5]int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func main() {
	array := [5]int{1, 8, 6, 3, 5}
	ints := sort(array)
	fmt.Println(ints)
}
