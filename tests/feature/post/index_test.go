package post

import (
	"github.com/hamedsz/go-blog/internal/controllers/post"
	post2 "github.com/hamedsz/go-blog/internal/domain"
	"github.com/hamedsz/go-blog/tests/feature/base"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostsIndexStatusCode(t *testing.T){
	controller := post.NewPostController()

	w := base.FeatureTest(controller.Index, "GET", nil , "")

	assert.Equal(t, w.Code , 200)
}

func insertSamplePost()  {
	model := post2.Model{
		Title: "test",
		Body: "test",
	}
	err := model.GetDB().Save()
	if err != nil{
		panic(err)
	}
}

func TestPostsIndexResponseStructure(t *testing.T){
	base.ResetDB()
	insertSamplePost()

	w := base.FeatureTest(post.NewPostController().Index, "GET", nil , "")

	jsonassert.New(t).Assertf(w.Body.String(), `
    {
		"items": [
			{
				"id": "<<PRESENCE>>",
				"title": "test",
				"body": "test",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			}
		],
		"page_count": "<<PRESENCE>>",
		"total_count": 1
	}
	`)
}

func TestPostsIndexPaginationPageOne(t *testing.T)  {
	base.ResetDB()

	const COUNT = 11;

	//insert posts
	for i := 1; i <= COUNT; i++  {
		insertSamplePost()
	}


	w := base.FeatureTest(
		post.NewPostController().Index,
		"GET",
		nil,
		"?page_count=10&page=1",
	)

	jsonassert.New(t).Assertf(w.Body.String(), `
    {
		"items": [
			{
				"id": "<<PRESENCE>>",
				"title": "test",
				"body": "test",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			},
			{
				"id": "<<PRESENCE>>",
				"title": "test",
				"body": "test",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			},
			{
				"id": "<<PRESENCE>>",
				"title": "test",
				"body": "test",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			},
			{
				"id": "<<PRESENCE>>",
				"title": "test",
				"body": "test",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			},
			{
				"id": "<<PRESENCE>>",
				"title": "test",
				"body": "test",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			},
			{
				"id": "<<PRESENCE>>",
				"title": "test",
				"body": "test",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			},
			{
				"id": "<<PRESENCE>>",
				"title": "test",
				"body": "test",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			},
			{
				"id": "<<PRESENCE>>",
				"title": "test",
				"body": "test",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			},
			{
				"id": "<<PRESENCE>>",
				"title": "test",
				"body": "test",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			},
			{
				"id": "<<PRESENCE>>",
				"title": "test",
				"body": "test",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			}
		],
		"page_count": 10,
		"total_count": %d
	}
	`, COUNT)

}

func TestPostsIndexPaginationPageTwo(t *testing.T)  {
	base.ResetDB()

	const COUNT = 11;

	//insert posts
	for i := 1; i <= COUNT; i++  {
		insertSamplePost()
	}


	w := base.FeatureTest(
		post.NewPostController().Index,
		"GET",
		nil,
		"?page_count=10&page=2",
	)

	jsonassert.New(t).Assertf(w.Body.String(), `
    {
		"items": [
			{
				"id": "<<PRESENCE>>",
				"title": "test",
				"body": "test",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			}
		],
		"page_count": 10,
		"total_count": %d
	}
	`, COUNT)
}