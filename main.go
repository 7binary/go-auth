package main

import (
	"github.com/gin-gonic/gin"
	"go-auth/controllers"
	"go-auth/initializers"
	"go-auth/middlewares"
	"os"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()

	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.Login)
	router.GET("/user/profile", middlewares.CheckAuth, controllers.GetUserProfile)

	router.Run(":" + os.Getenv("PORT"))
}
