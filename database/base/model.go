package base

import "time"

type Model interface {
	Name() string
	GetDB() Database
	SetCreatedAt(createdAt time.Time)
	SetUpdatedAt(updatedAt time.Time)
}
