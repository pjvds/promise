package main

import (
	"log"
	"time"
)

type PromiseRepository struct {
	promises     []PromiseTicket
	lastModified time.Time
}

func NewPromiseRepository(initialCapacity int) PromiseRepository {
	return PromiseRepository{
		promises:     make([]PromiseTicket, 0, initialCapacity),
		lastModified: time.Now(),
	}
}

func (r *PromiseRepository) Add(promise PromiseTicket) error {
	r.promises = append(r.promises, promise)
	r.lastModified = time.Now()

	log.Printf("Added new promise which makes the total %v", len(r.promises))
	return nil
}

func (r *PromiseRepository) All() []PromiseTicket {
	return r.promises
}

func (r *PromiseRepository) LastModified() time.Time {
	return r.lastModified
}
