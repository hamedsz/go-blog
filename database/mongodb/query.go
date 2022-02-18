package mongodb

import (
	"golang-blog/database/base"
	"golang-blog/database/mongodb/query"
)

func (db *Database) Query() base.Query {
	q := query.NewQuery()
	return q
}
