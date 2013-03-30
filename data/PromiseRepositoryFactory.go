package data

import ()

type PromiseRepositoryFactory interface {
	CreatePromiseTicketRepository() (*PromiseTicketRepository, error)
}
