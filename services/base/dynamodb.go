package base

import (
	"github.com/guregu/dynamo"
)

type Dynamodb interface {
	DB() *dynamo.DB
}

type DefaultDynamodb struct {
	Dynamodb

	ddb *dynamo.DB
}

func NewDynamodb(aws Aws) Dynamodb {
	return &DefaultDynamodb{
		ddb: dynamo.New(aws.Session()),
	}
}

func NewMockDynamodb() Dynamodb {
	aws, _ := NewAws()
	return NewDynamodb(aws)
}

func (d *DefaultDynamodb) DB() *dynamo.DB {
	return d.ddb
}
