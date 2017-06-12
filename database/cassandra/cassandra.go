package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/pwnartist/onestop/database"
)

type CassandraContext struct {
	sess  *gocql.Session
	table string
}

func New(host string, keyspace string, table string) (CassandraContext, error) {
	sess, err := connect(host, keyspace)

	return CassandraContext{sess: sess, table: table}, err
}

func (c CassandraContext) Session() interface{} {
	return c.sess
}

func (c CassandraContext) TableName() string {
	return c.table
}

func (c CassandraContext) Insert(data database.OnestopDataContext) error {
	//queryKey := strings.Join(data.Keys(), ",")
	//values := data.Values()
	//query := "INSERT INTO ? (" + queryKey + ") VALUES (" + strings.Repeat("?,", len(values)-1) + "?)"
	j, err := data.JsonString()
	if err != nil {
		return err
	}
	//query := "INSERT INTO ? JSON '?'"
	query := fmt.Sprintf("INSERT INTO %s JSON '%s'", c.TableName(), j)
	err = c.sess.Query(query).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (c CassandraContext) Select(data database.OnestopDataContext) ([]interface{}, error) {
	query := "SELECT JSON * FROM " + c.TableName()
	iter := c.sess.Query(query).Iter()
	defer iter.Close()
	//TODO: Converting iter into result array
	return data.Parse(iter)
}

func connect(host string, keyspace string) (*gocql.Session, error) {
	cluster := gocql.NewCluster(host)
	cluster.Keyspace = keyspace
	return cluster.CreateSession()
}
