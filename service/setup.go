package service

import (
	"scoreplay/env"
	"scoreplay/query"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	DB        *query.Query
	Validator *validator.Validate
}

func Setup() *Service {
	return &Service{
		DB:        query.Setup(env.POSTGRES_HOST, env.POSTGRES_PORT, env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_DB),
		Validator: validator.New(),
	}
}
