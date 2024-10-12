package main

import (
	"flag"
	"fmt"
	"os"
)

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

	task := NewTask()
	task.Description = c.description
	task.Status = "todo"

	TaskDataIns.Add(task)
	err = TaskDataIns.Save()
	if err != nil {
		return fmt.Errorf("error in saving the data: %s", err)
	}
	task.Print()
	return nil
}
