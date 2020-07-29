package main

import (
	"fmt"
	"os"
)

func case1() {
	var bytes [16]byte
	os.Stdin.Read(bytes[:])
	os.Stdout.WriteString(string(bytes[:]))
}

func case2() {
	a := 1 % 95
	b := 2 % 96
	c := 3 % 97
	d := 4 % 98
	e := 5 % 99
	f := 6 % 100
	g := 7 % 101
	h := 8 % 102
	i := 9 % 103
	j := 10 % 104
	k := 11 % 105
	l := 12 % 106
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}

func main() {
	//case1()
	case2()
}
