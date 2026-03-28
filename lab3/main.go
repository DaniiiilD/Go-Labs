package main

import (
	"database/sql"
	"log"

	"lab3/app"
	"lab3/database"

	fyneapp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	conn, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	database := database.New(conn)
	appInstance := app.New(database)

	message, err := appInstance.Run()
	if err != nil {
		log.Fatal(err)
	}

	myApp := fyneapp.New()
	myWindow := myApp.NewWindow("SQLite3 Test - Lab3")
	myWindow.SetContent(container.NewVBox(widget.NewLabel(message)))
	myWindow.ShowAndRun()
}
