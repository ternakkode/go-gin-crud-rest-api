package service

import (
	"github.com/ternakkode/go-gin-crud-rest-api/domain/user"
	"github.com/ternakkode/go-gin-crud-rest-api/utils/res"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	CreateUser(*user.User) (*user.User, *res.Err)
	GetUser() ([]user.User, *res.Err)
	FindUser(*int64) (*user.User, *res.Err)
	UpdateUser(*user.User) (*user.User, *res.Err)
	DeleteUser(*int64) *res.Err
}

func (u *userService) CreateUser(userReq *user.User) (userRes *user.User, err *res.Err) {
	if err := userReq.Validate(); err != nil {
		return nil, err
	}

	if err := userReq.Save(); err != nil {
		return nil, err
	}

	return userReq, nil
}

func (u *userService) GetUser() ([]user.User, *res.Err) {
	result, err := user.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userService) FindUser(userId *int64) (userRes *user.User, err *res.Err) {
	result := &user.User{Id: *userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userService) UpdateUser(userReq *user.User) (userRes *user.User, err *res.Err) {
	current, err := UserService.FindUser(&userReq.Id)
	if err != nil {
		return nil, err
	}

	current.FirstName = userReq.FirstName
	current.LastName = userReq.LastName
	current.Email = userReq.Email

	if err = current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (u *userService) DeleteUser(userId *int64) *res.Err {
	user := &user.User{Id: *userId}
	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}
