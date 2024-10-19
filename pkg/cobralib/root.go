package cobralib

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	tasks "github.com/DragoHex/task-tracker/pkg/tasks"
)

var (
	DataFile    = filepath.Join("data", "data.json")
	TaskDataIns *tasks.TaskData
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task-tracker",
	Short: "A cli tool for managing tasks",
	Long: `A cli tool for managing tasks.

Multiple tasks can be added.
Each can be assigned a satus to it.
Tasks can be listed and filtered as per their status.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	initData()
}

// initData initialises the json data store
func initData() {
	TaskDataIns = tasks.NewTaskData(DataFile)
	fmt.Println(os.Getwd())

	file, err := os.OpenFile(DataFile, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("error opening or creating file:", err)
		return
	}
	file.Close()

	data, err := os.ReadFile(DataFile)
	if err != nil {
		fmt.Println("error reading from the file:", err)
		return
	}

	// if there is no data in the file
	if data == nil {
		fmt.Println("Created a new json data file")
		return
	}

	// if data exist then import that to the var
	err = json.Unmarshal(data, TaskDataIns)
	if err != nil {
		fmt.Println("error unmarshalling data:", err)
		return
	}
}
