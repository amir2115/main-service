package config

import (
	"MainService/dao"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(mysql.Open(DbURL(BuildDBConfig())), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(
		&dao.GameSalesHistory{},
	)

}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("game_history_var"),
		Port:     3306,
		User:     "root",
		Password: "qweqwe",
		DBName:   "game_history",
	}
	return &dbConfig
}
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
