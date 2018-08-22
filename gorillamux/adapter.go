package gorillamux

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gorilla/mux"
)

type GorillaMuxAdapter struct {
	core.RequestAccessor
	router *httpadapter.HandlerAdapter
}

func New(router *mux.Router) *GorillaMuxAdapter {
	return &GorillaMuxAdapter{
		router: httpadapter.New(router),
	}
}

func (h *GorillaMuxAdapter) Proxy(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return h.router.Proxy(event)
}
