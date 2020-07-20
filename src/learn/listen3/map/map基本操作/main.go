package main

import (
	"fmt"
	"sort"
)

/**
判断map中的key是否存在
*/
func case1() {
	var kakaMap map[string]int = map[string]int{
		"age":  24,
		"like": 35,
	}
	fmt.Println(kakaMap)
	_, ok := kakaMap["123"]
	if ok {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
}

/**
map的遍历以及删除

*/
func case2() {
	var kakaMap map[string]int = map[string]int{
		"age":   24,
		"money": -1024,
	}
	for index, value := range kakaMap {
		fmt.Println(index)
		fmt.Println(value)
	}
	delete(kakaMap, "age")
	fmt.Println(kakaMap)
}

/**
检测map是否为引用类型
*/
func case3() {
	a := map[string]int{
		"kaka1":  100,
		"niuniu": 200,
	}
	a["lala"] = 300
	// map[kaka1:100 lala:300 niuniu:200]
	fmt.Println(a)
	a["lala"] = 400
	// map[kaka1:100 lala:400 niuniu:200]
	fmt.Println(a)
	b := a
	b["lala"] = 800
	// map[kaka1:100 lala:800 niuniu:200]
	fmt.Println(b)
	// map[kaka1:100 lala:800 niuniu:200]
	fmt.Println(a)
}

/**
排序
*/
func case4() {
	var a map[string]int = make(map[string]int, 10)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key:%d", i)
		a[key] = i
	}
	keys := make([]string, 0, 128)
	for key, _ := range a {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, value := range keys {
		//fmt.Printf("keys:%s  val:%d\n", value, a[value])
		fmt.Println(value, a[value])
	}
	//fmt.Println(a)
}

/**
切片类型map
*/
func case5() {
	var mapSlice []map[string]int
	mapSlice = make([]map[string]int, 5, 16)
	fmt.Println("before map init")
	for index, value := range mapSlice {
		fmt.Printf("slice[%d]=%v\n", index, value)
	}
	//rand.Seed(time.Now().UnixNano())
	mapSlice[0] = make(map[string]int, 16)
	mapSlice[0]["age"] = 24
	fmt.Println(mapSlice)
}

/**
map里边包含切片
*/
func case6() {
	// 定义map类型切片
	var mapSlice map[string][]int
	// 对map进行初始化
	mapSlice = make(map[string][]int, 16)
	// 声明一个key
	key := "stu01"
	value, ok := mapSlice[key]
	// 检测map中是否存在这个切片
	if !ok {
		mapSlice[key] = make([]int, 0, 16)
		value = mapSlice[key]
	}
	value = append(value, 100)
	value = append(value, 200)
	value = append(value, 300)
	mapSlice[key] = value
	fmt.Println(mapSlice)
}

func main() {
	//case1()
	//case2()
	//case3()
	case4()
	//case5()
	//case6()
}
