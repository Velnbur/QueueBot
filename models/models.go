package models

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"log"
)

var DB *sql.DB

type User struct {
	ID    sql.NullInt32
	Name  sql.NullString
	Phone sql.NullString
}

type Week struct {
	Date string
	ID   int
}

type Day struct {
	ID     int
	WeekID int
	Data   string
}

type Queue struct {
	ID   int
	Day  Day
	Time string
	User User
}

type NullString string

func (s *NullString) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	strVal, ok := value.(string)
	if !ok {
		return errors.New("Column is not a string")
	}
	*s = NullString(strVal)
	return nil
}

func (s NullString) Value() (driver.Value, error) {
	if len(s) == 0 { // if nil or empty string
		return nil, nil
	}
	return string(s), nil
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
	_, err = DB.Exec("INSERT INTO users(id, username) VALUES (?, ?)", ID, name)
	if err != nil {
		log.Fatal(err)
		return true
	}
	return true
}

func ListWeeks(weeks *[5]Week) {
	rows, err := DB.Query("SELECT dates, id FROM weeks LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var i = 0
	for rows.Next() {
		err = rows.Scan(&weeks[i].Date, &weeks[i].ID)
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

func ListDays(ID string, days *[6]Day) {
	rows, err := DB.Query("SELECT num, id FROM days WHERE week_id=?", ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var i = 0
	for rows.Next() {
		err = rows.Scan(&days[i].Data, &days[i].ID)
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

func ListQueue(ID string, times *[]Queue) {
	rows, err := DB.Query("SELECT days.*, t.id, t.time_pos, users.id, users.username FROM times t LEFT JOIN days ON days.id=day_id LEFT JOIN users ON users.id=user_id WHERE day_id=?;", ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var q Queue
		err = rows.Scan(
			&q.Day.ID,
			&q.Day.WeekID,
			&q.Day.Data,
			&q.ID,
			&q.Time,
			&q.User.ID,
			&q.User.Name,
		)
		if err != nil {
			log.Fatal(err)
		}
		*times = append(*times, q)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
