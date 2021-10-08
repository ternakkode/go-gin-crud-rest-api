package user

import (
	"strings"

	"github.com/ternakkode/go-gin-crud-rest-api/utils/res"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated int64  `json:"date_created"`
}

func (user *User) Validate() *res.Err {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return res.NewRestErr(400, "invalid email", "bad_request")
	}

	return nil
}
