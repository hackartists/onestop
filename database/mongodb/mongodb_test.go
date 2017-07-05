package mongodb

import (
	"encoding/json"
	"testing"

	"github.com/pwnartist/onestop/database"
)

type fakeMongoDBContext struct {
}

func (m fakeMongoDBContext) Session() interface{} {
	return nil
}

func (m fakeMongoDBContext) Insert(data database.OnestopDataContext) error {
	return nil
}

func (m fakeMongoDBContext) Select(data database.OnestopDataContext) ([]interface{}, error) {
	return nil, nil
}

type fakeDataContext struct {
	Name string `json:"name"`
}

func (d fakeDataContext) JsonString() (string, error) {
	return json.Marshal(d), nil
}

func (d fakeDataContext) Parse(interface{}) ([]interface{}, error) {
	return nil, nil
}

func (d fakeDataContext) Conditions() interface{} {
	return nil
}

func (d fakeDataContext) TableName() string {
	return ""
}

func (d fakeDataContext) Data() interface{} {
	return nil
}

func TestInsert(t *testing.T) {
	tests := []struct {
		input fakeDataContext
	}{
		{
			input: fakeDataContext{Name: "test"},
		},
	}

	m := fakeMongoDBContext{}
	for _, test := range tests {
		err := m.Insert(test.input)
		if err != nil {
			t.Errorf("TestInsert(%v): %v", m, err)
		}
	}
}

func TestSelect(t *testing.T) {
	tests := []struct {
		input fakeDataContext
	}{
		{
			input: fakeDataContext{Name: "test"},
		},
	}

	m := fakeMongoDBContext{}

	for _, test := range tests {
		_, err := m.Select(test.input)
		if err != nil {
			t.Errorf("TestInsert(%v): %v", m, err)
		}
	}
}
