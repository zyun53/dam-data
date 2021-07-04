package main

import (
	"github.com/zyun-i/dam-data/app"
	"github.com/zyun-i/dam-data/config"
	//	"github.com/zyun-i/dam-data/dammap"
)

func main() {

	//	dammap.Run()

	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
