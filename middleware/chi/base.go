package chi

import log "github.com/bayusamudra5502/go-backend-template/lib/log"

type ChiMiddleware struct {
	log log.Log
}

func NewChiMiddlware(log log.Log) (*ChiMiddleware) {
	return &ChiMiddleware{log}
}
