package events

import (
	"github.com/pjvds/promise/messaging"
	"github.com/pjvds/promise/model"
	"time"
)

type NewTicketCreated struct {
	messaging.Message

	Id   model.Identifier
	Name string
	Url  string
	When time.Time
}
