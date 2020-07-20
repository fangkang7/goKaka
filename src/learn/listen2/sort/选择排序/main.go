package main

import "fmt"

func sort(array [5]int) [5]int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
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
