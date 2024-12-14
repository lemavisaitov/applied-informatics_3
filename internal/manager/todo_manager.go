package manager

import (
	"errors"
	"time"
)

type Task struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Date        time.Time `json:"date"`
}

type TaskManager struct {
	Tasks []Task `json:"tasks"`
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (tm *TaskManager) AddTask(title, decription string, done bool, date time.Time) error {
	if title == "" {
		return errors.New("title is required")
	}

	tm.Tasks = append(tm.Tasks, Task{})
	return nil
}

func (tm *TaskManager) RemoveTask(index int) error {
	if index < 0 || index >= len(tm.Tasks) {
		return errors.New("index is out of range")
	}
	tm.Tasks = append(tm.Tasks[:index], tm.Tasks[index+1:]...)
	return nil
}

func (tm *TaskManager) EditTask(index int, title, decription string, done bool, date time.Time) error {
	if index < 0 || index >= len(tm.Tasks) {
		return errors.New("index is out of range")
	}
	if title == "" {
		return errors.New("title is required")
	}
	tm.Tasks[index] = Task{Title: title, Description: decription, Done: done, Date: date}
	return nil
}
func (tm *TaskManager) GetTasks() []Task {
	return tm.Tasks
}
