package post

import "golang-blog/services/post"

type PostController struct{
	Service post.Service
}

func NewPostController() *PostController {
	c := new(PostController)
	c.Service = post.Service{}
	return c
}