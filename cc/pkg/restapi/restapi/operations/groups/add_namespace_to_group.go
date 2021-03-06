// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddNamespaceToGroupHandlerFunc turns a function with the right signature into a add namespace to group handler
type AddNamespaceToGroupHandlerFunc func(AddNamespaceToGroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddNamespaceToGroupHandlerFunc) Handle(params AddNamespaceToGroupParams) middleware.Responder {
	return fn(params)
}

// AddNamespaceToGroupHandler interface for that can handle valid add namespace to group params
type AddNamespaceToGroupHandler interface {
	Handle(AddNamespaceToGroupParams) middleware.Responder
}

// NewAddNamespaceToGroup creates a new http.Handler for the add namespace to group operation
func NewAddNamespaceToGroup(ctx *middleware.Context, handler AddNamespaceToGroupHandler) *AddNamespaceToGroup {
	return &AddNamespaceToGroup{Context: ctx, Handler: handler}
}

/*AddNamespaceToGroup swagger:route POST /cc/v1/groups/namespaces Groups addNamespaceToGroup

add namespace to group.

Optional extended description in Markdown.

*/
type AddNamespaceToGroup struct {
	Context *middleware.Context
	Handler AddNamespaceToGroupHandler
}

func (o *AddNamespaceToGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddNamespaceToGroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
