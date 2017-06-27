package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pwnartist/onestop/database"

	a "github.com/pwnartist/onestop/aws"
)

type DynamoDBContext struct {
	sess *session.Session
}

func New(region string) (*DynamoDBContext, error) {
	// TODO: aws.Config have to be added.
	sess, err := a.Connect(region)

	return &DynamoDBContext{sess: sess}, err
}

func (c DynamoDBContext) Session() interface{} {
	return c.sess
}

func (c DynamoDBContext) Insert(data database.OnestopDataContext) error {
	sess := c.sess
	svc := dynamodb.New(sess)

	params := &dynamodb.PutItemInput{
		Item:      data.Data().(map[string]*dynamodb.AttributeValue),
		TableName: aws.String(data.TableName()),
	}

	_, err := svc.PutItem(params)

	return err
}

func (c DynamoDBContext) Select(data database.OnestopDataContext) ([]interface{}, error) {
	sess := c.sess

	svc := dynamodb.New(sess)

	params := &dynamodb.QueryInput{
		TableName: aws.String(data.TableName()),
		//Limit:                     aws.Long(10),
		//FilterExpression:          data.DynamoKeyConditionExpression(),
		//ExpressionAttributeValues: data.DynamoExpressionValues(),
		KeyConditions: data.Conditions().(map[string]*dynamodb.Condition),
	}

	resp, err := svc.Query(params)
	if err != nil {
		return nil, err
	}
	return data.Parse(resp)
}
