package mongodb

import (
	"context"
	"golang-blog/database/base"
)

func (db *Database) Count(query base.Query) (int64, error) {
	return GetDB().Collection(db.Model.Name()).CountDocuments(context.TODO(), query.GetFilter())
}
