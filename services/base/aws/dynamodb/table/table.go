package table

import "github.com/guregu/dynamo"

type IFace interface {
	Table() *dynamo.Table
}
