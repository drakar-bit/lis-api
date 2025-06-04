package main

import (
	"github.com/drakar-bit/lis-api/internal/interface/graphql"
	"github.com/gin-gonic/gin"
)

func main() {
	// Setting up Gin
	r := gin.Default()
	r.POST("/query", graphql.GraphqlHandler())
	r.GET("/", graphql.PlaygroundHandler())
	r.Run()
}
