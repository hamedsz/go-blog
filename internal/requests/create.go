package requests

type CreateRequest struct {
	Title     string `form:"title"`
	Body      string `form:"body"`
}

func NewCreateRequest() *CreateRequest {
	req := new(CreateRequest)
	return req
}