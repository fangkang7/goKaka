package main

import "fmt"

func testFun(a int, b int) (c int, d int) {
	c = a + b
	d = a - b
	return
}

func check(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

/**
求1-100质数
*/
func prime() {
	for i := 2; i < 100; i++ {
		if check(i) == true {
			fmt.Printf("%d is prime", i)
			fmt.Println()
		}
	}
}

/**
传参求1-100质数
*/
func prime1(n int) {
	if n <= 1 {
		fmt.Printf("这是个啥")
	}
	for i := 2; i < 100; i++ {
		if n%i == 0 {
			fmt.Printf("%d is prime", n)
		}
	}
}

func check_shuixianhua(n int) bool {
	first := n % 10
	second := (n / 10) % 10
	third := (n / 100) % 10
	//fmt.Printf("n:%d first:%d second:%d third:%d", n, first, second, third)
	//fmt.Println()
	if first*first*first+second*second*second+third*third*third == n {
		return true
	}
	return false
}

/**
实现100-1000水仙花数
*/
func shuixianhua() {
	for i := 100; i <= 1000; i++ {
		if check_shuixianhua(i) == true {
			fmt.Printf("%d is shuixianhua", i)
			fmt.Println()
		}
	}
}

/**
统计一个字符串
*/
func statistic(str string) {
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Print(r[i])
		if r[i] >= 'a' {
			fmt.Println("是字母")
		}
	}
	fmt.Println(string(r))
}

func main() {
	//c, d := testFun(1, 2)
	//fmt.Println(c)
	//fmt.Println(d)
	//prime()
	//prime1(2)
	//shuixianhua()
	var str string
	fmt.Printf("请输出字符串")
	fmt.Scan(&str)
	statistic(str)
}
