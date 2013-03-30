package main

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type PromiseTicket struct {
	Id           bson.ObjectId
	Name         string
	ExecuteAfter time.Time

	//callback HttpCallback
}

type HttpCallback struct {
	id  bson.ObjectId
	url string
}
