package main

import (
	"database/sql"
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	createTable := `CREATE TABLE IF NOT EXISTS messages(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text TEXT NOT NULL
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	_, err = db.Exec(`INSERT INTO messages (text) VALUES (?)`, "Hello Fyne from lab2!")
	if err != nil {
		log.Fatal("Failed to insert data:", err)
	}

	var message string
	err = db.QueryRow("SELECT text FROM messages ORDER BY id DESC LIMIT 1").Scan(&message)
	if err != nil {
		log.Fatal("Failed to query data:", err)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("SQLite3 Test - Lab2")

	label := widget.NewLabel(message)
	myWindow.SetContent(container.NewVBox(label))
	
	myWindow.ShowAndRun()
}