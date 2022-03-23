package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string              `bson:"title,omitempty" json:"title"`
	Body      string              `bson:"body,omitempty" json:"body"`
	CreatedAt time.Time           `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time           `bson:"updated_at,omitempty" json:"updated_at"`
}

func NewPost() *Post {
	post := new(Post)
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	return post
}