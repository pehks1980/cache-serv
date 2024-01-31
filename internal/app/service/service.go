package service

import (
	"github.com/pehks1980/cache-serv/internal/pkg/model"
)

type repo interface {
	Get(key string) (string, error)
	Put(putReq *model.PutValue) error
}

type Service struct {
	repo repo
}

// конструктор Service
// возвращает указатель на структуру с интерфейсом
func New(repo repo) *Service {
	return &Service{repo: repo}
}
