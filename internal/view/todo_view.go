package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/lemavisaitov/applied-informatics_3/internal/manager"
)

type TaskView struct {
	manager      *manager.TaskManager
	list         *widget.List
	selectedTask int
}

func NewTaskView(manager *manager.TaskManager) *TaskView {
	tv := &TaskView{
		manager:      manager,
		selectedTask: -1,
		list: widget.NewList(
			func() int {
				return len(manager.GetTasks())
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("")
			},
			func(id widget.ListItemID, item fyne.CanvasObject) {
				task := manager.GetTasks()[id]
				item.(*widget.Label).SetText(task.Title + " (" + task.Date.Format("04 Jul 2005") + ")")
			},
		),
	}
	tv.list.OnSelected = func(id widget.ListItemID) {
		tv.selectedTask = id
	}
	return tv
}

func (tv *TaskView) Refresh() {
	tv.list.Refresh()
}

func (tv *TaskView) GetSelectedTask() int {
	return tv.selectedTask
}

func (tv *TaskView) Widget() fyne.CanvasObject {
	return container.NewVBox(tv.list)
}
