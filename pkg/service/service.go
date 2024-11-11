package service

import "github.com/EvaZay/rest-api/pkg/repository"

type Authorization interface {
}
type FinanceList interface {
}

type FinanceItem interface {
}

type Service struct {
	Authorization
	FinanceList
	FinanceItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
