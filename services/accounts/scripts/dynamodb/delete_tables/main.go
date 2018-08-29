package main

import (
	"log"

	"github.com/taeho-io/family/services/accounts"
	"github.com/taeho-io/family/services/accounts/internal/repo"
	"github.com/taeho-io/family/services/base"
)

func main() {
	cfg := accounts.NewConfig(accounts.NewSettings())

	ddb, err := getDynamodb()
	if err != nil {
		log.Fatal(err)
	}

	fullTableName := base.FullDynamodbTableName(cfg, cfg.Settings().DynamodbAccountsTableName)

	accountsTable := repo.NewAccountsRepo(ddb, repo.NewAccountsRepoConfig(fullTableName)).Table()
	if err := accountsTable.DeleteTable().Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("`%s` table is being deleted.", accountsTable.Name())
}

func getDynamodb() (base.Dynamodb, error) {
	aws, err := base.NewAws()
	if err != nil {
		return nil, err
	}
	return base.NewDynamodb(aws), nil
}
