package persistant

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

}
