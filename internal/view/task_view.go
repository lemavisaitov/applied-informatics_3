package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/lemavisaitov/applied-informatics_3/internal/manager"
)

type TaskView struct {
	manager      *manager.TaskManager
	list         *widget.List
	selectedTask int
}

func New(taskManager *manager.TaskManager) *TaskView {
	tv := &TaskView{
		manager:      taskManager,
		selectedTask: -1,
		list: widget.NewListWithData(
			taskManager.GetTasks(),
			func() fyne.CanvasObject {
				return container.NewBorder(
					nil, nil, nil,
					// left of the border
					widget.NewCheck("", func(b bool) {}),
					// takes the rest of the space
					widget.NewLabel(""),
				)
			},
			func(di binding.DataItem, o fyne.CanvasObject) {
				ctr, _ := o.(*fyne.Container)

				l := ctr.Objects[0].(*widget.Label)
				c := ctr.Objects[1].(*widget.Check)

				diu, _ := di.(binding.Untyped).Get()
				todo := diu.(manager.Task)
				l.SetText(todo.Title)
				c.SetChecked(todo.Done)
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
	return container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		tv.list,
	)
}
