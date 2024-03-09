package routes

import (
	"github.com/devopsjester/weather-app-demo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Add your code here
	router.GET("/", controllers.HomeHandler)
	router.GET("/weather/zipcode/:zipcode", controllers.WeatherHandler)
	router.GET("/weather/city/:city/state/:state", controllers.WeatherHandler)
}
