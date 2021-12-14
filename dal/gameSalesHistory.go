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

func GetAllGameHistoriesByRank(rank uint) []dao.GameSalesHistory {
	var gameSalesHistories []dao.GameSalesHistory
	config.DB.Where(&dao.GameSalesHistory{Rank: rank}).Find(&gameSalesHistories)
	return gameSalesHistories
}

func GetAllGameHistoriesByPlatform(length int) []dao.GameSalesHistory {
	var gameSalesHistories []dao.GameSalesHistory
	config.DB.Limit(length).Offset(0).Order("platform desc").Find(&gameSalesHistories)
	return gameSalesHistories
}

func GetAllGameHistoriesByYear(length int) []dao.GameSalesHistory {
	var gameSalesHistories []dao.GameSalesHistory
	config.DB.Limit(length).Offset(0).Order("year desc").Find(&gameSalesHistories)
	return gameSalesHistories
}

func GetAllGameHistoriesByCategory(length int) []dao.GameSalesHistory {
	var gameSalesHistories []dao.GameSalesHistory
	config.DB.Limit(length).Offset(0).Order("genre").Find(&gameSalesHistories)
	return gameSalesHistories
}

func GetAllGameHistoriesByGlobalSales(platform string, year int) []dao.GameSalesHistory {
	var gameSalesHistories []dao.GameSalesHistory
	config.DB.Limit(5).Offset(0).Where(&dao.GameSalesHistory{Platform: platform, Year: year}).Order("global_sales desc").Find(&gameSalesHistories)
	return gameSalesHistories
}

func GetAllGameHistoriesByNaEuSales() []dao.GameSalesHistory {
	var gameSalesHistories []dao.GameSalesHistory
	config.DB.Limit(5).Offset(0).Order("global_sales desc").Find(&gameSalesHistories)
	return gameSalesHistories
}

func SearchGameHistoriesByName(name string) []dao.GameSalesHistory {
	var gameSalesHistories []dao.GameSalesHistory
	config.DB.Where("name LIKE ?", "%"+name+"%").Find(&gameSalesHistories)
	return gameSalesHistories
}
