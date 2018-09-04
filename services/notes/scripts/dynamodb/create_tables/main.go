package main

import (
	"log"
	"os"
	"strconv"

	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/notes"
	"github.com/taeho-io/family/services/notes/internal/model"
)

var (
	defaultReadUnits  int64 = 1
	defaultWriteUnits int64 = 1
)

func main() {
	cfg := notes.NewConfig(notes.NewSettings())

	ddb, err := getDynamodb()
	if err != nil {
		log.Fatal(err)
	}

	readUnits, writeUnits := loadProvisionUnits()

	dynamodbTodosFullTableName := base.FullDynamodbTableName(cfg, cfg.Settings().DynamodbNotesTableName)
	err = ddb.DB().CreateTable(dynamodbTodosFullTableName, model.Note{}).
		Provision(readUnits, writeUnits).
		Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("`%s` table is being created with readUnits:%d, writeUnits:%d.",
		dynamodbTodosFullTableName, readUnits, writeUnits)
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
