package main

import (
	"log"
	"task_manager_db/data"
	"task_manager_db/router"
)

func main() {
	if err := data.InitDB(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	r := router.SetupRouter()
	r.Run(":8080")
}
