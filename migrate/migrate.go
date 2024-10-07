package main

import (
	"go-auth/initializers"
	"go-auth/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
