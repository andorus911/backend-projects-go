# Backend projects: Go

## Task Tracker

Link on the project task: https://roadmap.sh/projects/task-tracker

```
# Adding a new task
go run main.go add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
go run main.go update 1 "Buy groceries and cook dinner"
go run main.go delete 1

# Marking a task as in progress or done
go run main.go mark-in-progress 1
go run main.go mark-done 1

# Listing all tasks
go run main.go list

# Listing tasks by status
go run main.go list done
go run main.go list todo
go run main.go list in-progress
```
