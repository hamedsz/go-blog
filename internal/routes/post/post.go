package post

import (
	"github.com/gin-gonic/gin"
	controller "github.com/hamedsz/go-blog/internal/controllers"
)

func Routes(route *gin.Engine)  {

	c := controller.NewPostController()

	home := route.Group("/posts")
	home.GET("/", c.Index)
	home.POST("/", c.Create)
}
