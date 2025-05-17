package main

import (
	"encoding/json"
	"os"
)

const fileName = "tasks.json"

func LoadTasks() ([]Task, error) {
	var tasks []Task

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return tasks, nil
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tasks)

	return tasks, err
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}
