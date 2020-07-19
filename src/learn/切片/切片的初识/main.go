package 切片的初识

import "fmt"

func case1() {
	var a []int
	if a == nil {
		fmt.Println("切片为空")
	} else {
		fmt.Println("切片不为空")
	}
}

/**
切片初始化
*/
func case2() {
	a := [5]int{1, 2, 3, 4, 5}
	var b []int
	b = a[1:4]
	fmt.Printf("%d\n", b)
	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(b[2])
}

/**
切片初始化
*/
func case3() {
	a := []int{1, 2, 3}
	fmt.Println(a)
}

/**
切片的修改
*/
func case4() {
	a := [...]int{1, 2, 3, 4, 5}
	b := a[2:5]
	fmt.Println(b)
	b[0] = b[0] + 1
	b[1] = b[1] + 1
	b[2] = b[2] + 1
	fmt.Println(b)
}

/**
同一个数组  切片不同也会互相影响值
*/
func case5() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[:]
	c := a[:]
	b[0] = 100
	c[0] = 200
	fmt.Println(b)
	fmt.Println(c)
}

/**
使用make动态定义切片
*/
func case6() {
	var a []int = make([]int, 2)
	//a := make([]int, 2)
	a[0] = 100
	a[1] = 100
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}

/**
定义切片的几种方式
*/
func case7() {
	//方式一
	var a []int = make([]int, 1)
	a[0] = 100
	fmt.Println(a)

	// 方式二
	b := make([]int, 1)
	b[0] = 100
	fmt.Println(b)

	// 方式三
	c := [4]int{1, 2, 3, 4}
	d := c[:]
	ints := append(d, 5)
	fmt.Println(ints)
}

/**
对切片的容量和长度认识
*/

func case8() {
	var a []int
	a = make([]int, 1, 10)
	a[0] = 10
	fmt.Printf("a=%v addr:%p len:%d  cap:%d\n", a, a, len(a), cap(a))
	a = append(a, 11)
	fmt.Printf("a=%v addr:%p len:%d  cap:%d\n", a, a, len(a), cap(a))
	for i := 0; i < 8; i++ {
		a = append(a, i)
		fmt.Printf("a=%v addr:%p len:%d  cap:%d\n", a, a, len(a), cap(a))
	}
	// 观察切片扩容的操作
	a = append(a, 1000)
	fmt.Printf("a=%v addr:%p len:%d  cap:%d\n", a, a, len(a), cap(a))
}

func main() {
	//case1()
	//case2()
	//case3()
	//case4()
	//case5()
	//case6()
	//case7()
	case8()
}
