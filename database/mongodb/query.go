package mongodb

import (
	"github.com/hamedsz/go-blog/database/base"
	"github.com/hamedsz/go-blog/database/mongodb/query"
)

func (db *Database) Query() base.Query {
	q := query.NewQuery()
	return q
}
