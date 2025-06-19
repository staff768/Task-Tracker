package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

const dataFile string = "data.json"

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(
	id int,
	description string,
	status string,
) Task {
	if reflect.TypeOf(id) != reflect.TypeOf(10) {
		return Task{}
	}

	if description == "" {
		return Task{}
	}

	return Task{
		Id:          id,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func checkFileExists(fname string) bool {
	info, err := os.Stat(fname)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Add(id int, description string, status string) {
	var tasks []Task

	if checkFileExists(dataFile) {
		data, err := os.ReadFile(dataFile)
		if err != nil {
			log.Fatal("Ошибка чтения файла при добавлении:", err)
		}
		if len(data) > 0 {
			err = json.Unmarshal(data, &tasks)
			if err != nil {
				log.Fatal("Ошибка парсинга JSON:", err)
			}
		}
	}

	task := NewTask(id, description, status)
	tasks = append(tasks, task)

	b, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		log.Fatal("Ошибка сериализации JSON:", err)
	}
	err = os.WriteFile(dataFile, b, 0644)
	if err != nil {
		log.Fatal("Ошибка записи файла:", err)
	} else {
		fmt.Println("Добавление задачи прошло удачно Id задачи", task.Id)
	}
}
func Delete(id int) {
	var tasks []Task
	data, err := os.ReadFile(dataFile)
	if err != nil {
		log.Fatal("Ошибка чтения файла при удалении:", err)
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal("Ошибка парсинга JSON:", err)
	}

	for i :=0 ; i < len(tasks); i++{
		if tasks[i].Id == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	b, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		log.Fatal("Ошибка сериализации JSON:", err)
	}
	err = os.WriteFile(dataFile, b, 0644)
	if err != nil {
		log.Fatal("Ошибка записи файла при удалении:", err)
	} else {
		fmt.Println("Удаление задачи прошло удачно Id удаленной задачи", id)
	}

}

func main() {
	//Add(1, "got to eat","todo" )
	//Add(2, "postrat", "in-progress")
	//Add(3, "do homework", "todo")
	//Add(4, "do homework", "todo")
	//Delete(2)
}
