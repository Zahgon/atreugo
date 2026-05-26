package atreugo

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/prefork"
)

func panicf(s string, args ...any) { _ = "STUB: not implemented"; return }

func viewToHandler(view View, errorView ErrorView) fasthttp.RequestHandler {
	_ = "STUB: not implemented"
	return *new(fasthttp.RequestHandler)
}

func isEqual(v1, v2 any) bool { _ = "STUB: not implemented"; return false }

func isNil(v any) bool { _ = "STUB: not implemented"; return false }

func middlewaresInclude(ms []Middleware, fn Middleware) bool {
	_ = "STUB: not implemented"
	return false
}

func appendMiddlewares(dst, src []Middleware, skip ...Middleware) []Middleware {
	_ = "STUB: not implemented"
	return nil
}

func newPreforkServerBase(s *Atreugo) *prefork.Prefork { _ = "STUB: not implemented"; return nil }
