package main

import (
	"github.com/gin-gonic/gin"
	"github.com/souvik03-136/JWT/controllers"
	"github.com/souvik03-136/JWT/initializers"
	"github.com/souvik03-136/JWT/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()

}
