///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package kovclientutils

import (
	"crypto/tls"
	"errors"
	"net/http"

	"github.com/go-openapi/runtime"
	httpTransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/supervised-io/kov/gen/client"
	"github.com/supervised-io/kov/gen/client/operations"
)

// GetKOVClient get Kov client
func GetKOVClient(host string) (*operations.Client, error) {
	if host == "" {
		return nil, errors.New("KOV endpoint not provided")
	}
	// get a new swagger client
	client := client.NewHTTPClient(nil)

	transport := client.Transport.(*httpTransport.Runtime)
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	transport.Transport = &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	transport.Host = host
	configureTransport(transport)
	kovClient := operations.New(transport, strfmt.Default)
	return kovClient, nil
}

func configureTransport(rt *httpTransport.Runtime) {
	rt.Consumers["application/json"] = runtime.JSONConsumer()
	rt.Consumers["application/octet-stream"] = runtime.ByteStreamConsumer()
	rt.Producers["application/json"] = runtime.JSONProducer()
	rt.Producers["application/octet-stream"] = runtime.ByteStreamProducer()
}
