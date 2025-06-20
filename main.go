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
				log.Fatal("Ошибка парсинга JSON при добавлении:", err)
			}
		}
	}

	task := NewTask(id, description, status)
	tasks = append(tasks, task)

	b, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		log.Fatal("Ошибка сериализации JSON при добавлении:", err)
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
		log.Fatal("Ошибка парсинга JSON при удалении:", err)
	}

	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	b, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		log.Fatal("Ошибка сериализации JSON при удалении:", err)
	}
	err = os.WriteFile(dataFile, b, 0644)
	if err != nil {
		log.Fatal("Ошибка записи файла при удалении:", err)
	} else {
		fmt.Println("Удаление задачи прошло удачно Id удаленной задачи", id)
	}

}
func Update(id int, new_description string) {
	var tasks []Task
	data, err := os.ReadFile(dataFile)
	if err != nil {
		log.Fatal("Ошибка чтения файла при обновлении:", err)
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal("Ошибка парсинга JSON при обнолвении:", err)
	}
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == id {
			tasks[i].Description = new_description
		}
	}
	b, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		log.Fatal("Ошибка сериализации JSON при обновлении :", err)
	}
	err = os.WriteFile(dataFile, b, 0644)
	if err != nil {
		log.Fatal("Ошибка записи файла при обновлении:", err)
	} else {
		fmt.Println("Обновление задачи прошло удачно Id обновленной задачи", id)
	}
}
func List(str string){
	var tasks []Task
	data,err := os.ReadFile(dataFile)
	if err != nil{
		log.Fatal("Ошибка чтения файла при выводе всех задач:", err)
	}
	err = json.Unmarshal(data,&tasks)
	if err !=nil {
		log.Fatal("Ошибка парсинга JSON при постановке статуса:", err)
	}
	switch str {
	case "":
		for i:=0; i < len(tasks);i++{
			fmt.Println(tasks[i]) 
		}
	case "todo":
		for i:=0; i < len(tasks);i++{
			if tasks[i].Status == "todo"{
				fmt.Println(tasks[i])
			}
		}
	case "in-progress":
		for i:=0; i < len(tasks);i++{
			if tasks[i].Status == "in-progress"{
				fmt.Println(tasks[i])
			}
		}
	case "done":
		for i:=0; i < len(tasks);i++{
			if tasks[i].Status == "done"{
				fmt.Println(tasks[i])
			}
		}
	}
}
func Mark_in_progress(id int){
	var tasks []Task
	data,err := os.ReadFile(dataFile)
	if err != nil {
		log.Fatal("Ошибка чтения файла при постановке статуса:", err)
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal("Ошибка парсинга JSON при постановке статуса:", err)
	}
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == id {
			tasks[i].Status = "in-progress"
		}
	}
	b, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		log.Fatal("Ошибка сериализации JSON при постановке статуса :", err)
	}
	err = os.WriteFile(dataFile, b, 0644)
	if err != nil {
		log.Fatal("Ошибка записи файла при постановке статуса:", err)
	} else {
		fmt.Println("Обновление задачи прошло удачно Id обновленной задачи", id)
	}
}
func Mark_done(id int){
	var tasks []Task
	data,err := os.ReadFile(dataFile)
	if err != nil {
		log.Fatal("Ошибка чтения файла при постановке статуса:", err)
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal("Ошибка парсинга JSON при постановке статуса:", err)
	}
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == id {
			tasks[i].Status = "done"
		}
	}
	b, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		log.Fatal("Ошибка сериализации JSON при постановке статуса :", err)
	}
	err = os.WriteFile(dataFile, b, 0644)
	if err != nil {
		log.Fatal("Ошибка записи файла при постановке статуса:", err)
	} else {
		fmt.Println("Обновление задачи прошло удачно Id обновленной задачи", id)
	}
}
func Mark_todo(id int){
	var tasks []Task
	data,err := os.ReadFile(dataFile)
	if err != nil {
		log.Fatal("Ошибка чтения файла при постановке статуса:", err)
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal("Ошибка парсинга JSON при постановке статуса:", err)
	}
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == id {
			tasks[i].Status = "todo"
		}
	}
	b, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		log.Fatal("Ошибка сериализации JSON при постановке статуса :", err)
	}
	err = os.WriteFile(dataFile, b, 0644)
	if err != nil {
		log.Fatal("Ошибка записи файла при постановке статуса:", err)
	} else {
		fmt.Println("Обновление задачи прошло удачно Id обновленной задачи", id)
	}
}

func main() {

	//Add(1, "got to eat","todo" )
	//Add(2, "postrat", "in-progress")
	//Mark_todo(1)
	//List("todo")
	//Update(1,"sleep until 12 pm")
	//Add(3, "do homework", "todo")
	//Add(4, "do homework", "todo")
	//Delete(1)
}
