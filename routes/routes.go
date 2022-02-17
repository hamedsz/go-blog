package routes

import (
	"github.com/gin-gonic/gin"
	"golang-blog/routes/home"
	"golang-blog/routes/post"
)

func LoadRoutes(route *gin.Engine){
	home.Routes(route)
	post.Routes(route)
}
