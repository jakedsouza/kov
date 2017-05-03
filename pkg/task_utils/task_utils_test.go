///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package taskutils

// TestPollWait test poller wait function
func TestPollWait(t *testing.T) {
	called := 3
	err := PollWait(func() (bool, error) {
		called = called - 1
		if called == 0 {
			return true, nil
		}
		return false, nil
	})
	assert.Nil(t, err)
	assert.Zero(t, called)
}
