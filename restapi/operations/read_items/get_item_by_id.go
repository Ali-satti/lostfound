// Code generated by go-swagger; DO NOT EDIT.

package read_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetItemByIDHandlerFunc turns a function with the right signature into a get item by Id handler
type GetItemByIDHandlerFunc func(GetItemByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetItemByIDHandlerFunc) Handle(params GetItemByIDParams) middleware.Responder {
	return fn(params)
}

// GetItemByIDHandler interface for that can handle valid get item by Id params
type GetItemByIDHandler interface {
	Handle(GetItemByIDParams) middleware.Responder
}

// NewGetItemByID creates a new http.Handler for the get item by Id operation
func NewGetItemByID(ctx *middleware.Context, handler GetItemByIDHandler) *GetItemByID {
	return &GetItemByID{Context: ctx, Handler: handler}
}

/* GetItemByID swagger:route GET /items/{itemId} Read Items getItemById

GetItemByID get item by Id API

*/
type GetItemByID struct {
	Context *middleware.Context
	Handler GetItemByIDHandler
}

func (o *GetItemByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetItemByIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
