package post

import (
	"github.com/gin-gonic/gin"
)

func (controller *PostController) Index(c *gin.Context)  {

	data, count, err := controller.Service.Index()

	if err != nil{
		panic(err)
	}

	c.JSON(200 , gin.H{
		"total_count": count,
		"page_count": 123,
		"items": data,
	})
}