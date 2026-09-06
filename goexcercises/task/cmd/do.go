package cmd

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/spf13/cobra"
)

func completeTaskKeys(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	repo := GetRepository(cmd)
	tasks, err := repo.AllTasks()
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	var keys []string
	for _, task := range tasks {
		if slices.Contains(args, strconv.Itoa(task.Key)) {
			continue // already given on the command line
		}
		keys = append(keys, fmt.Sprintf("%d\t%s", task.Key, task.Value))
	}
	return keys, cobra.ShellCompDirectiveNoFileComp
}

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:               "do",
	Short:             "do the command, so move it off the list",
	ValidArgsFunction: completeTaskKeys,
	Run: func(cmd *cobra.Command, args []string) {
		repo := GetRepository(cmd)
		var ids []int
		for _, arg := range args {
			intArg, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("failed to parse the argument: ", arg)
			} else {
				ids = append(ids, intArg)
			}
		}
		for _, id := range ids {
			fmt.Printf("marking id %d off your list\n", id)
			err := repo.DeleteTask(id)
			if err != nil {
				cmd.PrintErrf("Could not delete task %d: %v\n", id, err)
			}
		}
	},
}

func init() {
	RootCommand.AddCommand(doCmd)
}
