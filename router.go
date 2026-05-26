package atreugo

import (
	"io/fs"
	"net/http"

	"github.com/valyala/fasthttp"
)

func defaultErrorView(ctx *RequestCtx, err error, statusCode int) {
	_ = "STUB: not implemented"
	return
}

func emptyView(_ *RequestCtx) error { _ = "STUB: not implemented"; return nil }

func buildOptionsView(url string, fn View, paths map[string][]string) View {
	_ = "STUB: not implemented"
	return *new(View)
}

func newRouter(cfg Config) *Router { _ = "STUB: not implemented"; return nil }

func (r *Router) mutable(v bool) { _ = "STUB: not implemented"; return }

func (r *Router) buildMiddlewares(m Middlewares) Middlewares {
	_ = "STUB: not implemented"
	return *new(Middlewares)
}

func (r *Router) getGroupFullPath(path string) string { _ = "STUB: not implemented"; return "" }

func (r *Router) handler(fn View, middle Middlewares) fasthttp.RequestHandler {
	_ = "STUB: not implemented"
	return *new(fasthttp.RequestHandler)
}

func (r *Router) handleMiddlewareError(ctx *RequestCtx, err error) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) handlePath(p *Path) { _ = "STUB: not implemented"; return }

// NewGroupPath returns a new router to group paths.
func (r *Router) NewGroupPath(path string) *Router { _ = "STUB: not implemented"; return nil }

// ListPaths returns all registered routes grouped by method.
func (r *Router) ListPaths() map[string][]string { _ = "STUB: not implemented"; return nil }

// Middlewares defines the middlewares (before, after and skip) in the order in which you want to execute them
// for the view or group
//
// WARNING: The previous middlewares configuration could be overridden.
func (r *Router) Middlewares(middlewares Middlewares) *Router {
	_ = "STUB: not implemented"
	return nil
}

// UseBefore registers the middlewares in the order in which you want to execute them
// before the execution of the view or group.
func (r *Router) UseBefore(fns ...Middleware) *Router { _ = "STUB: not implemented"; return nil }

// UseAfter registers the middlewares in the order in which you want to execute them
// after the execution of the view or group.
func (r *Router) UseAfter(fns ...Middleware) *Router { _ = "STUB: not implemented"; return nil }

// UseFinal registers the given middlewares to be executed in the order in which they are added,
// after the view or group has been executed. These middlewares will always be executed,
// even if a previous middleware or the view/group returned a response.
func (r *Router) UseFinal(fns ...FinalMiddleware) *Router { _ = "STUB: not implemented"; return nil }

// SkipMiddlewares registers the middlewares that you want to skip when executing the view or group.
func (r *Router) SkipMiddlewares(fns ...Middleware) *Router { _ = "STUB: not implemented"; return nil }

// GET shortcut for router.Path("GET", url, viewFn).
func (r *Router) GET(url string, viewFn View) *Path { _ = "STUB: not implemented"; return nil }

// HEAD shortcut for router.Path("HEAD", url, viewFn).
func (r *Router) HEAD(url string, viewFn View) *Path { _ = "STUB: not implemented"; return nil }

// OPTIONS shortcut for router.Path("OPTIONS", url, viewFn).
func (r *Router) OPTIONS(url string, viewFn View) *Path { _ = "STUB: not implemented"; return nil }

// POST shortcut for router.Path("POST", url, viewFn).
func (r *Router) POST(url string, viewFn View) *Path { _ = "STUB: not implemented"; return nil }

// PUT shortcut for router.Path("PUT", url, viewFn).
func (r *Router) PUT(url string, viewFn View) *Path { _ = "STUB: not implemented"; return nil }

// PATCH shortcut for router.Path("PATCH", url, viewFn).
func (r *Router) PATCH(url string, viewFn View) *Path { _ = "STUB: not implemented"; return nil }

// DELETE shortcut for router.Path("DELETE", url, viewFn).
func (r *Router) DELETE(url string, viewFn View) *Path { _ = "STUB: not implemented"; return nil }

// ANY shortcut for router.Path("*", url, viewFn)
//
// WARNING: Use only for routes where the request method is not important.
func (r *Router) ANY(url string, viewFn View) *Path { _ = "STUB: not implemented"; return nil }

// RequestHandlerPath wraps fasthttp request handler to atreugo view and registers it to
// the given path and method.
func (r *Router) RequestHandlerPath(method, url string, handler fasthttp.RequestHandler) *Path {
	_ = "STUB: not implemented"
	return nil
}

// NetHTTPPath wraps net/http handler to atreugo view and registers it to
// the given path and method.
//
// While this function may be used for easy switching from net/http to fasthttp/atreugo,
// it has the following drawbacks comparing to using manually written fasthttp/atreugo,
// request handler:
//
//   - A lot of useful functionality provided by fasthttp/atreugo is missing
//     from net/http handler.
//   - net/http -> fasthttp/atreugo handler conversion has some overhead,
//     so the returned handler will be always slower than manually written
//     fasthttp/atreugo handler.
//
// So it is advisable using this function only for quick net/http -> fasthttp
// switching. Then manually convert net/http handlers to fasthttp handlers.
// according to https://github.com/valyala/fasthttp#switching-from-nethttp-to-fasthttp.
func (r *Router) NetHTTPPath(method, url string, handler http.Handler) *Path {
	_ = "STUB: not implemented"
	return nil
}

// Static serves static files from the given file system root path.
//
// Make sure your program has enough 'max open files' limit aka
// 'ulimit -n' if root folder contains many files.
func (r *Router) Static(url, rootPath string) *Path { _ = "STUB: not implemented"; return nil }

// StaticFS serves static files from the given file system.
//
// Make sure your program has enough 'max open files' limit aka
// 'ulimit -n' if filesystem contains many files.
func (r *Router) StaticFS(url string, filesystem fs.FS) *Path {
	_ = "STUB: not implemented"
	return nil
}

// StaticCustom serves static files from the given file system settings
//
// Make sure your program has enough 'max open files' limit aka
// 'ulimit -n' if root folder contains many files.
func (r *Router) StaticCustom(url string, fs *StaticFS) *Path {
	_ = "STUB: not implemented"
	return nil
}

// ServeFile returns HTTP response containing compressed file contents
// from the given path
//
// HTTP response may contain uncompressed file contents in the following cases:
//
//   - Missing 'Accept-Encoding: gzip' request header.
//   - No write access to directory containing the file.
//
// Directory contents is returned if path points to directory.
func (r *Router) ServeFile(url, filePath string) *Path { _ = "STUB: not implemented"; return nil }

// Path registers a new view with the given path and method
//
// This function is intended for bulk loading and to allow the usage of less
// frequently used, non-standardized or custom methods (e.g. for internal
// communication with a proxy).
func (r *Router) Path(method, url string, viewFn View) *Path { _ = "STUB: not implemented"; return nil }
