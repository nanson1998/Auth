package main

import (
	"jwt-todo/helper/redis"
	"jwt-todo/router"
	"log"
)

func main() {
	redis.ConnectRd()
	r := router.SetupRouter()
	log.Fatal(r.Run(":8080"))
	//mysql.Connect()

}
