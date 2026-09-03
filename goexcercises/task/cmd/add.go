package cmd

import (
	"fmt"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "adds a task to your task list.",
	RunE: func(cmd *cobra.Command, args []string) error {
		task := strings.Join(args, " ")
		key, err := db.CreateTask(task)
		if err != nil {
			return err
		}
		fmt.Printf("task created with key %d\n", key)
		return nil
	},
}

func init() {
	RootCommand.AddCommand(addCommand)
}
