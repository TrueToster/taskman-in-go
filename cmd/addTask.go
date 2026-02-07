package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
	"danbro/taskman/utils"
)

var addCmd = &cobra.Command{
	Use:   "addTask <name> <date> <priority> <details>",
	Args:  cobra.MinimumNArgs(4),
	RunE: func(cmd *cobra.Command, args []string) error {
		p, _ := strconv.Atoi(args[4])

		task := utils.Task{
			Id:       store.NextID(),
			Title:    args[0],
			Date:     args[1],
			Priority: p,
			Details:  args[3],
		}

		store.Save(task)
		fmt.Println("✔ Task added")
		return nil
	},
}
