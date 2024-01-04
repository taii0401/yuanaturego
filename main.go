package main

import (
	router "yuanaturego/routers"
)

func main() {
	router := router.InitRouter()
	//載入CSS、JS
	router.Static("/static", "./static")

	router.Run(":8099")
}
