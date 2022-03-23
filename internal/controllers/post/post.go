package post

import (
	"github.com/hamedsz/go-blog/internal/services/post"
)

type PostController struct{
	Service post.Service
}

func NewPostController() *PostController {
	c := new(PostController)
	c.Service = post.Service{}
	return c
}