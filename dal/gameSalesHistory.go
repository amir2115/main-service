package dal

import (
	"MainService/config"
	"MainService/dao"
	"MainService/utils"
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

func GetGameHistoryByRank(rank uint) []dao.GameSalesHistory {
	var gameSalesHistories []dao.GameSalesHistory
	config.DB.Where(&dao.GameSalesHistory{Rank: rank}).Find(&gameSalesHistories)
	return gameSalesHistories
}

func GetAllGameHistoriesByPlatform(length int) []dao.GameSalesHistory {
	var gameSalesHistories []dao.GameSalesHistory
	for _, platform := range utils.PlatformNames {
		var temp []dao.GameSalesHistory
		config.DB.Limit(length).Offset(0).Where("platform = ?", platform).Order("`rank`").Find(&temp)
		gameSalesHistories = append(gameSalesHistories, temp...)
	}
	return gameSalesHistories
}

func GetAllGameHistoriesByYear(length int, year int) []dao.GameSalesHistory {
	var gameSalesHistories []dao.GameSalesHistory
	config.DB.Limit(length).Offset(0).Where("year = ?", year).Order("`rank`").Find(&gameSalesHistories)
	return gameSalesHistories
}

func GetAllGameHistoriesByGenre(length int, genre string) []dao.GameSalesHistory {
	var gameSalesHistories []dao.GameSalesHistory
	config.DB.Limit(length).Offset(0).Where("genre = ?", genre).Order("`rank`").Find(&gameSalesHistories)
	return gameSalesHistories
}

func GetTopFiveGamesByYear(platform string, year int) []dao.GameSalesHistory {
	var gameSalesHistories []dao.GameSalesHistory
	config.DB.Limit(5).Offset(0).Where("year = ? AND platform = ?", year, platform).Order("global_sales desc").Find(&gameSalesHistories)
	return gameSalesHistories
}

func NAVsEU() []dao.GameSalesHistory {
	var gameSalesHistories []dao.GameSalesHistory
	config.DB.Where("eu_sales > na_sales").Find(&gameSalesHistories)
	return gameSalesHistories
}

func SearchGameHistoriesByName(name string) []dao.GameSalesHistory {
	var gameSalesHistories []dao.GameSalesHistory
	config.DB.Where("name LIKE ?", "%"+name+"%").Find(&gameSalesHistories)
	return gameSalesHistories
}
