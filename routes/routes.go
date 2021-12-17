package routes

import (
	"MainService/controllers"
	"MainService/middleware"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/csv/", middleware.Authentication(), controllers.InsertCSV)
	r.GET("/game_rank/:rank/", middleware.Authentication(), controllers.GetAllGameHistoriesByRank)
	r.GET("/game_name/:name/", middleware.Authentication(), controllers.SearchByName)
	r.GET("/game_platform/:length/", middleware.Authentication(), controllers.GetAllGameHistoriesByPlatform)
	r.GET("/game_year/:year/:length/", middleware.Authentication(), controllers.GetAllGameHistoriesByYear)
	r.GET("/game_genre/:genre/:length/", middleware.Authentication(), controllers.GetAllGameHistoriesByGenre)
	r.GET("/top_five/:year/:platform/", middleware.Authentication(), controllers.GetTopFiveGamesByYear)
	r.GET("/na_eu_sales/", middleware.Authentication(), controllers.NAVsEU)
	return r
}
