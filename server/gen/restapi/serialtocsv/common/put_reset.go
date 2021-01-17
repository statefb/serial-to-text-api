// Code generated by go-swagger; DO NOT EDIT.

package common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutResetHandlerFunc turns a function with the right signature into a put reset handler
type PutResetHandlerFunc func(PutResetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutResetHandlerFunc) Handle(params PutResetParams) middleware.Responder {
	return fn(params)
}

// PutResetHandler interface for that can handle valid put reset params
type PutResetHandler interface {
	Handle(PutResetParams) middleware.Responder
}

// NewPutReset creates a new http.Handler for the put reset operation
func NewPutReset(ctx *middleware.Context, handler PutResetHandler) *PutReset {
	return &PutReset{Context: ctx, Handler: handler}
}

/*PutReset swagger:route PUT /reset Common putReset

reset status.

*/
type PutReset struct {
	Context *middleware.Context
	Handler PutResetHandler
}

func (o *PutReset) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutResetParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
