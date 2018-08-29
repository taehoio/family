package main

import (
	"fmt"
	"log"

	"github.com/taeho-io/family/services/base"
)

func main() {
	ddb, err := getDynamodb()
	if err != nil {
		log.Fatal(err)
	}

	tableNames, err := ddb.DB().ListTables().All()
	if err != nil {
		log.Fatal(err)
	}

	for _, tableName := range tableNames {
		fmt.Println(tableName)
	}
}

func getDynamodb() (base.Dynamodb, error) {
	aws, err := base.NewAws()
	if err != nil {
		return nil, err
	}
	return base.NewDynamodb(aws), nil
}
