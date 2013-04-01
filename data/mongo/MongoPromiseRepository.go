package mongo

import (
	log "code.google.com/p/log4go"
	"github.com/pjvds/promise/model"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type MongoPromiseTicketRepository struct {
	session    *MongoPromiseSession
	collection *mgo.Collection
}

func NewMongoTicketPromiseRepository(session *MongoPromiseSession) (*MongoPromiseTicketRepository, error) {
	collection := session.database.C("Ticket")

	return &MongoPromiseTicketRepository{
		session:    session,
		collection: collection,
	}, nil
}

func (r *MongoPromiseTicketRepository) Add(promise *model.PromiseTicket) error {

	err := r.collection.Insert(&promise)

	if err != nil {
		log.Trace("error while inserting document in Mongo: %v", err)
	} else {
		log.Trace("new document inserted into mongo with id: %v", promise)
	}

	return err
}

func (r *MongoPromiseTicketRepository) All() []model.PromiseTicket {
	var tickets []model.PromiseTicket
	r.collection.Find(bson.D{}).All(&tickets)

	if tickets == nil {
		tickets = make([]model.PromiseTicket, 0, 0)
	}

	return tickets
}
