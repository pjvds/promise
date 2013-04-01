package data

import (
	"github.com/pjvds/promise/model"
)

type CallbackAttemptRepository interface {
	Add(attempt *model.CallbackAttempt) error
	All() []model.CallbackAttempt
}
