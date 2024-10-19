package lib

import (
	"flag"
	"fmt"
	tasks "github.com/DragoHex/task-tracker/pkg/tasks"
	"os"
)

// AddCmd used to add new task to the data store
type AddCmd struct {
	*flag.FlagSet
	description string
}

func NewAddCommand() *AddCmd {
	cmd := &AddCmd{
		FlagSet: flag.NewFlagSet("add", flag.ContinueOnError),
	}
	cmd.StringVar(&cmd.description, "task", "", "task description")
	return cmd
}

func (c *AddCmd) Execute() error {
	err := c.Parse(os.Args[2:])
	if err != nil {
		return fmt.Errorf("error in parsing the flags: %s", err)
	}

	task := tasks.NewTask()
	task.Description = c.description

	TaskDataIns.Add(task)
	err = TaskDataIns.Save()
	if err != nil {
		return fmt.Errorf("error in saving the data: %s", err)
	}
	task.Print()
	return nil
}
