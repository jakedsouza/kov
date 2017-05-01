///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supervised-io/kov"
)

func TestPrintVersion(t *testing.T) {
	c := NewCli()

	bufOut := &bytes.Buffer{}
	bufErr := &bytes.Buffer{}
	c.SetOutput(bufOut, bufErr)
	version := DevBuild
	printVersion(c)

	assert.NotNil(t, bufOut.String())
	assert.Contains(t, bufOut.String(), version)

	commit := "0.0.1-test"
	kov.Commit = "0.0.1-test"
	printVersion(c)
	assert.NotNil(t, bufOut.String())
	assert.Contains(t, bufOut.String(), commit)
}
