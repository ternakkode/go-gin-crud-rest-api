package app

import "github.com/ternakkode/go-gin-crud-rest-api/controller"

func mapUrls() {
	router.GET("/ping", controller.Ping)

	router.GET("/users", controller.GetUser)
	router.GET("users/:id", controller.FindUser)
	router.POST("users", controller.CreateUser)
	router.PUT("users/:id", controller.UpdateUser)
	router.DELETE("users/:id", controller.DeleteUser)
}
