package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hamedsz/go-blog/routes"
	_ "github.com/hamedsz/go-blog/docs"
)

func main() {
	r := gin.Default()
	routes.LoadRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}