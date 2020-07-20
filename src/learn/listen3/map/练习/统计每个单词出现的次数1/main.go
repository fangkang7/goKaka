package main

import (
	"fmt"
	"strings"
)

/**
统计"how do you do"这个字符串中，每个单词出现的次数
*/
func staticCode(s string) map[string]int {
	// 创建一个map并且进行初始化
	var result map[string]int = make(map[string]int, 128)
	words := strings.Split(s, " ")
	for _, value := range words {
		// 不存在加1  存在将这个值返回
		count, ok := result[value]
		if !ok {
			result[value] = 1
		} else {
			result[value] = count + 1
		}
	}
	return result
}

func main() {
	a := "how do you do do do"
	code := staticCode(a)
	fmt.Println(code)
}
