package main

import (
	_ "beer_dmeo/model"
	"beer_dmeo/router"
	"log"
)

func main() {
	log.Println("main init")
	srv := router.Init()
	if err := srv.Run(":8080"); err != nil {
		panic(err)
	}
}
