///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPollWait test poller wait function
func TestPollWait(t *testing.T) {
	called := 3
	err := PollWait(func() (bool, error) {
		called--
		if called == 0 {
			return true, nil
		}
		return false, nil
	})
	assert.Nil(t, err)
	assert.Zero(t, called)
}
