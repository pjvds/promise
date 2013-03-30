package mongo

import (
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

func (r *MongoPromiseTicketRepository) Add(promise model.PromiseTicket) error {
	return r.collection.Insert(&promise)
}

func (r *MongoPromiseTicketRepository) All() []model.PromiseTicket {
	var tickets []model.PromiseTicket
	r.collection.Find(bson.D{}).All(&tickets)

	return tickets
}
