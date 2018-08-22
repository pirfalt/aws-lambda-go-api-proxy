// Packge ginlambda add Gin support for the aws-severless-go-api library.
// Uses the core package behind the scenes and exposes the New method to
// get a new instance and Proxy method to send request to the Gin engine.
package ginadapter

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gin-gonic/gin"
)

// GinLambda makes it easy to send API Gateway proxy events to a Gin
// Engine. The library transforms the proxy event into an HTTP request and then
// creates a proxy response object from the http.ResponseWriter
type GinLambda struct {
	httpadapter.HandlerAdapter
	handler *httpadapter.HandlerAdapter
}

// New creates a new instance of the GinLambda object.
// Receives an initialized *gin.Engine object - normally created with gin.Default().
// It returns the initialized instance of the GinLambda object.
func New(gin *gin.Engine) *GinLambda {
	return &GinLambda{handler: httpadapter.New(gin)}
}

// Proxy receives an API Gateway proxy event, transforms it into an http.Request
// object, and sends it to the gin.Engine for routing.
// It returns a proxy response object gneerated from the http.ResponseWriter.
func (g *GinLambda) Proxy(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return g.handler.Proxy(req)
}
