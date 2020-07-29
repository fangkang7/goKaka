package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
使用scan读取只能读取一个字符

使用bufio可以读取全部字符
*/
func case1() {
	var str string
	fmt.Scan(&str)
	fmt.Println(str)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter some input")
	readString, e := reader.ReadString('\n')
	if e == nil {
		fmt.Printf("The input was %s\n", readString)
	}
}

func main() {
	case1()
}
