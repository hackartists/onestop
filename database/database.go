package database

type OnestopDatabaseContext interface {
	Session() interface{}
	Insert(OnestopDataContext) error
	Select(OnestopDataContext) ([]interface{}, error)
	TableName() string
}

type OnestopDataContext interface {
	JsonString() (string, error)
	Parse(interface{}) ([]interface{}, error)
	Conditions() interface{}
}
