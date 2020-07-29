package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func case1() {
	fmt.Println(os.Args[0])
	if len(os.Args) > 1 {
		for index, value := range os.Args {
			if index == 0 {
				continue
			}
			/**
			args[1]=kaka
			args[2]=niuniu
			args[3]=boss
			*/
			fmt.Printf("args[%d]=%v\n", index, value)
		}
	}
}

func case2() {
	app := cli.NewApp()
	app.Name = "kaka"
	app.Usage = "命令行的参数的简单说明"
	app.Action = func(c *cli.Context) error {
		var cmd string
		if c.NArg() > 0 {
			cmd = c.Args().First()
		}
		fmt.Println("Hello friend cmd:=", cmd)
		return nil
	}
	app.Run(os.Args)
}

func main() {
	//case1()
	case2()
}
