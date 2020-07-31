package main

import (
	"learn/logger"
	"time"
)

//func testTowParam() {
//	a := 1000 % 5
//	b := 1001 % 5
//	c := 1002 % 5
//	d := 1003 % 5
//	e := 1004 % 5
//	f := 1005 % 5
//	g := 1006 % 5
//	h := 1007 % 5
//	i := 1008 % 5
//	j := 1009 % 5
//	fmt.Println(a)
//	fmt.Println(b)
//	fmt.Println(c)
//	fmt.Println(d)
//	fmt.Println(e)
//	fmt.Println(f)
//	fmt.Println(g)
//	fmt.Println(h)
//	fmt.Println(i)
//	fmt.Println(j)
//}

var log logger.LogInterface

func initLogger(logPath, logName string, level int) {
	//log = logger.NewConsoleLogger(level)
	log = logger.NewFileLogger(level, logPath, logName)
	log.Debug("init logger success")
}

func Run() {
	for true {
		log.Debug("user server is running")
		time.Sleep(time.Second)
	}
}

func main() {
	//testTowParam()
	initLogger("D:/", "test", logger.LogLevelDebug)
	Run()
	return
}
