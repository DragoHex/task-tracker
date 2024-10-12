package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var TaskDataIns *TaskData

func init() {
	flag.Usage = func() {
		fmt.Println(`task-traker (tt) is a cli tool to manage tasks
it can be used to manage list of tasks
along with their current status`)
		flag.PrintDefaults()
	}

	initData()
}

func root() {
	help := flag.Bool("help", false, "to get more info on the tool")

	addCmd := NewAddCommand()
	updateCmd := NewUpdateCommand()
	deleteCmd := NewDeleteCmd()

	flag.Parse()

	if *help {
		flag.Usage()
	}

	switch os.Args[1] {
	case "add":
		err := addCmd.Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
	case "update":
		err := updateCmd.Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
	case "delete":
		err := deleteCmd.Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func initData() {
	TaskDataIns = NewTaskData()

	file, err := os.OpenFile("data.json", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error opening or creating file:", err)
		return
	}
	file.Close()

	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("error reading from the file:", err)
		return
	}

	// if there is no data in the file
	if data == nil {
		fmt.Println("Created a new json data file")
		return
	}

	// if data exist then import that to the var
	err = json.Unmarshal(data, TaskDataIns)
	if err != nil {
		fmt.Println("error unmarshalling data:", err)
		return
	}
}
