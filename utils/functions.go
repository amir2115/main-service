package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetCurrentTime() time.Time {
	utc := (3 * time.Hour) + (30 * time.Minute)
	return time.Now().Add(utc)
}

func GetParamUInt(c *gin.Context, param string) (bool, int) {
	paramToInt, err := strconv.ParseInt(c.Param(param), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messageCode": 403, "message": Messages[403]})
		return false, 0
	}
	return true, int(paramToInt)
}
