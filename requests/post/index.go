package post

type IndexRequest struct {
	Page      int64 `form:"page"`
	PageCount int64 `form:"page_count"`
}

func NewIndexRequest() *IndexRequest  {
	req := new(IndexRequest)
	req.Page = 1
	req.PageCount = 10
	return req
}