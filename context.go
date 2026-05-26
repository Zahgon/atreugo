package atreugo

import (
	"context"
	"fmt"
	"sync"

	"github.com/savsgio/gotils/bytes"
	"github.com/valyala/fasthttp"
)

var (
	attachedCtxKey = fmt.Sprintf("__attachedCtx::%s__", bytes.Rand(make([]byte, 15)))

	requestCtxPool = sync.Pool{
		New: func() any {
			ctx := new(RequestCtx)
			ctx.jsonMarshalFunc = defaultJSONMarshalFunc

			return ctx
		},
	}
)

// AcquireRequestCtx returns an empty RequestCtx instance from request context pool.
//
// The returned RequestCtx instance may be passed to ReleaseRequestCtx when it is
// no longer needed. This allows RequestCtx recycling, reduces GC pressure
// and usually improves performance.
func AcquireRequestCtx(ctx *fasthttp.RequestCtx) *RequestCtx { _ = "STUB: not implemented"; return nil }

// ReleaseRequestCtx returns ctx acquired via AcquireRequestCtx to request context pool.
//
// It is forbidden accessing ctx and/or its' members after returning
// it to request pool.
func ReleaseRequestCtx(ctx *RequestCtx) { _ = "STUB: not implemented"; return }

// RequestID returns the "X-Request-ID" header value.
func (ctx *RequestCtx) RequestID() []byte { _ = "STUB: not implemented"; return nil }

// Next pass control to the next middleware/view function.
func (ctx *RequestCtx) Next() error { _ = "STUB: not implemented"; return nil }

// SkipView sets flag to skip view execution in the current request
//
// Use it in before middlewares.
func (ctx *RequestCtx) SkipView() { _ = "STUB: not implemented"; return }

// AttachContext attach a context.Context to the RequestCtx
//
// WARNING: The extra context could not be itself.
func (ctx *RequestCtx) AttachContext(extraCtx context.Context) { _ = "STUB: not implemented"; return }

// AttachedContext returns the attached context.Context if exist.
func (ctx *RequestCtx) AttachedContext() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// MatchedRoutePath returns the matched route path
// if Atreugo.SaveMatchedRoutePath() is enabled.
func (ctx *RequestCtx) MatchedRoutePath() []byte { _ = "STUB: not implemented"; return nil }

// Value returns the value associated with attached context or this context for key,
// or nil if no value is associated with key. Successive calls to Value with
// the same key returns the same result.
//
// WARNING: The provided key should not be of type string or any other built-in
// to avoid extra allocating when assigning to an any, context keys often
// have concrete type struct{}. Alternatively, exported context key variables' static
// type should be a pointer or interface.
//
// If the key is of type string, try to use:
//
//	ctx.SetUserValue("myKey", "myValue")
//	ctx.UserValue("myKey")
//
// instead of:
//
//	ctx.AttachContext(context.WithValue(myCtx, "myKey", "myValue"))
//	ctx.Value("myKey")
//
// to avoid extra allocation.
func (ctx *RequestCtx) Value(key any) any { _ = "STUB: not implemented"; return *new(any) }
