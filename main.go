package main

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var userId int64
var currentStatusPage string = "incomplete"

type User struct {
	Id         int
	Name       string
	Username   string
	Email      string
	Password   string
	Created_at string
}

type Task struct {
	Id         int
	Task       string
	Status     string
	User       int
	Created_at string
}

func ConnectToSql() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "users",
	}
	fmt.Println(cfg.FormatDSN())
	var err error
	db, err = sql.Open("mysql", "root:password@/TodoGui")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}

func main() {
	var (
		TodoGui    = app.New()
		MainWindow = TodoGui.NewWindow("Todo List Manager - @Pineman834")
	)
	MainWindow.Resize(fyne.NewSize(400, 100))
	ConnectToSql()

	LoginPage(MainWindow)
	MainWindow.ShowAndRun()
}

func MainPage(MainWindow fyne.Window) {
	var (
		TodoList = GetTaskData()
		content  = PutTasksInLayout(TodoList, MainWindow)
	)
	MainWindow.SetContent(content)
}

func LoginPage(MainWindow fyne.Window) {
	var (
		uname   = widget.NewEntry()
		pword   = widget.NewEntry()
		newuser = widget.NewButton("Sign up", func() { NewUserPage(MainWindow) })
	)
	pword.Password = true

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Username:", Widget: uname}},
		OnSubmit: func() {
			if login(uname.Text, pword.Text) != nil {
				MainWindow.SetContent(widget.NewLabel("Successful login"))
				MainPage(MainWindow)
			} else {
				MainWindow.SetContent(widget.NewLabel("Inncorrect Username or Password"))
			}
			// MainWindow.Close()
		},
	}

	form.Append("Password:", pword)

	MainWindow.SetContent(container.New(layout.NewVBoxLayout(), form, newuser))
}

func NewUserPage(MainWindow fyne.Window) {
	var (
		name            = widget.NewEntry()
		uname           = widget.NewEntry()
		email           = widget.NewEntry()
		pword           = widget.NewEntry()
		sha256pword     = sha256.Sum256([]byte(uname.Text))
		loginpagebutton = widget.NewButton("Login", func() { LoginPage(MainWindow) })
	)
	pword.Password = true

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name:", Widget: name}},
		OnSubmit: func() {
			usrID, err := createUser(User{
				Name:     name.Text,
				Username: string(sha256pword[:]),
				Email:    email.Text,
				Password: pword.Text,
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("ID of added User: %v\n", usrID)
			if login(uname.Text, pword.Text) != nil {
				MainWindow.SetContent(container.New(layout.NewVBoxLayout(), widget.NewLabel("Successfuly created user, "+uname.Text), widget.NewButton("Main page", func() { MainPage(MainWindow) })))
			} else {
				MainWindow.SetContent(widget.NewLabel("Something went wrong"))
			}
			// MainWindow.Close()
		},
	}

	form.Append("Username:", uname)
	form.Append("Email:", email)
	form.Append("Password:", pword)

	MainWindow.SetContent(container.New(layout.NewVBoxLayout(), form, loginpagebutton))
}
