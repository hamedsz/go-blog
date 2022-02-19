package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hamedsz/go-blog/routes/home"
	"github.com/hamedsz/go-blog/routes/post"
	"github.com/hamedsz/go-blog/routes/swagger"
)

func LoadRoutes(route *gin.Engine){
	home.Routes(route)
	post.Routes(route)
	swagger.Routes(route)
}
