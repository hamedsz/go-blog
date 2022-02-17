package home

import (
	"github.com/gin-gonic/gin"
	controller "golang-blog/controllers/home"
)

func Routes(route *gin.Engine)  {

	c := controller.HomeController{}

	home := route.Group("/")
	home.GET("/", c.Show)
}