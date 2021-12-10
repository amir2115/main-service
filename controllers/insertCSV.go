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
	c.JSON(http.StatusBadRequest, gameSalesHistoriesDB)
}
