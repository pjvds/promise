package mongo

import (
	log "code.google.com/p/log4go"
	"github.com/pjvds/promise/model"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type MongoCallbackAttemptRepository struct {
	session    *MongoPromiseSession
	collection *mgo.Collection
}

func NewMongoCallbackAttemptRepository(session *MongoPromiseSession) (*MongoCallbackAttemptRepository, error) {
	collection := session.database.C("Attempt")

	return &MongoCallbackAttemptRepository{
		session:    session,
		collection: collection,
	}, nil
}

func (r *MongoCallbackAttemptRepository) Add(promise *model.CallbackAttempt) error {
	promise.Id = bson.NewObjectId()

	err := r.collection.Insert(&promise)

	if err != nil {
		log.Trace("error while inserting document in Mongo: %v", err)
	} else {
		log.Trace("new document inserted into mongo with id: %v", promise.Id)
	}

	return err
}

func (r *MongoCallbackAttemptRepository) All() []model.CallbackAttempt {
	var attempts []model.CallbackAttempt
	r.collection.Find(bson.D{}).All(&attempts)

	if attempts == nil {
		attempts = make([]model.CallbackAttempt, 0, 0)
	}

	return attempts
}
