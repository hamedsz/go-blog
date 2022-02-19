package base

type Query interface {
	WhereEqual(key string, value interface{}) Query
	WhereNotEqual(key string, value interface{}) Query
	GetFilter() map[string]interface{}
	Limit(limit int64) Query
	Skip(skip int64) Query
	GetLimit() int64
	GetSkip() int64
}
