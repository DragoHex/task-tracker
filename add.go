package main

import (
	"flag"
	"fmt"
	"os"
)

type AddCmd struct {
	*flag.FlagSet
}

func NewAddCommand() *AddCmd {
	return &AddCmd{
		flag.NewFlagSet("add", flag.ContinueOnError),
	}
}

func (a *AddCmd) AddLogic(des string) error {
	err := a.Parse(os.Args[2:])
	if err != nil {
		return fmt.Errorf("error in parsing the flags: %s", err)
	}

	task := NewTask()
	task.Description = des

	TaskDataIns.Add(task)
	err = TaskDataIns.Save()
	if err != nil {
		return fmt.Errorf("error in saving the data: %s", err)
	}
	task.Print()
	return nil
}
