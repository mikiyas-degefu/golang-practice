package main

import (
	"task_manager/router"
)

func main() {
	r := router.SetupRouter()
	// default listen on :8080
	r.Run(":8080")
}
