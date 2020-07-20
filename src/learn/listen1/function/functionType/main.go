package main

import "fmt"

/**
当有俩个返回值的时候，接受需要俩个变量接收
*/
func add(a int, b int) (sum int, reduce int) {
	sum = a + b
	reduce = +a - b
	return
}

func checkAddType() {
	f1 := add
	// 检测函数类型
	fmt.Printf("type if f1=%T", f1)
	// 函数有俩个返回值  所以需要俩个变量去接收
	sum1, sum2 := f1(100, 200)
	fmt.Println(sum1, sum2)
}

/**
匿名函数
*/
func anonymous() {
	f1 := func(a, b int) int {
		return a + b
	}
	fmt.Printf("func type of %T", f1)
	fmt.Println(f1(100, 200))
}

/**
defer 跟函数的使用
*/
func deferAddFunc() {
	var a int = 100
	defer fmt.Printf("defer a=%d", a)
	a = 1000
	fmt.Println(a)
}

/**
defer 跟匿名函数  闭包
*/
func deferAddAnonymous() int {
	var a int = 100
	defer func() {
		fmt.Printf("defer a=%d", a)
	}()
	a = 1000
	return a
}

/**
检测多个defer的运行顺序
检测结果  多个defer时，最后的defer先执行
*/
func deferTest() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

func test(a, b int) int {
	return a * b
}

/**
函数参数
*/
func funcParam(a, b int, kaka func(int, int) int) int {
	return kaka(a, b)
}

func main() {
	// 检测函数类型
	//checkAddType()
	// 匿名函数
	//anonymous()
	// defer的使用
	//deferAddFunc()
	//deferAddAnonymous()
	//deferTest()

	// 函数参数
	param := funcParam(2, 3, test)
	fmt.Println(param)
}
