package controller

import (
	"errors"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/lemavisaitov/applied-informatics_3/internal/manager"
	"github.com/lemavisaitov/applied-informatics_3/internal/view"
)

type AppController struct {
	app         fyne.App
	taskManager *manager.TaskManager
	view        *view.TaskView
}

func New() *AppController {
	todoManager := manager.NewTaskManager()
	vw := view.New(todoManager)
	todoManager.Logger.Infof("Create AppController")
	return &AppController{
		app:         app.New(),
		taskManager: todoManager,
		view:        vw,
	}
}

func (ac *AppController) Run() {
	mainWindow := ac.app.NewWindow("Day Planner")
	newtodoDescTxt := widget.NewEntry()
	newtodoDescTxt.PlaceHolder = "New Todo Description..."

	addTaskButton := widget.NewButton("Add", func() {
		ac.showAddTaskDialog(mainWindow)
	})

	editTaskButton := widget.NewButton("Edit", func() {
		selected := ac.view.GetSelectedTask()
		if selected >= 0 {
			di, err := ac.taskManager.GetTasks().GetItem(selected)
			if err != nil {
				ac.taskManager.Logger.Error(err)
			}
			task := ac.taskManager.NewTaskFromDataItem(di)
			ac.showEditTaskDialog(mainWindow, selected, task)
		}
	})

	removeTaskButton := widget.NewButton("Remove", func() {
		selected := ac.view.GetSelectedTask()
		if selected >= 0 {
			err := ac.taskManager.RemoveTask(selected)
			if err != nil {
				dialog.ShowError(err, mainWindow)
				return
			}
			ac.view.Refresh()
		}
	})

	mainWindow.SetContent(
		container.NewBorder(
			nil,
			container.NewGridWithColumns(
				3,
				addTaskButton,
				editTaskButton,
				removeTaskButton,
			),
			nil,
			nil,
			ac.view.Widget(),
		),
	)
	mainWindow.Resize(fyne.NewSize(400, 600))
	mainWindow.ShowAndRun()
}

func (ac *AppController) showAddTaskDialog(win fyne.Window) {
	//addTaskWindow := ac.app.NewWindow("Add Task")
	//addTaskWindow.Resize(fyne.NewSize(400, 300))
	//addTaskWindow.SetFixedSize(false)
	titleEntry := widget.NewEntry()
	descriptionEntry := widget.NewMultiLineEntry()
	datePicker := widget.NewEntry()
	datePicker.SetPlaceHolder("DD.MM.YYYY")

	form := widget.NewForm(
		widget.NewFormItem("Title", titleEntry),
		widget.NewFormItem("Date", datePicker),
		widget.NewFormItem("Description", descriptionEntry),
	)

	dlg := dialog.NewCustomConfirm("Add task", "Add", "Cancel", form, func(confirm bool) {
		if confirm {
			date, err := time.Parse("02.01.2006", datePicker.Text)
			if err != nil {
				dialog.ShowError(errors.New("invalid date format"), win)
				return
			}
			err = ac.taskManager.AddTask(titleEntry.Text, descriptionEntry.Text, date)

			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			ac.view.Refresh()
		}
	}, win)

	dlg.Resize(fyne.NewSize(400, 300))
	dlg.Show()

	//addButton := widget.NewButton("Add", func() {
	//	defer addTaskWindow.Close()
	//	date, err := time.Parse("02.01.2006", datePicker.Text)
	//	if err != nil {
	//		dialog.ShowError(errors.New("invalid date format"), win)
	//		return
	//	}
	//	err = ac.taskManager.AddTask(titleEntry.Text, descriptionEntry.Text, date)
	//	if err != nil {
	//		dialog.ShowError(err, win)
	//		return
	//	}
	//	//addedTask := ac.taskManager.Tasks[ac.view.GetSelectedTask()]
	//	//ac.logger.Infof("Added task ", addedTask.Title, addedTask.Description)
	//	ac.view.Refresh()
	//})
	//cancelButton := widget.NewButton("Cancel", func() {
	//	addTaskWindow.Close()
	//})
	//addTaskWindow.SetContent(
	//	container.NewBorder(
	//		nil,
	//		container.NewGridWithColumns(
	//			2,
	//			addButton,
	//			cancelButton,
	//		),
	//		nil,
	//		nil,
	//		form,
	//	),
	//)
	//addTaskWindow.Show()
}

func (ac *AppController) showEditTaskDialog(win fyne.Window, index int, task manager.Task) {
	titleEntry := widget.NewEntry()
	titleEntry.SetText(task.Title)
	descriptionEntry := widget.NewMultiLineEntry()
	descriptionEntry.SetText(task.Description)
	datePicker := widget.NewEntry()
	datePicker.SetText(task.Date.Format("02.01.2006"))

	form := widget.NewForm(
		widget.NewFormItem("Title", titleEntry),
		widget.NewFormItem("Description", descriptionEntry),
		widget.NewFormItem("Date", datePicker),
	)
	dlg := dialog.NewCustomConfirm("Edit Task", "Save", "Cancel", form, func(confirm bool) {
		if confirm {
			date, err := time.Parse("02.01.2006", datePicker.Text)
			if err != nil {
				dialog.ShowError(errors.New("invalid date format"), win)
				return
			}
			err = ac.taskManager.EditTask(index, titleEntry.Text, descriptionEntry.Text, date)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			ac.view.Refresh()
		}
	}, win)
	dlg.Resize(fyne.NewSize(400, 300))
	dlg.Show()
}
