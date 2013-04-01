package model

import (
	log "code.google.com/p/log4go"
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
	createEvent := NewTicketCreated{
		Name: name,
		Url:  url,
		When: when,
	}
	ticket := &PromiseTicket{
		EventSource: eventing.NewEventSource(),
	}
	ticket.EventSource.RegisterHandler(NewTicketCreated{}, func(event interface{}) error {
		e := event.(NewTicketCreated)
		return ticket.ApplyNewTicket(e)
	})

	ticket.Apply(createEvent)

	return ticket
}

func (ticket *PromiseTicket) ApplyNewTicket(e NewTicketCreated) error {
	ticket.Name = e.Name
	ticket.Url = e.Url
	ticket.When = e.When

	return nil
}
