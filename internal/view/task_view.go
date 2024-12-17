package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/lemavisaitov/applied-informatics_3/internal/manager"
)

type TaskView struct {
	manager      *manager.TaskManager
	list         *widget.List
	selectedTask int
}

func New(manager *manager.TaskManager) *TaskView {
	tv := &TaskView{
		manager:      manager,
		selectedTask: -1,
		list: widget.NewList(
			func() int {
				return manager.Tasks.Length()
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("")
			},
			func(id widget.ListItemID, item fyne.CanvasObject) {
				di, _ := manager.GetTasks().GetItem(id)
				task := manager.NewTaskFromDataItem(di)
				item.(*widget.Label).SetText(task.Title + " (" + task.Date.Format("02.01.2006") + ")")
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

func (tv *TaskView) Widget() binding.UntypedList {
	return tv.manager.Tasks
}

//func (tv *TaskView) Widget() fyne.CanvasObject {
//	return container.NewVBox(tv.list)
//}
