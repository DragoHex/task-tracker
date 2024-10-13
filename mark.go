package main

import (
	"flag"
	"fmt"
	"os"
)

type MarkTodoCmd struct {
	*flag.FlagSet
	id int
}

func NewMarkTodoCommand() *MarkTodoCmd {
	cmd := &MarkTodoCmd{
		FlagSet: flag.NewFlagSet("mark-todo", flag.ContinueOnError),
	}
	cmd.IntVar(&cmd.id, "id", 0, "id of the task to be updated")
	return cmd
}

func (c *MarkTodoCmd) Execute() error {
	err := c.Parse(os.Args[2:])
	if err != nil {
		return fmt.Errorf("error in parsing the flags: %s", err)
	}

	for i := range TaskDataIns.Tasks {
		if TaskDataIns.Tasks[i].Id == c.id {
			TaskDataIns.Tasks[i].Status = ToDo
			err = TaskDataIns.Save()
			if err != nil {
				return fmt.Errorf("error in updating task status: %s", err)
			}
			return nil
		}
	}
	return fmt.Errorf("error in updating task status: no task found with the ID(%d)", c.id)
}

type MarkInProgressCmd struct {
	*flag.FlagSet
	id int
}

func NewMarkInProgressCommand() *MarkInProgressCmd {
	cmd := &MarkInProgressCmd{
		FlagSet: flag.NewFlagSet("mark-in-progress", flag.ContinueOnError),
	}
	cmd.IntVar(&cmd.id, "id", 0, "id of the task to be updated")
	return cmd
}

func (c *MarkInProgressCmd) Execute() error {
	err := c.Parse(os.Args[2:])
	if err != nil {
		return fmt.Errorf("error in parsing the flags: %s", err)
	}

	for i := range TaskDataIns.Tasks {
		if TaskDataIns.Tasks[i].Id == c.id {
			TaskDataIns.Tasks[i].Status = InProgress
			err = TaskDataIns.Save()
			if err != nil {
				return fmt.Errorf("error in updating task status: %s", err)
			}
			return nil
		}
	}
	return fmt.Errorf("error in updating task status: no task found with the ID(%d)", c.id)
}

type MarkDoneCmd struct {
	*flag.FlagSet
	id int
}

func NewMarkDoneCommand() *MarkDoneCmd {
	cmd := &MarkDoneCmd{
		FlagSet: flag.NewFlagSet("mark-done", flag.ContinueOnError),
	}
	cmd.IntVar(&cmd.id, "id", 0, "id of the task to be updated")
	return cmd
}

func (c *MarkDoneCmd) Execute() error {
	err := c.Parse(os.Args[2:])
	if err != nil {
		return fmt.Errorf("error in parsing the flags: %s", err)
	}

	for i := range TaskDataIns.Tasks {
		if TaskDataIns.Tasks[i].Id == c.id {
			TaskDataIns.Tasks[i].Status = Done
			err = TaskDataIns.Save()
			if err != nil {
				return fmt.Errorf("error in updating task status: %s", err)
			}
			return nil
		}
	}
	return fmt.Errorf("error in updating task status: no task found with the ID(%d)", c.id)
}
