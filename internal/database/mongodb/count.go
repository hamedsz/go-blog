package mongodb

import (
	"context"
	"github.com/hamedsz/go-blog/internal/database/base"
)

func (db *Database) Count(query base.Query) (int64, error) {
	return GetDB().Collection(db.Model.Name()).CountDocuments(context.TODO(), query.GetFilter())
}
