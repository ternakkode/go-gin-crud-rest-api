package user

// TODO : add id, email validation when create

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	db "github.com/ternakkode/go-gin-crud-rest-api/datasource/mysql"
	"github.com/ternakkode/go-gin-crud-rest-api/utils/res"
)

const (
	queryInsertUser = "INSERT INTO users (firstname, lastname, email, created_at) VALUES (?, ?, ?, ?)"
	queryFindAll    = "SELECT * FROM users"
	queryGetById    = "SELECT id, firstname, lastname, email, created_at FROM users WHERE id=?"
	queryUpdate     = "UPDATE users SET firstname = ?, lastname = ?, email = ? WHERE id=?"
	queryDelete     = "DELETE FROM users WHERE id=?"
	errorNoRows     = "no rows in result set"
)

func GetAll() ([]User, *res.Err) {
	statement, err := db.Mysql.Prepare(queryFindAll)
	if err != nil {
		return nil, res.NewRestErr(http.StatusInternalServerError, err.Error(), "internal_server_error")
	}

	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		return nil, res.NewRestErr(http.StatusNotFound, err.Error(), "not found")
	}

	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
			return nil, res.NewRestErr(http.StatusInternalServerError, err.Error(), "internal_server_error")
		}

		users = append(users, user)
	}

	return users, nil
}

func (user *User) Get() *res.Err {
	statement, err := db.Mysql.Prepare(queryGetById)
	if err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return res.NewRestErr(http.StatusNotFound, fmt.Sprintf("user with id %d not found", user.Id), err.Error())
		}

		return res.NewRestErr(http.StatusInternalServerError, err.Error(), "internal_server_error")
	}

	defer statement.Close()

	result := statement.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return res.NewRestErr(http.StatusInternalServerError, err.Error(), "internal_server_error")
	}

	return nil
}

func (user *User) Save() *res.Err {
	statement, err := db.Mysql.Prepare(queryInsertUser)
	if err != nil {
		return res.NewRestErr(http.StatusInternalServerError, err.Error(), "internal_server_error")
	}

	defer statement.Close()

	user.DateCreated = time.Now().Unix()
	result, err := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return res.NewRestErr(http.StatusInternalServerError, err.Error(), "internal_server_error")
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return res.NewRestErr(http.StatusInternalServerError, err.Error(), "internal_server_error")
	}

	user.Id = userId

	return nil
}

func (user *User) Update() *res.Err {
	statement, err := db.Mysql.Prepare(queryUpdate)
	if err != nil {
		return res.NewRestErr(http.StatusInternalServerError, err.Error(), "internal_server_error")
	}

	defer statement.Close()
	_, err = statement.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return res.NewRestErr(http.StatusInternalServerError, err.Error(), "internal_server_error")
	}

	return nil
}

func (user *User) Delete() *res.Err {
	statement, err := db.Mysql.Prepare(queryDelete)
	if err != nil {
		return res.NewRestErr(http.StatusInternalServerError, err.Error(), "internal_server_error")
	}

	defer statement.Close()
	_, err = statement.Exec(user.Id)
	if err != nil {
		return res.NewRestErr(http.StatusInternalServerError, err.Error(), "internal_server_error")
	}

	return nil
}
