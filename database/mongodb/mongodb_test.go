package mongodb

import "github.com/pwnartist/onestop/database"

type fakeMongoDBContext struct {
}

func (m fakeMongoDBContext) Session() interface{} {
	return nil
}

func (m fakeMongoDBContext) Insert(data database.OnestopDataContext) error {
	return nil
}

func (m fakeMongoDBContext) Select(OnestopDataContext) ([]interface{}, error) {
	return nil, nil
}
