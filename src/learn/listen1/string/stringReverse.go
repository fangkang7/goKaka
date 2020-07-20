package main

import (
	"fmt"
	"strings"
)

func stringReverse() {
	var str = "Hello"
	// 字符串转字节
	var bytes []byte = []byte(str)
	for i := 0; i < len(str)/2; i++ {
		// 定义一个变量存放从后往前的值
		tmp := bytes[len(str)-i-1]
		// 从后往前的值跟从前往后的值调换
		bytes[len(str)-i-1] = bytes[i]
		// 从前往后的值跟从后往前的值进行调换
		bytes[i] = tmp
	}
	str = string(bytes)
	fmt.Println(str)
}

func stringReverse1() {
	var str = "hello"
	var bytes []byte = []byte(str)
	var build strings.Builder
	for i := 0; i < len(bytes); i++ {
		i2 := bytes[len(bytes)-i-1]
		build.WriteString(string(i2))
	}
	s3 := build.String()
	fmt.Println(s3)
}

func main() {
	stringReverse()
	stringReverse1()
}
