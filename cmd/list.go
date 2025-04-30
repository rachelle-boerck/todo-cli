package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		file := viper.GetString("todo_file")
		if file == "" {
			file = "todos.txt"
		}

		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("No tasks found.")
			return
		}

		tasks := strings.Split(string(data), "\n")
		for i, task := range tasks {
			if strings.TrimSpace(task) != "" {
				fmt.Printf("%d. %s\n", i+1, task)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
