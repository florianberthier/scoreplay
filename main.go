package main

import (
	"scoreplay/env"
	"scoreplay/server"

	"github.com/gin-gonic/gin"
)

func main() {
	env.Load()

	r := gin.Default()

	r.Use(gin.Recovery())
	server.SetupRouter(r)

	r.Run(":8080")
}
