package models

import (
	"database/sql"
)

var DB *sql.DB

type User struct {
	id int
	name string
	phone string
}

type Queue struct {
	id int
	day string
}

type QueuePos struct {
	Queue
	WorkTime string
	User
}

func (db DB) AddUser(ID, name string) error {


}