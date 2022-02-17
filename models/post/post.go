package post

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-blog/database/mongodb"
)

const NAME = "posts"

type Post struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title,omitempty"`
	Body  string             `bson:"body,omitempty"`
}

func (post *Post) Save()  {
	mongodb.Insert(NAME, post)
}

