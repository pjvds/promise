package memory

import (
	log "code.google.com/p/log4go"
	"github.com/pjvds/promise/model"
	"time"
)

type InMemoryPromiseRepository struct {
	promises     []model.PromiseTicket
	lastModified time.Time
}

type PromiseRepository interface {
	Add(promise model.PromiseTicket) error
	All() []PromiseTicket
}

func NewPromiseRepository(initialCapacity int) InMemoryPromiseRepository {
	return InMemoryPromiseRepository{
		promises:     make([]PromiseTicket, 0, initialCapacity),
		lastModified: time.Now(),
	}
}

func (r *InMemoryPromiseRepository) Add(promise PromiseTicket) error {
	r.promises = append(r.promises, promise)
	r.lastModified = time.Now()

	log.Printf("Added new promise which makes the total %v", len(r.promises))
	return nil
}

func (r *InMemoryPromiseRepository) All() []PromiseTicket {
	return r.promises
}

func (r *InMemoryPromiseRepository) LastModified() time.Time {
	return r.lastModified
}
