package model

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type PromiseTicket struct {
	Id        bson.ObjectId
	Name      string
	ExecuteAt time.Time

	//callback HttpCallback
}
