package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"io"

	"github.com/gin-gonic/gin"
)

type Location struct {
	Country string `json:"country"`
	Places  []struct {
		City  string `json:"place name"`
		State string `json:"state"`
	} `json:"places"`
}

type Weather struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func HomeHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func WeatherHandler(c *gin.Context) {
	zipcode := c.Query("zipcode")
	weatherInfo := getWeather(zipcode)
	c.String(200, "%s", weatherInfo)
}

func getWeather(zipcode string) string {
	city, state := getCityAndState(zipcode)
	temp := getTemperature(city, state)
	return fmt.Sprintf("The weather in %s, %s is %.2f degrees Fahrenheit.\n", city, state, temp)
}

func getCityAndState(zipcode string) (string, string) {
	resp, err := http.Get(fmt.Sprintf("http://api.zippopotam.us/us/%s", zipcode))
	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		fmt.Println(err)
		return "", ""
	}

	var location Location
	json.Unmarshal(body, &location)

	return location.Places[0].City, location.Places[0].State
}

func getTemperature(city, state string) float64 {
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s,%s,us&units=imperial&appid=%s", city, state, apiKey))
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	var weather Weather
	json.Unmarshal(body, &weather)

	return weather.Main.Temp
}
