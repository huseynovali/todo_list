package main

import (
	"fmt"
	"time"
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


