package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	hello "hello/proto/hello"
)

type Hello struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Hello) Call(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Info("Received Hello.Call request")
	rsp.Msg = "Hello " + req.Name
	rsp.Code  = 100
	return nil
}

