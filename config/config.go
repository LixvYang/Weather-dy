package config

import (
	"flag"
)

type Config struct {
	City string
	AirTips bool
	Unescape string
	Version string
	Appid	string 
	Appsecret string
}

func ParseFlags() (Config, error) {
	config := Config{}
	flag.StringVar(&config.City, "city", "", "Config your city.")
	flag.BoolVar(&config.AirTips, "tip" , false, "Air tips")
	flag.StringVar(&config.Unescape, "unescape", "1", "Unescape")
	flag.StringVar(&config.Version, "version", "v6", "Version")
	flag.StringVar(&config.Appid, "appid", "51369129", "Appid")
	flag.StringVar(&config.Appsecret, "appsecret", "fzfQ7Fnf", "Appsecret")


	

	flag.Parse()
	return config, nil
}
