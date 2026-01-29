package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:  "delTask <id>",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := strconv.Atoi(args[0])
		return store.Delete(id)
	},
}
