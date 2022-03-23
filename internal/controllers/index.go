package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hamedsz/go-blog/internal/requests"
)

// IndexPosts godoc
// @Summary index posts
// @Schemes
// @Description index posts
// @Tags Posts
// @Param page query int false "Page Number"
// @Param page_count query int false "Page Count"
// @Accept json
// @Produce json
// @Success 200 {string} Post
// @Router /posts [get]
func (controller *PostController) Index(c *gin.Context)  {

	req := requests.NewIndexRequest()

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