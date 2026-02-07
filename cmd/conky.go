package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"danbro/taskman/utils"
)

var conkyCmd = &cobra.Command{
	Use:   "conky [all]",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		showAll := false

		if len(args) == 1 {
			if args[0] != "all" {
				return fmt.Errorf("unknown arg: %s (use 'all' or nothing)", args[0])
			}
			showAll = true
		}

		return printTasksC(showAll)
	},
}

func printTasksC(showAll bool) error {
	tasks, err := store.LoadAll()
	if err != nil {
		return err
	}

	const (
		top = "╔══════════════════════════════════════════╗"
		sep = "╠══════════════════════════════════════════╣"
		bot = "╚══════════════════════════════════════════╝"
	)

	fmt.Println(top)
	fmt.Println("║                TO-DO LIST                ║")
	fmt.Println(sep)

	// фильтрация
	filtered := make([]utils.Task, 0)
	for _, t := range tasks {
		if showAll || !t.Done {
			filtered = append(filtered, t)
		}
	}

	if len(filtered) == 0 {
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║           There are no tasks             ║")
		fmt.Println("║                 relax                    ║")
		fmt.Println("╚══════════════════════════════════════════╝")
		fmt.Println(sep)
		fmt.Println("║ Total tasks: 0                           ║")
		fmt.Println(bot)
		return nil
	}

	fmt.Println(top)

	for i, t := range filtered {
		status := "×"
		if t.Done {
			status = "✓"
		}

		fmt.Printf("║ [%s] %-30s %5d ║\n", status, t.Title, t.Id)
		fmt.Printf(
			"║ To: %-25s %10s ║\n",
			t.Date,
			priorityText(t.Priority),
		)

		if i < len(filtered)-1 {
			fmt.Println(sep)
		}
	}

	fmt.Println(bot)
	fmt.Println(sep)
	fmt.Printf("║ Total tasks: %-27d ║\n", len(filtered))
	fmt.Println(bot)

	return nil
}