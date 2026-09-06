package cmd

import (
	"fmt"
	"task/db"

	"github.com/spf13/cobra"
)

var clearCommand = &cobra.Command{
	Use:   "clear",
	Short: "clears the database.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("clear the database...")
		err := db.ClearTasks()
		if err != nil {
			return err
		}
		fmt.Println("cleaned all tasks..")
		return err
	},
}

func init() {
	RootCommand.AddCommand(clearCommand)
}
