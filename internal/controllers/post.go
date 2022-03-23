package controllers

import (
	"github.com/hamedsz/go-blog/internal/services"
)

type PostController struct{
	Service services.Service
}

func NewPostController() *PostController {
	c := new(PostController)
	c.Service = services.Service{}
	return c
}