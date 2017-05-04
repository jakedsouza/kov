///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
package main

import (
	golog "log"
	"os"

	"github.com/casualjim/middlewares"
	"github.com/go-openapi/loads"
	"github.com/justinas/alice"
	"github.com/spf13/pflag"
	"github.com/supervised-io/kov"
	"github.com/supervised-io/kov/gen/restapi"
	"github.com/supervised-io/kov/gen/restapi/operations"
)

func main() {
	log := golog.New(os.Stderr, "[kovd] ", 0)
	log.Println("initializing")
	defer log.Println("exiting")

	pflag.Parse()

	app, err := kov.New("kovd", log)
	if err != nil {
		log.Fatalln(err)
	}

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewKovAPI(swaggerSpec)
	api.Logger = log.Printf

	server := restapi.NewServer(api)
	defer server.Shutdown()

	ml := &mwLogger{
		debug: golog.New(os.Stderr, "[DEBUG] [kovd] ", 0),
		info:  golog.New(os.Stderr, "[INFO] [kovd] ", 0),
		error: golog.New(os.Stderr, "[ERROR] [kovd] ", 0),
	}

	app.RegisterHandlers(api)

	handler := alice.New(
		middlewares.GzipMW(middlewares.DefaultCompression),
		middlewares.NewRecoveryMW(app.Info().Name, ml),
		middlewares.NewAuditMW(app.Info(), ml),
		middlewares.NewProfiler,
		middlewares.NewHealthChecksMW(app.Info().BasePath),
	).Then(api.Serve(nil))

	server.SetHandler(handler)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

type mwLogger struct {
	debug kov.Logger
	info  kov.Logger
	error kov.Logger
}

func (m *mwLogger) Debugf(msg string, args ...interface{}) {
	m.debug.Printf(msg, args...)
}
func (m *mwLogger) Infof(msg string, args ...interface{}) {
	m.info.Printf(msg, args...)
}
func (m *mwLogger) Errorf(msg string, args ...interface{}) {
	m.error.Printf(msg, args...)
}
