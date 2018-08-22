// Packge chilambda add Chi support for the aws-severless-go-api library.
// Uses the core package behind the scenes and exposes the New method to
// get a new instance and Proxy method to send request to the Chi mux.
package chiadapter

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/go-chi/chi"
)

// ChiLambda makes it easy to send API Gateway proxy events to a Chi
// Mux. The library transforms the proxy event into an HTTP request and then
// creates a proxy response object from the http.ResponseWriter
type ChiLambda struct {
	httpadapter.HandlerAdapter
	chiMux *httpadapter.HandlerAdapter
}

// New creates a new instance of the ChiLambda object.
// Receives an initialized *chi.Mux object - normally created with chi.NewRouter().
// It returns the initialized instance of the ChiLambda object.
func New(chi *chi.Mux) *ChiLambda {
	return &ChiLambda{chiMux: httpadapter.New(chi)}
}

// Proxy receives an API Gateway proxy event, transforms it into an http.Request
// object, and sends it to the chi.Mux for routing.
// It returns a proxy response object gneerated from the http.ResponseWriter.
func (g *ChiLambda) Proxy(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return g.chiMux.Proxy(req)
}
