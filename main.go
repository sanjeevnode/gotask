package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	tasks, _ := LoadTasks()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n\n--- TO-DO MENU ---")
		fmt.Println("[1] Add Task")
		fmt.Println("[2] List Tasks")
		fmt.Println("[3] Mark Task as Done")
		fmt.Println("[4] Delete Task")
		fmt.Println("[5] Exit")
		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println()
		switch choice {
		case "1":
			fmt.Print("Enter task title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			id := getNextID(tasks)
			tasks = append(tasks, Task{ID: id, Name: title, Done: false, CreatedAt: now, UpdatedAt: now})
			fmt.Println("✅ Task added.")

		case "2":
			if len(tasks) == 0 {
				fmt.Println("No tasks found.")
			} else {
				printTaskTable(tasks)
			}

		case "3":
			fmt.Print("Enter task ID to mark as done: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			found := false
			for i, t := range tasks {
				if t.ID == id {
					tasks[i].Done = true
					tasks[i].UpdatedAt = now
					fmt.Println("✅ Task marked as done.")
					found = true
					break
				}
			}
			if !found {
				fmt.Println("⚠️ Task not found.")
			}

		case "4":
			fmt.Print("Enter task ID to delete: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			found := false
			for i, t := range tasks {
				if t.ID == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					fmt.Println("🗑️ Task deleted.")
					found = true
					break
				}
			}
			if !found {
				fmt.Println("⚠️ Task not found.")
			}

		case "5":
			SaveTasks(tasks)
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("❌ Invalid choice. Please try again.")
		}
		println("\n")
	}
}

func getNextID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}
	maxID := tasks[0].ID
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

func printTaskTable(tasks []Task) {
	const maxTitleWidth = 30

	// Header
	fmt.Printf("\n%-5s | %-30s | %-6s | %-19s | %-19s\n", "ID", "Title", "Done", "Created At", "Updated At")
	fmt.Println(strings.Repeat("-", 90))

	for _, task := range tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}

		// Split title if too long
		lines := wrapText(task.Name, maxTitleWidth)

		for i, line := range lines {
			if i == 0 {
				fmt.Printf("%-5d | %-30s | %-6s | %-19s | %-19s\n", task.ID, line, status, task.CreatedAt, task.UpdatedAt)
			} else {
				fmt.Printf("      | %-30s | %-6s | %-19s | %-19s\n", line, "", "", "")
			}
		}
	}
}

func wrapText(text string, limit int) []string {
	var lines []string
	for len(text) > limit {
		lines = append(lines, text[:limit])
		text = text[limit:]
	}
	lines = append(lines, text)
	return lines
}
