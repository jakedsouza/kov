///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

//Package kov
package kov

// this code needs to be considered by go, so it can't be in a file that starts with _ or .
// NO TESTS

//go:generate swagger generate server -t gen -A kov -f ./swagger/swagger.yml --exclude-main --flag-strategy pflag
//go:generate swagger generate client -t gen -A kov -f ./swagger/swagger.yml
