package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	//connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "havabaaruser", "xba54gj5y976&Dn", "localhost", 3306, "havabaar")
	database, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic(err.Error())
	}

	database.AutoMigrate()
	DB = database
}
