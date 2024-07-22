package cmd

import (
	"fmt"
	"os"
	"strconv"
	"task-cli/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err == nil {
				ids = append(ids, id)
			} else {
				fmt.Printf("Failed to parse argument: %s\n", arg)
			}
		}

		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			os.Exit(1)
		}

		for _, id := range ids {
			if id < 1 || id > len(tasks) {
				fmt.Printf("Task number %d does not exist\n", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Something went wrong when deleting task number %d: %s\n", id, err.Error())
			} else {
				fmt.Printf("Marked task number %d: \"%s\" as completed\n", id, task.Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
