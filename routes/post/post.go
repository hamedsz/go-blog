package post

import (
	"github.com/gin-gonic/gin"
	controller "golang-blog/controllers/post"
)

func Routes(route *gin.Engine)  {

	c := controller.PostController{}

	home := route.Group("/posts")
	home.GET("/", c.Index)
}
