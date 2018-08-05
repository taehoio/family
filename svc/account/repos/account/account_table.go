package account

import (
	"strings"

	"github.com/guregu/dynamo"
	"github.com/taeho-io/family/svc/account/config"
	svcConfig "github.com/taeho-io/family/svc/account/config"
	"github.com/taeho-io/family/svc/srv/aws"
	"github.com/taeho-io/family/svc/srv/aws/dynamodb"
	"github.com/taeho-io/family/svc/srv/aws/dynamodb/table"
)

type Table struct {
	table.IFace

	accountTable *dynamo.Table
}

func New(ddb dynamodb.IFace, svcCfg config.IFace) table.IFace {
	fullTableName := fullTableName(svcCfg)
	accountTable := ddb.DB().Table(fullTableName)

	return &Table{
		accountTable: &accountTable,
	}
}

func NewMock() table.IFace {
	a, _ := aws.New()
	ddb := dynamodb.New(a)
	svcCfg := svcConfig.New(svcConfig.NewMockSettings())

	return New(ddb, svcCfg)

}

func fullTableName(svcCfg config.IFace) string {
	prefix := svcCfg.Prefix()
	tableName := svcCfg.Settings().DynamodbAccountTableName
	return strings.Join([]string{prefix, tableName}, "-")
}

func (t *Table) Table() *dynamo.Table {
	return t.accountTable
}
