package lib

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	tasks "github.com/DragoHex/task-tracker/pkg/tasks"
)

var (
	DataFile    = filepath.Join("data", "data.json")
	TaskDataIns *tasks.TaskData
)

func init() {
	flag.Usage = func() {
		fmt.Println(`A cli tool for managing tasks.

Multiple tasks can be added.
Each can be assigned a satus to it.
Tasks can be listed and filtered as per their status.

Usage:
  task-tracker [command]

Available Commands:
  add              A command to add new tasks
  completion       Generate the autocompletion script for the specified shell
  delete           A command to delete existing tasks
  help             Help about any command
  list             A command to list existing tasks
  mark-done        A command to mark a task as done
  mark-in-progress A command to mark a task as In Progress
  mark-todo        A command to mark a task as done
  update           A command to update description for existing tasks

Flags:
  -h, --help     help for task-tracker
  -t, --toggle   Help message for toggle

Use "task-tracker [command] --help" for more information about a command.`)
	}

	initData()
}

func Root() {
	help := flag.Bool("help", false, "to get more info on the tool")

	addCmd := NewAddCommand()
	updateCmd := NewUpdateCommand()
	deleteCmd := NewDeleteCmd()
	listCmd := NewListCommand()
	markTodoCmd := NewMarkTodoCommand()
	markInProgressCmd := NewMarkInProgressCommand()
	markDoneCmd := NewMarkDoneCommand()

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
	case "list":
		err := listCmd.Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
	case "mark-todo":
		err := markTodoCmd.Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
	case "mark-in-progress":
		err := markInProgressCmd.Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
	case "mark-done":
		err := markDoneCmd.Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func initData() {
	TaskDataIns = tasks.NewTaskData(DataFile)
	fmt.Println(os.Getwd())

	file, err := os.OpenFile(DataFile, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("error opening or creating file:", err)
		return
	}
	file.Close()

	data, err := os.ReadFile(DataFile)
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
