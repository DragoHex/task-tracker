package lib

import (
	"flag"
	"fmt"
	tasks "github.com/DragoHex/task-tracker/pkg/tasks"
	"os"
	"text/tabwriter"
)

type ListCmd struct {
	*flag.FlagSet
	Status string
}

func NewListCommand() *ListCmd {
	cmd := &ListCmd{
		FlagSet: flag.NewFlagSet("list", flag.ContinueOnError),
	}
	cmd.StringVar(&cmd.Status, "status", "", "status for list to be filtered with")
	return cmd
}

func (c *ListCmd) Execute() error {
	err := c.Parse(os.Args[2:])
	if err != nil {
		return fmt.Errorf("error in parsing the flags: %s", err)
	}

	if c.Status != "" {
		if c.Status != tasks.ToDo.String() && c.Status != tasks.InProgress.String() &&
			c.Status != tasks.Done.String() {
			return fmt.Errorf(
				"error in listing task: %s",
				"invalid status value passed | only 'todo', 'in-progress' & 'done' are supported",
			)
		}
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "ID\tTask\tStatus")
	for _, task := range TaskDataIns.Tasks {
		if c.Status != "" {
			if task.Status.String() == c.Status {
				fmt.Fprintf(w, "%d\t%q\t%s\n", task.Id, task.Description, task.Status.String())
			}
		} else {
			fmt.Fprintf(w, "%d\t%q\t%s\n", task.Id, task.Description, task.Status.String())
		}
	}
	w.Flush()
	return nil
}
