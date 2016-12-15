package store

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/wantedly/github-username-converter/models"
)

const (
	githubUsersTable = "GitHubUsers"
)

type DynamoDB struct {
	db *dynamodb.DynamoDB
}

func NewDynamoDB() *DynamoDB {
	db := dynamodb.New(session.New(&aws.Config{}))

	return &DynamoDB{
		db: db,
	}
}

func (d *DynamoDB) ListUsers() ([]*models.User, error) {
	resp, err := d.db.Scan(&dynamodb.ScanInput{
		TableName: aws.String(githubUsersTable),
	})
	if err != nil {
		return nil, err
	}

	var users []*models.User

	for _, item := range resp.Items {
		users = append(users, models.NewUser(*item["LoginName"].S, *item["GitHubUSername"].S))
	}

	return users, nil
}

