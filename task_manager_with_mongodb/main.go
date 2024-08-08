package main

import (
	"log"
	"task_manager/router"
)

func main() {
	r := router.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err.Error())
	}
}
