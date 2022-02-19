package post

import (
	"golang-blog/models/post"
	post2 "golang-blog/requests/post"
)

type Service struct {

}

func (s *Service) Index(request post2.IndexRequest) ([]post.Model, *int64 , error) {

	model := post.Model{}

	query := model.GetDB().Query().Limit(request.PageCount).Skip((request.Page-1) * request.PageCount);

	results, err := model.GetDB().Find(query)

	if err != nil{
		return nil, nil, err
	}

	count, err := model.GetDB().Count(query)

	if err != nil{
		return nil, nil, err
	}

	data , err := post.List(results)

	if err != nil{
		return nil, nil, err
	}

	return data, &count , nil
}