package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var clearCommand = &cobra.Command{
	Use:   "clear",
	Short: "clears the database.",
	RunE: func(cmd *cobra.Command, args []string) error {
		repo := GetRepository(cmd)
		fmt.Println("clear the database...")
		err := repo.ClearTasks()
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
