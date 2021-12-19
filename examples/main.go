package main

import (
	"encoding/json"
	"fmt"
	// "flag"

	"io/ioutil"
	"net/http"

	"weather-dy/api"
)

type LiveWeather struct {
	Cityid     string `json:"cityid"`
	Date       string `json:"date"`
	Week       string `json:"week"`
	UpdateTime string `json:"update_time"`
	City       string `json:"city"`
	CityEn     string `json:"cityEn"`
	Country    string `json:"country"`
	CountryEn  string `json:"countryEn"`
	Wea        string `json:"wea"`
	WeaImg     string `json:"wea_img"`
	Tem        string `json:"tem"`
	Tem1       string `json:"tem1"`
	Tem2       string `json:"tem2"`
	Win        string `json:"win"`
	WinSpeed   string `json:"win_speed"`
	WinMeter   string `json:"win_meter"`
	Humidity   string `json:"humidity"`
	Visibility string `json:"visibility"`
	Pressure   string `json:"pressure"`
	Air        string `json:"air"`
	AirPm25    string `json:"air_pm25"`
	AirLevel   string `json:"air_level"`
	AirTips    string `json:"air_tips"`
}

func main() {

	url := api.Init()

	// fmt.Println("1. Performing Http Get...")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	// bodyString := string(bodyBytes)
	// fmt.Println("API Response as String:\n" + bodyString)

	var liveweather LiveWeather
	// LiveWeather := string(body)
	json.Unmarshal(bodyBytes, &liveweather)

	// fmt.Printf("API Response as struct %+v\n", liveweather)
	fmt.Printf("城市    :%s\n", liveweather.City)
	fmt.Printf("实时温度:%s°C\n", liveweather.Tem)
	fmt.Printf("空气质量:%s\n", liveweather.AirLevel)
	fmt.Printf("小贴士  :%s\n", liveweather.AirTips)
	fmt.Printf("天气    :%s\n", liveweather.Wea)

}
