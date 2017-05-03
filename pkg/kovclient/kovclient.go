///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package kovclient

import (
	"errors"

	"github.com/go-openapi/runtime"
	httpTransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/supervised-io/kov/gen/client"
	"github.com/supervised-io/kov/gen/client/operations"
)

// GetClient get Kov client
func GetClient(host string) (*operations.Client, error) {
	if host == "" {
		return nil, errors.New("KOV endpoint not provided")
	}
	// get a new swagger client
	client := client.NewHTTPClient(nil)
	transport := client.Transport.(*httpTransport.Runtime)
	transport.Host = host
	configureTransport(transport)
	kovClient := operations.New(transport, strfmt.Default)
	return kovClient, nil
}

func configureTransport(rt *httpTransport.Runtime) {
	rt.Consumers["application/json"] = runtime.JSONConsumer()
	rt.Producers["application/json"] = runtime.JSONProducer()
}
