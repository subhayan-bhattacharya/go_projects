package cmd

import (
	"cmp"
	"fmt"
	"slices"
	"task/db"

	"github.com/spf13/cobra"
)

func sortTasksByPriority(tasks []db.Task) []db.Task {
	copiedTasks := append([]db.Task(nil), tasks...)
	slices.SortFunc(copiedTasks, func(a, b db.Task) int {
		return cmp.Compare(a.Priority, b.Priority)
	})
	return copiedTasks
}

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
		tasks = sortTasksByPriority(tasks)
		if len(tasks) == 0 {
			fmt.Println("no task to complete...")
		}
		for _, task := range tasks {
			fmt.Printf("%d. Priority: %d, task : %s\n", task.Key, task.Priority, task.Value)
		}
		return nil
	},
}

func init() {
	RootCommand.AddCommand(listCmd)
}
