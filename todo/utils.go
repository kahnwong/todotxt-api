package todo

import (
	"fmt"
	"os"
	"strings"
)

func sanitizeTodo() {
	// sanitize
	todoBytes, err := os.ReadFile(os.Getenv("TODO_PATH"))
	if err != nil {
		panic(err)
	}

	todoStr := string(todoBytes)
	todoSanitized := strings.ReplaceAll(todoStr, "https://", "")

	// write to file
	err = os.WriteFile(todotxtSanitizedPath, []byte(todoSanitized), 0644)
	if err != nil {
		fmt.Println("Error writing to file")
	}
}
