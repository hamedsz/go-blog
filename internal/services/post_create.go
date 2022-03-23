package services

import (
	"github.com/hamedsz/go-blog/internal/domain"
	post2 "github.com/hamedsz/go-blog/internal/requests"
)


func (s *Service) Create(request post2.CreateRequest) (*domain.Post, error){
	post := domain.NewPost()
	post.Title = request.Title
	post.Body = request.Body

	err := post.GetDB().Save()
	if err != nil {
		return nil, err
	}

	return post, nil
}