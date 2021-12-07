package controllers

import (
	"MainService/dal"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InsertCSV(c *gin.Context) {
	filepath := "assets/vgsales.csv"
	gameHistoryList := dal.ReadCSVFile(filepath)
	dal.CreateAllGameHistories(gameHistoryList)
	gameSalesHistoriesDB := dal.GetAllGameHistories()
	c.JSON(http.StatusBadRequest, gameSalesHistoriesDB)
}
