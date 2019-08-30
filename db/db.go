package db

import (
	"fmt"
	"github.com/golang-gin/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Init() {
	c := config.GetConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		c.GetString("dbserver"), c.GetInt("dbport"), c.GetString("dbusername"), c.GetString("dbname"), c.GetString("dbpassword"))
	var err error
	db, err := gorm.Open("postgres",psqlInfo)

	if err != nil {
		fmt.Println("db err: ", err)
	}
	DB =db
}

func GetDB() *gorm.DB {
	 return DB
}