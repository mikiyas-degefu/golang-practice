package data

import (
	"errors"
	"sync"
	"task_manager/models"
)

// In-memory store
var (
	mu     sync.Mutex
	tasks  = map[int]models.Task{}
	nextID = 1
)

// GetAll returns all tasks
func GetAll() []models.Task {
	mu.Lock()
	defer mu.Unlock()
	list := make([]models.Task, 0, len(tasks))
	for _, t := range tasks {
		list = append(list, t)
	}
	return list
}

// GetByID returns a task by id
func GetByID(id int) (models.Task, error) {
	mu.Lock()
	defer mu.Unlock()
	t, ok := tasks[id]
	if !ok {
		return models.Task{}, errors.New("not found")
	}
	return t, nil
}

// Create a new task
func Create(t models.Task) models.Task {
	mu.Lock()
	defer mu.Unlock()
	t.ID = nextID
	nextID++
	tasks[t.ID] = t
	return t
}

// Update an existing task
func Update(id int, updated models.Task) (models.Task, error) {
	mu.Lock()
	defer mu.Unlock()
	_, ok := tasks[id]
	if !ok {
		return models.Task{}, errors.New("not found")
	}
	updated.ID = id
	tasks[id] = updated
	return updated, nil
}

// Delete removes a task
func Delete(id int) error {
	mu.Lock()
	defer mu.Unlock()
	_, ok := tasks[id]
	if !ok {
		return errors.New("not found")
	}
	delete(tasks, id)
	return nil
}
