package main

import (
	"encoding/json"
	"fmt"
	// "flag"

	"log"
	"io/ioutil"
	"net/http"

	api "weather-dy/api"
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
	// help := flag.Bool("h", false, "Display Help")
	
	
	// if *help {
	// 	fmt.Println("Out a help")
	// 	fmt.Println("")
	// 	fmt.Println("Usage:...")
	// 	flag.PrintDefaults()
	// }
	// flag.Parse()


	url := api.Init()

	fmt.Println("1. Performing Http Get...")
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println("API Response as String:\n" + bodyString)

	var liveweather LiveWeather
	// LiveWeather := string(body)
	json.Unmarshal(bodyBytes, &liveweather)
	
	log.Printf("API Response as struct %+v\n", liveweather)
	log.Printf("你的城市是:%s\n",liveweather.City)

	// if config.AirTips {
	// 	fmt.Printf("小贴士：%s\n", liveweather.AirTips)
	// }
}
