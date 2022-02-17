package post

import (
	"github.com/gin-gonic/gin"
	post2 "golang-blog/models/post"
)

func (controller *PostController) Index(c *gin.Context)  {
	post := post2.Post{
		Title: "hello",
		Body: "how are u",
	}

	post.Save()

	c.String(200 , "hell world")
}