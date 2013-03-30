package data

import (
	"github.com/pjvds/promise/model"
)

type PromiseTicketRepository interface {
	Add(promise model.PromiseTicket) error
	All() []model.PromiseTicket
}
