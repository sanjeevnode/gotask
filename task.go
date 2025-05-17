package main

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Done      bool   `json:"done"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
