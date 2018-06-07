package cmd

import (
	"strings"

	"github.com/simongibbons/gophercises/task/store"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:       "add",
	Short:     "Add a task to the TODO list",
	ValidArgs: []string{"task"},
	Run: func(cmd *cobra.Command, args []string) {
		db, err := store.New("bolt.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		db.AddTask(store.Task{
			Content: strings.Join(args, " "),
		})
	},
}
