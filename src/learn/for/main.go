package main

import (
	"fmt"
	"time"
)

func test1() {
	var a int
	for {
		a++
		fmt.Println(a)
		time.Sleep(time.Millisecond * 2000)
		if a > 5 {
			return
		}
	}
}

func test2() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d", j, i, j*i)
		}
		fmt.Println()
	}
}

func main() {
	//test1()
	test2()
}
