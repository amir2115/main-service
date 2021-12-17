package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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

func GetToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.Request.Header.Get("Authorization")
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 {
		return "", errors.New(Messages[401])
	}
	return tokenParts[1], nil
}
