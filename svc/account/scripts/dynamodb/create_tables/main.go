package main

import (
	"log"

	"os"
	"strconv"

	"github.com/taeho-io/family/svc/account/config"
	"github.com/taeho-io/family/svc/account/repos/account"
	"github.com/taeho-io/family/svc/account/repos/account_email"
	"github.com/taeho-io/family/svc/srv/aws"
	"github.com/taeho-io/family/svc/srv/aws/dynamodb"
)

var (
	defaultReadUnits  int64 = 1
	defaultWriteUnits int64 = 1
)

func main() {
	svcCfg := config.New(config.NewSettings())

	ddb, err := getDynamodb()
	if err != nil {
		log.Fatal(err)
	}

	readUnits, writeUnits := loadProvisionUnits()

	accountTableName := account.New(ddb, svcCfg).Table().Name()
	err = ddb.DB().CreateTable(accountTableName, account.Account{}).
		Provision(readUnits, writeUnits).
		Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("`%s` table is being created with readUnits:%d, writeUnits:%d.",
		accountTableName, readUnits, writeUnits)

	accountEmailTableName := account_email.New(ddb, svcCfg).Table().Name()
	err = ddb.DB().CreateTable(accountEmailTableName, account_email.AccountEmail{}).
		Provision(readUnits, writeUnits).
		Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("`%s` table is beingcreated with readUnits:%d, writeUnits:%d.",
		accountEmailTableName, readUnits, writeUnits)
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
