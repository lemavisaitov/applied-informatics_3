package manager

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2/data/binding"
	"github.com/lemavisaitov/applied-informatics_3/pkg/logging"
	"time"
)

type Task struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Date        time.Time `json:"date"`
}

func (t *Task) String() string {
	return fmt.Sprintf("%s - %s", t.Title, t.Date)
}

type TaskManager struct {
	Tasks  binding.UntypedList `json:"tasks"`
	Logger *logging.Logger
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks:  binding.NewUntypedList(),
		Logger: logging.GetLogger(),
	}
}

func (tm *TaskManager) AddTask(title, description string, date time.Time) error {
	if title == "" {
		return errors.New("title is required")
	}
	task := Task{Title: title, Description: description, Date: date}
	if err := tm.Tasks.Append(task); err != nil {
		return err
	}
	tm.Logger.Infof("Added task: %s", task.String())
	return nil
}

func (tm *TaskManager) RemoveTask(index int) error {
	if index < 0 || index >= tm.Tasks.Length() {
		return errors.New("index is out of range")
	}
	di, _ := tm.Tasks.GetItem(index)
	task := tm.NewTaskFromDataItem(di)
	tm.Logger.Infof("task: %s - removed", task.String())
	err := tm.Tasks.Remove(index)
	if err != nil {
		return err
	}
	return nil
}

func (tm *TaskManager) EditTask(index int, title, description string, date time.Time) error {
	if index < 0 || index >= tm.Tasks.Length() {
		return errors.New("index is out of range")
	}
	if title == "" {
		return errors.New("title is required")
	}
	if err := tm.Tasks.SetValue(index, Task{Title: title, Description: description, Date: date}); err != nil {
		return err
	}
	return nil
}

func (tm *TaskManager) GetTasks() binding.UntypedList {
	return tm.Tasks
}

func (tm *TaskManager) NewTaskFromDataItem(di binding.DataItem) Task {
	v, _ := di.(binding.Untyped).Get()
	return v.(Task)
}
