package model

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type CallbackAttempt struct {
	Id             bson.ObjectId
	TicketId       bson.ObjectId
	StartedAt      time.Time
	EndedAt        time.Time
	Duration       time.Duration
	Attempt        int
	Url            string
	HttpStatusCode string
	Log            string
}
