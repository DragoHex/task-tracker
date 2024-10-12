package main

import (
	"flag"
	"fmt"
	"os"
)

type DeleteCmd struct {
	*flag.FlagSet
	id int
}

func NewDeleteCmd() *DeleteCmd {
	cmd := &DeleteCmd{
		FlagSet: flag.NewFlagSet("delete", flag.ContinueOnError),
	}
	cmd.IntVar(&cmd.id, "id", -1, "id of the task to be deleted")
	return cmd
}

func (c *DeleteCmd) Execute() error {
	err := c.Parse(os.Args[2:])
	if err != nil {
		return fmt.Errorf("error in parsing flags: %s", err)
	}
	idx := -1
	for i := range TaskDataIns.Tasks {
		if TaskDataIns.Tasks[i].Id == c.id {
			idx = i
			break
		}
	}
	if idx < 0 {
		return fmt.Errorf("error in deleting the task: the task with ID: %d doesn't exist", c.id)
	}

	fmt.Printf("deleting the task with ID: %d\n", c.id)
	fmt.Println(TaskDataIns.Tasks[idx].Description)

	TaskDataIns.Tasks = append(TaskDataIns.Tasks[:idx], TaskDataIns.Tasks[idx+1:]...)
	err = TaskDataIns.Save()
	if err != nil {
		return fmt.Errorf("error in deleting the task: %s", err)
	}
	return nil
}
