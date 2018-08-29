package base

import "github.com/guregu/dynamo"

type DynamodbRepo interface {
	Table() *dynamo.Table
}
