package cmd

import (
	"context"
	"task/db"

	"github.com/spf13/cobra"
)

// repositoryKey is used to store the repository in context
type repositoryKey struct{}

var RootCommand = &cobra.Command{
	Use:   "task",
	Short: "Task is a cli task manager.",
}

// SetRepository stores the repository in the root command's context
func SetRepository(repo db.Repository) {
	RootCommand.SetContext(context.WithValue(context.Background(), repositoryKey{}, repo))
}

// GetRepository retrieves the repository from a command's context
func GetRepository(cmd *cobra.Command) db.Repository {
	return cmd.Context().Value(repositoryKey{}).(db.Repository)
}
