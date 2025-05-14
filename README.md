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
## GitHub user activity

Link on the project task: https://roadmap.sh/projects/github-user-activity

```
# Checking activity summary
go run main.go Qaw
Output:
 - Created 3 branches/tags at Qaw/rafales
 - Created 1 branches/tags at ZigmundKreud/cof
 - Created 2 branches/tags at 12-Monkeys-Developers/tenebris
 - Created 3 branches/tags at Fonderie-VTT/fr-core
 - Created/updated 1 releases at ZigmundKreud/cof
 - Created/updated 1 releases at 12-Monkeys-Developers/tenebris
 - Created/updated 3 releases at Fonderie-VTT/fr-core
 - Created/updated 2 releases at Qaw/rafales
 - Pushed 1 commits to ZigmundKreud/cof
 - Pushed 3 commits to 12-Monkeys-Developers/tenebris
 - Pushed 5 commits to Fonderie-VTT/fr-core
 - Pushed 2 commits to Qaw/rafales
 - Updated 1 comments on PR/issue at Fonderie-VTT/fr-core
 - Updated 2 PRs to Fonderie-VTT/fr-core
```
