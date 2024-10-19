package cobralib

import (
	"fmt"

	tasks "github.com/DragoHex/task-tracker/pkg/tasks"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A command to add new tasks",
	Long: `A command to add new tasks.
Requires one 'task' flag to add the description.`,
	Run: func(cmd *cobra.Command, args []string) {
		desc, err := cmd.Flags().GetString("task")
		if err != nil {
			fmt.Printf("error in adding task: %s\n", err)
			return
		}

		task := tasks.NewTask()
		task.Description = desc

		TaskDataIns.Add(task)
		err = TaskDataIns.Save()
		if err != nil {
			fmt.Printf("error in adding task: %s\n", err)
			return
		}
		task.Print()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().String("task", "", "task description")
	err := addCmd.MarkFlagRequired("task")
	if err != nil {
		fmt.Printf("error in marking flag as required: %s", err)
		return
	}
}
