package config

import (
	"flag"
)

type Config struct {
	City string
	Unescape string
	Version string
	Appid	string 
	Appsecret string
}

func ParseFlags() (Config, error) {
	// help := flag.Bool("h", false, "Display Help")
	
	
	// if *help {
	// 	fmt.Println("Out a help")
	// 	fmt.Println("")
	// 	fmt.Println("Usage:...")
	// 	flag.PrintDefaults()
	// }
	config := Config{}
	flag.StringVar(&config.City, "city", "", "Config your city.")
	flag.StringVar(&config.Unescape, "unescape", "1", "Unescape")
	flag.StringVar(&config.Version, "version", "v6", "Version")
	flag.StringVar(&config.Appid, "appid", "51369129", "Appid")
	flag.StringVar(&config.Appsecret, "appsecret", "fzfQ7Fnf", "Appsecret")


	

	flag.Parse()
	return config, nil
}
