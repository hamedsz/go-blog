package query

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang-blog/database/base"
)

type QueryBuilder struct {
	Filter map[string]interface{}
	NumLimit int64
	NumSkip int64
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

func (query *QueryBuilder) Limit(limit int64) base.Query {

	query.NumLimit = limit

	return query
}

func (query *QueryBuilder) Skip(skip int64) base.Query {

	query.NumSkip = skip

	return query
}

func (query *QueryBuilder) GetLimit() int64 {
	return query.NumLimit
}

func (query *QueryBuilder) GetSkip() int64 {
	return query.NumSkip
}