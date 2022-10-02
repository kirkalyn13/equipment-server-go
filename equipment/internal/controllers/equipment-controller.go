package controllers

import (
	"fmt"
	"kirky/equipment-server/equipment/internal/config"
	"kirky/equipment-server/equipment/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var equipment struct {
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
	Certificate       string
	Image             string
}

// Fetch All Equipment
func GetEquipment(c *gin.Context) {
	var equipments []models.Equipment

	config.DB.Unscoped().Find(&equipments)

	c.JSON(http.StatusOK, gin.H{
		"equipment": equipments,
	})
}

// Fetch Equipment By ID
func GetEquipmentById(c *gin.Context) {
	var equipment models.Equipment

	id := c.Param("id")

	config.DB.Unscoped().Find(&equipment, id)

	c.JSON(http.StatusOK, gin.H{
		"equipment": equipment,
	})
}

// Add New Equipment
func AddEquipment(c *gin.Context) {
	c.Bind(&equipment)

	newEquipment := models.Equipment{
		Name:              equipment.Name,
		Type:              equipment.Type,
		Model:             equipment.Model,
		Serial:            equipment.Serial,
		Description:       equipment.Description,
		Brand:             equipment.Brand,
		Price:             equipment.Price,
		Manufacturer:      equipment.Manufacturer,
		Expiration:        equipment.Expiration,
		PurchaseDate:      equipment.PurchaseDate,
		CalibrationDate:   equipment.CalibrationDate,
		CalibrationMethod: equipment.CalibrationMethod,
		NextCalibration:   equipment.NextCalibration,
		Location:          equipment.Location,
		IssuedBy:          equipment.IssuedBy,
		IssuedTo:          equipment.IssuedTo,
		Remarks:           equipment.Remarks,
		Status:            equipment.Status,
		Certificate:       equipment.Certificate,
		Image:             equipment.Image,
	}

	result := config.DB.Unscoped().Create(&newEquipment)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"equipment": equipment,
	})
}

// Add New Equipment
func EditEquipment(c *gin.Context) {
	id := c.Param("id")

	c.Bind(&equipment)

	var toEdit models.Equipment
	config.DB.First(&toEdit, id)

	config.DB.Model(&toEdit).Updates(models.Equipment{
		Name:              equipment.Name,
		Type:              equipment.Type,
		Model:             equipment.Model,
		Serial:            equipment.Serial,
		Description:       equipment.Description,
		Brand:             equipment.Brand,
		Price:             equipment.Price,
		Manufacturer:      equipment.Manufacturer,
		Expiration:        equipment.Expiration,
		PurchaseDate:      equipment.PurchaseDate,
		CalibrationDate:   equipment.CalibrationDate,
		CalibrationMethod: equipment.CalibrationMethod,
		NextCalibration:   equipment.NextCalibration,
		Location:          equipment.Location,
		IssuedBy:          equipment.IssuedBy,
		IssuedTo:          equipment.IssuedTo,
		Remarks:           equipment.Remarks,
		Status:            equipment.Status,
		Certificate:       equipment.Certificate,
		Image:             equipment.Image,
	})

	c.JSON(http.StatusOK, gin.H{
		"equipment": equipment,
	})

}

// Delete Equipment By ID
func DeleteEquipmentById(c *gin.Context) {
	id := c.Param("id")

	config.DB.Delete(&models.Equipment{}, id)

	c.Status(http.StatusOK)
}

// Download Certificate
func DownloadCertificate(c *gin.Context) {
	var cert models.Certificate

	idx := c.Param("idx")

	config.DB.Find(&models.Equipment{}, idx)
	config.DB.Model(&models.Equipment{}).Find(&cert)

	filename := fmt.Sprintf("certificate_%v.pdf", idx)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.Writer.Write([]byte(cert.Certificate))

	c.JSON(http.StatusOK, gin.H{})
}
