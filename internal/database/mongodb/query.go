package mongodb

import (
	"github.com/hamedsz/go-blog/internal/database/base"
	"github.com/hamedsz/go-blog/internal/database/mongodb/query"
)

func (db *Database) Query() base.Query {
	q := query.NewQuery()
	return q
}
