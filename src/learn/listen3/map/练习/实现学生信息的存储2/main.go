package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
)

func testInterface() {
	var a interface{}
	b := 100
	c := "kaka"
	a = b
	fmt.Println(a)
	a = c
	fmt.Println(a)
}

func studentData() {
	var stuMap map[int]map[string]interface{}
	stuMap = make(map[int]map[string]interface{}, 16)
	var id = 1
	var name = "ka"
	var score = 78.9
	var age = 24
	value, ok := stuMap[id]
	if !ok {
		value = make(map[string]interface{}, 8)
	}
	value["name"] = name
	value["score"] = score
	value["age"] = age
	stuMap[id] = value
	//fmt.Println(stuMap)

	for i := 0; i < 10; i++ {
		value, ok := stuMap[i]
		if !ok {
			value = make(map[string]interface{}, 16)
		}
		value["name"] = fmt.Sprintf("stu%d", i)
		value["score"] = rand.Float32() * 100.0
		value["age"] = rand.Intn(100)
		value["id"] = i
		stuMap[i] = value
	}
	fmt.Println()

	// 声明一个key用来保存stuMap的key值
	var keys []int = make([]int, 0)
	for index, _ := range stuMap {
		keys = append(keys, index)
	}
	// 把key重新进行排序
	sort.Ints(keys)
	for k, v := range keys {
		fmt.Printf("id=%d stu info=%#v\n", k, stuMap[v])
	}
	// 序列化
	bytes, _ := json.Marshal(stuMap)
	fmt.Println(string(bytes))
}

func main() {
	//testInterface()
	studentData()
}
