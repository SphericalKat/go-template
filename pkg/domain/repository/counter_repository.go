package repository

import "sample/ent"

type CounterRepository interface {
}

type counterRepositoryImpl struct {
	db *ent.Client
}
