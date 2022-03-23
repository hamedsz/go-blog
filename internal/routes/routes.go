package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hamedsz/go-blog/internal/routes/post"
	"github.com/hamedsz/go-blog/internal/routes/swagger"
)

func LoadRoutes(route *gin.Engine){
	post.Routes(route)
	swagger.Routes(route)
}
