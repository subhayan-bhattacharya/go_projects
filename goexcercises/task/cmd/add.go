package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("add task %s to the list\n", task)
	},
}

func init() {
	RootCommand.AddCommand(addCommand)
}
