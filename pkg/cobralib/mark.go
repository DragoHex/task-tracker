package cobralib

import (
	"fmt"

	tasks "github.com/DragoHex/task-tracker/pkg/tasks"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(markDoneCmd)
	rootCmd.AddCommand(markInProgressCmd)
	rootCmd.AddCommand(markTodoCmd)
	markDoneCmd.Flags().Int("id", -1, "id of the task to be marked done")
	err := markDoneCmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Printf("error in marking flag as required: %s", err)
		return
	}
	markInProgressCmd.Flags().Int("id", -1, "id of the task to be marked in-progress")
	err = markInProgressCmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Printf("error in marking flag as required: %s", err)
		return
	}
	markTodoCmd.Flags().Int("id", -1, "id of the task to be marked todo")
	err = markTodoCmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Printf("error in marking flag as required: %s", err)
		return
	}
}

// markDoneCmd represents the markDone command
var markDoneCmd = &cobra.Command{
	Use:   "mark-done",
	Short: "A command to mark a task as done",
	Long: `A command to mark a task as done.
It accepts an 'id' flag for the task.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			fmt.Printf("error in updating task status: %s", err)
			return
		}
		for i := range TaskDataIns.Tasks {
			if TaskDataIns.Tasks[i].Id == id {
				TaskDataIns.Tasks[i].Status = tasks.Done
				err = TaskDataIns.Save()
				if err != nil {
					fmt.Printf("error in updating task status: %s", err)
					return
				}
				fmt.Printf("%q\t%s\n", TaskDataIns.Tasks[i].Description, TaskDataIns.Tasks[i].Status)
				return
			}
		}
		fmt.Printf("error in updating task status: no task found with the ID(%d)", id)
	},
}

// markInProgressCmd represents the markInProgress command
var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "A command to mark a task as In Progress",
	Long: `A command to mark a task as In Progress.
It accepts an 'id' flag for the task.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			fmt.Printf("error in updating task status: %s", err)
			return
		}
		for i := range TaskDataIns.Tasks {
			if TaskDataIns.Tasks[i].Id == id {
				TaskDataIns.Tasks[i].Status = tasks.InProgress
				err = TaskDataIns.Save()
				if err != nil {
					fmt.Printf("error in updating task status: %s", err)
					return
				}
				fmt.Printf("%q\t%s\n", TaskDataIns.Tasks[i].Description, TaskDataIns.Tasks[i].Status)
				return
			}
		}
		fmt.Printf("error in updating task status: no task found with the ID(%d)", id)
	},
}

// markTodoCmd represents the markTodo command
var markTodoCmd = &cobra.Command{
	Use:   "mark-todo",
	Short: "A command to mark a task as done",
	Long: `A command to mark a task as done.
It accepts an 'id' flag for the task.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			fmt.Printf("error in updating task status: %s", err)
			return
		}
		for i := range TaskDataIns.Tasks {
			if TaskDataIns.Tasks[i].Id == id {
				TaskDataIns.Tasks[i].Status = tasks.ToDo
				err = TaskDataIns.Save()
				if err != nil {
					fmt.Printf("error in updating task status: %s", err)
					return
				}
				fmt.Printf("%q\t%s\n", TaskDataIns.Tasks[i].Description, TaskDataIns.Tasks[i].Status)
				return
			}
		}
		fmt.Printf("error in updating task status: no task found with the ID(%d)", id)
	},
}
