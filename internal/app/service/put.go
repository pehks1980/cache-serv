package service

import (
	"log"

	"github.com/pehks1980/cache-serv/internal/pkg/model"

	cache_serv "github.com/pehks1980/cache-serv/pkg/cache-serv"
)

func (s *Service) Put(req *cache_serv.PutValueReq) error {
	if err := s.repo.Put(&model.PutValue{
		Key:   req.Key,
		Value: req.Value,
	}); err != nil {
		log.Printf("service/Put: put repo err: %v", err)
		return err
	}

	return nil
}
