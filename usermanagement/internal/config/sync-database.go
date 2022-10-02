package config

import "kirky/equipment-server/usermanagement/internal/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
