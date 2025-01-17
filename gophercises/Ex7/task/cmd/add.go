package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"task-cli/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.Create(task)
		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Added \"%s\" to your task list", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
