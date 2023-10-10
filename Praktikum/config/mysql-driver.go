package config

import (
	"fmt"
	"project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&tls=true",
		appConfig.DBUSER,
		appConfig.DBPASS,
		appConfig.DBHOST,
		appConfig.DBPORT,
		appConfig.DBNAME,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
}
