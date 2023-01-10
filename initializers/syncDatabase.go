package initializers

import "github/loa212/go-todo/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Todo{}, &models.User{})
}