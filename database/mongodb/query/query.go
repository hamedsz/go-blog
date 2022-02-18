package query

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang-blog/database/base"
)

type QueryBuilder struct {
	Filter map[string]interface{}
}


func NewQuery() *QueryBuilder {
	q := new(QueryBuilder)
	q.Filter = map[string]interface{}{}
	return q
}


func (query *QueryBuilder) GetFilter() map[string]interface{}{
	return query.Filter
}

func (query *QueryBuilder) WhereEqual(key string, value interface{}) base.Query{

	query.Filter[key] = value

	return query
}

func (query *QueryBuilder) WhereNotEqual(key string, value interface{}) base.Query{

	query.Filter[key] = bson.M{
		"$ne": value,
	}

	return query
}

func (query *QueryBuilder) Get() {

}