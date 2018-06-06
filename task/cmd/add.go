package cmd

import "github.com/spf13/cobra"

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
