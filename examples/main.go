package main

import (
	"beer_dmeo/router"
)

func main() {
	srv := router.Init()
	if err := srv.Run(":8080"); err != nil {
		panic(err)
	}
}
