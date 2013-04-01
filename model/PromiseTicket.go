package model

import (
	"time"
)

type PromiseTicket struct {
	Id        Identifier
	Name      string
	Url       string
	When      time.Time
	Finalized bool
}
