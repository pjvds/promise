package mongo

import (
	"labix.org/v2/mgo"
)

type MongoPromiseSession struct {
	session  *mgo.Session
	database *mgo.Database
}

func NewMongoPromiseSession(session *mgo.Session, database *mgo.Database) *MongoPromiseSession {
	return &MongoPromiseSession{
		session:  session,
		database: database,
	}
}
