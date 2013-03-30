package mongo

import ()

type MongoPromiseRepositoryFactory struct {
	session *MongoPromiseSession
}

func NewMongoPromiseRepositoryFactory(session *MongoPromiseSession) *MongoPromiseRepositoryFactory {
	return &MongoPromiseRepositoryFactory{
		session: session,
	}
}

func (fac *MongoPromiseRepositoryFactory) CreatePromiseTicketRepository() *MongoPromiseTicketRepository {
	repo, err := NewMongoTicketPromiseRepository(fac.session)

	if err != nil {
		panic(err)
	}

	return repo
}
