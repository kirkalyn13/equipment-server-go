package app

import (
	"kirky/equipment-server/equipment/internal/config"
	"kirky/equipment-server/equipment/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RunApp() {
	r := gin.Default()

	// Equipment APIs
	r.GET("/equipment/all", controllers.GetEquipment)
	r.GET("/equipment/:id", controllers.GetEquipmentById)
	r.POST("/equipment/add", controllers.AddEquipment)
	r.PUT("/equipment/edit/:id", controllers.EditEquipment)
	r.DELETE("/equipment/delete/:id", controllers.DeleteEquipmentById)
	r.GET("/equipment/certificate/:id", controllers.DownloadCertificate)

	// Changelog APIs
	r.GET("/changelogs/:id", controllers.GetChangelogsById)
	r.POST("/changelogs/add/:id", controllers.AddChangelog)
	r.GET("/changelogs/certificate/:idx", controllers.DownloadCertificateByLog)

	// Run Service
	r.Run() // listen and serve on localhost:3005
}

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}
