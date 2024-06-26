package cache_serv

type (
	GetValueReq struct {
		Key string
	}

	PutValueReq struct {
		Key   	string `json:"key"`
		Value	string `json:"value"`
		Expiry 	string `json:"expiry"`
	}
)
