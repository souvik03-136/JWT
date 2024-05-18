package initializers

import "github.com/souvik03-136/JWT/models"

func SyncDatabase() {

	DB.AutoMigrate(&models.User{})

}
