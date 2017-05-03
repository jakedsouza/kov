///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package configfileutils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supervised-io/kov/gen/models"
	"github.com/supervised-io/kov/pkg/test_utils"
)

func TestReadConfigFile(t *testing.T) {
	sampleJSON := `
	{
		"name":"kube",
		"thumbprint":"123ASDF",
		"minNodes":1,
		"maxNodes":3,
		"noOfMasters":1,
		"masterSize":"small",
		"nodeSize":"small",
		"credentials":
			{
				"username":"testuser",
				"password":"testpassword!23"
			},
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
thumbprint: 123ASDF
minNodes: 1
maxNodes: 3
noOfMasters: 1
masterSize: small
nodeSize: small
credentials:
   username: testuser
   password: testpassword!23
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
		file, err := testutils.NewTempFile("", "testReadConfigFile", tc.extension)
		defer os.Remove(file.Name())
		assert.NoError(t, err)
		_, err1 := file.Write([]byte(tc.fileData))
		assert.NoError(t, err1)

		// read the file
		configBytes, err2 := ReadConfigFile(file.Name())
		assert.NoError(t, err2)

		clusterCreateConfig, err3 := ParseClusterCreateConfig(configBytes)
		assert.NoError(t, err3)

		sampleCreateClusterConfig, err4 := ParseClusterCreateConfig([]byte(sampleJSON))
		assert.NoError(t, err4)

		assert.Equal(t, clusterCreateConfig, sampleCreateClusterConfig)

	}
}

func TestParseClusterCreateConfig(t *testing.T) {
	sampleJSON := `
	{
		"name":"kube",
		"thumbprint":"123ASDF",
		"minNodes":1,
		"maxNodes":3,
		"noOfMasters":1,
		"masterSize":"small",
		"nodeSize":"small",
		"credentials":
			{
				"username":"testuser",
				"password":"testpassword!23"
			},
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
	var minNodeNum int32 = 1
	var masterNum int32 = 1
	var name = "kube"
	var managementNetwork = "testNetwork"

	testcases := []struct {
		extension         string
		sample            string
		name              *string
		thumbprint        string
		minNodes          *int32
		maxNodes          int32
		noOfMasters       *int32
		masterSize        models.InstanceSize
		nodeSize          models.InstanceSize
		resourcePool      string
		nodeResourcePools []string
		managementNetwork *string
		nodeNetwork       string
		publicNetwork     string
	}{
		{
			extension:         ".json",
			sample:            sampleJSON,
			name:              &name,
			thumbprint:        "123ASDF",
			minNodes:          &minNodeNum,
			maxNodes:          3,
			noOfMasters:       &masterNum,
			masterSize:        models.InstanceSize("small"),
			nodeSize:          models.InstanceSize("small"),
			resourcePool:      "pool",
			nodeResourcePools: []string{"pool1"},
			managementNetwork: &managementNetwork,
			nodeNetwork:       "testNodeNetwork",
			publicNetwork:     "testPublicNetwork",
		},
	}

	for _, tc := range testcases {
		// create the temp file
		file, err := testutils.NewTempFile("", "testReadConfigFile", tc.extension)
		defer os.Remove(file.Name())
		assert.NoError(t, err)
		_, err1 := file.Write([]byte(tc.sample))
		assert.NoError(t, err1)

		// read the file
		configBytes, err2 := ReadConfigFile(file.Name())
		assert.NoError(t, err2)

		clusterCreateConfig, err3 := ParseClusterCreateConfig(configBytes)
		assert.NoError(t, err3)

		assert.Equal(t, *clusterCreateConfig.Name, *tc.name)
		assert.Equal(t, clusterCreateConfig.Thumbprint, tc.thumbprint)
		assert.Equal(t, *clusterCreateConfig.MinNodes, *tc.minNodes)
		assert.Equal(t, clusterCreateConfig.MaxNodes, tc.maxNodes)
		assert.Equal(t, *clusterCreateConfig.NoOfMasters, *tc.noOfMasters)
		assert.Equal(t, clusterCreateConfig.MasterSize, tc.masterSize)
		assert.Equal(t, clusterCreateConfig.NodeSize, tc.nodeSize)
		assert.Equal(t, clusterCreateConfig.ResourcePool, tc.resourcePool)
		assert.Equal(t, clusterCreateConfig.NodeResourcePools, tc.nodeResourcePools)
		assert.Equal(t, *clusterCreateConfig.ManagementNetwork, *tc.managementNetwork)
		assert.Equal(t, clusterCreateConfig.NodeNetwork, tc.nodeNetwork)
		assert.Equal(t, clusterCreateConfig.PublicNetwork, tc.publicNetwork)
	}
}
