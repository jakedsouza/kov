///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package testutils

import (
	"fmt"
	"os"
	"path/filepath"
)

// NewTempFile new tempary files
func NewTempFile(dir, filename, suffix string) (*os.File, error) {
	if dir == "" {
		dir = os.TempDir()
	}

	path := filepath.Join(dir, fmt.Sprintf("%s%s", filename, suffix))
	if _, err := os.Stat(path); err != nil {
		return os.Create(path)
	}

	// Give up if error
	return nil, fmt.Errorf("could not create file of the form %s", path)
}
