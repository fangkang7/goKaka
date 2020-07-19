package main

import "fmt"

func sum(a [5]int) int {
	var sum int
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func sum1(a [5]int) int {
	var sum int
	for index, _ := range a {
		sum += a[index]
	}
	return sum
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	ints := sum(a)
	i := sum1(a)
	fmt.Println(ints)
	fmt.Println(i)
}
