package main

import (
	"fmt"
	"os"
	"path/filepath"
	"task-cli/cmd"
	"task-cli/db"
)

func main() {
	ex, err := os.Executable()
	must(err)
	dbPath := filepath.Join(filepath.Dir(ex), "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
