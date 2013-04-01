package model

import (
	"labix.org/v2/mgo/bson"
)

type Identifier bson.ObjectId

func NewIdentifier() Identifier {
	return Identifier(bson.NewObjectId())
}
