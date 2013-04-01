package events

import (
	"github.com/pjvds/model"
	"time"
)

type NewTicketCreated struct {
	Id   model.Identifier
	Name string
	Url  string
	When time.Time
}
