package main

import (
	"log"

	"os"
	"strconv"

	"github.com/taeho-io/family/services/accounts/config"
	"github.com/taeho-io/family/services/accounts/models"
	"github.com/taeho-io/family/services/accounts/repos/accounts_repo"
	"github.com/taeho-io/family/services/base/aws"
	"github.com/taeho-io/family/services/base/aws/dynamodb"
)

var (
	defaultReadUnits  int64 = 1
	defaultWriteUnits int64 = 1
)

func main() {
	cfg := config.New(config.NewSettings())

	ddb, err := getDynamodb()
	if err != nil {
		log.Fatal(err)
	}

	readUnits, writeUnits := loadProvisionUnits()

	accountsTableName := accounts_repo.New(ddb, cfg).Table().Name()
	err = ddb.DB().CreateTable(accountsTableName, models.Account{}).
		Provision(readUnits, writeUnits).
		Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("`%s` table is being created with readUnits:%d, writeUnits:%d.",
		accountsTableName, readUnits, writeUnits)
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

func getDynamodb() (dynamodb.IFace, error) {
	a, err := aws.New()
	if err != nil {
		return nil, err
	}
	return dynamodb.New(a), nil
}
