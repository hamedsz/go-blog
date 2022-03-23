package application

import (
	"github.com/gin-gonic/gin"
	"github.com/hamedsz/go-blog/internal/database/mongodb"
	"github.com/hamedsz/go-blog/internal/routes"
)

type Service struct {}

func NewService() *Service {
	service := new(Service)
	return service
}

func (service *Service) Run()  {
	r := gin.Default()
	routes.LoadRoutes(r)
	mongodb.GetDB()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}