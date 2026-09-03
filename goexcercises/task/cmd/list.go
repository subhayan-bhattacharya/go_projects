package cmd

import (
	"fmt"
	"task/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the tasks that we have",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("listing all tasks...")
		tasks, err := db.AllTasks()
		if err != nil {
			return err
		}
		if len(tasks) == 0 {
			fmt.Println("no task to complete...")
		}
		for _, task := range tasks {
			fmt.Printf("%d. %s\n", task.Key, task.Value)
		}
		return nil
	},
}

func init() {
	RootCommand.AddCommand(listCmd)
}
