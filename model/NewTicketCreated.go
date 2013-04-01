package model

import (
	"github.com/pjvds/promise/eventing"
	"github.com/pjvds/promise/identification"
	"time"
)

type NewTicketCreated struct {
	eventing.EventMessage

	Id   identification.Identifier
	Name string
	Url  string
	When time.Time
}
