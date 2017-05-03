///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package reader

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/loads/fmts"
	"github.com/supervised-io/kov/gen/models"

	"regexp"

	yaml "gopkg.in/yaml.v2"
)

// ReadConfigFile reads a config file at path
func readConfigFile(path string) ([]byte, error) {
	if path == "" {
		return nil, errors.New("Configfile path not provided")
	}

	var jsonDoc json.RawMessage
	// read the file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	dataStr := strings.TrimSpace(string(data))

	// if string starts with {  - assume json else yaml
	if strings.HasPrefix(dataStr, "{") {
		jsonDoc = json.RawMessage(data)
	} else {
		var yamlDoc map[interface{}]interface{}

		if err := yaml.Unmarshal(data, &yamlDoc); err != nil {
			return nil, err
		}
		jsonDoc, err = fmts.YAMLToJSON(yamlDoc)
		if err != nil {
			return nil, err
		}
	}

	return jsonDoc, nil
}

// ParseClusterCreateConfig parse bytes as ClusterConfig
func ParseClusterCreateConfig(path string) (*models.ClusterConfig, error) {
	bytes, err := readConfigFile(path)
	if err != nil {
		return nil, err
	}
	config := &models.ClusterConfig{}
	err = json.Unmarshal(bytes, &config)

	if err != nil {
		return nil, err
	}

	// check required field
	err = validateClusterCreateConfig(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// validateClusterCreateConfig validate the clusterCreateConfig
func validateClusterCreateConfig(config *models.ClusterConfig) error {
	// check credentials
	if config.Credentials == nil {
		return errors.New("Credential not provided")
	} else if config.Credentials.Username == "" {
		return errors.New("Username not provided")
	} else if config.Credentials.Password == "" {
		return errors.New("Password not provided")
	} else if len(config.Credentials.Password) < 6 {
		return errors.New("Password length should be at least 6")
	}

	// check ManagementNetwork
	if config.ManagementNetwork == nil {
		return errors.New("Management Network not provided")
	}

	// check MasterSize
	sizeTypeList := map[models.InstanceSize]bool{
		models.InstanceSizeSmall:     true,
		models.InstanceSizeMedium:    true,
		models.InstanceSizeLarge:     true,
		models.InstanceSizeHuge:      true,
		models.InstanceSizeGinormous: true,
	}
	if config.MasterSize != "" {
		if !sizeTypeList[config.MasterSize] {
			return errors.New("Master Size invalid")
		}
	}

	// check MinNodes
	if *config.MinNodes < 1 {
		return errors.New("Minimum nodes number should be at least 1")
	}

	// check MaxNodes
	if config.MaxNodes == 0 {
		config.MaxNodes = *config.MinNodes
	} else if config.MaxNodes < *config.MinNodes {
		return errors.New("Maximum nodes number should be larger or equal than Minimum node number")
	}

	// check Name
	if config.Name == nil {
		return errors.New("Name not provided")
	}
	matched, err := regexp.MatchString("^[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?(\\.[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?)*$", *config.Name)
	if err != nil || !matched {
		return errors.New("Name format invalid")
	}

	// check NodeOfMasters
	if config.NoOfMasters == nil {
		return errors.New("Number of master nodes to create not provided")
	} else if *config.NoOfMasters < 1 {
		return errors.New("Number of master nodes should be at least 1")
	}

	// check NodeNetwork
	// defaults to management network
	if config.NodeNetwork == "" {
		config.NodeNetwork = *config.ManagementNetwork
	}

	// check NodeSize
	if config.NodeSize != "" {
		if !sizeTypeList[config.NodeSize] {
			return errors.New("Node Size invalid")
		}
	}

	// check PublicNetwork
	// defaults to management network
	if config.PublicNetwork == "" {
		config.PublicNetwork = *config.ManagementNetwork
	}

	// check ResourcePool
	if len(config.ResourcePool) < 1 {
		return errors.New("Resource Pool not provided")
	}

	// check Thumbprint
	if config.Thumbprint == "" {
		return errors.New("Thumbprint not provided")
	}
	if len(config.Thumbprint) < 57 {
		return errors.New("Length of Thumbprint should be at least 57")
	}
	if matched, err := regexp.MatchString("[a-fA-F0-9:]+", config.Thumbprint); err != nil || !matched {
		return errors.New("Thumbprint format invalid")
	}

	return nil
}
