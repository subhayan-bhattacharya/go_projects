package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "adds a task to your task list.",
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := GetRepository(cmd)
		priority, err := cmd.Flags().GetInt("priority")
		if err != nil {
			return err
		}
		task := strings.Join(args, " ")
		key, err := repo.CreateTask(task, priority)
		if err != nil {
			return err
		}
		fmt.Printf("task created	 with key %d\n", key)
		return nil
	},
}

func init() {
	addCommand.Flags().Int("priority", 1, "what is the priority of the command")
	RootCommand.AddCommand(addCommand)
}
