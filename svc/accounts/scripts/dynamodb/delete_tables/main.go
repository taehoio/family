package main

import (
	"log"

	"github.com/taeho-io/family/svc/accounts/config"
	"github.com/taeho-io/family/svc/accounts/repos/account_email_repo"
	"github.com/taeho-io/family/svc/accounts/repos/account_repo"
	"github.com/taeho-io/family/svc/srv/aws"
	"github.com/taeho-io/family/svc/srv/aws/dynamodb"
)

func main() {
	svcCfg := config.New(config.NewSettings())

	ddb, err := getDynamodb()
	if err != nil {
		log.Fatal(err)
	}

	accountTable := account_repo.New(ddb, svcCfg).Table()
	if err := accountTable.DeleteTable().Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("`%s` table is being deleted.", accountTable.Name())

	accountEmailTable := account_email_repo.New(ddb, svcCfg).Table()
	if err := accountEmailTable.DeleteTable().Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("`%s` table is beingdeleted.", accountEmailTable.Name())
}

func getDynamodb() (dynamodb.IFace, error) {
	a, err := aws.New()
	if err != nil {
		return nil, err
	}
	return dynamodb.New(a), nil
}