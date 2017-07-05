package mongodb

import (
	"github.com/pwnartist/onestop/database"
	mgo "gopkg.in/mgo.v2"
)

type MongoDBContext struct {
	sess *mgo.Session
	db   string
}

func New(db, url string) (*MongoDBContext, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	return &MongoDBContext{
		sess: session,
		db:   db,
	}, nil
}

func (m MongoDBContext) Session() interface{} {
	return m.sess
}

func (m MongoDBContext) Insert(data database.OnestopDataContext) error {
	c := m.sess.DB(m.db).C(data.TableName())
	c.Insert(data.Data())
	return nil
}

func (m MongoDBContext) Select(data database.OnestopDataContext) ([]interface{}, error) {
	return nil, nil
}
