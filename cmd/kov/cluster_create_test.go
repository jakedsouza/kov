///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"

	"net/http/httptest"

	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/assert"
	"github.com/supervised-io/kov/pkg/test_utils"
)

func TestCreateCluster(t *testing.T) {

	bufOut := &bytes.Buffer{}
	bufErr := &bytes.Buffer{}

	cli := NewCli()

	const (
		response      = http.StatusAccepted
		taskID        = `"1234-5678"`
		taskResp      = `{"context":{"cellId":"cellname"},"id":"testjob","state":"completed"}`
		configContent = `{"name":"kube","thumbprint":"123ASDF", "minNodes":1, "maxNodes": 3, "noOfMasters":1, "masterSize":"small","nodeSize":"small", "credentials":{"username":"testuser", "password":"testpassword!23"}, "resourcePool":"pool", "nodeResourcePools":["pool1"], "managementNetwork": "testNetwork", "nodeNetwork":"testNodeNetwork", "publicNetwork":"testPublicNetwork"}`
		clusterName   = "kube"
	)

	// test server
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add(runtime.HeaderContentType, "application/json")
		// Post for run cluster
		if req.Method == http.MethodPost {
			rw.WriteHeader(http.StatusAccepted)
			rw.Write([]byte(taskID))
		}
		// Get for job status
		if req.Method == http.MethodGet {
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(taskResp))
		}
	}))
	defer server.Close()

	u, _ := url.Parse(server.URL)

	f, err := testutils.NewTempFile("", "", ".json")
	ioutil.WriteFile(f.Name(), []byte(configContent), 0777)
	defer os.Remove(f.Name())
	assert.NoError(t, err)

	cli.SetOutput(bufOut, bufErr)

	cli.clusterCmd = &clusterCmd{
		configFile: f.Name(),
		url:        u.Host,
	}

	err = createClusterRun(cli)
	assert.NoError(t, err)
	assert.Contains(t, bufOut.String(), fmt.Sprintf("Creating cluster %s", clusterName))
	assert.Contains(t, bufOut.String(), fmt.Sprintf("Created cluster %s", clusterName))
}
