package main

import (
	"log"

	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/todos"
	"github.com/taeho-io/family/services/todos/internal/repo"
)

func main() {
	cfg := todos.NewConfig(todos.NewSettings())

	ddb, err := getDynamodb()
	if err != nil {
		log.Fatal(err)
	}

	dynamodbTodosTableName := base.FullDynamodbTableName(cfg, cfg.Settings().DynamodbTodosTableName)
	todosTable := repo.NewTodosRepo(ddb, repo.NewTodosRepoConfig(dynamodbTodosTableName)).(base.DynamodbRepo).Table()
	if err := todosTable.DeleteTable().Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("`%s` table is being deleted.", todosTable.Name())
}

func getDynamodb() (base.Dynamodb, error) {
	aws, err := base.NewAws()
	if err != nil {
		return nil, err
	}
	return base.NewDynamodb(aws), nil
}
