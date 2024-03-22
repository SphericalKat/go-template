package service

import (
	"sample/pkg/domain/repository"
)

type CounterService interface {
}

type counterServiceImpl struct {
	counterRepository repository.CounterRepository
}
