package cmd

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
	"danbro/taskman/utils"
)

var (
	rootCmd = &cobra.Command{
		Use:   "taskman",
		Short: "Simple CLI manager tasks",
	}
	store *utils.Storage
)

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	baseDir := getBaseDir()
	tasksDir := filepath.Join(baseDir, "tasks")

	store = utils.NewStorage(tasksDir)
	_ = store.Init()

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(delCmd)
	rootCmd.AddCommand(viewCmd)
	rootCmd.AddCommand(conkyCmd)
}

func getBaseDir() string {
	home, _ := os.UserHomeDir()

	switch runtime.GOOS {
	case "windows":
		if appdata := os.Getenv("APPDATA"); appdata != "" {
			return filepath.Join(appdata, "taskman")
		}
		return filepath.Join(home, "taskman")

	case "darwin":
		return filepath.Join(home, "Library", "Application Support", "taskman")

	default: // linux
		return filepath.Join(home, ".local", "share", "taskman")
	}
}
