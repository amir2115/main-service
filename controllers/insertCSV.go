package controllers

import (
	"MainService/dal"
	"MainService/dao"
	"MainService/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InsertCSV(c *gin.Context) {
	filepath := utils.CSV_ADDRESS
	gameHistoryList := dal.ReadCSVFile(filepath)
	dal.CreateAllGameHistories(gameHistoryList)
	gameSalesHistoriesDB := dal.GetAllGameHistories()
	c.JSON(http.StatusBadRequest, gameSalesHistoriesDB)
}

func GetAllGameHistoriesByRank(c *gin.Context) {
	parsed, rank := utils.GetParamUInt(c, "rank")
	if !parsed {
		return
	}
	gameSalesHistoriesDB := dal.GetAllGameHistoriesByRank(uint(rank))
	c.JSON(http.StatusBadRequest, gameSalesHistoriesDB)
}

func GetAllGameHistoriesByPlatform(c *gin.Context) {
	parsed, length := utils.GetParamUInt(c, "length")
	if !parsed {
		return
	}
	gameSalesHistoriesDB := dal.GetAllGameHistoriesByPlatform(length)
	c.JSON(http.StatusBadRequest, gameSalesHistoriesDB)
}

func GetAllGameHistoriesByYear(c *gin.Context) {
	parsed, length := utils.GetParamUInt(c, "length")
	if !parsed {
		return
	}
	gameSalesHistoriesDB := dal.GetAllGameHistoriesByYear(length)
	c.JSON(http.StatusBadRequest, gameSalesHistoriesDB)
}

func GetAllGameHistoriesByCategory(c *gin.Context) {
	parsed, length := utils.GetParamUInt(c, "length")
	if !parsed {
		return
	}
	gameSalesHistoriesDB := dal.GetAllGameHistoriesByCategory(length)
	c.JSON(http.StatusBadRequest, gameSalesHistoriesDB)
}

func GetAllGameHistoriesByGlobalSales(c *gin.Context) {
	var input dao.GlobalSales
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 403, "message": utils.Messages[403], "error": err.Error()})
		return
	}
	gameSalesHistoriesDB := dal.GetAllGameHistoriesByGlobalSales(input.Platform, input.Year)
	c.JSON(http.StatusBadRequest, gameSalesHistoriesDB)
}

func GetAllGameHistoriesByNaEuSales(c *gin.Context) {
	gameSalesHistoriesDB := dal.GetAllGameHistoriesByNaEuSales()
	c.JSON(http.StatusBadRequest, gameSalesHistoriesDB)
}

func SearchByName(c *gin.Context) {
	var input dao.SearchByName
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 403, "message": utils.Messages[403], "error": err.Error()})
		return
	}
	gameSalesHistoriesDB := dal.SearchGameHistoriesByName(input.Name)
	c.JSON(http.StatusBadRequest, gameSalesHistoriesDB)
}
