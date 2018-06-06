package cmd

import "github.com/spf13/cobra"

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

