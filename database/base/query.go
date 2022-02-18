package base

type Query interface {
	WhereEqual(key string, value interface{}) Query
	WhereNotEqual(key string, value interface{}) Query
	GetFilter() map[string]interface{}
}
