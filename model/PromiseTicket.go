package model

import (
	"github.com/pjvds/promise/eventing"
	"time"
)

type PromiseTicket struct {
	eventing.EventSource

	Name      string
	Url       string
	When      time.Time
	Finalized bool
}

func NewTicket(name string, url string, when time.Time) *PromiseTicket {
	createEvent := &NewTicketCreated{
		Name: name,
		Url:  url,
		When: when,
	}
	ticket := &PromiseTicket{}
	ticket.Apply(&createEvent.EventMessage)

	return ticket
}

func (ticket *PromiseTicket) ApplyNewTicket(e *NewTicketCreated) {
	ticket.Name = e.Name
	ticket.Url = e.Url
	ticket.When = e.When
}
