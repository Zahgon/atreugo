package atreugo

import (
	"io"
)

func defaultJSONMarshalFunc(w io.Writer, body any) error { _ = "STUB: not implemented"; return nil }

// nolint:wrapcheck

// JSONResponse return response with body in json format.
func (ctx *RequestCtx) JSONResponse(body any, statusCode ...int) error {
	_ = "STUB: not implemented"
	return nil
}

// nolint:wrapcheck

// HTTPResponse return response with body in html format.
func (ctx *RequestCtx) HTTPResponse(body string, statusCode ...int) error {
	_ = "STUB: not implemented"
	return nil
}

// HTTPResponseBytes return response with body in html format.
func (ctx *RequestCtx) HTTPResponseBytes(body []byte, statusCode ...int) error {
	_ = "STUB: not implemented"
	return nil
}

// TextResponse return response with body in text format.
func (ctx *RequestCtx) TextResponse(body string, statusCode ...int) error {
	_ = "STUB: not implemented"
	return nil
}

// TextResponseBytes return response with body in text format.
func (ctx *RequestCtx) TextResponseBytes(body []byte, statusCode ...int) error {
	_ = "STUB: not implemented"
	return nil
}

// RawResponse returns response without encoding the body.
func (ctx *RequestCtx) RawResponse(body string, statusCode ...int) error {
	_ = "STUB: not implemented"
	return nil
}

// RawResponseBytes returns response without encoding the body.
func (ctx *RequestCtx) RawResponseBytes(body []byte, statusCode ...int) error {
	_ = "STUB: not implemented"
	return nil
}

// FileResponse return a streaming response with file data.
func (ctx *RequestCtx) FileResponse(fileName, filePath, mimeType string) error {
	_ = "STUB: not implemented"
	return nil
}

// nolint:errcheck

// RedirectResponse redirect request to an especific url.
func (ctx *RequestCtx) RedirectResponse(url string, statusCode int) error {
	_ = "STUB: not implemented"
	return nil
}

// ErrorResponse returns an error response.
func (ctx *RequestCtx) ErrorResponse(err error, statusCode ...int) error {
	_ = "STUB: not implemented"
	return nil
}
