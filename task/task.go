package task

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)


type Task struct {
	id int
	description string
	status string
	createdAt time.Time 
	updatedAt time.Time
}

func NewTask(
	id int,
	description string,
	status string,
) Task {
	if  reflect.TypeOf(id) != reflect.TypeOf(10) {
		return Task{}
	}
	
	if description  == "" {
		return Task{}
	} 

	return Task {
		id: id,
		description:   description,
		status: status,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func (t Task) Add (id int, description string, status string) {
	task := NewTask(id ,description, status)
	b, err := json.Marshal(task)
	if err != nil {
		log.Fatal("error while trying to add")
	}
	fmt.Println(b)
	fmt.Println("Task added successfully")

} 
