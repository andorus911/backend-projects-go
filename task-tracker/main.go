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

func main() {
	Args := os.Args[1:]

	var jsonBlob = []byte(`[
		{"id": 1, "description": "Monotremata", "status": 0, "createdAt": "0001-01-01T00:00:00Z", "updatedAt": "0001-01-01T00:00:00Z"},
		{"id": 2, "description": "Monotremata", "status": 0, "createdAt": "0001-01-01T00:00:00Z", "updatedAt": "0001-01-01T00:00:00Z"}
	]`)
	var taskMap []Task
	err := json.Unmarshal(jsonBlob, &taskMap)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("taskMap:", taskMap)

	if len(Args) == 0 {
		err = fmt.Errorf("No arguments!")
		fmt.Println("error:", err)
		return
	}

	switch Args[0] {
	case "add":
		// task := new(Task)
		// task.Id = 1
	case "update":
	case "delete":
	case "mark-in-progress":
	case "mark-done":
	case "list":
	}
}
