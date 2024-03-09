package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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
	zipcode := c.Param("zipcode")
	city := c.Param("city")
	state := c.Param("state")

	var weatherInfo string
	if zipcode != "" {
		weatherInfo = getWeatherByZipcode(zipcode)
	} else if city != "" && state != "" {
		weatherInfo = getWeatherByCityAndState(city, state)
	} else {
		c.String(http.StatusBadRequest, "Invalid parameters. Please provide either a zipcode or a city and state.")
		return
	}

	c.String(200, "%s", weatherInfo)
}

func getWeatherByZipcode(zipcode string) string {
	city, state := getCityAndState(zipcode)
	return getWeather(city, state)
}

func getWeatherByCityAndState(city, state string) string {
	return getWeather(city, state)
}

func getWeather(city, state string) string {
	temp := getTemperature(city, state)
	return fmt.Sprintf("The weather in %s, %s is %.2f degrees Fahrenheit.\n", title(city), title(state), temp)
}

func title(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	title, _, _ := transform.String(t, s)
	return strings.Title(title)
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

	city := location.Places[0].City
	state := location.Places[0].State

	return city, state
}

func getTemperature(city, state string) float64 {
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	city = url.QueryEscape(city)
	state = url.QueryEscape(state)
	resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s,%s,us&units=imperial&appid=%s", city, state, apiKey))
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		fmt.Println(err)
		return 0
	}

	var weather Weather
	json.Unmarshal(body, &weather)

	return weather.Main.Temp
}
