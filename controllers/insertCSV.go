package controllers

import (
	"MainService/dal"
	"MainService/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InsertCSV(c *gin.Context) {
	filepath := utils.CSV_ADDRESS
	gameHistoryList := dal.ReadCSVFile(filepath)
	dal.CreateAllGameHistories(gameHistoryList)
	gameSalesHistoriesDB := dal.GetAllGameHistories()
	c.JSON(http.StatusOK, gameSalesHistoriesDB)
}

func GetAllGameHistoriesByRank(c *gin.Context) {
	parsed, rank := utils.GetParamUInt(c, "rank")
	if !parsed {
		return
	}
	gameSalesHistoriesDB := dal.GetGameHistoryByRank(uint(rank))
	c.JSON(http.StatusOK, gameSalesHistoriesDB)
}

func GetAllGameHistoriesByPlatform(c *gin.Context) {
	parsed, length := utils.GetParamUInt(c, "length")
	if !parsed {
		return
	}
	gameSalesHistoriesDB := dal.GetAllGameHistoriesByPlatform(length)
	c.JSON(http.StatusOK, gameSalesHistoriesDB)
}

func GetAllGameHistoriesByYear(c *gin.Context) {
	parsed, length := utils.GetParamUInt(c, "length")
	if !parsed {
		return
	}
	parsed, year := utils.GetParamUInt(c, "year")
	if !parsed {
		return
	}
	gameSalesHistoriesDB := dal.GetAllGameHistoriesByYear(length, year)
	c.JSON(http.StatusOK, gameSalesHistoriesDB)
}

func GetAllGameHistoriesByGenre(c *gin.Context) {
	parsed, length := utils.GetParamUInt(c, "length")
	if !parsed {
		return
	}
	genre := c.Param("genre")
	gameSalesHistoriesDB := dal.GetAllGameHistoriesByGenre(length, genre)
	c.JSON(http.StatusOK, gameSalesHistoriesDB)
}

func GetTopFiveGamesByYear(c *gin.Context) {
	parsed, year := utils.GetParamUInt(c, "year")
	if !parsed {
		return
	}
	platform := c.Param("platform")
	gameSalesHistoriesDB := dal.GetTopFiveGamesByYear(platform, year)
	c.JSON(http.StatusOK, gameSalesHistoriesDB)
}

func NAVsEU(c *gin.Context) {
	gameSalesHistoriesDB := dal.NAVsEU()
	c.JSON(http.StatusOK, gameSalesHistoriesDB)
}

func SearchByName(c *gin.Context) {
	name := c.Param("name")
	gameSalesHistoriesDB := dal.SearchGameHistoriesByName(name)
	c.JSON(http.StatusOK, gameSalesHistoriesDB)
}
