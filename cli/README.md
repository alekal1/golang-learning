## CLI app
Command-line tool in Go that manages a list of tasks.


1. Add a task - add "Task Description"
   
	Example: `go run main.go add "Buy Groceries"`
 3. List tasks - list

	Example: `go run main.go list`

	Output:
```
1. [ ] Buy groceries
2. [x] Pay bills
```
    
 5. Mark Task as Done - `go run main.go done 2`
 6. Delete task - `go run main.go delete 2`

 Persistance - *JSON* file named tasks.json
