package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:  "view <id>",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := strconv.Atoi(args[0])
		t, err := store.Get(id)
		if err != nil {
			return err
		}

		fmt.Println("ID:", t.Id)
		fmt.Println("Title:", t.Title)
		fmt.Println("To:", t.Date)
		fmt.Println("Priority:", t.Priority)
		fmt.Println("Status:", t.Done)
		fmt.Println("Details:", t.Details)
		return nil
	},
}
