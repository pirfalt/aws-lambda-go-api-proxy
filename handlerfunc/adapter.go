package handlerfunc

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

type HandlerFuncAdapter struct {
	httpadapter.HandlerAdapter
	handler *httpadapter.HandlerAdapter
}

func New(handlerFunc http.HandlerFunc) *HandlerFuncAdapter {
	return &HandlerFuncAdapter{
		handler: httpadapter.New(handlerFunc),
	}
}

func (h *HandlerFuncAdapter) Proxy(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return h.handler.Proxy(event)
}
