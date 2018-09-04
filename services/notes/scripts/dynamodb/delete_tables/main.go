package main

import (
	"log"

	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/notes"
	"github.com/taeho-io/family/services/notes/internal/repo"
)

func main() {
	cfg := notes.NewConfig(notes.NewSettings())

	ddb, err := getDynamodb()
	if err != nil {
		log.Fatal(err)
	}

	dynamodbNotesTableName := base.FullDynamodbTableName(cfg, cfg.Settings().DynamodbNotesTableName)
	notesTable := repo.NewNotesRepo(ddb, repo.NewNotesRepoConfig(dynamodbNotesTableName)).(base.DynamodbRepo).Table()
	if err := notesTable.DeleteTable().Run(); err != nil {
		log.Fatal(err)
	}
	log.Printf("`%s` table is being deleted.", notesTable.Name())
}

func getDynamodb() (base.Dynamodb, error) {
	aws, err := base.NewAws()
	if err != nil {
		return nil, err
	}
	return base.NewDynamodb(aws), nil
}
