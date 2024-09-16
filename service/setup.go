package service

import (
	"scoreplay/query"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	DB        *query.Query
	Validator *validator.Validate
}

func Setup() *Service {
	return &Service{
		DB:        query.Setup(),
		Validator: validator.New(),
	}
}
