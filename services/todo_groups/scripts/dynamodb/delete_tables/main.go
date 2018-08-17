package main

import (
	"log"

	"github.com/taeho-io/family/services/base/aws"
	"github.com/taeho-io/family/services/base/aws/dynamodb"
	"github.com/taeho-io/family/services/todos/config"
	"github.com/taeho-io/family/services/todos/repos/todo_group_permits_repo"
	"github.com/taeho-io/family/services/todos/repos/todo_groups_repo"
)

func main() {
	cfg := config.New(config.NewSettings())

	ddb, err := getDynamodb()
	if err != nil {
		log.Fatal(err)
	}

	todoGroupsTable := todo_groups_repo.New(ddb, cfg).Table()
	if err := todoGroupsTable.DeleteTable().Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("`%s` table is being deleted.", todoGroupsTable.Name())

	todoGroupPermitsTable := todo_group_permits_repo.New(ddb, cfg).Table()
	if err := todoGroupPermitsTable.DeleteTable().Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("`%s` table is being deleted.", todoGroupPermitsTable.Name())
}

func getDynamodb() (dynamodb.IFace, error) {
	a, err := aws.New()
	if err != nil {
		return nil, err
	}
	return dynamodb.New(a), nil
}
