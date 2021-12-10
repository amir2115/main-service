package dal

import (
	"MainService/config"
	"MainService/dao"
	"fmt"
	"strconv"
)

func CreateAllGameHistories(gameSalesHistories []dao.GameSalesHistory) {
	for i, history := range gameSalesHistories {
		config.DB.Create(&history)
		fmt.Println(strconv.Itoa(i) + "/" + strconv.Itoa(len(gameSalesHistories)))
	}
}

func GetAllGameHistories() []dao.GameSalesHistory {
	var gameSalesHistories []dao.GameSalesHistory
	config.DB.Find(&gameSalesHistories)
	return gameSalesHistories
}
