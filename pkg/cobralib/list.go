package cobralib

import (
	"fmt"
	"os"
	"text/tabwriter"

	tasks "github.com/DragoHex/task-tracker/pkg/tasks"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A command to list existing tasks",
	Long: `A command to list existing tasks.
It also excepts filters based on the 'status' flag value.
Supported status values are 'todo', 'in-progress' & 'done'`,
	Run: func(cmd *cobra.Command, args []string) {
		status, err := cmd.Flags().GetString("status")
		if err != nil {
			fmt.Printf("error in listing task: %s\n", err)
			return
		}

		if status != "" {
			if status != tasks.ToDo.String() && status != tasks.InProgress.String() &&
				status != tasks.Done.String() {
				fmt.Printf("error in listing tasks: %s",
					"invalid status value passed | only 'todo', 'in-progress' & 'done' are supported\n")
				return
			}
		}

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		fmt.Fprintln(w, "ID\tTask\tStatus")
		for _, task := range TaskDataIns.Tasks {
			if status != "" {
				if task.Status.String() == status {
					fmt.Fprintf(w, "%d\t%q\t%s\n", task.Id, task.Description, task.Status.String())
				}
			} else {
				fmt.Fprintf(w, "%d\t%q\t%s\n", task.Id, task.Description, task.Status.String())
			}
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().String("status", "", "status for list to be filtered with")
}
