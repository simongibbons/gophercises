package cmd

import (
	"fmt"
	"github.com/simongibbons/gophercises/task/store"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List current tasks",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := store.New("bolt.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		tasks, err := db.GetTasks()
		if err != nil {
			panic(err)
		}

		for _, t := range tasks {
			fmt.Printf("%v - %s\n", *t.Id, t.Content)
		}
	},
}
