package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		file := viper.GetString("todo_file")
		if file == "" {
			file = "todos.txt"
		}

		f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		cobra.CheckErr(err)
		defer f.Close()

		_, err = f.WriteString(task + "\n")
		cobra.CheckErr(err)

		fmt.Println("Added task:", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
