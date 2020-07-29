package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Id   int
	Sex  string
}

type Class struct {
	Name     string
	Count    int
	Students []*Student
}

func testJson() {
	class := &Class{
		Name:  "101",
		Count: 0,
	}

	for i := 0; i <= 10; i++ {
		stu := &Student{
			Name: fmt.Sprintf("stu%d", i),
			Id:   i,
			Sex:  fmt.Sprintf("%d", i),
		}
		class.Students = append(class.Students, stu)
	}
	bytes, e := json.Marshal(class)
	if e != nil {
		fmt.Println("格式化失败")
	}
	fmt.Println(string(bytes))
}

/**
反序列化
*/
func testUnMarshal() {
	var jsonString = `{"Name":"101","Count":0,"Students":[{"Name":"stu0","Id":0,"Sex":"0"},{"Name":"stu1","Id":1,"Sex":"1"},{"Name":"stu2","Id":2,"Sex":"2"},{"Name":"stu3","Id":3,"Sex":"3"},{"Name":"stu4","Id":4,"Sex":"4"},{"Name":"stu5","Id":5,"Sex":"5"},{"Name":"stu6","Id":6,"Sex":"6"},{"Name":"stu7","Id":7,"Sex":"7"},{"Name":"stu8","Id":8,"Sex":"8"},{"Name":"stu9","Id":9,"Sex":"9"},{"Name":"stu10","Id":10,"Sex":"10"}]}
`
	cla := Class{}
	err := json.Unmarshal([]byte(jsonString), &cla)
	if err != nil {
		fmt.Println("反序列失败")
	}
	//fmt.Println(cla)
	fmt.Printf("cla:%#v\n", cla)

	for _, v := range cla.Students {
		fmt.Printf("stu%#v\n", v)
		fmt.Println(v)
	}
}

func main() {
	//testJson()
	testUnMarshal()
}
