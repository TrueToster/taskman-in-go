package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Показать список задач",
	RunE: func(cmd *cobra.Command, args []string) error {
		return printTasks()
	},
}

func printTasks() error {
	tasks, err := store.LoadAll()
	if err != nil {
		return err
	}

	const (
		top = "╔══════════════════════════════════════════════════╗"
		sep = "╠══════════════════════════════════════════════════╣"
		bot = "╚══════════════════════════════════════════════════╝"
	)

	fmt.Println(top)
	fmt.Println("║                   СПИСОК ЗАДАЧ                   ║")
	fmt.Println(sep)

	if len(tasks) == 0 {
		fmt.Println("╔══════════════════════════════════════════════════╗")
		fmt.Println("║ Добавь задачу командой:                          ║")
		fmt.Println("║ taskman add \"имя\" дата приоритет детали          ║")
		fmt.Println("╚══════════════════════════════════════════════════╝")
		fmt.Println(sep)
		fmt.Println("║ Всего задач: 0                                   ║")
		fmt.Println(bot)
		return nil
	}

	fmt.Println(top)

	for i, t := range tasks {
		status := "×"
		if t.Done {
			status = "✓"
		}

		fmt.Printf("║ [%s] %-38s %5d ║\n", status, t.Title, t.Id)
		fmt.Printf(
			"║ To: %-32s %s%11s%s ║\n",
			t.Date,
			priorityColor(t.Priority),
			priorityText(t.Priority),
			colorReset,
		)

		if i < len(tasks)-1 {
			fmt.Println(sep)
		}
	}

	fmt.Println(bot)
	fmt.Println(sep)
	fmt.Printf("║ Всего задач: %-35d ║\n", len(tasks))
	fmt.Println(bot)

	return nil
}

const colorReset = "\033[0m"

func priorityText(p int) string {
	switch p {
	case 0:
		return "special"
	case 1:
		return "lowest"
	case 2:
		return "low"
	case 3:
		return "medium"
	case 4:
		return "high"
	case 5:
		return "highest"
	default:
		return "invalid"
	}
}

func priorityColor(p int) string {
	switch p {
	case 0:
		return "\033[35m" // purple
	case 1:
		return "\033[37m" // gray
	case 2:
		return "\033[32m" // green
	case 3:
		return "\033[34m" // blue
	case 4:
		return "\033[33m" // yellow
	case 5:
		return "\033[31m" // red
	default:
		return ""
	}
}