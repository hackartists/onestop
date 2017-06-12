package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pwnartist/onestop/database"

	a "github.com/pwnartist/onestop/aws"
)

type DynamoDBContext struct {
	sess  *session.Session
	table string
}

func New(region string, table string) (DynamoDBContext, error) {
	// TODO: aws.Config have to be added.
	sess, err := a.Connect(region)

	return DynamoDBContext{sess: sess, table: table}, err
}

func (c DynamoDBContext) Session() interface{} {
	return c.sess
}

func (c DynamoDBContext) TableName() string {
	return c.table
}

func (c DynamoDBContext) Insert(data database.OnestopDataContext) error {
	sess := c.sess
	svc := dynamodb.New(sess)

	params := &dynamodb.PutItemInput{
		Item:      data.DynamoAttribute(),
		TableName: aws.String(d.TableName()),
	}

	_, err := svc.PutItem(params)

	return err
}

func (c DynamoDBContext) Select(data data.baseOnestopDataContext) ([]interface{}, error) {
	sess := c.sess

	svc := dynamodb.New(sess)

	params := &dynamodb.QueryInput{
		TableName: aws.String(c.TableName()),
		//Limit:                     aws.Long(10),
		//FilterExpression:          data.DynamoKeyConditionExpression(),
		//ExpressionAttributeValues: data.DynamoExpressionValues(),
		KeyConditions: data.Conditions(),
	}

	resp, err := svc.Query(params)
	if err != nil {
		return nil, err
	}
	return data.Parse(resp)
}
