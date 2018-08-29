package main

import (
	"log"
	"os"
	"strconv"

	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/todogroups"
	"github.com/taeho-io/family/services/todogroups/internal/model"
)

var (
	defaultReadUnits  int64 = 1
	defaultWriteUnits int64 = 1
)

func main() {
	cfg := todogroups.NewConfig(todogroups.NewSettings())

	ddb, err := getDynamodb()
	if err != nil {
		log.Fatal(err)
	}

	readUnits, writeUnits := loadProvisionUnits()

	dynamodbTodoGroupsFullTableName := base.FullDynamodbTableName(cfg, cfg.Settings().DynamodbTodoGroupsTableName)
	err = ddb.DB().CreateTable(dynamodbTodoGroupsFullTableName, model.TodoGroup{}).
		Provision(readUnits, writeUnits).
		Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("`%s` table is being created with readUnits:%d, writeUnits:%d.",
		dynamodbTodoGroupsFullTableName, readUnits, writeUnits)

	dynamodbTodoGroupPermitsFullTableName := base.FullDynamodbTableName(cfg, cfg.Settings().DynamodbTodoGroupPermitsTableName)
	err = ddb.DB().CreateTable(dynamodbTodoGroupPermitsFullTableName, model.TodoGroup{}).
		Provision(readUnits, writeUnits).
		Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("`%s` table is being created with readUnits:%d, writeUnits:%d.",
		dynamodbTodoGroupPermitsFullTableName, readUnits, writeUnits)
}

func loadProvisionUnits() (readUnits int64, writeUnits int64) {
	readUnits, err := strconv.ParseInt(os.Getenv("READ_UNITS"), 10, 64)
	if err != nil {
		readUnits = defaultReadUnits
	}

	writeUnits, err = strconv.ParseInt(os.Getenv("WRITE_UNITS"), 10, 64)
	if err != nil {
		writeUnits = defaultWriteUnits
	}

	return readUnits, writeUnits
}

func getDynamodb() (base.Dynamodb, error) {
	aws, err := base.NewAws()
	if err != nil {
		return nil, err
	}
	return base.NewDynamodb(aws), nil
}
