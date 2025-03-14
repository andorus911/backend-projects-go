package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

func (tm *TaskManager) Update(index int, desc string) {
	task := tm.taskList[index]
	task.Description = desc
	task.UpdatedAt = time.Now()
	tm.taskList[index] = task
}

func (tm *TaskManager) List() {
	for index, task := range tm.taskList {
		var status string
		switch task.Status {
		case todo:
			status = "to do"
		case inprogress:
			status = "in progress"
		case done:
			status = "done"
		}

		fmt.Printf("ID %d : *%s* \"%s\" (cr %v, up %v)\n", index, status, task.Description, task.CreatedAt, task.UpdatedAt)
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
		id, err := strconv.Atoi(args[1])
		if err != nil {
			err = fmt.Errorf("incorrect ID: %s", args[1])
			check(err)
		}

		desc := strings.Join(args[2:], " ")

		s := []rune(desc)
		if len(s) > 0 && s[0] == '"' {
			s = s[1:]
		}
		if len(s) > 0 && s[len(s)-1] == '"' {
			s = s[:len(s)-1]
		}

		desc = string(s)
		tm.Update(id, desc)
		tm.WriteFile()
	case "delete":
		// delete one
		// write all
	case "mark-in-progress":
		// update status
		// write all
	case "mark-done":
		// update status
		// write all
	case "list":
		tm.List()
	}
}
