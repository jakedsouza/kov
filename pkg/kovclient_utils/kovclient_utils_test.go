///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package kovclientutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetLCMClient test lcm client
func TestGetKOVClient(t *testing.T) {
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
		client, err := GetKOVClient(tc.host)
		assert.Equal(t, tc.outerr, err)
		if err != nil {
			assert.NotNil(t, client)
		}
	}
}
