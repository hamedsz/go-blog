package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hamedsz/go-blog/internal/requests"
)

// CreatePost godoc
// @Summary create posts
// @Schemes
// @Description create posts
// @Tags Posts
// @Param page query int false "Page Number"
// @Param page_count query int false "Page Count"
// @Accept json
// @Produce json
// @Success 200 {string} Post
// @Router /posts [get]
func (controller *PostController) Create(c *gin.Context)  {
	req := requests.NewCreateRequest()

	err := c.ShouldBindJSON(&req)

	if err != nil{
		panic(err)
	}

	post, err := controller.Service.Create(*req)

	if err != nil{
		panic(err)
	}

	c.JSON(200 , post)
}
