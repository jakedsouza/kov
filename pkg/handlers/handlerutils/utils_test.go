///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package handlerutils

import (
	"bytes"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supervised-io/kov"
)

func getTestItems() []*http.Request {
	var res []*http.Request
	req1, _ := http.NewRequest("", "", nil)
	req2, _ := http.NewRequest("", "", nil)
	req2.Header.Set("X-Request-ID", "dummy-request-id")
	res = append(res, req1)
	res = append(res, req2)
	return res
}

func TestNewContextFromParentContext(t *testing.T) {
	req := getTestItems()
	var buf bytes.Buffer
	logger := log.New(&buf, "test-kov-handlerutils: ", log.Lshortfile)
	app, err := kov.New("test-kov-handlerutils", logger)
	assert.NoError(t, err)

	for _, v := range req {
		newCtx := NewContextWithRequestScopedValues(app, v)
		if v == nil {
			assert.Nil(t, newCtx)
		}
		if hdr := v.Header.Get("X-Request-ID"); hdr != "" {
			assert.Equal(t, hdr, newCtx.Value(RequestIDKey).(string))
		} else {
			assert.NotNil(t, newCtx.Value(RequestIDKey))
		}
	}
}
