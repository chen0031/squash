// Code generated by go-swagger; DO NOT EDIT.

package debugattchment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddDebugAttachmentHandlerFunc turns a function with the right signature into a add debug attachment handler
type AddDebugAttachmentHandlerFunc func(AddDebugAttachmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddDebugAttachmentHandlerFunc) Handle(params AddDebugAttachmentParams) middleware.Responder {
	return fn(params)
}

// AddDebugAttachmentHandler interface for that can handle valid add debug attachment params
type AddDebugAttachmentHandler interface {
	Handle(AddDebugAttachmentParams) middleware.Responder
}

// NewAddDebugAttachment creates a new http.Handler for the add debug attachment operation
func NewAddDebugAttachment(ctx *middleware.Context, handler AddDebugAttachmentHandler) *AddDebugAttachment {
	return &AddDebugAttachment{Context: ctx, Handler: handler}
}

/*AddDebugAttachment swagger:route POST /debugattachment debugattchment addDebugAttachment

Request squash to attach to a running container.

A debug attachment instructs squash to attach to a container. Debug attachment is made of
  - image: The container image we are debugging. this is used for extra validation, as placing breakpoints on the wrong binary can lead to unexpected results. if not provided huerisrtics will be used to identify it.
  - debugger: Type of debugger to use. "dlv" and "gdb" are supported now.
  - match_request: Whether to match this attachment to a debug request. This is used in automated use-cases to guarantee that the attachment will be noticed.


*/
type AddDebugAttachment struct {
	Context *middleware.Context
	Handler AddDebugAttachmentHandler
}

func (o *AddDebugAttachment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddDebugAttachmentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
