package lib

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type UpdateCmd struct {
	*flag.FlagSet
	id          int
	description string
}

func NewUpdateCommand() *UpdateCmd {
	cmd := &UpdateCmd{
		FlagSet: flag.NewFlagSet("update", flag.ContinueOnError),
	}
	cmd.IntVar(&cmd.id, "id", -1, "task id")
	cmd.StringVar(&cmd.description, "task", "", "task description")
	return cmd
}

func (c *UpdateCmd) Execute() error {
	err := c.Parse(os.Args[2:])
	if err != nil {
		return fmt.Errorf("error in parsing the flags: %s", err)
	}

	for i, task := range TaskDataIns.Tasks {
		if task.Id == c.id {
			TaskDataIns.Tasks[i].Description = c.description
			TaskDataIns.Tasks[i].UpdatedAt = time.Now()
			TaskDataIns.Tasks[i].Print()
			err = TaskDataIns.Save()
			if err != nil {
				return fmt.Errorf("error in updating task: %s", err)
			}
			return nil
		}
	}

	return fmt.Errorf("error in updating task: no task found with ID: %d", c.id)
}
