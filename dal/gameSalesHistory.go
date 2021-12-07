package dal

import (
	"MainService/config"
	"MainService/dao"
)

func CreateAllGameHistories(gameSalesHistories []dao.GameSalesHistory) {
	for _, history := range gameSalesHistories {
		config.DB.Create(&history)
	}
}

func GetAllGameHistories() []dao.GameSalesHistory {
	var gameSalesHistories []dao.GameSalesHistory
	config.DB.Find(&gameSalesHistories)
	return gameSalesHistories
}
