package repository

import "github.com/goalistdeveloper/go-template/ent"

type CounterRepository interface {
}

type counterRepositoryImpl struct {
	db *ent.Client
}
