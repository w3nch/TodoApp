package main

import (
	"todo/cmd"
	"todo/storage"
)

func main() {
	todos := cmd.Todos{}
	storage := storage.NewStroage[cmd.Todos]("data/todos.json")
	storage.Load(&todos)

	cmdFlags := cmd.NewCmdFlags()
	cmdFlags.Execute(&todos)

	storage.Save(todos)
}
