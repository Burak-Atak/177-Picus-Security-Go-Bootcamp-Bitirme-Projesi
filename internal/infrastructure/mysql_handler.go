package infrastructure

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySqlDB(conString string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Cannot connect to database : %s", err.Error()))
	}

	return db
}
