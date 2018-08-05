package main

import (
	"fmt"
	"log"

	"github.com/taeho-io/family/svc/srv/aws"
	"github.com/taeho-io/family/svc/srv/aws/dynamodb"
)

func main() {
	ddb, err := getDynamodb()
	if err != nil {
		log.Fatal(err)
	}

	tableNames, err := ddb.DB().ListTables().All()
	if err != nil {
		log.Fatal(err)
	}

	for _, tableName := range tableNames {
		fmt.Println(tableName)
	}
}

func getDynamodb() (dynamodb.IFace, error) {
	a, err := aws.New()
	if err != nil {
		return nil, err
	}
	return dynamodb.New(a), nil
}
