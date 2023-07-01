/*
Package db represents database and stores methods to work with database
*/

package db

import (
	"GoProject/model"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"log"
)

var DB Database

type Database struct {
	Db *sql.DB
}

func NewDatabase() error {
	dsn := mysql.Config{
		User:   "root",
		Passwd: "1234",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "go_project",
	}
	db, err := sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatal("Couldn't connect to database!")
		return err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		return pingErr
	}

	DB.Db = db
	addTablesToDataBase()

	return nil
}

func addTablesToDataBase() {
	table := "CREATE TABLE IF NOT EXISTS tasks(id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, title VARCHAR(20), description TEXT)"
	_, err := DB.Db.Exec(table)
	if err != nil {
		log.Fatal("Couldn't crate table")
	}
}

// GetTasks returns all tasks from database
func (db *Database) GetTasks() []model.Task {
	var tasks []model.Task
	rows, err := db.Db.Query("SELECT * FROM tasks")
	if err != nil {
		return tasks
	}

	defer rows.Close()

	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description); err != nil {
			return tasks
		}

		tasks = append(tasks, task)
	}

	return tasks
}

func (db *Database) CreateTAsk(task model.Task, ctx *gin.Context) {
	tx, err := db.Db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Print("Couldn't start transaction!")
		return
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "INSERT INTO tasks (title, description) VALUES (?, ?)", task.Title, task.Description)
	if err != nil {
		fmt.Print("Couldn't execute context!")
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Print("Couldn't commit!")
		return
	}
}

func (db *Database) DeleteTask(id int, ctx *gin.Context) {
	tx, err := db.Db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Print("Couldn't start transaction!")
		return
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		fmt.Printf("Error while deleting row %d", id)
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Print("Couldn't commit!")
		return
	}
}
