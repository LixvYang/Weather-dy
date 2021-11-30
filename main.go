// Package main provides main
package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "weather"
	app.Usage = "An application showing the weather system."
	app.Version = "v0.0.1"
	app.Flags = []cli.Flag {
			&cli.StringFlag{
				Name:	"local",
				Aliases:[]string{"l"},
				Value:	"上海",
				Usage: "Your location",
			},
	} 

	app.EnableBashCompletion = true

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}

}