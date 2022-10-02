package app

import (
	"kirky/equipment-server/usermanagement/internal/config"
	"kirky/equipment-server/usermanagement/internal/controllers"
	"kirky/equipment-server/usermanagement/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RunApp() {
	r := gin.Default()

	r.POST("/user/signup", controllers.Signup)
	r.POST("/user/login", controllers.Login)
	r.PUT("/user/edit/:id", controllers.EditUser)
	r.GET("/user/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
	config.SyncDatabase()
}
