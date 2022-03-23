package services

import (
	"github.com/hamedsz/go-blog/internal/domain"
	post2 "github.com/hamedsz/go-blog/internal/requests"
)

func (s *Service) Index(request post2.IndexRequest) ([]domain.Post, *int64 , error) {

	model := domain.Post{}

	query := model.GetDB().Query().Limit(request.PageCount).Skip((request.Page-1) * request.PageCount);

	results, err := model.GetDB().Find(query)

	if err != nil{
		return nil, nil, err
	}

	count, err := model.GetDB().Count(query)

	if err != nil{
		return nil, nil, err
	}

	data , err := domain.List(results)

	if err != nil{
		return nil, nil, err
	}

	return data, &count , nil
}