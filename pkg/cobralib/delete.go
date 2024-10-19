package cobralib

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A command to delete existing tasks",
	Long: `A command to delete existing tasks.
Requires task 'id' as an input flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			fmt.Printf("failed deleting the task with: %s\n", err)
			return
		}

		idx := -1
		for i := range TaskDataIns.Tasks {
			if TaskDataIns.Tasks[i].Id == id {
				idx = i
				break
			}
		}
		if idx < 0 {
			fmt.Printf("error in deleting the task: the task with ID: %d doesn't exist\n", id)
		}

		fmt.Printf("deleting the task with ID: %d\n", id)
		fmt.Println(TaskDataIns.Tasks[idx].Description)

		TaskDataIns.Tasks = append(TaskDataIns.Tasks[:idx], TaskDataIns.Tasks[idx+1:]...)
		err = TaskDataIns.Save()
		if err != nil {
			fmt.Printf("error in deleting the task: %s", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Int("id", -1, "id of the task to be deleted")
	err := deleteCmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Printf("error in marking flag as required: %s\n", err)
		return
	}
}
