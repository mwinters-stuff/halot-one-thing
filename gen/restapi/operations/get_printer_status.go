// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetPrinterStatusHandlerFunc turns a function with the right signature into a get printer status handler
type GetPrinterStatusHandlerFunc func(GetPrinterStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPrinterStatusHandlerFunc) Handle(params GetPrinterStatusParams) middleware.Responder {
	return fn(params)
}

// GetPrinterStatusHandler interface for that can handle valid get printer status params
type GetPrinterStatusHandler interface {
	Handle(GetPrinterStatusParams) middleware.Responder
}

// NewGetPrinterStatus creates a new http.Handler for the get printer status operation
func NewGetPrinterStatus(ctx *middleware.Context, handler GetPrinterStatusHandler) *GetPrinterStatus {
	return &GetPrinterStatus{Context: ctx, Handler: handler}
}

/*
	GetPrinterStatus swagger:route GET /printer-status getPrinterStatus

GetPrinterStatus get printer status API
*/
type GetPrinterStatus struct {
	Context *middleware.Context
	Handler GetPrinterStatusHandler
}

func (o *GetPrinterStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPrinterStatusParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
