package cmd

import (
	"github.com/spf13/cobra"
	"strconv"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
	},
}
