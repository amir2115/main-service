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
	r.GET("/rank/:rank/", middleware.Authentication(), controllers.GetAllGameHistoriesByRank)
	r.POST("/search_by_name/", middleware.Authentication(), controllers.SearchByName)
	r.GET("/platform/:length/", middleware.Authentication(), controllers.GetAllGameHistoriesByPlatform)
	r.GET("/year/:length/", middleware.Authentication(), controllers.GetAllGameHistoriesByYear)
	r.GET("/category/:length/", middleware.Authentication(), controllers.GetAllGameHistoriesByCategory)
	r.POST("/global_sales/", middleware.Authentication(), controllers.GetAllGameHistoriesByGlobalSales)
	r.GET("/na_eu_sales/", middleware.Authentication(), controllers.GetAllGameHistoriesByNaEuSales)
	return r
}
