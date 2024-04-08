package service

import (
	"github.com/goalistdeveloper/go-template/pkg/domain/repository"
)

type CounterService interface {
}

type counterServiceImpl struct {
	counterRepository repository.CounterRepository
	emailService      EmailService
}
