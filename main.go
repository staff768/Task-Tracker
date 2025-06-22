package main

import (
	"fmt"
	"os"
	"task/task"
)
const helpLine string = "Список конманд:\ntask-cli.exe add <Description> (todo done or in-progress)\ntask-cli.exe delete <id>\ntask-cli.exe update <id> <Description>\ntask-cli.exe mark-in-progress <id> or mark-done <id>\ntask-cli.exe list or list <status>"
func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Ошибка. Для помощи введитe <имя программы> help")
	}
	switch args[1]{
	case "help":
		fmt.Println(helpLine)
	case "add":
		id := task.ConverterOfArgs(args[2])
		task.Add(id,args[3])
	case "delete":
		id := task.ConverterOfArgs(args[2])
		task.Delete(id)
	case "update":
		id := task.ConverterOfArgs(args[2])
		task.Update(id, args[3])
	case "list":
		task.List(args[2])
	case "mark-in-progress":
		id := task.ConverterOfArgs(args[2])
		task.Mark_in_progress(id)
	case "mark-done":
		id := task.ConverterOfArgs(args[2])
		task.Mark_done(id)
	default:
		fmt.Println("Не известная команда или аргумент help для списка команд")
	}
}
