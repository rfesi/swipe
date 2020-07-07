//+build !swipe

// Code generated by Swipe v1.22.3. DO NOT EDIT.

//go:generate swipe
package jsonrpc

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/l-vitaly/go-kit/transport/http/jsonrpc"
	"github.com/swipe-io/swipe/fixtures/service"
	"net/http"
)

type httpError struct {
	code int
}

func (e httpError) Error() string {
	return http.StatusText(e.code)
}
func (e httpError) StatusCode() int {
	return e.code
}
func ErrorDecode(code int) (_ error) {
	switch code {
	default:
		return httpError{code: code}
	case -32001:
		return service.ErrUnauthorized{}
	}
}

func middlewareChain(middlewares []endpoint.Middleware) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		if len(middlewares) == 0 {
			return next
		}
		outer := middlewares[0]
		others := middlewares[1:]
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return outer(next)
	}
}

type ServiceInterfaceServerOption func(*serverServiceInterfaceOpts)
type serverServiceInterfaceOpts struct {
	genericServerOption           []jsonrpc.ServerOption
	genericEndpointMiddleware     []endpoint.Middleware
	createServerOption            []jsonrpc.ServerOption
	createEndpointMiddleware      []endpoint.Middleware
	deleteServerOption            []jsonrpc.ServerOption
	deleteEndpointMiddleware      []endpoint.Middleware
	getServerOption               []jsonrpc.ServerOption
	getEndpointMiddleware         []endpoint.Middleware
	getAllServerOption            []jsonrpc.ServerOption
	getAllEndpointMiddleware      []endpoint.Middleware
	testMethodServerOption        []jsonrpc.ServerOption
	testMethodEndpointMiddleware  []endpoint.Middleware
	testMethod2ServerOption       []jsonrpc.ServerOption
	testMethod2EndpointMiddleware []endpoint.Middleware
}

func ServiceInterfaceGenericServerOptions(v ...jsonrpc.ServerOption) (_ ServiceInterfaceServerOption) {
	return func(o *serverServiceInterfaceOpts) { o.genericServerOption = v }
}

func ServiceInterfaceGenericServerEndpointMiddlewares(v ...endpoint.Middleware) (_ ServiceInterfaceServerOption) {
	return func(o *serverServiceInterfaceOpts) { o.genericEndpointMiddleware = v }
}

func ServiceInterfaceCreateServerOptions(opt ...jsonrpc.ServerOption) (_ ServiceInterfaceServerOption) {
	return func(c *serverServiceInterfaceOpts) { c.createServerOption = opt }
}

func ServiceInterfaceCreateServerEndpointMiddlewares(opt ...endpoint.Middleware) (_ ServiceInterfaceServerOption) {
	return func(c *serverServiceInterfaceOpts) { c.createEndpointMiddleware = opt }
}

func ServiceInterfaceDeleteServerOptions(opt ...jsonrpc.ServerOption) (_ ServiceInterfaceServerOption) {
	return func(c *serverServiceInterfaceOpts) { c.deleteServerOption = opt }
}

func ServiceInterfaceDeleteServerEndpointMiddlewares(opt ...endpoint.Middleware) (_ ServiceInterfaceServerOption) {
	return func(c *serverServiceInterfaceOpts) { c.deleteEndpointMiddleware = opt }
}

func ServiceInterfaceGetServerOptions(opt ...jsonrpc.ServerOption) (_ ServiceInterfaceServerOption) {
	return func(c *serverServiceInterfaceOpts) { c.getServerOption = opt }
}

func ServiceInterfaceGetServerEndpointMiddlewares(opt ...endpoint.Middleware) (_ ServiceInterfaceServerOption) {
	return func(c *serverServiceInterfaceOpts) { c.getEndpointMiddleware = opt }
}

func ServiceInterfaceGetAllServerOptions(opt ...jsonrpc.ServerOption) (_ ServiceInterfaceServerOption) {
	return func(c *serverServiceInterfaceOpts) { c.getAllServerOption = opt }
}

func ServiceInterfaceGetAllServerEndpointMiddlewares(opt ...endpoint.Middleware) (_ ServiceInterfaceServerOption) {
	return func(c *serverServiceInterfaceOpts) { c.getAllEndpointMiddleware = opt }
}

func ServiceInterfaceTestMethodServerOptions(opt ...jsonrpc.ServerOption) (_ ServiceInterfaceServerOption) {
	return func(c *serverServiceInterfaceOpts) { c.testMethodServerOption = opt }
}

func ServiceInterfaceTestMethodServerEndpointMiddlewares(opt ...endpoint.Middleware) (_ ServiceInterfaceServerOption) {
	return func(c *serverServiceInterfaceOpts) { c.testMethodEndpointMiddleware = opt }
}

func ServiceInterfaceTestMethod2ServerOptions(opt ...jsonrpc.ServerOption) (_ ServiceInterfaceServerOption) {
	return func(c *serverServiceInterfaceOpts) { c.testMethod2ServerOption = opt }
}

func ServiceInterfaceTestMethod2ServerEndpointMiddlewares(opt ...endpoint.Middleware) (_ ServiceInterfaceServerOption) {
	return func(c *serverServiceInterfaceOpts) { c.testMethod2EndpointMiddleware = opt }
}
