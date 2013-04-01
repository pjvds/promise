package main

import (
	"github.com/pjvds/promise/data"
	"github.com/pjvds/promise/model"
	"sort"
)

var (
	tickets []*model.PromiseTicket
)

func main() {
	database := server.DB("promise")
	session := mongo.NewMongoPromiseSession(server, database)
	repoFac := mongo.NewMongoPromiseRepositoryFactory(session)
	repo := repoFac.CreatePromiseTicketRepository()

	for {
		repo
	}
}
