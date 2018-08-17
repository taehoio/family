package main

import (
	"log"

	"github.com/taeho-io/family/services/accounts/config"
	"github.com/taeho-io/family/services/accounts/repos/accounts_repo"
	"github.com/taeho-io/family/services/base/aws"
	"github.com/taeho-io/family/services/base/aws/dynamodb"
)

func main() {
	cfg := config.New(config.NewSettings())

	ddb, err := getDynamodb()
	if err != nil {
		log.Fatal(err)
	}

	accountsTable := accounts_repo.New(ddb, cfg).Table()
	if err := accountsTable.DeleteTable().Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("`%s` table is being deleted.", accountsTable.Name())
}

func getDynamodb() (dynamodb.IFace, error) {
	a, err := aws.New()
	if err != nil {
		return nil, err
	}
	return dynamodb.New(a), nil
}
