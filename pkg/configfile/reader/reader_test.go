///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package reader

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supervised-io/kov/gen/models"
)

func TestReadConfigFile(t *testing.T) {
	sampleJSON := `
	{
		"name":"kube",
		"minNodes":1,
		"maxNodes":3,
		"noOfMasters":1,
		"masterSize":"small",
		"nodeSize":"small",
		"resourcePool":"pool",
		"nodeResourcePools":
			[
				"pool1"
			],
		"managementNetwork":"testNetwork",
		"nodeNetwork":"testNodeNetwork",
		"publicNetwork":"testPublicNetwork"
	}
	`

	sampleYML := `---
name: kube
minNodes: 1
maxNodes: 3
noOfMasters: 1
masterSize: small
nodeSize: small
resourcePool: pool
nodeResourcePools:
- pool1
managementNetwork: testNetwork
nodeNetwork: testNodeNetwork
publicNetwork: testPublicNetwork`

	testcases := []struct {
		extension   string
		fileData    string
		expectedErr error
	}{
		{
			extension:   ".json",
			fileData:    sampleJSON,
			expectedErr: nil,
		},
		{
			extension:   ".yaml",
			fileData:    sampleYML,
			expectedErr: nil,
		},
		{
			extension:   ".foo",
			fileData:    sampleYML,
			expectedErr: nil,
		},
		{
			extension:   ".bar",
			fileData:    sampleJSON,
			expectedErr: nil,
		},
		{
			extension:   "",
			fileData:    sampleJSON,
			expectedErr: nil,
		},
	}

	for _, tc := range testcases {
		// create the temp file
		file, err := ioutil.TempFile("", "testReadConfigFile"+tc.extension)
		defer os.Remove(file.Name())
		assert.NoError(t, err)
		_, err1 := file.Write([]byte(tc.fileData))
		assert.NoError(t, err1)

		// read the file
		configBytes, err2 := readConfigFile(file.Name())
		assert.NoError(t, err2)

		clusterCreateConfig := &models.ClusterConfig{}
		err3 := json.Unmarshal(configBytes, &clusterCreateConfig)
		assert.NoError(t, err3)

		sampleCreateClusterConfig := &models.ClusterConfig{}
		err4 := json.Unmarshal([]byte(sampleJSON), &sampleCreateClusterConfig)
		assert.NoError(t, err4)

		assert.Equal(t, clusterCreateConfig, sampleCreateClusterConfig)
	}
}

func TestParseClusterCreateConfig(t *testing.T) {
	sampleJSON := `
	{
		"name":"kube",
		"minNodes":1,
		"maxNodes":3,
		"noOfMasters":1,
		"masterSize":"small",
		"nodeSize":"small",
		"resourcePool":"pool",
		"nodeResourcePools":
			[
				"pool1"
			],
		"managementNetwork":"testNetwork",
		"nodeNetwork":"testNodeNetwork",
		"publicNetwork":"testPublicNetwork"
	}
	`

	var (
		minNodeNum            int32 = 1
		masterNum             int32 = 1
		testMaxNodes          int32 = 3
		testName                    = "kube"
		testManagementNetwork       = "testNetwork"

		name              = &testName
		minNodes          = &minNodeNum
		maxNodes          = &testMaxNodes
		noOfMasters       = &masterNum
		masterSize        = models.InstanceSize("small")
		nodeSize          = models.InstanceSize("small")
		resourcePool      = "pool"
		nodeResourcePools = []string{"pool1"}
		managementNetwork = &testManagementNetwork
		nodeNetwork       = "testNodeNetwork"
		publicNetwork     = "testPublicNetwork"
	)

	testcases := []struct {
		extension   string
		fileData    string
		expectedErr error
	}{
		{
			extension:   ".json",
			fileData:    sampleJSON,
			expectedErr: nil,
		},
	}

	for _, tc := range testcases {
		// create the temp file
		file, err := ioutil.TempFile("", "testReadConfigFile"+tc.extension)
		defer os.Remove(file.Name())
		assert.NoError(t, err)
		_, err1 := file.Write([]byte(tc.fileData))
		assert.NoError(t, err1)

		// read the file
		clusterCreateConfig, err2 := ParseClusterCreateConfig(file.Name())
		assert.NoError(t, err2)

		assert.Equal(t, *clusterCreateConfig.Name, *name)
		assert.Equal(t, *clusterCreateConfig.MinNodes, *minNodes)
		assert.Equal(t, clusterCreateConfig.MaxNodes, *maxNodes)
		assert.Equal(t, *clusterCreateConfig.NoOfMasters, *noOfMasters)
		assert.Equal(t, clusterCreateConfig.MasterSize, masterSize)
		assert.Equal(t, clusterCreateConfig.NodeSize, nodeSize)
		assert.Equal(t, clusterCreateConfig.ResourcePool, resourcePool)
		assert.Equal(t, clusterCreateConfig.NodeResourcePools, nodeResourcePools)
		assert.Equal(t, *clusterCreateConfig.ManagementNetwork, *managementNetwork)
		assert.Equal(t, clusterCreateConfig.NodeNetwork, nodeNetwork)
		assert.Equal(t, clusterCreateConfig.PublicNetwork, publicNetwork)

	}
}
