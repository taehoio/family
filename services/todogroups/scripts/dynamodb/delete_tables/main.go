package main

import (
	"log"

	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/todogroups"
	"github.com/taeho-io/family/services/todogroups/internal/repo"
)

func main() {
	cfg := todogroups.NewConfig(todogroups.NewSettings())

	ddb, err := getDynamodb()
	if err != nil {
		log.Fatal(err)
	}

	dynamodbTodoGroupsTableName := base.FullDynamodbTableName(cfg, cfg.Settings().DynamodbTodoGroupsTableName)
	todoGroupsTable := repo.NewTodoGroupsRepo(ddb, repo.NewTodoGroupPermitsRepoConfig(dynamodbTodoGroupsTableName)).Table()
	if err := todoGroupsTable.DeleteTable().Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("`%s` table is being deleted.", todoGroupsTable.Name())

	dynamodbTodoGroupPermitsTableName := base.FullDynamodbTableName(cfg, cfg.Settings().DynamodbTodoGroupPermitsTableName)
	todoGroupPermitsTable := repo.NewTodoGroupPermitsRepo(ddb, repo.NewTodoGroupPermitsRepoConfig(dynamodbTodoGroupPermitsTableName)).Table()
	if err := todoGroupPermitsTable.DeleteTable().Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("`%s` table is being deleted.", todoGroupPermitsTable.Name())
}

func getDynamodb() (base.Dynamodb, error) {
	aws, err := base.NewAws()
	if err != nil {
		return nil, err
	}
	return base.NewDynamodb(aws), nil
}
