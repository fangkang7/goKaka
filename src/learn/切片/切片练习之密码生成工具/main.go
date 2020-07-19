package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// 声明用户输入的密码长度和类型
var (
	length  int
	charset string
)

// 定义三种密码类型
const (
	Int     = "0123456789"
	String  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Advance = "$**^*^*#&%*!&*&&"
)

/**
参数解析
*/
func parseArgs() {
	flag.IntVar(&length, "l", 16, "-l 生成密码长度")
	flag.StringVar(&charset, "t", "num", "`t 制定密码生成的字符集 , num:只使用数据[0-9],chat:只使用英文字母[a-zA-Z],mix:使用数字和字母,advance:使用数字、字母以及特殊字符")
	flag.Parse()
}

/**
生成密码
*/
func generatePassword() string {
	// 定义密码切片
	var password []byte = make([]byte, length, length)
	var sourceStr string
	if charset == "num" {
		sourceStr = Int
	} else if charset == "chat" {
		sourceStr = String
	} else if charset == "mix" {
		sourceStr = Int + String
	} else if charset == "advance" {
		sourceStr = Int + String + Advance
	} else {
		sourceStr = Int
	}
	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		password[i] = sourceStr[index]
	}
	return string(password)
}

func main() {
	// 这个的含义
	rand.Seed(time.Now().UnixNano())
	parseArgs()
	fmt.Printf("length:%d  chatset:%s\n", length, charset)
	password := generatePassword()
	fmt.Println(password)
}
