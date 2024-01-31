package repository

import (
	"errors"
	"sync"
	"github.com/pehks1980/cache-serv/internal/pkg/model"
)

type cache struct {
	//defaultExpiration time.Duration
	items       map[string]string
}

type MemRepo struct{
	sync.RWMutex
	Cache cache
}

func NewMemRepo() *MemRepo {
	var lCache = &cache{
		items: make(map[string]string),
	}

	return &MemRepo{
		Cache: *lCache,
	}
}

func (mr *MemRepo) New (filename string) RepoIf{
	// init - filename is not used here
	var lCache = &cache{
		items: make(map[string]string),
	}

	return &MemRepo{
		Cache: *lCache,
	}
}

func (mr *MemRepo) Get(getReq string) (string, error) {
	mr.RWMutex.RLock()
	defer mr.RWMutex.RUnlock()

	if Value, ok := mr.Cache.items[getReq]; ok {
		return Value, nil
	}

	return "", errors.New("key not found")
}

func (mr *MemRepo) Put(putReq *model.PutValue) error {
	
	mr.RWMutex.Lock()
	mr.Cache.items[putReq.Key] = putReq.Value
	mr.RWMutex.Unlock()

	return nil
}
