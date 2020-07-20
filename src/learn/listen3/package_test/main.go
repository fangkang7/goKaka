package main

import (
	"fmt"
	custom2 "learn/listen3/custom"
)

/**
测试包引用
引用的包是custom
*/
func testPackage() {
	sum := custom2.TowSum(1, 2)
	fmt.Println(sum)
}

func main() {
	testPackage()
}
