package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Status byte

const (
	todo Status = iota
	inprogress
	done
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadLog(fileLogName string) map[int]Task {
	taskMap := make(map[int]Task)
	b, err := os.ReadFile(fileLogName)
	if err != nil {
		b, err = json.Marshal(taskMap)
		check(err)
		os.WriteFile(fileLogName, b, os.ModeAppend)
	}
	err = json.Unmarshal(b, &taskMap)
	check(err)
	return taskMap
}

func main() {

	const fileLogName = "/tmp/task_tracker_log"

	args := os.Args[1:]

	if len(args) == 0 {
		err := fmt.Errorf("No arguments!")
		fmt.Println("error:", err)
		return
	}

	switch args[0] {
	case "add":
		// create new task
		// read list
		// add new
		// write all
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
		tasks := ReadLog(fileLogName)
		for _, t := range tasks {
			fmt.Println(t)
		}
	}
}
