///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package kovclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetLCMClient test lcm client
func TestGetClient(t *testing.T) {
	testcases := []struct {
		host   string
		outerr error
	}{
		{
			host:   "localhost",
			outerr: nil,
		},
	}
	for _, tc := range testcases {
		client, err := GetClient(tc.host)
		assert.Equal(t, tc.outerr, err)
		if err != nil {
			assert.NotNil(t, client)
		}
	}
}
