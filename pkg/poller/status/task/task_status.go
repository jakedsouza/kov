///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package task

import "time"

// PollWait takes a poller function as an argument which should return a boolean(done) or error.
// If done, function returns nil else returns the error
func PollWait(poller func() (bool, error)) error {
	for {
		waitStart := time.Now()
		if done, err := poller(); err != nil {
			return err
		} else if done {
			return nil
		}

		rest := time.Second*3 - time.Since(waitStart)
		time.Sleep(rest)
	}
}
