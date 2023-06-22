package initializers

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var dsn = "root:himanshu123@@@tcp(localhost:3306)/empgolang?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	//gorm.Open(mysql.Open(urlDSN), &gorm.Config{})

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Sorry we are not able to conected")
	}

}
