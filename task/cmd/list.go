package cmd

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List current tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
