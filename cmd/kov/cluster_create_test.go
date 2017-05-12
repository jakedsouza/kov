///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/supervised-io/kov/pkg/cluster"
	"github.com/supervised-io/kov/pkg/configfile/reader"
)

func TestCreateCluster(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	bufOut := &bytes.Buffer{}
	bufErr := &bytes.Buffer{}

	cli := NewCli()

	var (
		taskID        = `"1234-5678"`
		configContent = `{"name":"kube", "minNodes":1, "maxNodes": 3, "noOfMasters":1, "masterSize":"small","nodeSize":"small", "resourcePool":"pool", "nodeResourcePools":["pool1"], "managementNetwork": "testNetwork", "nodeNetwork":"testNodeNetwork", "publicNetwork":"testPublicNetwork"}`
		clusterName   = "kube"
		url           = "TestURL"
	)

	f, err := ioutil.TempFile("", "testCreateClusterTmp.json")
	ioutil.WriteFile(f.Name(), []byte(configContent), 0777)
	defer os.Remove(f.Name())
	assert.NoError(t, err)

	cli.SetOutput(bufOut, bufErr)

	cli.clusterCmd = &clusterCmd{
		configFile: f.Name(),
		url:        url,
	}

	clusterClient := cluster.NewMockClusterAPI(controller)
	cli.SetCluster(clusterClient)
	clusterConfig, err := reader.ParseClusterCreateConfig(f.Name())

	clusterClient.EXPECT().CreateCluster(gomock.Eq(clusterConfig)).Times(1).Return(&taskID, nil)
	clusterClient.EXPECT().GetTaskStatus(gomock.Eq(taskID)).Times(1).Return(true, nil)

	err = createClusterRun(cli)
	assert.NoError(t, err)
	assert.Contains(t, bufOut.String(), fmt.Sprintf("Creating cluster %s", clusterName))
	assert.Contains(t, bufOut.String(), fmt.Sprintf("Created cluster %s", clusterName))
}
