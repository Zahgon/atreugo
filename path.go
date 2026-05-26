package atreugo

import (
	"time"
)

// Middlewares defines the middlewares (before, after and skip) in the order in which you want to execute them
// only for the view
//
// WARNING: The previous middlewares configuration could be overridden.
func (p *Path) Middlewares(middlewares Middlewares) *Path { _ = "STUB: not implemented"; return nil }

// UseBefore registers the middlewares in the order in which you want to execute them
// only before the execution of the view.
func (p *Path) UseBefore(fns ...Middleware) *Path { _ = "STUB: not implemented"; return nil }

// UseAfter registers the middlewares in the order in which you want to execute them
// only after the execution of the view.
func (p *Path) UseAfter(fns ...Middleware) *Path { _ = "STUB: not implemented"; return nil }

// UseFinal registers the given middlewares to be executed in the order in which they are added,
// after the view or group has been executed. These middlewares will always be executed,
// even if a previous middleware or the view/group returned a response.
func (p *Path) UseFinal(fns ...FinalMiddleware) *Path { _ = "STUB: not implemented"; return nil }

// SkipMiddlewares registers the middlewares that you want to skip only when executing the view.
func (p *Path) SkipMiddlewares(fns ...Middleware) *Path { _ = "STUB: not implemented"; return nil }

// Timeout sets the timeout and the error message to the view, which returns StatusRequestTimeout
// error with the given msg to the client if view didn't return during
// the given duration.
//
// The returned view may return StatusTooManyRequests error with the given
// msg to the client if there are more concurrent views are running
// at the moment than Server.Concurrency.
func (p *Path) Timeout(timeout time.Duration, msg string) *Path {
	_ = "STUB: not implemented"
	return nil
}

// TimeoutCode sets the timeout and the error message to the view, which returns an error with
// the given msg and status code to the client if view didn't return during
// the given duration.
//
// The returned view may return StatusTooManyRequests error with the given
// msg to the client if there are more concurrent views are running
// at the moment than Server.Concurrency.
func (p *Path) TimeoutCode(timeout time.Duration, msg string, statusCode int) *Path {
	_ = "STUB: not implemented"
	return nil
}
