package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

var DB *sql.DB

type User struct {
	id    int
	name  string
	phone string
}

type Queue struct {
	id  int
	day string
}

type QueuePos struct {
	Queue
	WorkTime string
	User
}

type Week struct {
	Date string
	ID   int
}

type UserExistes struct {
	When    time.Time
	Existes bool
}

func (u *UserExistes) Error() string {
	return fmt.Sprintf("at %s user existes %t", u.When, u.Existes)
}

func AddUser(ID int, name string) bool {
	result, err := DB.Query("SELECT id FROM users WHERE id=?", ID)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if result.Next() {
		return false
	}
	_, err = DB.Exec("INSERT INTO users(id, name) VALUES (?, ?)", ID, name)
	if err != nil {
		log.Fatal(err)
		return true
	}
	return true
}

func ListWeeks(weeks *[5]Week) {
	rows, err := DB.Query("SELECT dates, id FROM week LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var i = 0
	for rows.Next() {
		err = rows.Scan(weeks[i].Date, weeks[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		i++
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
