///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package handlerutils

import (
	"context"
	"net/http"

	"github.com/pborman/uuid"
	"github.com/supervised-io/kov"
)

type key uint8

const (
	// RequestIDKey key for retrieving the request id from the context
	RequestIDKey key = iota
)

// NewContextWithRequestScopedValues returns a new context with the new logger and requestID setup
func NewContextWithRequestScopedValues(app kov.Application, req *http.Request) context.Context {
	if req == nil {
		app.Log().Println("no request scoped context present, skipping")
		return nil
	}

	var reqID string
	if req.Header != nil {
		reqID = req.Header.Get("X-Request-ID")
	}
	if reqID == "" {
		reqID = uuid.New()
	}

	ctx := req.Context()
	reqIDCtx := context.WithValue(ctx, RequestIDKey, reqID)
	return reqIDCtx
}

// RequestIDFromContext returns the request id given a context
func RequestIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return uuid.New()
	}
	if str, ok := ctx.Value(RequestIDKey).(string); ok {
		return str
	}
	return uuid.New()
}
