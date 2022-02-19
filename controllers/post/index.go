package post

import (
	"github.com/gin-gonic/gin"
	"golang-blog/requests/post"
)

func (controller *PostController) Index(c *gin.Context)  {

	req := post.NewIndexRequest()

	err := c.BindQuery(&req)

	if err != nil{
		panic(err)
	}

	data, count, err := controller.Service.Index(*req)

	if err != nil{
		panic(err)
	}

	c.JSON(200 , gin.H{
		"total_count": count,
		"page_count": 10,
		"items": data,
	})
}