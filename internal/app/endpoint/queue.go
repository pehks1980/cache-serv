package endpoint

import (
	"net/http"

	cache_serv "github.com/pehks1980/cache-serv/pkg/cache-serv"

	"github.com/gorilla/mux"
)

type queueSvc interface {
	Put(req *cache_serv.PutValueReq) error // записиать json ключ:значение во хранилище
	Get(req *cache_serv.GetValueReq) (*cache_serv.GetValueResp, error) // получить из хранилища значение по ключу
}

// регистрация роутинга путей
func RegisterPublicHTTP(queueSvc queueSvc) *mux.Router {
	//
	r := mux.NewRouter()
	r.HandleFunc("/api", putToQueue(queueSvc) ).Methods(http.MethodPut)
	r.HandleFunc("/api", getFromQueue(queueSvc) ).Methods(http.MethodGet)
	return r
}

// вьюха для put
func putToQueue(queueSvc queueSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		// TODO: parse req and call queueSvc.Put(...)
	}
}
// вьюха для get
func getFromQueue(queueSvc queueSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		// TODO: parse req and call queueSvc.Get(...)
	}
}
