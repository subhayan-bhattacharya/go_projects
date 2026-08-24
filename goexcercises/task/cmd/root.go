package cmd

import "github.com/spf13/cobra"

var RootCommand = &cobra.Command{
	Use:   "task",
	Short: "Task is a cli task manager.",
}
