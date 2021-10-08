package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ternakkode/go-gin-crud-rest-api/domain/user"
	"github.com/ternakkode/go-gin-crud-rest-api/service"
	"github.com/ternakkode/go-gin-crud-rest-api/utils/res"
)

func CreateUser(c *gin.Context) {
	var user *user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("invalid json input ", err)
		c.JSON(http.StatusBadRequest, res.NewRestErr(http.StatusBadRequest, "invalid json input", "bad_request"))
		return
	}

	result, saveErr := service.UserService.CreateUser(user)
	if saveErr != nil {
		log.Println("failed create user", saveErr.Message)
		c.JSON(saveErr.Code, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func FindUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		log.Println("failed parse user id from user")
		c.JSON(http.StatusBadRequest, res.NewRestErr(http.StatusBadRequest, userErr.Error(), "bad_request"))
		return
	}

	user, getUserErr := service.UserService.FindUser(&userId)
	if getUserErr != nil {
		c.JSON(getUserErr.Code, getUserErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	users, getUsersErr := service.UserService.GetUser()
	if getUsersErr != nil {
		c.JSON(getUsersErr.Code, getUsersErr)
	}

	c.JSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		log.Println("failed parse user id from user")
		c.JSON(http.StatusBadRequest, res.NewRestErr(http.StatusBadRequest, userErr.Error(), "bad_request"))
		return
	}

	var user *user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("invalid json input ", err)
		c.JSON(http.StatusBadRequest, res.NewRestErr(http.StatusBadRequest, "invalid json input", "bad_request"))
		return
	}

	user.Id = userId

	res, err := service.UserService.UpdateUser(user)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		log.Println("failed parse user id from user")
		c.JSON(http.StatusBadRequest, res.NewRestErr(http.StatusBadRequest, userErr.Error(), "bad_request"))
		return
	}

	if err := service.UserService.DeleteUser(&userId); err != nil {
		log.Println("failed to delete user data")
		c.JSON(http.StatusBadRequest, res.NewRestErr(http.StatusBadRequest, err.Error, "bad_request"))
	}

	c.JSON(http.StatusOK, map[string]string{"deleted": "ok"})
}
