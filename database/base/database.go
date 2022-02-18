package base


type Database interface {
	Insert(table string, data interface{}) (*string , error) //returns id
	Update(table string, filter interface{}, data interface{}) error
	Save() error
	Query() Query
	Find(query Query) ([]map[string]interface{}, error)
	Count(query Query) (int64, error)
}


func GetDatabase(db Database) Database{
	return db
}