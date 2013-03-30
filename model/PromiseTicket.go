package model

import (
	"labix.org/v2/mgo/bson"
)

type PromiseTicket struct {
	Id   bson.ObjectId
	Name string

	//callback HttpCallback
}
