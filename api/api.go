// Package api provides api
package api

import (
	cli "weather-dy/config"
)

func Init() string {
	config, err := cli.ParseFlags()
	if err != nil {
		panic(err)
	}

	var (
		WeatherAPI = "https://tianqiapi.com/api?" +
			"unescape=" + config.Unescape + "&" +
			"version=" + config.Version + "&" +
			"appid=" + config.Appid + "&" +
			"appsecret=" + config.Appsecret + "&" +
			"city=" + config.City
	)

	return WeatherAPI

}
