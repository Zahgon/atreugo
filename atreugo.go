package atreugo

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/valyala/fasthttp"
)

var (
	tcpNetworks   = []string{"tcp", "tcp4", "tcp6"}
	validNetworks = append(tcpNetworks, "unix")

	defaultLogger Logger = log.New(os.Stderr, "", log.LstdFlags)
)

// New create a new instance of Atreugo Server.
func New(cfg Config) *Atreugo { _ = "STUB: not implemented"; return nil }

func newFasthttpServer(cfg Config) *fasthttp.Server { _ = "STUB: not implemented"; return nil }

func (s *Atreugo) handler() fasthttp.RequestHandler {
	_ = "STUB: not implemented"
	return *new(fasthttp.RequestHandler)
}

// IsPreforkChild checks if the current thread/process is a child.
func IsPreforkChild() bool { _ = "STUB: not implemented"; return false }

// SaveMatchedRoutePath if enabled, adds the matched route path onto the ctx.UserValue context
// before invoking the handler.
// The matched route path is only added to handlers of routes that were
// registered when this option was enabled.
//
// It's deactivated by default.
func (s *Atreugo) SaveMatchedRoutePath(v bool) { _ = "STUB: not implemented"; return }

// RedirectTrailingSlash enables/disables automatic redirection if the current route
// can't be matched but a handler for the path with (without) the trailing slash exists.
// For example if /foo/ is requested but a route only exists for /foo, the
// client is redirected to /foo with http status code 301 for GET requests
// and 307 for all other request methods.
//
// It's activated by default.
func (s *Atreugo) RedirectTrailingSlash(v bool) { _ = "STUB: not implemented"; return }

// RedirectFixedPath if enabled, the router tries to fix the current request path, if no
// handle is registered for it.
// First superfluous path elements like ../ or // are removed.
// Afterwards the router does a case-insensitive lookup of the cleaned path.
// If a handle can be found for this route, the router makes a redirection
// to the corrected path with status code 301 for GET requests and 307 for
// all other request methods.
// For example /FOO and /..//Foo could be redirected to /foo.
// RedirectTrailingSlash is independent of this option.
//
// It's activated by default.
func (s *Atreugo) RedirectFixedPath(v bool) { _ = "STUB: not implemented"; return }

// HandleMethodNotAllowed if enabled, the router checks if another method is allowed for the
// current route, if the current request can not be routed.
// If this is the case, the request is answered with 'Method Not Allowed'
// and HTTP status code 405.
// If no other Method is allowed, the request is delegated to the NotFound
// handler.
//
// It's activated by default.
func (s *Atreugo) HandleMethodNotAllowed(v bool) { _ = "STUB: not implemented"; return }

// HandleOPTIONS if enabled, the router automatically replies to OPTIONS requests.
// Custom OPTIONS handlers take priority over automatic replies.
//
// It's activated by default.
func (s *Atreugo) HandleOPTIONS(v bool) { _ = "STUB: not implemented"; return }

// ServeConn serves HTTP requests from the given connection.
//
// ServeConn returns nil if all requests from the c are successfully served.
// It returns non-nil error otherwise.
//
// Connection c must immediately propagate all the data passed to Write()
// to the client. Otherwise requests' processing may hang.
//
// ServeConn closes c before returning.
func (s *Atreugo) ServeConn(c net.Conn) error { _ = "STUB: not implemented"; return nil }

// nolint:wrapcheck

// Serve serves incoming connections from the given listener.
//
// Serve blocks until the given listener returns permanent error.
func (s *Atreugo) Serve(ln net.Listener) error { _ = "STUB: not implemented"; return nil }

// nolint:wrapcheck

// nolint:wrapcheck

// NewVirtualHost returns a new sub-router for running more than one web site
// (such as company1.example.com and company2.example.com) on a single atreugo instance.
// Virtual hosts can be "IP-based", meaning that you have a different IP address
// for every web site, or "name-based", meaning that you have multiple names
// running on each IP address.
//
// The fact that they are running on the same atreugo instance is not apparent to the end user.
//
// If you pass multiples hostnames, all of them will have the same behaviour.
func (s *Atreugo) NewVirtualHost(hostnames ...string) *Router {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown gracefully shuts down the server without interrupting any active connections.
// Shutdown works by first closing all open listeners and then waiting indefinitely for
// all connections to return to idle and then shut down.
//
// When Shutdown is called, Serve, ListenAndServe, and ListenAndServeTLS immediately return
// nil. Make sure the program doesn't exit and waits instead for Shutdown to return.
//
// Shutdown does not close keepalive connections so it's recommended to set ReadTimeout
// and IdleTimeout to something else than 0.
func (s *Atreugo) Shutdown() (err error) { _ = "STUB: not implemented"; return nil }

// ShutdownWithContext gracefully shuts down the server without interrupting any active
// connections. ShutdownWithContext works by first closing all open listeners and then
// waiting for all connections to return to idle or context timeout and then shut down.
//
// When ShutdownWithContext is called, Serve, ListenAndServe, and ListenAndServeTLS
// immediately return nil. Make sure the program doesn't exit and waits instead for
// Shutdown to return.
//
// ShutdownWithContext does not close keepalive connections so it's recommended to set
// ReadTimeout and IdleTimeout to something else than 0.
func (s *Atreugo) ShutdownWithContext(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}
