package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do the command, so move it off the list",
	Run: func(cmd *cobra.Command, args []string) {
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
		}
	},
}

func init() {
	RootCommand.AddCommand(doCmd)
}
