package identification

import (
	"labix.org/v2/mgo/bson"
)

type Identifier bson.ObjectId

func NewId() Identifier {
	return Identifier(bson.NewObjectId())
}
