package todo

import (
	"fmt"
	"strings"

	todo "github.com/1set/todotxt"
)

var todotxtSanitizedPath = "/tmp/todo.txt"

type Todo struct {
	ID   int    `json:"id"`
	Todo string `json:"todo"`
}

func parseTodos(tasks todo.TaskList) []Todo {
	var todos []Todo

	for _, t := range tasks {
		// context
		var context string
		if len(t.Contexts) > 0 {
			context = fmt.Sprintf("@%s", t.Contexts[0])
		}

		// project
		var project string
		if len(t.Projects) > 0 {
			project = fmt.Sprintf("+%s", t.Projects[0])
		}

		// append
		todos = append(todos, Todo{
			t.ID,
			strings.TrimSpace(fmt.Sprintf("%s %s %s", context, project, t.Todo)),
		})
	}

	return todos
}

func getTodos() []Todo {
	var todos []Todo

	sanitizeTodo() // strip leading `https://` which results in the todo body returning null
	if tasklist, err := todo.LoadFromPath(todotxtSanitizedPath); err != nil {
		fmt.Printf("Error reading todo.txt: %s", err)
	} else {
		todayTasks := tasklist.Filter(todo.FilterNotCompleted).Filter(todo.FilterDueToday, todo.FilterOverdue)
		_ = todayTasks.Sort(todo.SortPriorityAsc, todo.SortProjectAsc)
		todos = parseTodos(todayTasks)

	}
	return todos
}
