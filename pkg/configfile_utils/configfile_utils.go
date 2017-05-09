///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package configfileutils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/loads/fmts"
	"github.com/supervised-io/kov/gen/models"

	yaml "gopkg.in/yaml.v2"
)

// ReadConfigFile reads a config file at path
func ReadConfigFile(path string) ([]byte, error) {
	if path == "" {
		return nil, errors.New("Config not found")
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
func ParseClusterCreateConfig(bytes []byte) (*models.ClusterConfig, error) {
	resp := &models.ClusterConfig{}
	err := json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
