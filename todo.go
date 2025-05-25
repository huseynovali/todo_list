package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) Delete(index int) error {
	t := *todos

	if index < 0 || index >= len(t) {
		return fmt.Errorf("index out of range")
	}

	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) Toogle(index int) error {
	t := *todos

	if index < 0 || index >= len(t) {
		return fmt.Errorf("index out of range")
	}

	isCompleted := t[index].Completed
	if !isCompleted {
		now := time.Now()
		t[index].CompletedAt = &now
	}

	t[index].Completed = !isCompleted

	return nil
}



func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if index < 0 || index >= len(t) {
		return fmt.Errorf("index out of range")
	}

	t[index].Title = title
	return nil
}

func (todos *Todos) Print(){
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for i, todo := range *todos {
		complated := "❌"
		complatedAt := "N/A"

		if todo.Completed {
			complated = "✅"
			if todo.CompletedAt != nil {
				complatedAt = todo.CompletedAt.Format(time.RFC3339)
			}
		}

		table.AddRow(strconv.Itoa(i+1), todo.Title, complated, todo.CreatedAt.Format(time.RFC3339), complatedAt)

}

	table.Render()
}

