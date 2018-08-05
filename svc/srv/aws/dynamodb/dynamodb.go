package dynamodb

import (
	"github.com/guregu/dynamo"
	"github.com/taeho-io/family/svc/srv/aws"
)

type IFace interface {
	DB() *dynamo.DB
}

type Dynamodb struct {
	IFace

	db *dynamo.DB
}

func New(a aws.IFace) IFace {
	return &Dynamodb{
		db: dynamo.New(a.Session()),
	}
}

func NewMock() IFace {
	a, _ := aws.New()
	return New(a)
}

func (d *Dynamodb) DB() *dynamo.DB {
	return d.db
}
