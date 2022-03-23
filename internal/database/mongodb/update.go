package mongodb

import (
	"context"
	"time"
)

func (db *Database) Update(table string, filter interface{}, data interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_ , err := GetDB().Collection(table).UpdateOne(ctx, filter, data)

	return err
}
