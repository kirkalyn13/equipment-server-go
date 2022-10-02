package controllers

import (
	"fmt"
	"kirky/equipment-server/equipment/internal/config"
	"kirky/equipment-server/equipment/internal/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var changelog struct {
	ID                int
	Name              string
	Type              string
	Model             string
	Serial            string
	Description       string
	Brand             string
	Price             string
	Manufacturer      string
	Expiration        string
	PurchaseDate      string
	CalibrationDate   string
	CalibrationMethod string
	NextCalibration   string
	Location          string
	IssuedBy          string
	IssuedTo          string
	Remarks           string
	Status            string
	ModifiedBy        string
}

// Fetch All Change Logs
func GetChangelogsById(c *gin.Context) {
	var changelogs []models.Changelog

	id := c.Param("id")
	config.DB.Unscoped().Find(&changelogs, id)

	c.JSON(http.StatusOK, gin.H{
		"changelogs": changelogs,
	})
}

// Add New Change Log
func AddChangelog(c *gin.Context) {
	c.Bind(&changelog)

	id := c.Param("id")
	logId, err := strconv.Atoi(id)

	if err != nil {
		log.Fatal("Unable to convert string ID to int.")
	}

	newChangelog := models.Changelog{
		ID:                logId,
		Name:              changelog.Name,
		Type:              changelog.Type,
		Model:             changelog.Model,
		Serial:            changelog.Serial,
		Description:       changelog.Description,
		Brand:             changelog.Brand,
		Price:             changelog.Price,
		Manufacturer:      changelog.Manufacturer,
		Expiration:        changelog.Expiration,
		PurchaseDate:      changelog.PurchaseDate,
		CalibrationDate:   changelog.CalibrationDate,
		CalibrationMethod: changelog.CalibrationMethod,
		NextCalibration:   changelog.NextCalibration,
		Location:          changelog.Location,
		IssuedBy:          changelog.IssuedBy,
		IssuedTo:          changelog.IssuedTo,
		Remarks:           changelog.Remarks,
		Status:            changelog.Status,
		ModifiedBy:        changelog.ModifiedBy,
	}

	result := config.DB.Omit("IndexNum").Omit("Timestamp").Create(&newChangelog)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"changelog": changelog,
	})
}

// Download Certificate by Changelogs
func DownloadCertificateByLog(c *gin.Context) {
	var cert models.Certificate

	id := c.Param("id")

	config.DB.Find(&models.Changelog{}, id)
	config.DB.Model(&models.Changelog{}).Find(&cert)

	filename := fmt.Sprintf("certificate_%v.pdf", id)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.Writer.Write([]byte(cert.Certificate))

	c.JSON(http.StatusOK, gin.H{})
}
