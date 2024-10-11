package main

import (
	"flag"
	"fmt"
)

func init() {
	flag.Usage = func() {
		fmt.Println(`task-traker (tt) is a cli tool to manage tasks
it can be used to manage list of tasks
along with their current status`)
		flag.PrintDefaults()
	}
}

func main() {
	help := flag.Bool("help", false, "to get more info on the tool")
	flag.Parse()

	if help != nil {
		flag.Usage()
	}
}
