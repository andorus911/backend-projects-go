package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type TaskManager struct {
	path     string
	taskList map[int]Task
	nextID   int
}

type Status byte

const (
	todo Status = iota
	inprogress
	done
)

type Task struct {
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTaskManager(path string) *TaskManager {
	tm := &TaskManager{
		path:     path,
		taskList: make(map[int]Task),
		nextID:   0,
	}

	return tm
}

func (tm *TaskManager) ReadFile() {
	b, err := os.ReadFile(tm.path)
	if err != nil {
		b, err = json.Marshal(tm.taskList)
		check(err)
		os.WriteFile(tm.path, b, os.ModeAppend)
	}
	err = json.Unmarshal(b, &tm.taskList)
	check(err)

	for i := range tm.taskList {
		if i >= tm.nextID {
			tm.nextID = i + 1
		}
	}
}

func (tm *TaskManager) AddTask(desc string) {
	newTask := Task{
		Description: desc,
		Status:      todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tm.taskList[tm.nextID] = newTask
	tm.nextID++
}

func (tm *TaskManager) List() {
	for index, task := range tm.taskList {
		fmt.Printf("ID %d : %s (cr %v, up %v)", index, task.Description, task.CreatedAt, task.UpdatedAt)
		fmt.Println(task)
	}
}

func (tm *TaskManager) WriteFile() {
	b, err := json.Marshal(tm.taskList)
	check(err)
	err = os.WriteFile(tm.path, b, os.ModeAppend)
	check(err)
}

func main() {

	const fileName = "/tmp/task_tracker_log"
	tm := NewTaskManager(fileName)
	tm.ReadFile()

	args := os.Args[1:]

	if len(args) == 0 {
		err := fmt.Errorf("no arguments")
		panic(err)
	}

	switch args[0] {
	case "add":
		desc := strings.Join(args[1:], " ")

		s := []rune(desc)
		if len(s) > 0 && s[0] == '"' {
			s = s[1:]
		}
		if len(s) > 0 && s[len(s)-1] == '"' {
			s = s[:len(s)-1]
		}

		desc = string(s)
		tm.AddTask(desc)
		tm.WriteFile()
	case "update":
		// read list
		// update
		// write all
	case "delete":
		// read list
		// delete one
		// write all
	case "mark-in-progress":
		// read list
		// update status
		// write all
	case "mark-done":
		// read list
		// update status
		// write all
	case "list":
		tm.List()
	}
}
