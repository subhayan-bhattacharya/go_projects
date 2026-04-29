/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	generator "concurrent-file-procesing/internal/generator"

	"github.com/spf13/cobra"
)

// gendataCmd represents the gendata command
var gendataCmd = &cobra.Command{
	Use:   "gendata",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: runGendata,
}

func runGendata(cmd *cobra.Command, args []string) error {
	files, err := cmd.Flags().GetInt("files")
	if err != nil {
		return err
	}
	rows, err := cmd.Flags().GetInt("rows")
	if err != nil {
		return err
	}
	outDir, err := cmd.Flags().GetString("outDir")
	if err != nil {
		return err
	}
	config := generator.CommandConfig{
		Rows:   rows,
		Files:  files,
		OutDir: outDir,
	}
	return config.Run()
}

func init() {
	gendataCmd.Flags().Int("files", 0, "how many files to create")
	gendataCmd.Flags().Int("rows", 0, "how many rows per file")
	gendataCmd.Flags().String("outDir", "", "Path to the output directory")
	rootCmd.AddCommand(gendataCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gendataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gendataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
