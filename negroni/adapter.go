package negroniadapter

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/urfave/negroni"
)

type NegroniAdapter struct {
	core.RequestAccessor
	n *httpadapter.HandlerAdapter
}

func New(n *negroni.Negroni) *NegroniAdapter {
	return &NegroniAdapter{
		n: httpadapter.New(n),
	}
}

func (h *NegroniAdapter) Proxy(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return h.n.Proxy(event)
}
