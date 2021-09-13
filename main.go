package main

import (
	"kalika-be/config"
	"kalika-be/router"
)

func main()  {
	r := router.Routes()

	err := r.Start(":" + config.Get("APP_PORT").String())
	if err != nil {
		r.Logger.Fatal(err)
	}
}