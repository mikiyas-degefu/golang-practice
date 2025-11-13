package main

import (
	"log"
	"task_manager_db/data"
	"task_manager_db/router"
)

func main() {
	// Initialize MongoDB connection. Reads MONGODB_URI and MONGODB_DB from env (defaults provided).
	if err := data.InitDB(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	r := router.SetupRouter()
	// default listen on :8080
	r.Run(":8080")
}
