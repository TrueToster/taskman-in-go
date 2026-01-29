package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:  "setTask <id> <done|undone>",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := strconv.Atoi(args[0])
		task, err := store.Get(id)
		if err != nil {
			return err
		}

		task.Done = args[1] == "done"
		store.Save(*task)
		fmt.Println("✔ Status set")
		return nil
	},
}
