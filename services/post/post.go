package post

import "golang-blog/models/post"

type Service struct {

}

func (s *Service) Index() ([]post.Model, *int64 , error) {

	model := post.Model{}

	query := model.GetDB().Query().WhereNotEqual("title", "hello2");

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