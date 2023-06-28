package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GetTaskData() []Task {
	var tasks []Task

	rows, err := db.Query("SELECT * FROM task;")
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var tsk Task
		if err := rows.Scan(&tsk.Id, &tsk.Task, &tsk.Status, &tsk.User, &tsk.Created_at); err != nil {
			return nil
		}
		tasks = append(tasks, tsk)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return tasks
}

func RemoveTask(Id int, MainWindow fyne.Window) (bool, error) {
	rows, err := db.Query("DELETE FROM task WHERE id = ?", Id)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	RefreshTaskList(MainWindow)
	return true, nil
}

func CreateTask(task string, status string, MainWindow fyne.Window) int64 {
	NewTask := Task{
		Task:   task,
		Status: status,
		User:   int(userId),
	}
	result, err := db.Exec("INSERT INTO task (task, status, user) VALUES (?, ?, ?)", NewTask.Task, NewTask.Status, NewTask.User)
	if err != nil {
		fmt.Println(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	RefreshTaskList(MainWindow)
	return id
}

func createTaskHBox(task Task, MainWindow fyne.Window) fyne.CanvasObject {
	remove := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		dialog.ShowConfirm("", "Delete task "+task.Task, func(decision bool) {
			if decision {
				confirm, err := RemoveTask(task.Id, MainWindow)
				if err != nil {
					fmt.Println(err)
				}
				if confirm {
					dialog.ShowInformation("", "Successfully deleted "+task.Task, MainWindow)
				}
			}
		}, MainWindow)
	})
	taskhbox := container.New(layout.NewHBoxLayout(), widget.NewLabel(task.Task), layout.NewSpacer(), remove)
	return taskhbox
}

func PutTasksInLayout(TodoList []Task, MainWindow fyne.Window) *fyne.Container {
	content := container.New(layout.NewVBoxLayout())
	taskFormEntry := widget.NewEntry()
	etaskFormEntry := widget.NewFormItem("Task:", taskFormEntry)
	createNewTask := widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
		// dialog.ShowEntryDialog("New Task", "Task:", func(s string) { CreateNewTask(s, "incomplete") }, MainWindow) //deprecated
		dialog.ShowForm("New Task", "Ok", "Cancel", []*widget.FormItem{etaskFormEntry}, func(b bool) {
			if b {
				CreateTask(taskFormEntry.Text, "incomplete", MainWindow)
			}
		}, MainWindow)
	})
	newTaskButtonContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), createNewTask)
	content.Add(newTaskButtonContainer)
	for i := 0; i < len(TodoList); i++ {
		if TodoList[i].Status == "incomplete" && TodoList[i].User == int(userId) {
			taskhbox := createTaskHBox(TodoList[i], MainWindow)
			content.Add(taskhbox)
		}
	}
	return content
}

func RefreshTaskList(MainWindow fyne.Window) {
	content := container.New(layout.NewVBoxLayout())
	taskFormEntry := widget.NewEntry()
	etaskFormEntry := widget.NewFormItem("Task:", taskFormEntry)
	createNewTask := widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
		// dialog.ShowEntryDialog("New Task", "Task:", func(s string) { CreateNewTask(s, "incomplete") }, MainWindow) //deprecated
		dialog.ShowForm("New Task", "Ok", "Cancel", []*widget.FormItem{etaskFormEntry}, func(b bool) {
			if b {
				CreateTask(taskFormEntry.Text, "incomplete", MainWindow)
			}
		}, MainWindow)
	})
	newTaskButtonContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), createNewTask)
	content.Add(newTaskButtonContainer)
	TodoList := GetTaskData()
	MainWindow.SetContent(PutTasksInLayout(TodoList, MainWindow))
}
