package models

import (
	"errors"

	"github.com/JuanDavidLC/Go_Rest_Api/database"
)

type User struct {
	id_user   int
	Name      string
	Last_name string
}

func CreateUser(user *User) error {

	q := `INSERT INTO 
			users (name,last_name)
			VALUES($1,$2)`

	db := database.Connect()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(user.Name, user.Last_name)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {

		return errors.New("Error: More than one row was affected")

	}

	return nil

}

func GetUsers() ([]User, error) {

	q := `
		SELECT
			users.id_user,  
			users.name,
			users.last_name
		FROM public.users`

	db := database.Connect()
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}

	users := []User{}
	for rows.Next() {

		user := User{}

		err = rows.Scan(&user.id_user, &user.Name, &user.Last_name)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}
